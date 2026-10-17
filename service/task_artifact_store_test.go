package service

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/config"
	"github.com/QuantumNous/new-api/setting/system_setting"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setTaskArtifactStoreTestConfig(t *testing.T, configValue system_setting.TaskArtifactStoreConfig) {
	t.Helper()
	original := system_setting.GetTaskArtifactStoreConfig()
	t.Cleanup(func() {
		applyTaskArtifactStoreTestConfig(t, original)
	})
	applyTaskArtifactStoreTestConfig(t, configValue)
}

func applyTaskArtifactStoreTestConfig(t *testing.T, configValue system_setting.TaskArtifactStoreConfig) {
	t.Helper()
	registered := config.GlobalConfig.Get(system_setting.TaskArtifactStoreConfigName)
	require.NotNil(t, registered)
	require.NoError(t, config.UpdateConfigFromMap(registered, map[string]string{
		"mode":       configValue.Mode,
		"endpoint":   configValue.Endpoint,
		"bucket":     configValue.Bucket,
		"region":     configValue.Region,
		"access_key": configValue.AccessKey,
		"secret_key": configValue.SecretKey,
		"prefix":     configValue.Prefix,
		"path_style": strings.ToLower(strings.TrimSpace(map[bool]string{true: "true", false: "false"}[configValue.PathStyle])),
	}))
}

func TestTaskArtifactStoreDisabledHasNoBehavior(t *testing.T) {
	setTaskArtifactStoreTestConfig(t, system_setting.TaskArtifactStoreConfig{
		Mode:   system_setting.TaskArtifactStoreModeUpstream,
		Region: system_setting.DefaultTaskArtifactStoreRegion,
	})
	store := GetTaskArtifactStore()
	require.NotNil(t, store)
	assert.False(t, store.Enabled())

	task := &model.Task{TaskID: "task-disabled-store"}
	ref, err := store.Resolve(t.Context(), task, "video")
	require.NoError(t, err)
	assert.Nil(t, ref)

	ref, err = store.Persist(t.Context(), task, types.TaskArtifact{Key: "video", Type: "video"}, strings.NewReader("content"))
	assert.Nil(t, ref)
	assert.ErrorIs(t, err, ErrTaskArtifactStoreDisabled)
	assert.ErrorIs(t, store.Serve(&gin.Context{}, task, &StoredArtifactRef{Backend: "s3"}), ErrTaskArtifactStoreDisabled)
	assert.Same(t, store, GetTaskArtifactStore())
}

func TestTaskArtifactObjectKey(t *testing.T) {
	task := &model.Task{ID: 42}
	assert.Equal(t, "tasks/42/video", taskArtifactObjectKey(system_setting.TaskArtifactStoreConfig{Prefix: "tasks/"}, task, "video"))
	assert.Equal(t, "tasks/42/video", taskArtifactObjectKey(system_setting.TaskArtifactStoreConfig{Prefix: "/tasks/"}, task, "video"))

	task = &model.Task{TaskID: "task/id with spaces"}
	assert.Equal(t, "task_id_with_spaces/video", taskArtifactObjectKey(system_setting.TaskArtifactStoreConfig{}, task, "video"))
}

func TestTaskArtifactDownloadDisposition(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  bool
	}{
		{name: "absent", query: "", want: false},
		{name: "one", query: "?download=1", want: true},
		{name: "true", query: "?download=true", want: true},
		{name: "bare", query: "?download", want: true},
		{name: "false", query: "?download=0", want: false},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request = httptest.NewRequest(http.MethodGet, "/artifact"+testCase.query, nil)
			assert.Equal(t, testCase.want, TaskArtifactDownloadRequested(c))
		})
	}

	assert.Equal(t, "image-main.png", taskArtifactDownloadFilename("image-main", "image/png"))
	assert.Equal(t, "clip.mp4", taskArtifactDownloadFilename("clip.mp4", "video/mp4"))

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodGet, "/artifact?download=1", nil)
	SetTaskArtifactDownloadDisposition(c, "image-main", "image/png")
	disposition := recorder.Header().Get("Content-Disposition")
	assert.Contains(t, disposition, "attachment;")
	assert.Contains(t, disposition, "image-main.png")

	recorder = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodGet, "/artifact", nil)
	SetTaskArtifactDownloadDisposition(c, "image-main", "image/png")
	assert.Empty(t, recorder.Header().Get("Content-Disposition"))
}

