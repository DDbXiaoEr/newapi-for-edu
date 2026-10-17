package service

import (
	"context"
	"errors"
	"io"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
)

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
