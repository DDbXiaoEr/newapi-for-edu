package system_setting

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/config"
)

const TaskArtifactStoreConfigName = "task_artifact_store"

const (
	TaskArtifactStoreModeUpstream = "upstream"
	TaskArtifactStoreModeS3       = "s3"

	DefaultTaskArtifactStoreRegion = "us-east-1"
)

const (
	TaskArtifactStoreModeEnv          = "TASK_ARTIFACT_STORE_MODE"
	TaskArtifactStoreS3EndpointEnv    = "TASK_ARTIFACT_STORE_S3_ENDPOINT"
	TaskArtifactStoreS3BucketEnv      = "TASK_ARTIFACT_STORE_S3_BUCKET"
	TaskArtifactStoreS3RegionEnv      = "TASK_ARTIFACT_STORE_S3_REGION"
	TaskArtifactStoreS3AccessKeyEnv   = "TASK_ARTIFACT_STORE_S3_ACCESS_KEY"
	TaskArtifactStoreS3SecretKeyEnv   = "TASK_ARTIFACT_STORE_S3_SECRET_KEY"
	TaskArtifactStoreS3PrefixEnv      = "TASK_ARTIFACT_STORE_S3_PREFIX"
	TaskArtifactStoreS3PathStyleEnv   = "TASK_ARTIFACT_STORE_S3_PATH_STYLE"
	DefaultTaskArtifactStorePathStyle = true
)

var (
	taskArtifactStoreBucketPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`)
	taskArtifactStoreRegionPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$`)
)

