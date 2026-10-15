package middleware

import (
	"net/http"
	"strings"
)

var openaiCompatPrefixes = []string{
	"/alpha/search",
	"/audio/",
	"/chat/completions",
	"/completions",
	"/edits",
	"/embeddings",
	"/engines/",
	"/files",
	"/fine-tunes",
	"/images",
	"/messages",
	"/models",
	"/moderations",
	"/realtime",
	"/rerank",
	"/responses",
	"/tasks",
	"/video/",
	"/videos",
}

func WrapOpenAICompatHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		applyOpenAICompatPath(r)
		next.ServeHTTP(w, r)
	})
}

func applyOpenAICompatPath(r *http.Request) {
	if r == nil || r.URL == nil {
		return
	}
	rewritten, ok := rewriteOpenAICompatPath(r.URL.Path)
	if !ok {
		return
	}
	r.URL.Path = rewritten
	r.URL.RawPath = ""
	if r.URL.RawQuery != "" {
		r.RequestURI = rewritten + "?" + r.URL.RawQuery
	} else {
		r.RequestURI = rewritten
	}
}

func rewriteOpenAICompatPath(path string) (string, bool) {
	if path == "" || path == "/" {
		return path, false
	}
	if strings.HasPrefix(path, "/openai/") {
		rest := strings.TrimPrefix(path, "/openai")
		if rest == "/v1" || strings.HasPrefix(rest, "/v1/") {
			return rest, true
		}
		path = rest
	}
	if path == "/v1" || strings.HasPrefix(path, "/v1/") ||
		strings.HasPrefix(path, "/api/") || path == "/api" ||
		strings.HasPrefix(path, "/v1beta") ||
		strings.HasPrefix(path, "/healthz") ||
		strings.HasPrefix(path, "/pg/") ||
		strings.HasPrefix(path, "/mj") ||
		strings.HasPrefix(path, "/dashboard") {
		return path, false
	}
	for _, prefix := range openaiCompatPrefixes {
		if path == strings.TrimSuffix(prefix, "/") || strings.HasPrefix(path, prefix) {
			return "/v1" + path, true
		}
	}
	return path, false
}
