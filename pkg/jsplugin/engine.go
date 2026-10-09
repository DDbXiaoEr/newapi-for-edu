package jsplugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"math/big"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/grafana/sobek"
	"github.com/grafana/sobek/parser"
)

const (
	DefaultCallTimeout = 5 * time.Second
	DefaultConcurrency = 8
)

var ErrCallAdmissionTimeout = errors.New("plugin call admission timed out")

const hookErrorMessageLimit = 512

// HookError reports a JavaScript exception thrown by a plugin hook. Message
// is the sanitized JS error message with engine prefixes stripped; it is safe
// to surface to API callers.
type HookError struct {
	Hook    string
	Message string
	wrapped error
}

func (e *HookError) Error() string {
	if e == nil || e.wrapped == nil {
		return "plugin hook failed"
	}
	return e.wrapped.Error()
}

func (e *HookError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.wrapped
}

// ResultError reports a hook result that CallInto could not store: one the
// host codec cannot encode (a NaN, an infinity, a cycle) or one the target's
// type rejects. Its message is the codec's.
type ResultError struct {
	err error
}

func (e *ResultError) Error() string {
	return e.err.Error()
}

func (e *ResultError) Unwrap() error {
	return e.err
}

func newHookError(hook, rawMessage string, wrapped error) *HookError {
	var b strings.Builder
	b.Grow(len(rawMessage))
	count := 0
	for _, r := range rawMessage {
		if count >= hookErrorMessageLimit {
			break
		}
		if r < 0x20 || (r >= 0x80 && r <= 0x9F) {
			r = ' '
		}
		b.WriteRune(r)
		count++
	}
	message := b.String()
	if message == "" {
		message = "plugin hook failed"
	}
	return &HookError{Hook: hook, Message: message, wrapped: wrapped}
}

func hookErrorFromException(hook string, exc *sobek.Exception, wrapped error) (hookErr *HookError) {
	// Reading message/toString executes plugin getters, which can throw again
	// and panic sobek. By this point the caller's recover is already consumed,
	// so a second panic would crash the process; fall back to a blank message.
	defer func() {
		if recover() != nil {
			hookErr = newHookError(hook, "", wrapped)
		}
	}()
	raw := ""
	if exc != nil {
		if val := exc.Value(); val != nil && !sobek.IsUndefined(val) && !sobek.IsNull(val) {
			gotMessage := false
			if obj, ok := val.(*sobek.Object); ok {
				if msg := obj.Get("message"); msg != nil && !sobek.IsUndefined(msg) && !sobek.IsNull(msg) {
					raw = msg.String()
					gotMessage = true
				}
			}
			if !gotMessage {
				if exported, ok := val.Export().(string); ok {
					raw = exported
				} else {
					raw = val.String()
				}
			}
		}
	}
	return newHookError(hook, raw, wrapped)
}

var forbiddenSyntax = regexp.MustCompile(`(?m)(^|[^A-Za-z0-9_$])(async|await|import)([^A-Za-z0-9_$]|$)`)

type Options struct {
	Key         string
	Version     string
	Timeout     time.Duration
	Concurrency int
	Now         func() time.Time
	Log         func(string)
}

type Engine struct {
	key       string
	version   string
	timeout   time.Duration
	now       func() time.Time
	log       func(string)
	module    *sobek.SourceTextModuleRecord
	pool      chan *runtimeInstance
	semaphore chan struct{}
}

type runtimeInstance struct {
	runtime    *sobek.Runtime
	module     sobek.ModuleInstance
	logContext *runtimeLogContext
}

type runtimeLogContext struct {
	context context.Context
}

