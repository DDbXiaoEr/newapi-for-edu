package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/system_setting"
	"github.com/QuantumNous/new-api/types"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/gin-gonic/gin"
)

const taskArtifactS3Backend = "s3"

type s3ArtifactStore struct {
	mutex       sync.Mutex
	fingerprint string
	client      *s3.Client
	uploader    *manager.Uploader
	config      system_setting.TaskArtifactStoreConfig
}

func newS3ArtifactStore() *s3ArtifactStore {
	return &s3ArtifactStore{}
}

func (s *s3ArtifactStore) Enabled() bool {
	return system_setting.GetTaskArtifactStoreConfig().S3Enabled()
}

func (s *s3ArtifactStore) Resolve(ctx context.Context, task *model.Task, artifactKey string) (*StoredArtifactRef, error) {
	if task == nil || artifactKey == "" || !s.Enabled() {
		return nil, nil
	}
	configValue, client, err := s.backend()
	if err != nil {
		return nil, err
	}
	if stored, ok := task.PrivateData.Artifacts[artifactKey]; ok {
		if stored.Backend != "" && stored.Backend != taskArtifactS3Backend {
			return nil, nil
		}
		bucket := stored.Bucket
		if bucket == "" {
			bucket = configValue.Bucket
		}
		objectKey := stored.ObjectKey
		if objectKey == "" {
			objectKey = taskArtifactObjectKey(configValue, task, artifactKey)
		}
		return &StoredArtifactRef{
			Backend:   taskArtifactS3Backend,
			Bucket:    bucket,
			ObjectKey: objectKey,
			MimeType:  stored.MimeType,
			Size:      stored.Size,
		}, nil
	}

	objectKey := taskArtifactObjectKey(configValue, task, artifactKey)
	head, err := client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(configValue.Bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		if isS3NotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("%w: %v", ErrTaskArtifactStoreUnavailable, err)
	}
	ref := &StoredArtifactRef{
		Backend:   taskArtifactS3Backend,
		Bucket:    configValue.Bucket,
		ObjectKey: objectKey,
	}
	if head.ContentType != nil {
		ref.MimeType = *head.ContentType
	}
	if head.ContentLength != nil {
		ref.Size = *head.ContentLength
	}
	return ref, nil
}

func (s *s3ArtifactStore) Persist(ctx context.Context, task *model.Task, artifact types.TaskArtifact, content io.Reader) (*StoredArtifactRef, error) {
	if task == nil || strings.TrimSpace(artifact.Key) == "" || content == nil {
		return nil, errors.New("invalid task artifact")
	}
	if !s.Enabled() {
		return nil, ErrTaskArtifactStoreDisabled
	}
	configValue, _, uploader, err := s.backendWithUploader()
	if err != nil {
		return nil, err
	}
	objectKey := taskArtifactObjectKey(configValue, task, artifact.Key)
	contentType := sanitizeArtifactContentType(artifact.MimeType)
	counter := &countingReader{reader: content}
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(configValue.Bucket),
		Key:         aws.String(objectKey),
		Body:        counter,
		ContentType: aws.String(contentType),
		Metadata: map[string]string{
			"artifact-key":  artifact.Key,
			"artifact-type": artifact.Type,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTaskArtifactStoreUnavailable, err)
	}
	return &StoredArtifactRef{
		Backend:   taskArtifactS3Backend,
		Bucket:    configValue.Bucket,
		ObjectKey: objectKey,
		MimeType:  contentType,
		Size:      counter.count,
	}, nil
}

