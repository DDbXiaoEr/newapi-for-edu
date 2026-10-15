package router

import (
	"net/http"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/controller"
	"github.com/QuantumNous/new-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetPureBackendRouter(router *gin.Engine, pluginDispatcher gin.HandlerFunc) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "ok",
			"data": gin.H{
				"mode":    "backend",
				"version": common.Version,
			},
		})
	})
	router.NoRoute(pluginDispatcher, func(c *gin.Context) {
		c.Set(middleware.RouteTagKey, "web")
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/v1") || strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/images/") {
			controller.RelayNotFound(c)
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"message": "API endpoint not found. Use OpenAI-compatible paths such as /v1/chat/completions or /chat/completions.",
				"type":    "invalid_request_error",
			},
		})
	})
}