// Compile performs upload-time syntax checks and compiles an ESM plugin once.
// All Sobek-specific module and runtime handling is intentionally kept here.
func Compile(source string, options Options) (*Engine, error) {
	if match := forbiddenSyntax.FindString(sourceWithoutCommentsAndStrings(source)); match != "" {
		return nil, fmt.Errorf("unsupported plugin syntax %q: plugins must be synchronous and cannot import modules", strings.TrimSpace(match))
	}

	resolve := func(_ any, specifier string) (sobek.ModuleRecord, error) {
		return nil, fmt.Errorf("plugin imports are disabled: %s", specifier)
	}
	// Plugin source is untrusted; without this option a sourceMappingURL
	// comment makes the parser read arbitrary server files via os.ReadFile.
	module, err := sobek.ParseModule(options.Key+".js", source, resolve, parser.WithDisableSourceMaps)
	if err != nil {
		return nil, fmt.Errorf("compile plugin: %w", err)
	}
	if err = module.Link(); err != nil {
		return nil, fmt.Errorf("link plugin: %w", err)
	}

	timeout := options.Timeout
	if timeout <= 0 {
		timeout = DefaultCallTimeout
	}
	concurrency := options.Concurrency
	if concurrency <= 0 {
		concurrency = DefaultConcurrency
	}
	now := options.Now
	if now == nil {
		now = time.Now
	}

	engine := &Engine{
		key:       options.Key,
		version:   options.Version,
		timeout:   timeout,
		now:       now,
		log:       options.Log,
		module:    module,
		semaphore: make(chan struct{}, concurrency),
		pool:      make(chan *runtimeInstance, concurrency),
	}
	instance, err := engine.newRuntime(context.Background())
	if err != nil {
		return nil, err
	}
	instance.logContext.context = nil
	engine.putRuntime(instance)
	return engine, nil
}