func TestTaskArtifactStoreSigningRegion(t *testing.T) {
	assert.Equal(t, system_setting.DefaultTaskArtifactStoreRegion, taskArtifactStoreSigningRegion(system_setting.TaskArtifactStoreConfig{}))
	assert.Equal(t, system_setting.DefaultTaskArtifactStoreRegion, taskArtifactStoreSigningRegion(system_setting.TaskArtifactStoreConfig{Region: "   "}))
	assert.Equal(t, "eu-west-1", taskArtifactStoreSigningRegion(system_setting.TaskArtifactStoreConfig{Region: "eu-west-1"}))
}

func TestS3TaskArtifactStorePersistResolveServe(t *testing.T) {
	var mutex sync.Mutex
	objects := make(map[string][]byte)
	contentTypes := make(map[string]string)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := strings.TrimPrefix(r.URL.Path, "/")
		mutex.Lock()
		defer mutex.Unlock()
		switch r.Method {
		case http.MethodPut:
			body, _ := io.ReadAll(r.Body)
			objects[key] = body
			contentTypes[key] = r.Header.Get("Content-Type")
			w.Header().Set("ETag", `"etag-1"`)
			w.WriteHeader(http.StatusOK)
		case http.MethodHead, http.MethodGet:
			body, exists := objects[key]
			if !exists {
				w.Header().Set("Content-Type", "application/xml")
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?><Error><Code>NoSuchKey</Code><Message>not found</Message></Error>`))
				return
			}
			w.Header().Set("Content-Type", contentTypes[key])
			w.Header().Set("ETag", `"etag-1"`)
			w.Header().Set("Accept-Ranges", "bytes")
			w.Header().Set("Last-Modified", "Wed, 01 Jan 2025 00:00:00 GMT")
			if r.Method == http.MethodHead {
				w.Header().Set("Content-Length", strconv.Itoa(len(body)))
				w.WriteHeader(http.StatusOK)
				return
			}
			_, _ = w.Write(body)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	setTaskArtifactStoreTestConfig(t, system_setting.TaskArtifactStoreConfig{
		Mode:      system_setting.TaskArtifactStoreModeS3,
		Endpoint:  server.URL,
		Bucket:    "bucket",
		AccessKey: "access",
		SecretKey: "secret",
		PathStyle: true,
	})

	store := GetTaskArtifactStore()
	require.True(t, store.Enabled())
	task := &model.Task{ID: 7, TaskID: "task-7", Status: model.TaskStatusSuccess}
	ref, err := store.Persist(t.Context(), task, types.TaskArtifact{Key: "video", Type: "video", MimeType: "video/mp4"}, strings.NewReader("content"))
	require.NoError(t, err)
	require.NotNil(t, ref)
	assert.Equal(t, "7/video", ref.ObjectKey)
	assert.Equal(t, int64(7), ref.Size)

	resolved, err := store.Resolve(t.Context(), task, "video")
	require.NoError(t, err)
	require.NotNil(t, resolved)
	assert.Equal(t, "7/video", resolved.ObjectKey)

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodGet, "/artifact", nil)
	require.NoError(t, store.Serve(context, task, resolved))
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "content", recorder.Body.String())
	assert.Equal(t, "video/mp4", recorder.Header().Get("Content-Type"))
	assert.Equal(t, "private, no-store", recorder.Header().Get("Cache-Control"))

	downloadRecorder := httptest.NewRecorder()
	downloadContext, _ := gin.CreateTestContext(downloadRecorder)
	downloadContext.Params = gin.Params{{Key: "artifact_key", Value: "video"}}
	downloadContext.Request = httptest.NewRequest(http.MethodGet, "/artifact?download=1", nil)
	require.NoError(t, store.Serve(downloadContext, task, resolved))
	assert.Equal(t, http.StatusOK, downloadRecorder.Code)
	disposition := downloadRecorder.Header().Get("Content-Disposition")
	assert.Contains(t, disposition, "attachment;")
	assert.Contains(t, disposition, "video.mp4")
}
