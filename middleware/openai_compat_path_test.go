package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRewriteOpenAICompatPath(t *testing.T) {
	tests := []struct {
		path      string
		rewritten string
		ok        bool
	}{
		{path: "/", rewritten: "/", ok: false},
		{path: "/v1/chat/completions", rewritten: "/v1/chat/completions", ok: false},
		{path: "/chat/completions", rewritten: "/v1/chat/completions", ok: true},
		{path: "/chat/completions/", rewritten: "/v1/chat/completions/", ok: true},
		{path: "/openai/v1/chat/completions", rewritten: "/v1/chat/completions", ok: true},
		{path: "/openai/chat/completions", rewritten: "/v1/chat/completions", ok: true},
		{path: "/models", rewritten: "/v1/models", ok: true},
		{path: "/models/MiniMax-M2.7", rewritten: "/v1/models/MiniMax-M2.7", ok: true},
		{path: "/embeddings", rewritten: "/v1/embeddings", ok: true},
		{path: "/messages", rewritten: "/v1/messages", ok: true},
		{path: "/responses", rewritten: "/v1/responses", ok: true},
		{path: "/images/generations", rewritten: "/v1/images/generations", ok: true},
		{path: "/api/status", rewritten: "/api/status", ok: false},
		{path: "/healthz", rewritten: "/healthz", ok: false},
		{path: "/pg/chat/completions", rewritten: "/pg/chat/completions", ok: false},
		{path: "/index.html", rewritten: "/index.html", ok: false},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got, ok := rewriteOpenAICompatPath(tt.path)
			assert.Equal(t, tt.ok, ok)
			assert.Equal(t, tt.rewritten, got)
		})
	}
}

func TestWrapOpenAICompatHandlerRewritesBeforeRouting(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.POST("/v1/chat/completions", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path)
	})
	engine.GET("/v1/models", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path+"?"+c.Request.URL.RawQuery)
	})

	handler := WrapOpenAICompatHandler(engine)

	t.Run("chat completions", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/chat/completions", nil)
		handler.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "/v1/chat/completions", recorder.Body.String())
	})

	t.Run("openai prefix", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/openai/v1/chat/completions", nil)
		handler.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "/v1/chat/completions", recorder.Body.String())
	})

	t.Run("models with query", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/models?key=abc", nil)
		handler.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "/v1/models?key=abc", recorder.Body.String())
	})
}
