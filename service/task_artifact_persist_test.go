package service

import (
	"context"
	"encoding/base64"
	"io"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type recordingArtifactStore struct {
	persisted map[string]string
}

func (store *recordingArtifactStore) Enabled() bool { return true }

func (store *recordingArtifactStore) Resolve(context.Context, *model.Task, string) (*StoredArtifactRef, error) {
	return nil, nil
}

func (store *recordingArtifactStore) Persist(_ context.Context, _ *model.Task, artifact types.TaskArtifact, content io.Reader) (*StoredArtifactRef, error) {
	body, err := io.ReadAll(content)
	if err != nil {
		return nil, err
	}
	if store.persisted == nil {
		store.persisted = make(map[string]string)
	}
	store.persisted[artifact.Key] = string(body)
	return &StoredArtifactRef{
		Backend:   "s3",
		Bucket:    "bucket",
		ObjectKey: "object/" + artifact.Key,
		MimeType:  artifact.MimeType,
		Size:      int64(len(body)),
	}, nil
}

func (store *recordingArtifactStore) Serve(*gin.Context, *model.Task, *StoredArtifactRef) error {
	return nil
}

type fakeArtifactSource struct {
	artifacts []types.TaskArtifact
	bodies    map[string]string
}

func (source *fakeArtifactSource) ListArtifacts() ([]types.TaskArtifact, error) {
	return source.artifacts, nil
}

func (source *fakeArtifactSource) BuildContentRequest(key string) (*TaskArtifactFetchRequest, error) {
	body := source.bodies[key]
	return &TaskArtifactFetchRequest{URL: "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(body))}, nil
}

func TestPersistTaskArtifactsUsesSourceAndManifest(t *testing.T) {
	originalStore := taskArtifactStore
	originalFactory := TaskArtifactSourceFactory
	t.Cleanup(func() {
		taskArtifactStore = originalStore
		TaskArtifactSourceFactory = originalFactory
	})

	store := &recordingArtifactStore{}
	taskArtifactStore = store
	source := &fakeArtifactSource{
		artifacts: []types.TaskArtifact{
			{Key: "image-main", Type: "image"},
			{Key: "invalid key", Type: "image"},
		},
		bodies: map[string]string{"image-main": "image-bytes"},
	}
	TaskArtifactSourceFactory = func(*model.Task) (TaskArtifactSource, error) {
		return source, nil
	}

	task := &model.Task{TaskID: "task-persist", Status: model.TaskStatusSuccess}
	require.NoError(t, PersistTaskArtifacts(t.Context(), task))
	assert.Equal(t, "image-bytes", store.persisted["image-main"])
	assert.Contains(t, task.PrivateData.Artifacts, "image-main")
	assert.NotContains(t, task.PrivateData.Artifacts, "invalid key")
	stored := task.PrivateData.Artifacts["image-main"]
	assert.Equal(t, "s3", stored.Backend)
	assert.Equal(t, "object/image-main", stored.ObjectKey)
}

func TestLegacyVideoTaskAvailable(t *testing.T) {
	task := &model.Task{
		Status:      model.TaskStatusSuccess,
		Action:      constant.TaskActionTextToVideo,
		PrivateData: model.TaskPrivateData{ResultURL: "https://example.com/video.mp4"},
	}
	assert.True(t, LegacyVideoTaskAvailable(task))

	task.PrivateData.Execution = &model.TaskExecutionSnapshot{TaskPlugin: &model.TaskPluginSnapshot{Key: "plugin"}}
	assert.False(t, LegacyVideoTaskAvailable(task))

	task.PrivateData.Execution = nil
	task.PrivateData.ResultURL = "data:video/mp4;base64,AAAA"
	assert.False(t, LegacyVideoTaskAvailable(task))
	assert.True(t, strings.HasPrefix(task.PrivateData.ResultURL, "data:"))
}
