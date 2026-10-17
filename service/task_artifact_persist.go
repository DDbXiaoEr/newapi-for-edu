package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/system_setting"
	"github.com/QuantumNous/new-api/types"
	"github.com/bytedance/gopkg/util/gopool"
	"golang.org/x/net/http/httpguts"
)

const (
	taskArtifactPersistTimeout     = 30 * time.Minute
	taskArtifactPersistConcurrency = 4
	taskArtifactFetchTimeout       = 15 * time.Minute
	taskArtifactDataURLMaxEncoded  = 96 << 20
)

var taskArtifactKeyPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._~-]{0,127}$`)

// TaskArtifactFetchRequest describes how to fetch one artifact from upstream.
// It is intentionally transport-neutral so service code does not depend on the
// relay/channel package.
type TaskArtifactFetchRequest struct {
	URL            string
	Method         string
	Headers        map[string]string
	Body           []byte
	Credentialless bool
}

// TaskArtifactSource lists artifacts and builds upstream content requests for
// one task. It is implemented by the relay layer to avoid a service -> relay
// import cycle.
type TaskArtifactSource interface {
	ListArtifacts() ([]types.TaskArtifact, error)
	BuildContentRequest(artifactKey string) (*TaskArtifactFetchRequest, error)
}

// TaskArtifactSourceFactory is wired from main so polling and download
// handlers can resolve task plugin artifacts without importing relay.
var TaskArtifactSourceFactory func(task *model.Task) (TaskArtifactSource, error)

var taskArtifactPersistSemaphore = make(chan struct{}, taskArtifactPersistConcurrency)

// QueueTaskArtifactPersistence saves a completed task's artifacts in the
// background. It never blocks the caller and drops work when the background
// queue is saturated; the download path can persist artifacts lazily instead.
func QueueTaskArtifactPersistence(task *model.Task) {
	if task == nil || task.Status != model.TaskStatusSuccess || !GetTaskArtifactStore().Enabled() {
		return
	}
	taskCopy := *task
	if task.PrivateData.Artifacts != nil {
		taskCopy.PrivateData.Artifacts = make(map[string]model.TaskStoredArtifact, len(task.PrivateData.Artifacts))
		for key, value := range task.PrivateData.Artifacts {
			taskCopy.PrivateData.Artifacts[key] = value
		}
	}
	gopool.Go(func() {
		select {
		case taskArtifactPersistSemaphore <- struct{}{}:
		default:
			logger.LogWarn(context.Background(), fmt.Sprintf("task artifact persistence queue is full; skipping task %s", taskCopy.TaskID))
			return
		}
		defer func() { <-taskArtifactPersistSemaphore }()

		ctx, cancel := context.WithTimeout(context.Background(), taskArtifactPersistTimeout)
		defer cancel()
		if err := PersistTaskArtifacts(ctx, &taskCopy); err != nil {
			logger.LogWarn(ctx, fmt.Sprintf("failed to persist artifacts for task %s: %v", taskCopy.TaskID, err))
		}
	})
}

// PersistTaskArtifacts stores every artifact produced by a successful task.
// Individual failures are logged so one broken artifact cannot block others.
func PersistTaskArtifacts(ctx context.Context, task *model.Task) error {
	if task == nil || task.Status != model.TaskStatusSuccess || !GetTaskArtifactStore().Enabled() {
		return nil
	}
	if ctx == nil {
		ctx = context.Background()
	}
	artifacts, source, err := taskArtifactsForPersistence(task)
	if err != nil {
		return err
	}
	var firstErr error
	for _, artifact := range artifacts {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if err := persistTaskArtifact(ctx, task, source, artifact); err != nil {
			logger.LogWarn(ctx, fmt.Sprintf("failed to persist task %s artifact %q: %v", task.TaskID, artifact.Key, err))
			if firstErr == nil {
				firstErr = err
			}
		}
	}
	return firstErr
}

// PersistTaskArtifact stores one artifact. It is used by the download path when
// background persistence has not completed yet.
func PersistTaskArtifact(ctx context.Context, task *model.Task, artifactKey string) error {
	if task == nil || task.Status != model.TaskStatusSuccess || !GetTaskArtifactStore().Enabled() {
		return nil
	}
	if ctx == nil {
		ctx = context.Background()
	}
	artifactKey = strings.TrimSpace(artifactKey)
	if !taskArtifactKeyPattern.MatchString(artifactKey) {
		return errors.New("invalid artifact key")
	}
	artifacts, source, err := taskArtifactsForPersistence(task)
	if err != nil {
		return err
	}
	for _, artifact := range artifacts {
		if artifact.Key == artifactKey {
			return persistTaskArtifact(ctx, task, source, artifact)
		}
	}
	return nil
}

func taskArtifactsForPersistence(task *model.Task) ([]types.TaskArtifact, TaskArtifactSource, error) {
	var source TaskArtifactSource
	if TaskArtifactSourceFactory != nil {
		resolved, err := TaskArtifactSourceFactory(task)
		if err != nil {
			logger.LogWarn(context.Background(), fmt.Sprintf("task %s artifact source unavailable: %v", task.TaskID, err))
		} else if resolved != nil {
			source = resolved
			artifacts, listErr := source.ListArtifacts()
			if listErr != nil {
				logger.LogWarn(context.Background(), fmt.Sprintf("task %s artifact listing failed: %v", task.TaskID, listErr))
				source = nil
			} else {
				valid := make([]types.TaskArtifact, 0, len(artifacts))
				for _, artifact := range artifacts {
					if validatePersistableTaskArtifact(artifact) {
						valid = append(valid, artifact)
					}
				}
				if len(valid) > 0 {
					return valid, source, nil
				}
			}
		}
	}
	if LegacyVideoTaskAvailable(task) {
		return []types.TaskArtifact{{Key: "video", Type: "video"}}, nil, nil
	}
	return nil, nil, nil
}

func validatePersistableTaskArtifact(artifact types.TaskArtifact) bool {
	if len(artifact.Key) == 0 || !taskArtifactKeyPattern.MatchString(artifact.Key) {
		return false
	}
	if artifact.Key != strings.TrimSpace(artifact.Key) || artifact.Type != strings.TrimSpace(artifact.Type) {
		return false
	}
	switch artifact.Type {
	case "video", "audio", "image", "file":
	default:
		return false
	}
	return len(artifact.MimeType) <= 255 && !strings.ContainsAny(artifact.MimeType, "\r\n")
}

func persistTaskArtifact(ctx context.Context, task *model.Task, source TaskArtifactSource, artifact types.TaskArtifact) error {
	if task == nil {
		return errors.New("task is nil")
	}
	if stored, ok := task.PrivateData.Artifacts[artifact.Key]; ok && stored.ObjectKey != "" {
		return nil
	}
	store := GetTaskArtifactStore()
	if ref, err := store.Resolve(ctx, task, artifact.Key); err == nil && ref != nil {
		mergeStoredTaskArtifact(ctx, task, artifact, ref)
		return nil
	} else if err != nil && !errors.Is(err, ErrTaskArtifactNotFound) {
		return err
	}

	request, err := buildTaskArtifactFetchRequest(task, source, artifact.Key)
	if err != nil {
		return err
	}
	proxy := taskArtifactChannelProxy(task)
	body, contentType, _, err := fetchTaskArtifact(ctx, request, proxy)
	if err != nil {
		return err
	}
	defer body.Close()
	if contentType != "" {
		artifact.MimeType = contentType
	}
	ref, err := store.Persist(ctx, task, artifact, body)
	if err != nil {
		return err
	}
	mergeStoredTaskArtifact(ctx, task, artifact, ref)
	return nil
}

func buildTaskArtifactFetchRequest(task *model.Task, source TaskArtifactSource, artifactKey string) (*TaskArtifactFetchRequest, error) {
	if source != nil {
		request, err := source.BuildContentRequest(artifactKey)
		if err != nil {
			return nil, err
		}
		if request != nil {
			return request, nil
		}
	}
	if artifactKey == "video" && LegacyVideoTaskAvailable(task) {
		return &TaskArtifactFetchRequest{
			URL:            task.GetResultURL(),
			Method:         http.MethodGet,
			Credentialless: true,
		}, nil
	}
	return nil, errors.New("artifact content request is unavailable")
}

func mergeStoredTaskArtifact(ctx context.Context, task *model.Task, artifact types.TaskArtifact, ref *StoredArtifactRef) {
	if task == nil || ref == nil {
		return
	}
	stored := model.TaskStoredArtifact{
		Key:       artifact.Key,
		Type:      artifact.Type,
		MimeType:  ref.MimeType,
		Size:      ref.Size,
		Backend:   ref.Backend,
		Bucket:    ref.Bucket,
		ObjectKey: ref.ObjectKey,
	}
	if stored.MimeType == "" {
		stored.MimeType = artifact.MimeType
	}
	if task.PrivateData.Artifacts == nil {
		task.PrivateData.Artifacts = make(map[string]model.TaskStoredArtifact, 1)
	}
	task.PrivateData.Artifacts[stored.Key] = stored
	if task.ID > 0 {
		if err := model.MergeStoredArtifacts(task.ID, []model.TaskStoredArtifact{stored}); err != nil {
			logger.LogWarn(ctx, fmt.Sprintf("failed to persist task %s artifact manifest: %v", task.TaskID, err))
		}
	}
}

func taskArtifactChannelProxy(task *model.Task) string {
	if task == nil || task.ChannelId == 0 {
		return ""
	}
	channel, err := model.CacheGetChannel(task.ChannelId)
	if err != nil || channel == nil {
		return ""
	}
	return strings.TrimSpace(channel.GetSetting().Proxy)
}

// LegacyVideoTaskAvailable reports whether a successful non-plugin task has a
// downloadable legacy video result.
func LegacyVideoTaskAvailable(task *model.Task) bool {
	if task == nil || task.Status != model.TaskStatusSuccess ||
		(task.PrivateData.Execution != nil && task.PrivateData.Execution.TaskPlugin != nil && strings.TrimSpace(task.PrivateData.Execution.TaskPlugin.Key) != "") ||
		task.Platform == constant.TaskPlatformSuno {
		return false
	}
	resultURL := strings.TrimSpace(task.GetResultURL())
	if !strings.HasPrefix(resultURL, "http://") && !strings.HasPrefix(resultURL, "https://") {
		return false
	}
	switch constant.NormalizeTaskAction(task.Action) {
	case constant.TaskActionImageToVideo,
		constant.TaskActionTextToVideo,
		constant.TaskActionFirstTailToVideo,
		constant.TaskActionReferenceToVideo,
		constant.TaskActionRemix:
		return true
	default:
		return false
	}
}

func fetchTaskArtifact(ctx context.Context, request *TaskArtifactFetchRequest, proxy string) (io.ReadCloser, string, int64, error) {
	if request == nil {
		return nil, "", 0, errors.New("artifact request is nil")
	}
	rawURL := strings.TrimSpace(request.URL)
	if rawURL == "" {
		return nil, "", 0, errors.New("artifact URL is empty")
	}
	if strings.HasPrefix(rawURL, "data:") {
		return decodeTaskArtifactDataURL(rawURL)
	}
	if len(rawURL) > 64<<10 {
		return nil, "", 0, errors.New("artifact URL is too long")
	}
	if taskArtifactFetchLoop(rawURL) {
		return nil, "", 0, errors.New("artifact proxy loop was rejected")
	}
	method := strings.ToUpper(strings.TrimSpace(request.Method))
	if method == "" || method == http.MethodHead {
		// Persistence needs bytes, so HEAD descriptors are fetched with GET.
		method = http.MethodGet
	}
	switch method {
	case http.MethodGet, http.MethodPost:
	default:
		return nil, "", 0, errors.New("artifact method is not allowed")
	}
	if len(request.Body) > 1<<20 {
		return nil, "", 0, errors.New("artifact request body is too large")
	}
	if request.Credentialless && (len(request.Headers) != 0 || len(request.Body) != 0) {
		return nil, "", 0, errors.New("credentialless artifact requests cannot contain headers or a body")
	}
	if err := validateTaskArtifactFetchURL(rawURL); err != nil {
		return nil, "", 0, err
	}
	if _, err := url.Parse(rawURL); err != nil {
		return nil, "", 0, fmt.Errorf("invalid artifact URL: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, method, rawURL, bytes.NewReader(request.Body))
	if err != nil {
		return nil, "", 0, err
	}
	if err := applyTaskArtifactFetchHeaders(httpRequest.Header, request.Headers); err != nil {
		return nil, "", 0, err
	}

	var client *http.Client
	if proxy != "" {
		client, err = GetHttpClientWithProxy(proxy)
		if err != nil {
			return nil, "", 0, err
		}
	} else {
		client = GetSSRFProtectedHTTPClient()
	}
	if client == nil {
		client = http.DefaultClient
	}
	client = taskArtifactFetchClient(client, request.Credentialless)
	response, err := client.Do(httpRequest)
	if err != nil {
		return nil, "", 0, err
	}
	switch response.StatusCode {
	case http.StatusOK, http.StatusPartialContent:
	case http.StatusUnauthorized, http.StatusForbidden:
		response.Body.Close()
		return nil, "", 0, errors.New("artifact upstream authentication failed")
	case http.StatusNotFound, http.StatusGone:
		response.Body.Close()
		return nil, "", 0, ErrTaskArtifactNotFound
	default:
		response.Body.Close()
		return nil, "", 0, fmt.Errorf("artifact upstream returned status %d", response.StatusCode)
	}
	contentType := response.Header.Get("Content-Type")
	if len(contentType) > 255 || strings.ContainsAny(contentType, "\r\n") {
		contentType = ""
	}
	return response.Body, contentType, response.ContentLength, nil
}

func decodeTaskArtifactDataURL(rawURL string) (io.ReadCloser, string, int64, error) {
	if len(rawURL) > taskArtifactDataURLMaxEncoded {
		return nil, "", 0, errors.New("artifact data URL is too large")
	}
	parts := strings.SplitN(rawURL, ",", 2)
	if len(parts) != 2 {
		return nil, "", 0, errors.New("invalid artifact data URL")
	}
	header := parts[0]
	payload := parts[1]
	if !strings.HasPrefix(header, "data:") || !strings.Contains(header, ";base64") {
		return nil, "", 0, errors.New("unsupported artifact data URL")
	}
	mimeType := strings.TrimSuffix(strings.TrimPrefix(header, "data:"), ";base64")
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	if len(mimeType) > 255 || strings.ContainsAny(mimeType, "\r\n") {
		return nil, "", 0, errors.New("invalid artifact data URL media type")
	}
	var decoded []byte
	var decodeErr error
	for _, candidate := range []*base64.Encoding{base64.StdEncoding, base64.RawStdEncoding} {
		decoded, decodeErr = io.ReadAll(base64.NewDecoder(candidate, strings.NewReader(payload)))
		if decodeErr == nil {
			break
		}
	}
	if decodeErr != nil {
		return nil, "", 0, errors.New("invalid artifact base64 payload")
	}
	return io.NopCloser(bytes.NewReader(decoded)), mimeType, int64(len(decoded)), nil
}

func applyTaskArtifactFetchHeaders(destination http.Header, headers map[string]string) error {
	if len(headers) > 64 {
		return errors.New("too many artifact request headers")
	}
	for name, value := range headers {
		name = strings.TrimSpace(name)
		if !httpguts.ValidHeaderFieldName(name) || !httpguts.ValidHeaderFieldValue(value) || len(value) > 8192 {
			return errors.New("invalid artifact request header")
		}
		switch strings.ToLower(name) {
		case "host", "content-length", "accept-encoding", "connection", "proxy-connection", "keep-alive",
			"proxy-authorization", "te", "trailer", "transfer-encoding", "upgrade", "range", "if-range":
			return errors.New("artifact request header is not allowed")
		}
		destination.Set(name, value)
	}
	return nil
}

func validateTaskArtifactFetchURL(rawURL string) error {
	fetchSetting := system_setting.GetFetchSetting()
	return common.ValidateURLWithFetchSetting(
		rawURL,
		fetchSetting.EnableSSRFProtection,
		fetchSetting.AllowPrivateIp,
		fetchSetting.DomainFilterMode,
		fetchSetting.IpFilterMode,
		fetchSetting.DomainList,
		fetchSetting.IpList,
		fetchSetting.AllowedPorts,
		fetchSetting.ApplyIPFilterForDomain,
	)
}

func taskArtifactFetchClient(base *http.Client, credentialless bool) *http.Client {
	cloned := *base
	cloned.Timeout = taskArtifactFetchTimeout
	cloned.CheckRedirect = func(request *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("too many artifact redirects")
		}
		if request.URL == nil || (request.URL.Scheme != "http" && request.URL.Scheme != "https") ||
			request.URL.Host == "" || request.URL.User != nil || request.URL.Fragment != "" {
			return errors.New("invalid artifact redirect URL")
		}
		if taskArtifactFetchLoop(request.URL.String()) {
			return errors.New("artifact redirect loop was rejected")
		}
		if err := validateTaskArtifactFetchURL(request.URL.String()); err != nil {
			return err
		}
		if len(via) > 0 && !sameTaskArtifactOrigin(via[len(via)-1].URL, request.URL) {
			if !credentialless {
				return errors.New("credentialed cross-origin artifact redirect was rejected")
			}
			for name := range request.Header {
				request.Header.Del(name)
			}
			request.Body = http.NoBody
			request.GetBody = nil
			request.ContentLength = 0
		}
		return nil
	}
	return &cloned
}

func taskArtifactFetchLoop(rawURL string) bool {
	parsedURL, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil || parsedURL == nil {
		return false
	}
	path, err := url.PathUnescape(parsedURL.EscapedPath())
	if err != nil {
		path = parsedURL.Path
	}
	artifactPath := strings.HasPrefix(path, "/v1/videos/") && strings.HasSuffix(path, "/content") ||
		strings.HasPrefix(path, "/v1/tasks/") && strings.Contains(path, "/artifacts/") && strings.HasSuffix(path, "/content")
	if !artifactPath {
		return false
	}
	targetHost := normalizeTaskArtifactHost(parsedURL.Scheme, parsedURL.Host)
	if targetHost == "" {
		return false
	}
	for _, base := range []string{system_setting.ServerAddress, system_setting.TaskPublicAddress} {
		baseURL, parseErr := url.Parse(strings.TrimSpace(base))
		if parseErr != nil || baseURL == nil || baseURL.Host == "" {
			continue
		}
		if targetHost == normalizeTaskArtifactHost(baseURL.Scheme, baseURL.Host) {
			return true
		}
	}
	return false
}

func sameTaskArtifactOrigin(left, right *url.URL) bool {
	if left == nil || right == nil {
		return false
	}
	return strings.EqualFold(left.Scheme, right.Scheme) &&
		strings.EqualFold(normalizeTaskArtifactHost(left.Scheme, left.Host), normalizeTaskArtifactHost(right.Scheme, right.Host))
}

func normalizeTaskArtifactHost(scheme, host string) string {
	host = strings.ToLower(strings.TrimSpace(host))
	hostname, port, err := net.SplitHostPort(host)
	if err != nil {
		return strings.TrimSuffix(host, ".")
	}
	hostname = strings.TrimSuffix(strings.ToLower(hostname), ".")
	if (strings.EqualFold(scheme, "http") && port == "80") || (strings.EqualFold(scheme, "https") && port == "443") {
		return hostname
	}
	return net.JoinHostPort(hostname, port)
}
