package service

import (
	"context"
	"errors"
	"io"
	"mime"
	"path"
	"strings"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
)

const TaskArtifactDownloadQueryParameter = "download"

var taskArtifactExtensionByMimeType = map[string]string{
	"application/pdf": ".pdf",
	"audio/mpeg":      ".mp3",
	"audio/ogg":       ".ogg",
	"audio/wav":       ".wav",
	"audio/x-wav":     ".wav",
	"image/gif":       ".gif",
	"image/jpeg":      ".jpg",
	"image/png":       ".png",
	"image/webp":      ".webp",
	"text/plain":      ".txt",
	"video/mp4":       ".mp4",
	"video/quicktime": ".mov",
	"video/webm":      ".webm",
}

// StoredArtifactRef describes a persisted artifact object.
type StoredArtifactRef struct {
	Backend   string
	Bucket    string
	ObjectKey string
	MimeType  string
	Size      int64
}

// TaskArtifactStore is the persistence boundary for generated artifact bytes.
// types.TaskArtifact is re-exported by relay/channel as channel.TaskArtifact.
type TaskArtifactStore interface {
	Enabled() bool
	Resolve(ctx context.Context, task *model.Task, artifactKey string) (*StoredArtifactRef, error)
	Persist(ctx context.Context, task *model.Task, artifact types.TaskArtifact, content io.Reader) (*StoredArtifactRef, error)
	Serve(c *gin.Context, task *model.Task, ref *StoredArtifactRef) error
}

var (
	ErrTaskArtifactStoreDisabled       = errors.New("task artifact store is disabled")
	ErrTaskArtifactNotFound            = errors.New("task artifact not found")
	ErrTaskArtifactStoreUnavailable    = errors.New("task artifact store is unavailable")
	ErrTaskArtifactRangeNotSatisfiable = errors.New("task artifact range is not satisfiable")
)

var taskArtifactStore TaskArtifactStore = newS3ArtifactStore()

// GetTaskArtifactStore returns the process-wide artifact storage backend. The
// returned store is inert while S3 artifact storage is disabled or incomplete.
func GetTaskArtifactStore() TaskArtifactStore {
	return taskArtifactStore
}

func TaskArtifactDownloadRequested(c *gin.Context) bool {
	if c == nil || c.Request == nil {
		return false
	}
	values, exists := c.Request.URL.Query()[TaskArtifactDownloadQueryParameter]
	if !exists {
		return false
	}
	if len(values) == 0 {
		return true
	}
	for _, value := range values {
		switch strings.ToLower(strings.TrimSpace(value)) {
		case "", "1", "true":
			return true
		}
	}
	return false
}

func SetTaskArtifactDownloadDisposition(c *gin.Context, artifactKey, contentType string) {
	if !TaskArtifactDownloadRequested(c) {
		return
	}
	filename := taskArtifactDownloadFilename(artifactKey, contentType)
	disposition := mime.FormatMediaType("attachment", map[string]string{"filename": filename})
	if disposition != "" {
		c.Header("Content-Disposition", disposition)
	}
}

func taskArtifactDownloadFilename(artifactKey, contentType string) string {
	filename := strings.TrimSpace(artifactKey)
	if filename == "" {
		filename = "artifact"
	}
	if path.Ext(filename) != "" {
		return filename
	}
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		mediaType = strings.TrimSpace(contentType)
	}
	if extension := taskArtifactExtensionByMimeType[mediaType]; extension != "" {
		return filename + extension
	}
	if extensions, err := mime.ExtensionsByType(mediaType); err == nil && len(extensions) > 0 {
		return filename + extensions[0]
	}
	return filename
}