func (s *s3ArtifactStore) Serve(c *gin.Context, task *model.Task, ref *StoredArtifactRef) error {
	if c == nil || ref == nil {
		return errors.New("invalid task artifact request")
	}
	if !s.Enabled() {
		return ErrTaskArtifactStoreDisabled
	}
	configValue, client, err := s.backend()
	if err != nil {
		return err
	}
	bucket := strings.TrimSpace(ref.Bucket)
	if bucket == "" {
		bucket = configValue.Bucket
	}
	objectKey := strings.TrimSpace(ref.ObjectKey)
	if objectKey == "" {
		objectKey = taskArtifactObjectKey(configValue, task, "")
	}
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}
	rangeHeader := strings.TrimSpace(c.Request.Header.Get("Range"))
	if rangeHeader != "" && len(rangeHeader) <= 256 {
		input.Range = aws.String(rangeHeader)
	}

	if c.Request.Method == http.MethodHead && input.Range == nil {
		head, headErr := client.HeadObject(c.Request.Context(), &s3.HeadObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(objectKey),
		})
		if headErr != nil {
			if isS3NotFound(headErr) {
				return ErrTaskArtifactNotFound
			}
			return fmt.Errorf("%w: %v", ErrTaskArtifactStoreUnavailable, headErr)
		}
		header := c.Writer.Header()
		if head.ContentType != nil && *head.ContentType != "" {
			header.Set("Content-Type", *head.ContentType)
		}
		if head.ContentLength != nil && *head.ContentLength > 0 {
			header.Set("Content-Length", strconv.FormatInt(*head.ContentLength, 10))
		}
		header.Set("Accept-Ranges", "bytes")
		if head.ETag != nil && *head.ETag != "" {
			header.Set("ETag", *head.ETag)
		}
		if head.LastModified != nil {
			header.Set("Last-Modified", head.LastModified.UTC().Format(http.TimeFormat))
		}
		if head.ContentDisposition != nil && *head.ContentDisposition != "" {
			header.Set("Content-Disposition", *head.ContentDisposition)
		}
		setStoredArtifactSecurityHeaders(header)
		c.Status(http.StatusOK)
		return nil
	}

	output, err := client.GetObject(c.Request.Context(), input)
	if err != nil {
		if isS3NotFound(err) {
			return ErrTaskArtifactNotFound
		}
		if isS3InvalidRange(err) {
			return ErrTaskArtifactRangeNotSatisfiable
		}
		return fmt.Errorf("%w: %v", ErrTaskArtifactStoreUnavailable, err)
	}
	defer output.Body.Close()

	header := c.Writer.Header()
	if output.ContentType != nil && *output.ContentType != "" {
		header.Set("Content-Type", *output.ContentType)
	}
	if output.ContentLength != nil && *output.ContentLength > 0 {
		header.Set("Content-Length", strconv.FormatInt(*output.ContentLength, 10))
	}
	if output.ContentRange != nil && *output.ContentRange != "" {
		header.Set("Content-Range", *output.ContentRange)
	}
	if output.AcceptRanges != nil && *output.AcceptRanges != "" {
		header.Set("Accept-Ranges", *output.AcceptRanges)
	} else {
		header.Set("Accept-Ranges", "bytes")
	}
	if output.ETag != nil && *output.ETag != "" {
		header.Set("ETag", *output.ETag)
	}
	if output.LastModified != nil {
		header.Set("Last-Modified", output.LastModified.UTC().Format(http.TimeFormat))
	}
	if output.ContentDisposition != nil && *output.ContentDisposition != "" {
		header.Set("Content-Disposition", *output.ContentDisposition)
	}
	setStoredArtifactSecurityHeaders(header)

	status := http.StatusOK
	if output.ContentRange != nil && *output.ContentRange != "" {
		status = http.StatusPartialContent
	}
	c.Status(status)
	c.Writer.WriteHeaderNow()
	if c.Request.Method == http.MethodHead {
		return nil
	}
	if _, err := io.Copy(c.Writer, output.Body); err != nil {
		return fmt.Errorf("stream stored artifact: %w", err)
	}
	return nil
}

func (s *s3ArtifactStore) backend() (system_setting.TaskArtifactStoreConfig, *s3.Client, error) {
	configValue, client, _, err := s.backendWithUploader()
	return configValue, client, err
}