// Export returns one module export without exposing Sobek values outside the
// engine boundary. It is used for declarative exports such as meta.
func (e *Engine) Export(ctx context.Context, exportName string) (result any, err error) {
	select {
	case e.semaphore <- struct{}{}:
		defer func() { <-e.semaphore }()
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	instance, err := e.getRuntime(ctx)
	if err != nil {
		return nil, err
	}
	reusable := true
	defer func() {
		instance.runtime.ClearInterrupt()
		instance.logContext.context = nil
		if reusable {
			e.putRuntime(instance)
		}
	}()
	timedOut := errors.New("plugin export timed out")
	stopInterrupt := watchRuntimeContext(instance.runtime, ctx, e.timeout, timedOut)
	defer stopInterrupt()
	defer func() {
		if recovered := recover(); recovered != nil {
			reusable = false
			switch value := recovered.(type) {
			case *sobek.InterruptedError:
				err = fmt.Errorf("plugin %s@%s export %s interrupted: %v", e.key, e.version, exportName, value.Value())
			case *sobek.Exception:
				err = fmt.Errorf("plugin %s@%s export %s failed: %v", e.key, e.version, exportName, value)
			default:
				panic(recovered)
			}
		}
	}()
	value := instance.module.GetBindingValue(exportName)
	if value == nil || sobek.IsUndefined(value) {
		return nil, fmt.Errorf("plugin export %q not found", exportName)
	}
	return value.Export(), nil
}

// HasExport reports whether a module export exists. Optional contract hooks
// should be detected with this method instead of relying on engine errors.
func (e *Engine) HasExport(ctx context.Context, exportName string) (bool, error) {
	select {
	case e.semaphore <- struct{}{}:
		defer func() { <-e.semaphore }()
	case <-ctx.Done():
		return false, ctx.Err()
	}
	instance, err := e.getRuntime(ctx)
	if err != nil {
		return false, err
	}
	defer func() {
		instance.logContext.context = nil
		e.putRuntime(instance)
	}()
	value := instance.module.GetBindingValue(exportName)
	return value != nil && !sobek.IsUndefined(value), nil
}

// HasCallablePath reports whether an exported value, or a nested member below
// it, exists and is callable.
func (e *Engine) HasCallablePath(ctx context.Context, exportName string, members ...string) (found bool, err error) {
	select {
	case e.semaphore <- struct{}{}:
		defer func() { <-e.semaphore }()
	case <-ctx.Done():
		return false, ctx.Err()
	}
	instance, err := e.getRuntime(ctx)
	if err != nil {
		return false, err
	}
	reusable := true
	defer func() {
		instance.runtime.ClearInterrupt()
		instance.logContext.context = nil
		if reusable {
			e.putRuntime(instance)
		}
	}()
	timedOut := errors.New("plugin inspection timed out")
	stopInterrupt := watchRuntimeContext(instance.runtime, ctx, e.timeout, timedOut)
	defer stopInterrupt()
	defer func() {
		if recovered := recover(); recovered != nil {
			reusable = false
			hookName := strings.Join(append([]string{exportName}, members...), ".")
			switch value := recovered.(type) {
			case *sobek.InterruptedError:
				err = fmt.Errorf("plugin %s@%s hook %s inspection interrupted: %v", e.key, e.version, hookName, value.Value())
			case *sobek.Exception:
				err = fmt.Errorf("plugin %s@%s hook %s inspection failed: %v", e.key, e.version, hookName, value)
			default:
				panic(recovered)
			}
		}
	}()
	value, _, found := resolveExportPath(instance, exportName, members)
	if !found {
		return false, nil
	}
	_, callable := sobek.AssertFunction(value)
	return callable, nil
}

// RawJSON is a hook argument given as JSON text: the hook receives JSON.parse
// of it, members in order, whether it is an argument or inside one, with no Go
// values built in between; text nested too deep for JSON.parse is decoded with
// the host codec instead. Strings of the call's result may share the bytes, so
// the caller must never change them once passed.
type RawJSON []byte

// Call invokes one named module export and returns its JSON-compatible value.
func (e *Engine) Call(ctx context.Context, exportName string, args ...any) (result any, err error) {
	result, _, err = e.call(ctx, callOptions{}, exportName, nil, args...)
	return result, err
}

// CallInto invokes one named module export and stores its result in target as
// the host codec decodes the encoding of the value Call returns, without that
// round trip for plain results. A result the codec cannot encode or target's
// type rejects is a *ResultError, and target may then hold part of it.
func (e *Engine) CallInto(ctx context.Context, target any, exportName string, args ...any) error {
	_, _, err := e.call(ctx, callOptions{target: target}, exportName, nil, args...)
	return err
}

// CallJSONInto is CallInto with the semantics of JSON.stringify's text of the
// result instead of the host codec's encoding: objects keep their member
// order, which a json.RawMessage in target receives as text, members whose
// value is undefined are left out and toJSON runs.
func (e *Engine) CallJSONInto(ctx context.Context, target any, exportName string, args ...any) error {
	_, _, err := e.call(ctx, callOptions{target: target, stringify: true}, exportName, nil, args...)
	return err
}

// CallMember invokes a function stored on an exported object, such as a
// renderer in the renderers export.
func (e *Engine) CallMember(ctx context.Context, exportName, memberName string, args ...any) (result any, err error) {
	result, _, err = e.call(ctx, callOptions{}, exportName, []string{memberName}, args...)
	return result, err
}

// CallPath invokes a function nested below an exported object. It is used for
// protocol hooks such as protocols.openai_responses.renderEvents.
func (e *Engine) CallPath(ctx context.Context, exportName string, members []string, args ...any) (result any, err error) {
	result, _, err = e.call(ctx, callOptions{}, exportName, members, args...)
	return result, err
}

// CallPathWithAdmissionTimeout gives long-lived observers a separate bound for
// waiting on JavaScript capacity. Once admitted, the hook receives the
// engine's full execution timeout instead of inheriting time already spent in
// the semaphore queue.
func (e *Engine) CallPathWithAdmissionTimeout(
	ctx context.Context,
	admissionTimeout time.Duration,
	exportName string,
	members []string,
	args ...any,
) (result any, err error) {
	result, _, err = e.call(ctx, callOptions{admissionTimeout: admissionTimeout}, exportName, members, args...)
	return result, err
}

// CallPathWithMemberJSON is CallPathWithAdmissionTimeout that also returns the
// result's member as the text JSON.stringify writes for it, which keeps the
// member order of its objects. The text is nil when member is empty or the
// result has no such member.
func (e *Engine) CallPathWithMemberJSON(
	ctx context.Context,
	admissionTimeout time.Duration,
	member string,
	exportName string,
	members []string,
	args ...any,
) (result any, text json.RawMessage, err error) {
	return e.call(ctx, callOptions{admissionTimeout: admissionTimeout, jsonMember: member}, exportName, members, args...)
}

type callOptions struct {
	admissionTimeout time.Duration
	target           any
	stringify        bool
	jsonMember       string
}

func (e *Engine) call(
	ctx context.Context,
	opts callOptions,
	exportName string,
	members []string,
	args ...any,
) (result any, text json.RawMessage, err error) {
	if err = e.acquireCallSlot(ctx, opts.admissionTimeout); err != nil {
		return nil, nil, err
	}
	defer func() { <-e.semaphore }()

	instance, err := e.getRuntime(ctx)
	if err != nil {
		return nil, nil, err
	}
	reusable := true
	defer func() {
		instance.runtime.ClearInterrupt()
		instance.logContext.context = nil
		if reusable {
			e.putRuntime(instance)
		}
	}()

	hookName := strings.Join(append([]string{exportName}, members...), ".")
	timedOut := errors.New("plugin call timed out")
	stopInterrupt := watchRuntimeContext(instance.runtime, ctx, e.timeout, timedOut)
	defer stopInterrupt()
	defer func() {
		if recovered := recover(); recovered != nil {
			reusable = false
			switch value := recovered.(type) {
			case *sobek.InterruptedError:
				err = fmt.Errorf("plugin %s@%s hook %s interrupted: %v", e.key, e.version, hookName, value.Value())
			case *sobek.Exception:
				wrapped := fmt.Errorf("plugin %s@%s hook %s failed: %v", e.key, e.version, hookName, value)
				err = hookErrorFromException(hookName, value, wrapped)
			default:
				panic(recovered)
			}
		}
	}()

	value, resolvedHookName, found := resolveExportPath(instance, exportName, members)
	hookName = resolvedHookName
	if !found {
		if len(members) == 0 {
			return nil, nil, fmt.Errorf("plugin export %q not found", exportName)
		}
		return nil, nil, fmt.Errorf("plugin hook %q not found", hookName)
	}
	if value == nil || sobek.IsUndefined(value) {
		return nil, nil, fmt.Errorf("plugin export %q not found", exportName)
	}
	callable, ok := sobek.AssertFunction(value)
	if !ok {
		return nil, nil, fmt.Errorf("plugin hook %q is not a function", hookName)
	}

	callArgs := make([]sobek.Value, len(args))
	for index, arg := range args {
		converted, convErr := argumentValue(instance.runtime, arg)
		if convErr != nil {
			return nil, nil, fmt.Errorf("plugin %s@%s hook %s argument %d: %w", e.key, e.version, hookName, index+1, convErr)
		}
		callArgs[index] = converted
	}

	value, err = callable(sobek.Undefined(), callArgs...)
	if err != nil {
		var interrupted *sobek.InterruptedError
		if errors.As(err, &interrupted) {
			reusable = false
			return nil, nil, fmt.Errorf("plugin %s@%s hook %s failed: %w", e.key, e.version, hookName, err)
		}
		wrapped := fmt.Errorf("plugin %s@%s hook %s failed: %w", e.key, e.version, hookName, err)
		var exc *sobek.Exception
		if errors.As(err, &exc) {
			return nil, nil, hookErrorFromException(hookName, exc, wrapped)
		}
		return nil, nil, wrapped
	}

	var memberText json.RawMessage
	switch {
	case opts.target != nil && opts.stringify:
		encoded, stringifyErr := stringifyJSON(instance.runtime, value)
		if stringifyErr != nil {
			return nil, nil, &ResultError{err: stringifyErr}
		}
		if err = common.Unmarshal(encoded, opts.target); err != nil {
			return nil, nil, &ResultError{err: err}
		}
	case opts.target != nil:
		exported := value.Export()
		encoded, marshalErr := common.Marshal(exported)
		if marshalErr != nil {
			return nil, nil, &ResultError{err: marshalErr}
		}
		if err = common.Unmarshal(encoded, opts.target); err != nil {
			return nil, nil, &ResultError{err: err}
		}
		result = exported
	default:
		result = value.Export()
	}
	if opts.jsonMember != "" && !sobek.IsUndefined(value) && !sobek.IsNull(value) {
		member := value.ToObject(instance.runtime).Get(opts.jsonMember)
		if member != nil && !sobek.IsUndefined(member) {
			memberText, err = stringifyJSON(instance.runtime, member)
			if err != nil {
				return nil, nil, &ResultError{err: err}
			}
		}
	}
	return result, memberText, nil
}

func argumentValue(rt *sobek.Runtime, arg any) (sobek.Value, error) {
	if raw, ok := arg.(RawJSON); ok {
		return parseJSONValue(rt, raw)
	}
	converted, _ := pluginValue(rt, arg, 0)
	return rt.ToValue(converted), nil
}

func parseJSONValue(rt *sobek.Runtime, raw RawJSON) (sobek.Value, error) {
	parsed, err := jsonParse(rt, raw)
	if err == nil {
		return parsed, nil
	}
	var decoded any
	if unmarshalErr := common.Unmarshal(raw, &decoded); unmarshalErr != nil {
		return sobek.Undefined(), err
	}
	return rt.ToValue(decoded), nil
}

func jsonParse(rt *sobek.Runtime, raw RawJSON) (sobek.Value, error) {
	parse, err := jsonFunction(rt, "parse")
	if err != nil {
		return sobek.Undefined(), err
	}
	return parse(sobek.Undefined(), rt.ToValue(string(raw)))
}

func stringifyJSON(rt *sobek.Runtime, value sobek.Value) ([]byte, error) {
	stringify, err := jsonFunction(rt, "stringify")
	if err != nil {
		return nil, err
	}
	text, err := stringify(sobek.Undefined(), value)
	if err != nil {
		return nil, err
	}
	if text == nil || sobek.IsUndefined(text) || sobek.IsNull(text) {
		return nil, nil
	}
	return []byte(text.String()), nil
}

func jsonFunction(rt *sobek.Runtime, name string) (sobek.Callable, error) {
	jsonValue := rt.Get("JSON")
	if jsonValue == nil || sobek.IsUndefined(jsonValue) || sobek.IsNull(jsonValue) {
		return nil, fmt.Errorf("JSON is not available")
	}
	fn, ok := sobek.AssertFunction(jsonValue.ToObject(rt).Get(name))
	if !ok {
		return nil, fmt.Errorf("JSON.%s is not a function", name)
	}
	return fn, nil
}

const maxPluginValueDepth = 512

func pluginValue(rt *sobek.Runtime, v any, depth int) (converted any, changed bool) {
	if depth > maxPluginValueDepth {
		return v, false
	}
	switch typed := v.(type) {
	case nil, bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
		float32, float64, json.Number, *big.Int, []byte, []string, map[string]string, map[string][]string,
		sobek.Value:
		return v, false
	case RawJSON:
		if parsed, err := jsonParse(rt, typed); err == nil {
			return parsed, true
		}
		var decoded any
		if err := common.Unmarshal(typed, &decoded); err != nil {
			return v, false
		}
		return decoded, true
	case map[string]any:
		var copied map[string]any
		for key, item := range typed {
			if item, changed = pluginValue(rt, item, depth+1); changed {
				if copied == nil {
					copied = maps.Clone(typed)
				}
				copied[key] = item
			}
		}
		if copied != nil {
			return copied, true
		}
		return typed, false
	case []any:
		var copied []any
		for index, item := range typed {
			if item, changed = pluginValue(rt, item, depth+1); changed {
				if copied == nil {
					copied = slices.Clone(typed)
				}
				copied[index] = item
			}
		}
		if copied != nil {
			return copied, true
		}
		return typed, false
	default:
		if reflect.TypeOf(v) == nil {
			return v, false
		}
		return v, false
	}
}

func (e *Engine) acquireCallSlot(ctx context.Context, admissionTimeout time.Duration) error {
	if admissionTimeout <= 0 {
		select {
		case e.semaphore <- struct{}{}:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	timer := time.NewTimer(admissionTimeout)
	defer timer.Stop()
	select {
	case e.semaphore <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		if err := ctx.Err(); err != nil {
			return err
		}
		return fmt.Errorf("%w: plugin %s@%s", ErrCallAdmissionTimeout, e.key, e.version)
	}
}

func resolveExportPath(instance *runtimeInstance, exportName string, members []string) (sobek.Value, string, bool) {
	value := instance.module.GetBindingValue(exportName)
	hookName := exportName
	if value == nil || sobek.IsUndefined(value) || sobek.IsNull(value) {
		return nil, hookName, false
	}
	for _, member := range members {
		hookName += "." + member
		object := value.ToObject(instance.runtime)
		own := slices.Contains(object.GetOwnPropertyNames(), member)
		if !own {
			return nil, hookName, false
		}
		value = object.Get(member)
		if value == nil || sobek.IsUndefined(value) || sobek.IsNull(value) {
			return nil, hookName, false
		}
	}
	return value, hookName, true
}

// Idle runtimes are bounded by the execution limit and survive GC. Start with
// one instance and grow only when concurrent work actually needs more.
func (e *Engine) putRuntime(instance *runtimeInstance) {
	select {
	case e.pool <- instance:
	default:
	}
}

func (e *Engine) getRuntime(ctx context.Context) (*runtimeInstance, error) {
	select {
	case instance := <-e.pool:
		instance.logContext.context = ctx
		return instance, nil
	default:
		return e.newRuntime(ctx)
	}
}

// A timeout callback must finish before its runtime can be reused. Merely
// stopping a timer does not wait for an already-started Interrupt call.
func watchRuntimeContext(runtime *sobek.Runtime, ctx context.Context, timeout time.Duration, timeoutError error) func() {
	callContext, cancel := context.WithTimeoutCause(ctx, timeout, timeoutError)
	interrupted := make(chan struct{})
	stop := context.AfterFunc(callContext, func() {
		runtime.Interrupt(context.Cause(callContext))
		close(interrupted)
	})
	return func() {
		if !stop() {
			<-interrupted
		}
		cancel()
	}
}

func (e *Engine) newRuntime(ctx context.Context) (instance *runtimeInstance, err error) {
	runtime := sobek.New()
	logContext := &runtimeLogContext{context: ctx}
	logOutput := e.log
	if logOutput == nil {
		logOutput = func(message string) {
			logger.LogDebug(logContext.context, "task_plugin subsystem=runtime event=console message=%q", message)
		}
	}
	if err := injectGlobals(runtime, func() string {
		return fmt.Sprintf("[plugin:%s@%s]", e.key, e.version)
	}, e.now, logOutput); err != nil {
		return nil, fmt.Errorf("inject plugin utils: %w", err)
	}
	timedOut := errors.New("plugin initialization timed out")
	stopInterrupt := watchRuntimeContext(runtime, ctx, e.timeout, timedOut)
	defer func() {
		stopInterrupt()
		runtime.ClearInterrupt()
	}()
	defer func() {
		if recovered := recover(); recovered != nil {
			if interrupted, ok := recovered.(*sobek.InterruptedError); ok {
				instance = nil
				err = fmt.Errorf("initialize plugin %s@%s: %v", e.key, e.version, interrupted.Value())
				return
			}
			panic(recovered)
		}
	}()
	promise := runtime.CyclicModuleRecordEvaluate(e.module, func(_ any, specifier string) (sobek.ModuleRecord, error) {
		return nil, fmt.Errorf("plugin imports are disabled: %s", specifier)
	})
	if promise.State() != sobek.PromiseStateFulfilled {
		return nil, fmt.Errorf("evaluate plugin: %v", promise.Result().Export())
	}
	return &runtimeInstance{
		runtime:    runtime,
		module:     runtime.GetModuleInstance(e.module),
		logContext: logContext,
	}, nil
}

func sourceWithoutCommentsAndStrings(source string) string {
	var output strings.Builder
	output.Grow(len(source))
	quote := byte(0)
	escaped := false
	lineComment := false
	blockComment := false
	for i := 0; i < len(source); i++ {
		current := source[i]
		next := byte(0)
		if i+1 < len(source) {
			next = source[i+1]
		}
		if lineComment {
			if current == '\n' {
				lineComment = false
				output.WriteByte('\n')
			} else {
				output.WriteByte(' ')
			}
			continue
		}
		if blockComment {
			if current == '*' && next == '/' {
				blockComment = false
				output.WriteString("  ")
				i++
			} else {
				output.WriteByte(' ')
			}
			continue
		}
		if quote != 0 {
			output.WriteByte(' ')
			if escaped {
				escaped = false
			} else if current == '\\' {
				escaped = true
			} else if current == quote {
				quote = 0
			}
			continue
		}
		if current == '/' && next == '/' {
			lineComment = true
			output.WriteString("  ")
			i++
			continue
		}
		if current == '/' && next == '*' {
			blockComment = true
			output.WriteString("  ")
			i++
			continue
		}
		if current == '\'' || current == '"' || current == '`' {
			quote = current
			output.WriteByte(' ')
			continue
		}
		output.WriteByte(current)
	}
	return output.String()
}