// TaskArtifactStoreConfig configures generated-artifact persistence. S3 mode
// keeps artifact bytes in object storage while downloads stay proxied by this
// service, so object-store credentials are never exposed to clients.
type TaskArtifactStoreConfig struct {
	Mode      string `json:"mode"`
	Endpoint  string `json:"endpoint"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Prefix    string `json:"prefix"`
	PathStyle bool   `json:"path_style"`
}

func defaultTaskArtifactStoreConfig() TaskArtifactStoreConfig {
	return TaskArtifactStoreConfig{
		Mode:      common.GetEnvOrDefaultString(TaskArtifactStoreModeEnv, TaskArtifactStoreModeUpstream),
		Endpoint:  common.GetEnvOrDefaultString(TaskArtifactStoreS3EndpointEnv, ""),
		Bucket:    common.GetEnvOrDefaultString(TaskArtifactStoreS3BucketEnv, ""),
		Region:    common.GetEnvOrDefaultString(TaskArtifactStoreS3RegionEnv, ""),
		AccessKey: common.GetEnvOrDefaultString(TaskArtifactStoreS3AccessKeyEnv, ""),
		SecretKey: common.GetEnvOrDefaultString(TaskArtifactStoreS3SecretKeyEnv, ""),
		Prefix:    common.GetEnvOrDefaultString(TaskArtifactStoreS3PrefixEnv, ""),
		PathStyle: common.GetEnvOrDefaultBool(TaskArtifactStoreS3PathStyleEnv, DefaultTaskArtifactStorePathStyle),
	}
}

var taskArtifactStoreConfig = defaultTaskArtifactStoreConfig()

func init() {
	config.GlobalConfig.Register(TaskArtifactStoreConfigName, &taskArtifactStoreConfig)
}

// GetTaskArtifactStoreConfig returns a copy of the current configuration.
func GetTaskArtifactStoreConfig() TaskArtifactStoreConfig {
	configValue := taskArtifactStoreConfig
	normalizeTaskArtifactStoreConfig(&configValue)
	return configValue
}

// S3Enabled reports whether S3 storage is selected and fully configured.
func (configValue TaskArtifactStoreConfig) S3Enabled() bool {
	normalizeTaskArtifactStoreConfig(&configValue)
	return configValue.Mode == TaskArtifactStoreModeS3 &&
		ValidateTaskArtifactStoreConfig(configValue) == nil
}

func normalizeTaskArtifactStoreConfig(configValue *TaskArtifactStoreConfig) {
	if configValue == nil {
		return
	}
	configValue.Mode = strings.TrimSpace(strings.ToLower(configValue.Mode))
	if configValue.Mode == "" {
		configValue.Mode = TaskArtifactStoreModeUpstream
	}
}

// LoadTaskArtifactStoreConfig returns the current configuration. Invalid
// values are reported and leave artifact storage in upstream mode.
func LoadTaskArtifactStoreConfig() TaskArtifactStoreConfig {
	configValue := GetTaskArtifactStoreConfig()
	if err := ValidateTaskArtifactStoreConfig(configValue); err != nil {
		common.SysError("invalid task artifact store configuration: " + err.Error() + "; using upstream mode")
		configValue.Mode = TaskArtifactStoreModeUpstream
	}
	return configValue
}

// ValidateTaskArtifactStoreOption validates one admin option update against the
// current configuration. It performs syntax checks only; it never resolves
// hosts, contacts an endpoint, or verifies credentials.
func ValidateTaskArtifactStoreOption(key, value string) error {
	if !strings.HasPrefix(key, TaskArtifactStoreConfigName+".") {
		return nil
	}
	configValue := GetTaskArtifactStoreConfig()
	switch strings.TrimPrefix(key, TaskArtifactStoreConfigName+".") {
	case "mode":
		mode := strings.TrimSpace(strings.ToLower(value))
		if mode != TaskArtifactStoreModeUpstream && mode != TaskArtifactStoreModeS3 {
			return fmt.Errorf("unsupported mode %q", value)
		}
		if mode == TaskArtifactStoreModeUpstream {
			return nil
		}
		configValue.Mode = mode
	case "endpoint":
		configValue.Endpoint = value
	case "bucket":
		configValue.Bucket = value
	case "region":
		configValue.Region = value
	case "access_key":
		configValue.AccessKey = value
	case "secret_key":
		configValue.SecretKey = value
	case "prefix":
		configValue.Prefix = value
	case "path_style":
		configValue.PathStyle = strings.EqualFold(strings.TrimSpace(value), "true")
	default:
		return nil
	}
	normalizeTaskArtifactStoreConfig(&configValue)
	return ValidateTaskArtifactStoreConfig(configValue)
}

// ValidateTaskArtifactStoreConfig performs syntax checks only. It never
// resolves hosts, contacts an endpoint, or verifies credentials.
func ValidateTaskArtifactStoreConfig(configValue TaskArtifactStoreConfig) error {
	if configValue.Mode != TaskArtifactStoreModeUpstream && configValue.Mode != TaskArtifactStoreModeS3 {
		return fmt.Errorf("unsupported mode %q", configValue.Mode)
	}

	requireS3Fields := configValue.Mode == TaskArtifactStoreModeS3
	if requireS3Fields && configValue.Endpoint == "" {
		return errors.New("S3 endpoint is required")
	}
	if configValue.Endpoint != "" {
		if configValue.Endpoint != strings.TrimSpace(configValue.Endpoint) {
			return errors.New("S3 endpoint must not contain surrounding whitespace")
		}
		endpoint, err := url.Parse(configValue.Endpoint)
		if err != nil || endpoint == nil || endpoint.Host == "" || endpoint.User != nil || endpoint.Opaque != "" {
			return errors.New("S3 endpoint must be an absolute URL without userinfo")
		}
		if endpoint.Scheme != "http" && endpoint.Scheme != "https" {
			return errors.New("S3 endpoint must use http or https")
		}
		if endpoint.RawQuery != "" || endpoint.ForceQuery || endpoint.Fragment != "" {
			return errors.New("S3 endpoint must not contain a query or fragment")
		}
	}

	if requireS3Fields && configValue.Bucket == "" {
		return errors.New("S3 bucket is required")
	}
	if configValue.Bucket != "" {
		if !taskArtifactStoreBucketPattern.MatchString(configValue.Bucket) ||
			strings.Contains(configValue.Bucket, "..") || net.ParseIP(configValue.Bucket) != nil {
			return errors.New("S3 bucket syntax is invalid")
		}
	}

	if configValue.Region != "" && !taskArtifactStoreRegionPattern.MatchString(configValue.Region) {
		return errors.New("S3 region syntax is invalid")
	}
	if err := validateTaskArtifactStoreCredential("access key", configValue.AccessKey, 256, requireS3Fields); err != nil {
		return err
	}
	if err := validateTaskArtifactStoreCredential("secret key", configValue.SecretKey, 1024, requireS3Fields); err != nil {
		return err
	}

	if configValue.Prefix != "" {
		if configValue.Prefix != strings.TrimSpace(configValue.Prefix) || len(configValue.Prefix) > 512 ||
			strings.HasPrefix(configValue.Prefix, "/") || strings.Contains(configValue.Prefix, "\\") {
			return errors.New("S3 prefix syntax is invalid")
		}
		for part := range strings.SplitSeq(configValue.Prefix, "/") {
			if part == "." || part == ".." {
				return errors.New("S3 prefix must not contain dot segments")
			}
		}
		for _, character := range configValue.Prefix {
			if unicode.IsControl(character) {
				return errors.New("S3 prefix must not contain control characters")
			}
		}
	}
	return nil
}

func validateTaskArtifactStoreCredential(name, value string, maxLength int, required bool) error {
	if required && value == "" {
		return fmt.Errorf("S3 %s is required", name)
	}
	if value == "" {
		return nil
	}
	if value != strings.TrimSpace(value) || len(value) > maxLength {
		return fmt.Errorf("S3 %s syntax is invalid", name)
	}
	for _, character := range value {
		if unicode.IsControl(character) {
			return fmt.Errorf("S3 %s syntax is invalid", name)
		}
	}
	return nil
}