func (s *s3ArtifactStore) backendWithUploader() (system_setting.TaskArtifactStoreConfig, *s3.Client, *manager.Uploader, error) {
	configValue := system_setting.GetTaskArtifactStoreConfig()
	if !configValue.S3Enabled() {
		return configValue, nil, nil, ErrTaskArtifactStoreDisabled
	}
	fingerprint := taskArtifactStoreFingerprint(configValue)

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.client == nil || s.fingerprint != fingerprint {
		awsConfig := aws.Config{
			Region:      taskArtifactStoreSigningRegion(configValue),
			Credentials: credentials.NewStaticCredentialsProvider(configValue.AccessKey, configValue.SecretKey, ""),
		}
		client := s3.NewFromConfig(awsConfig, func(options *s3.Options) {
			options.UsePathStyle = configValue.PathStyle
			if configValue.Endpoint != "" {
				options.BaseEndpoint = aws.String(configValue.Endpoint)
			}
		})
		s.client = client
		s.uploader = manager.NewUploader(client)
		s.fingerprint = fingerprint
		s.config = configValue
	}
	return s.config, s.client, s.uploader, nil
}

func taskArtifactStoreSigningRegion(configValue system_setting.TaskArtifactStoreConfig) string {
	region := strings.TrimSpace(configValue.Region)
	if region == "" {
		return system_setting.DefaultTaskArtifactStoreRegion
	}
	return region
}

func taskArtifactStoreFingerprint(configValue system_setting.TaskArtifactStoreConfig) string {
	return strings.Join([]string{
		configValue.Endpoint,
		configValue.Bucket,
		configValue.Region,
		configValue.AccessKey,
		configValue.SecretKey,
		configValue.Prefix,
		strconv.FormatBool(configValue.PathStyle),
	}, "\x00")
}

func taskArtifactObjectKey(configValue system_setting.TaskArtifactStoreConfig, task *model.Task, artifactKey string) string {
	taskIdentity := "unknown"
	if task != nil {
		if task.ID > 0 {
			taskIdentity = strconv.FormatInt(task.ID, 10)
		} else if safe := sanitizeTaskArtifactIdentity(task.TaskID); safe != "" {
			taskIdentity = safe
		}
	}
	parts := make([]string, 0, 3)
	if prefix := strings.Trim(configValue.Prefix, "/"); prefix != "" {
		parts = append(parts, prefix)
	}
	parts = append(parts, taskIdentity)
	if artifactKey = strings.TrimSpace(artifactKey); artifactKey != "" {
		parts = append(parts, artifactKey)
	}
	return path.Join(parts...)
}

func sanitizeTaskArtifactIdentity(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if len(value) > 64 {
		value = value[:64]
	}
	var builder strings.Builder
	for _, character := range value {
		switch {
		case character >= 'a' && character <= 'z',
			character >= 'A' && character <= 'Z',
			character >= '0' && character <= '9',
			character == '.', character == '_', character == '-':
			builder.WriteRune(character)
		default:
			builder.WriteByte('_')
		}
	}
	return strings.Trim(builder.String(), "._-")
}

func sanitizeArtifactContentType(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || len(value) > 255 || strings.ContainsAny(value, "\r\n") {
		return "application/octet-stream"
	}
	return value
}

func setStoredArtifactSecurityHeaders(header http.Header) {
	header.Set("Cache-Control", "private, no-store")
	header.Set("Content-Security-Policy", "sandbox; default-src 'none'")
	header.Set("Referrer-Policy", "no-referrer")
	header.Set("X-Content-Type-Options", "nosniff")
}

func isS3NotFound(err error) bool {
	if err == nil {
		return false
	}
	var notFound *s3types.NotFound
	var noSuchKey *s3types.NoSuchKey
	if errors.As(err, &notFound) || errors.As(err, &noSuchKey) {
		return true
	}
	var apiError smithy.APIError
	if errors.As(err, &apiError) {
		switch apiError.ErrorCode() {
		case "NotFound", "NoSuchKey", "404":
			return true
		}
	}
	return strings.Contains(strings.ToLower(err.Error()), "status code: 404")
}

func isS3InvalidRange(err error) bool {
	if err == nil {
		return false
	}
	var apiError smithy.APIError
	if errors.As(err, &apiError) {
		switch apiError.ErrorCode() {
		case "InvalidRange", "RequestedRangeNotSatisfiable":
			return true
		}
	}
	return false
}

type countingReader struct {
	reader io.Reader
	count  int64
}

func (reader *countingReader) Read(p []byte) (int, error) {
	n, err := reader.reader.Read(p)
	reader.count += int64(n)
	return n, err
}
