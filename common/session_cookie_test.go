package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func resetSessionCookieSettingsAfterTest(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		SessionCookieSecure = false
		SessionCookieTrustedURLs = nil
	})
}

func TestInitSessionCookieSettingsDefaultsToInsecure(t *testing.T) {
	resetSessionCookieSettingsAfterTest(t)
	t.Setenv("SESSION_COOKIE_SECURE", "")
	t.Setenv("SESSION_COOKIE_TRUSTED_URL", "")

	require.NoError(t, InitSessionCookieSettings())
	assert.False(t, SessionCookieSecure)
	assert.Empty(t, SessionCookieTrustedURLs)
}

func TestInitSessionCookieSettingsRequiresBothEnvVars(t *testing.T) {
	t.Run("secure without trusted url", func(t *testing.T) {
		resetSessionCookieSettingsAfterTest(t)
		t.Setenv("SESSION_COOKIE_SECURE", "true")
		t.Setenv("SESSION_COOKIE_TRUSTED_URL", "")

		require.Error(t, InitSessionCookieSettings())
	})

	t.Run("trusted url without secure", func(t *testing.T) {
		resetSessionCookieSettingsAfterTest(t)
		t.Setenv("SESSION_COOKIE_SECURE", "")
		t.Setenv("SESSION_COOKIE_TRUSTED_URL", "https://example.com")

		require.Error(t, InitSessionCookieSettings())
	})
}

func TestInitSessionCookieSettingsRequiresHTTPSURL(t *testing.T) {
	resetSessionCookieSettingsAfterTest(t)
	t.Setenv("SESSION_COOKIE_SECURE", "true")
	t.Setenv("SESSION_COOKIE_TRUSTED_URL", "http://example.com")

	require.Error(t, InitSessionCookieSettings())
}

func TestInitSessionCookieSettingsEnablesSecureCookie(t *testing.T) {
	resetSessionCookieSettingsAfterTest(t)
	t.Setenv("SESSION_COOKIE_SECURE", "true")
	t.Setenv("SESSION_COOKIE_TRUSTED_URL", "https://example.com")

	require.NoError(t, InitSessionCookieSettings())
	assert.True(t, SessionCookieSecure)
	assert.Equal(t, []string{"https://example.com"}, SessionCookieTrustedURLs)
}

func TestInitSessionCookieSettingsAllowsMultipleTrustedURLs(t *testing.T) {
	resetSessionCookieSettingsAfterTest(t)
	t.Setenv("SESSION_COOKIE_SECURE", "true")
	t.Setenv("SESSION_COOKIE_TRUSTED_URL", "https://example.com, https://admin.example.com")

	require.NoError(t, InitSessionCookieSettings())
	assert.True(t, SessionCookieSecure)
	assert.Equal(t, []string{"https://example.com", "https://admin.example.com"}, SessionCookieTrustedURLs)
}

func TestInitSessionCookieSettingsRejectsEmptyTrustedURLInList(t *testing.T) {
	resetSessionCookieSettingsAfterTest(t)
	t.Setenv("SESSION_COOKIE_SECURE", "true")
	t.Setenv("SESSION_COOKIE_TRUSTED_URL", "https://example.com,")

	require.Error(t, InitSessionCookieSettings())
}

func TestInitSessionCookieSettingsNormalizesExactOrigins(t *testing.T) {
	resetSessionCookieSettingsAfterTest(t)
	t.Setenv("SESSION_COOKIE_SECURE", "true")
	t.Setenv("SESSION_COOKIE_TRUSTED_URL", "https://EXAMPLE.com:443,https://admin.example.com:8443/")

	require.NoError(t, InitSessionCookieSettings())
	assert.Equal(t, []string{"https://example.com", "https://admin.example.com:8443"}, SessionCookieTrustedURLs)
}

func TestInitSessionCookieSettingsRejectsNonOriginURLs(t *testing.T) {
	for _, trustedURL := range []string{
		"https://*.example.com",
		"https://user@example.com",
		"https://example.com/admin",
		"https://example.com?next=admin",
		"https://example.com#admin",
	} {
		t.Run(trustedURL, func(t *testing.T) {
			resetSessionCookieSettingsAfterTest(t)
			t.Setenv("SESSION_COOKIE_SECURE", "true")
			t.Setenv("SESSION_COOKIE_TRUSTED_URL", trustedURL)
			require.Error(t, InitSessionCookieSettings())
		})
	}
}
