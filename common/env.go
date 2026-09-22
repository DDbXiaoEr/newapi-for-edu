package common

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var defaultRequiredEnvVars = []string{"SQL_DSN", "LOG_SQL_DSN"}

// CheckRequiredEnvVars refuses startup when a required variable is missing or
// empty. By default SQL_DSN and LOG_SQL_DSN must be set so an unconfigured
// process fails fast instead of silently falling back to the local SQLite file.
// Override the list with REQUIRED_ENV_VARS (comma/space separated). Set
// REQUIRED_ENV_VARS="" to disable the check.
func CheckRequiredEnvVars() error {
	names := requiredEnvVarNames()
	if len(names) == 0 {
		return nil
	}
	missing := make([]string, 0, len(names))
	for _, name := range names {
		if strings.TrimSpace(os.Getenv(name)) == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return fmt.Errorf(
		"missing required environment variables: %s\n"+
			"configure them in a .env file, via `docker run -e %s=...`, or the `environment:` section of docker-compose.\n"+
			"to disable this check, set REQUIRED_ENV_VARS=\"\"",
		strings.Join(missing, ", "), missing[0])
}

func requiredEnvVarNames() []string {
	raw, set := os.LookupEnv("REQUIRED_ENV_VARS")
	if !set {
		return append([]string(nil), defaultRequiredEnvVars...)
	}
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return nil
	}
	parts := strings.FieldsFunc(raw, func(r rune) bool {
		return r == ',' || r == ';' || r == ' ' || r == '\t' || r == '\n'
	})
	names := make([]string, 0, len(parts))
	for _, name := range parts {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		names = append(names, name)
	}
	return names
}

func GetEnvOrDefault(env string, defaultValue int) int {
	if env == "" || os.Getenv(env) == "" {
		return defaultValue
	}
	num, err := strconv.Atoi(os.Getenv(env))
	if err != nil {
		SysError(fmt.Sprintf("failed to parse %s: %s, using default value: %d", env, err.Error(), defaultValue))
		return defaultValue
	}
	return num
}

func GetEnvOrDefaultString(env string, defaultValue string) string {
	if env == "" || os.Getenv(env) == "" {
		return defaultValue
	}
	return os.Getenv(env)
}

func GetEnvOrDefaultBool(env string, defaultValue bool) bool {
	if env == "" || os.Getenv(env) == "" {
		return defaultValue
	}
	b, err := strconv.ParseBool(os.Getenv(env))
	if err != nil {
		SysError(fmt.Sprintf("failed to parse %s: %s, using default value: %t", env, err.Error(), defaultValue))
		return defaultValue
	}
	return b
}
