package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// NormalizeOpenAIImagePath rewrites Cherry Studio-style OpenAI image paths
// that omit the /v1 prefix so later relay matching, channel selection, and
// upstream URL construction keep using the canonical /v1/images routes.
func NormalizeOpenAIImagePath() gin.HandlerFunc {
	return func(c *gin.Context) {
		if rewritten, ok := rewriteOpenAIImagePath(c.Request.URL.Path); ok {
			c.Request.URL.Path = rewritten
		}
		c.Next()
	}
}

func rewriteOpenAIImagePath(path string) (string, bool) {
	if path == "/images/generations" || strings.HasPrefix(path, "/images/generations/") ||
		path == "/images/edits" || strings.HasPrefix(path, "/images/edits/") {
		return "/v1" + path, true
	}
	return path, false
}
