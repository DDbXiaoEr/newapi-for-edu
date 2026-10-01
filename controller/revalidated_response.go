package controller

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/gin-gonic/gin"
)

// etagVersionPublicContent namespaces the public-content ETag; bump it when
// the JSON envelope served by serveRevalidatedJSON changes shape.
const etagVersionPublicContent = "public-content:v1"

// etagVersionAboutContent namespaces the about-page ETag; bump it when
// the JSON envelope served by serveRevalidatedAboutJSON changes shape.
const etagVersionAboutContent = "about-content:v1"

type publicContentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type aboutContentResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Data        string `json:"data"`
	ContentType string `json:"content_type"`
}

// serveRevalidatedJSON writes public content as JSON with a weak
// content-derived ETag and answers conditional requests with 304 Not
// Modified. The ETag is a weak validator derived from the content, so it is
// stable across replicas and JSON encodings, and a new one is issued when
// the content changes. Cache-Control: no-cache forces revalidation before
// reuse, so an admin edit takes effect on the next request; Vary:
// Accept-Encoding keeps the gzip and identity encodings apart in shared
// caches.
func serveRevalidatedJSON(c *gin.Context, content string) {
	body, err := common.Marshal(publicContentResponse{
		Success: true,
		Message: "",
		Data:    content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	etag := common.ETagFor(etagVersionPublicContent, content)

	c.Header("ETag", etag)
	c.Header("Cache-Control", "no-cache")
	c.Header("Vary", "Accept-Encoding")

	if common.ETagMatches(c.GetHeader("If-None-Match"), etag) {
		c.Status(http.StatusNotModified)
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", body)
}

func serveRevalidatedAboutJSON(c *gin.Context, content, contentType string) {
	normalizedType := normalizeAboutContentType(contentType, content)
	body, err := common.Marshal(aboutContentResponse{
		Success:     true,
		Message:     "",
		Data:        content,
		ContentType: normalizedType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	etag := common.ETagFor(etagVersionAboutContent, content+"\n"+normalizedType)

	c.Header("ETag", etag)
	c.Header("Cache-Control", "no-cache")
	c.Header("Vary", "Accept-Encoding")

	if common.ETagMatches(c.GetHeader("If-None-Match"), etag) {
		c.Status(http.StatusNotModified)
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", body)
}

func normalizeAboutContentType(contentType, content string) string {
	switch strings.ToLower(strings.TrimSpace(contentType)) {
	case "html":
		return "html"
	case "markdown":
		return "markdown"
	default:
		trimmed := strings.TrimSpace(content)
		if isAboutHttpURL(trimmed) || looksLikeHTML(trimmed) {
			return "html"
		}
		return "markdown"
	}
}

func isAboutHttpURL(value string) bool {
	parsed, err := url.Parse(value)
	if err != nil {
		return false
	}
	return parsed.Scheme == "http" || parsed.Scheme == "https"
}

func looksLikeHTML(value string) bool {
	lower := strings.ToLower(value)
	if strings.Contains(lower, "<!doctype html") ||
		strings.Contains(lower, "<html") ||
		strings.Contains(lower, "<head") ||
		strings.Contains(lower, "<body") ||
		strings.Contains(lower, "<style") ||
		strings.Contains(lower, "<script") {
		return true
	}
	return strings.Contains(value, "<") && strings.Contains(value, ">")
}
