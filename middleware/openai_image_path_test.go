package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRewriteOpenAIImagePath(t *testing.T) {
	tests := []struct {
		path      string
		rewritten string
		ok        bool
	}{
		{path: "/images/generations", rewritten: "/v1/images/generations", ok: true},
		{path: "/images/edits", rewritten: "/v1/images/edits", ok: true},
		{path: "/v1/images/generations", rewritten: "/v1/images/generations", ok: false},
		{path: "/v1/chat/completions", rewritten: "/v1/chat/completions", ok: false},
		{path: "/images", rewritten: "/images", ok: false},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got, ok := rewriteOpenAIImagePath(tt.path)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.rewritten, got)
		})
	}
}

func TestNormalizeOpenAIImagePathRewritesRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(NormalizeOpenAIImagePath())
	engine.POST("/images/generations", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path)
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/images/generations", nil)
	engine.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "/v1/images/generations", recorder.Body.String())
}
