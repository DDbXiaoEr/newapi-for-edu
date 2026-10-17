package system_setting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateTaskArtifactStoreConfig(t *testing.T) {
	valid := TaskArtifactStoreConfig{
		Mode:      TaskArtifactStoreModeS3,
		Endpoint:  "https://objects.example.com/storage",
		Bucket:    "task-artifacts",
		Region:    "us-east-1",
		AccessKey: "access-key",
		SecretKey: "secret-key",
		Prefix:    "tasks/v1/",
		PathStyle: true,
	}
	require.NoError(t, ValidateTaskArtifactStoreConfig(valid))
	require.NoError(t, ValidateTaskArtifactStoreConfig(TaskArtifactStoreConfig{
		Mode: TaskArtifactStoreModeUpstream,
	}))
	withoutRegion := valid
	withoutRegion.Region = ""
	require.NoError(t, ValidateTaskArtifactStoreConfig(withoutRegion))

	tests := []struct {
		name   string
		mutate func(*TaskArtifactStoreConfig)
		match  string
	}{
		{name: "mode", mutate: func(config *TaskArtifactStoreConfig) { config.Mode = "filesystem" }, match: "unsupported mode"},
		{name: "endpoint scheme", mutate: func(config *TaskArtifactStoreConfig) { config.Endpoint = "ftp://objects.example.com" }, match: "http or https"},
		{name: "endpoint credentials", mutate: func(config *TaskArtifactStoreConfig) { config.Endpoint = "https://user:pass@objects.example.com" }, match: "without userinfo"},
		{name: "endpoint query", mutate: func(config *TaskArtifactStoreConfig) { config.Endpoint = "https://objects.example.com?token=secret" }, match: "query or fragment"},
		{name: "bucket", mutate: func(config *TaskArtifactStoreConfig) { config.Bucket = "Invalid_Bucket" }, match: "bucket syntax"},
		{name: "IP bucket", mutate: func(config *TaskArtifactStoreConfig) { config.Bucket = "192.168.1.1" }, match: "bucket syntax"},
		{name: "region", mutate: func(config *TaskArtifactStoreConfig) { config.Region = "bad region" }, match: "region syntax"},
		{name: "access key", mutate: func(config *TaskArtifactStoreConfig) { config.AccessKey = " access-key" }, match: "access key syntax"},
		{name: "secret key", mutate: func(config *TaskArtifactStoreConfig) { config.SecretKey = "secret\nkey" }, match: "secret key syntax"},
		{name: "prefix root", mutate: func(config *TaskArtifactStoreConfig) { config.Prefix = "/tasks" }, match: "prefix syntax"},
		{name: "prefix traversal", mutate: func(config *TaskArtifactStoreConfig) { config.Prefix = "tasks/../private" }, match: "dot segments"},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			config := valid
			testCase.mutate(&config)
			assert.ErrorContains(t, ValidateTaskArtifactStoreConfig(config), testCase.match)
		})
	}
}

func TestValidateTaskArtifactStoreConfigRequiresS3Fields(t *testing.T) {
	assert.ErrorContains(t, ValidateTaskArtifactStoreConfig(TaskArtifactStoreConfig{
		Mode: TaskArtifactStoreModeS3,
	}), "S3 endpoint is required")
}

func TestLoadTaskArtifactStoreConfig(t *testing.T) {
	t.Setenv(TaskArtifactStoreModeEnv, TaskArtifactStoreModeS3)
	t.Setenv(TaskArtifactStoreS3EndpointEnv, "https://objects.example.com")
	t.Setenv(TaskArtifactStoreS3BucketEnv, "task-artifacts")
	t.Setenv(TaskArtifactStoreS3RegionEnv, "")
	t.Setenv(TaskArtifactStoreS3AccessKeyEnv, "access-key")
	t.Setenv(TaskArtifactStoreS3SecretKeyEnv, "secret-key")
	t.Setenv(TaskArtifactStoreS3PrefixEnv, "tasks/v1")
	t.Setenv(TaskArtifactStoreS3PathStyleEnv, "true")

	config := defaultTaskArtifactStoreConfig()
	assert.Equal(t, TaskArtifactStoreModeS3, config.Mode)
	assert.Equal(t, "https://objects.example.com", config.Endpoint)
	assert.Empty(t, config.Region)
	assert.True(t, config.PathStyle)
	require.NoError(t, ValidateTaskArtifactStoreConfig(config))
}

func TestValidateTaskArtifactStoreOption(t *testing.T) {
	original := taskArtifactStoreConfig
	t.Cleanup(func() { taskArtifactStoreConfig = original })
	taskArtifactStoreConfig = TaskArtifactStoreConfig{
		Mode:      TaskArtifactStoreModeUpstream,
		Region:    DefaultTaskArtifactStoreRegion,
		PathStyle: true,
	}

	require.NoError(t, ValidateTaskArtifactStoreOption("task_artifact_store.endpoint", "https://objects.example.com"))
	require.NoError(t, ValidateTaskArtifactStoreOption("task_artifact_store.region", ""))
	assert.ErrorContains(t, ValidateTaskArtifactStoreOption("task_artifact_store.endpoint", "ftp://objects.example.com"), "http or https")
	assert.ErrorContains(t, ValidateTaskArtifactStoreOption("task_artifact_store.mode", TaskArtifactStoreModeS3), "S3 endpoint is required")
	assert.ErrorContains(t, ValidateTaskArtifactStoreOption("task_artifact_store.bucket", "Invalid_Bucket"), "bucket syntax")
	assert.NoError(t, ValidateTaskArtifactStoreOption("unrelated.option", "value"))
}
