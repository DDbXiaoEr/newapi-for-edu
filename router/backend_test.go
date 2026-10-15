package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPureBackendRootIsAvailable(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	SetPureBackendRouter(engine, func(c *gin.Context) { c.Next() })

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	engine.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), `"mode":"backend"`)
	assert.NotContains(t, recorder.Body.String(), "Frontend not available")
}

func TestPureBackendCompatPathsReachRelayAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	SetRelayRouter(engine)
	SetPureBackendRouter(engine, func(c *gin.Context) { c.Next() })
	handler := middleware.WrapOpenAICompatHandler(engine)

	tests := []string{
		"/chat/completions",
		"/openai/v1/chat/completions",
		"/v1/chat/completions",
		"/models",
		"/v1/models",
	}
	for _, path := range tests {
		t.Run(path, func(t *testing.T) {
			method := http.MethodGet
			if path != "/models" && path != "/v1/models" {
				method = http.MethodPost
			}
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(method, path, nil)
			handler.ServeHTTP(recorder, request)
			assert.Equal(t, http.StatusUnauthorized, recorder.Code)
			assert.Contains(t, recorder.Body.String(), "token.invalid")
			assert.NotContains(t, recorder.Body.String(), "Frontend not available")
		})
	}
}
