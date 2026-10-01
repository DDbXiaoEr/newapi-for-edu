package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeAboutContentType(t *testing.T) {
	assert.Equal(t, "html", normalizeAboutContentType("html", "# About"))
	assert.Equal(t, "markdown", normalizeAboutContentType("markdown", "<p>About</p>"))
	assert.Equal(t, "html", normalizeAboutContentType("", "https://example.com"))
	assert.Equal(t, "html", normalizeAboutContentType("", "<p>About us</p>"))
	assert.Equal(t, "markdown", normalizeAboutContentType("", "# About us"))
	assert.Equal(t, "markdown", normalizeAboutContentType("", ""))
}
