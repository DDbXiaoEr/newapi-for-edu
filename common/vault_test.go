package common

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVaultEnvDisabledByDefault(t *testing.T) {
	t.Setenv(VaultEnabledEnv, "")
	t.Setenv("SQL_DSN", "keep-me")

	require.False(t, VaultEnvEnabled())
	require.NoError(t, LoadEnvFromVault())
	require.Equal(t, "keep-me", os.Getenv("SQL_DSN"))
}

func TestLoadEnvFromVaultRequiresToken(t *testing.T) {
	t.Setenv(VaultEnabledEnv, "true")
	t.Setenv(VaultAddrEnv, "http://127.0.0.1:8200")
	t.Setenv(VaultTokenEnv, "")

	err := LoadEnvFromVault()
	require.Error(t, err)
	require.Contains(t, err.Error(), VaultTokenEnv)
}

func TestLoadEnvFromVaultOverridesEnv(t *testing.T) {
	var gotToken, gotNamespace string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotToken = r.Header.Get("X-Vault-Token")
		gotNamespace = r.Header.Get("X-Vault-Namespace")
		require.Equal(t, "/v1/secret/data/newapi_edu", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"data":{"SQL_DSN":"from-vault","PORT":3000,"ERROR_LOG_ENABLED":true}}}`))
	}))
	defer server.Close()

	t.Setenv(VaultEnabledEnv, "true")
	t.Setenv(VaultAddrEnv, server.URL)
	t.Setenv(VaultTokenEnv, "test-token")
	t.Setenv(VaultNamespaceEnv, "team-a")
	t.Setenv("SQL_DSN", "from-dotenv")

	require.NoError(t, LoadEnvFromVault())

	require.Equal(t, "test-token", gotToken)
	require.Equal(t, "team-a", gotNamespace)
	require.Equal(t, "from-vault", os.Getenv("SQL_DSN"))
	require.Equal(t, "3000", os.Getenv("PORT"))
	require.Equal(t, "true", os.Getenv("ERROR_LOG_ENABLED"))
}

func TestLoadEnvFromVaultFailsOnNonOK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "permission denied", http.StatusForbidden)
	}))
	defer server.Close()

	t.Setenv(VaultEnabledEnv, "true")
	t.Setenv(VaultAddrEnv, server.URL)
	t.Setenv(VaultTokenEnv, "test-token")

	err := LoadEnvFromVault()
	require.Error(t, err)
	require.Contains(t, err.Error(), "403")
}
