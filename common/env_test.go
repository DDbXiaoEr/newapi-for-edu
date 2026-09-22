package common

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func unsetRequiredEnvVarsList(t *testing.T) {
	t.Helper()
	original, had := os.LookupEnv("REQUIRED_ENV_VARS")
	require.NoError(t, os.Unsetenv("REQUIRED_ENV_VARS"))
	t.Cleanup(func() {
		if had {
			require.NoError(t, os.Setenv("REQUIRED_ENV_VARS", original))
			return
		}
		require.NoError(t, os.Unsetenv("REQUIRED_ENV_VARS"))
	})
}

func TestCheckRequiredEnvVars(t *testing.T) {
	t.Run("defaults to SQL_DSN and LOG_SQL_DSN", func(t *testing.T) {
		t.Setenv("SQL_DSN", "")
		t.Setenv("LOG_SQL_DSN", "")
		unsetRequiredEnvVarsList(t)

		err := CheckRequiredEnvVars()
		require.Error(t, err)
		require.Contains(t, err.Error(), "SQL_DSN")
		require.Contains(t, err.Error(), "LOG_SQL_DSN")
	})

	t.Run("passes when defaults are set", func(t *testing.T) {
		t.Setenv("SQL_DSN", "postgresql://root:123456@localhost:5432/newapi?sslmode=disable")
		t.Setenv("LOG_SQL_DSN", "clickhouse://default:@localhost:9000/newapi_audit")
		unsetRequiredEnvVarsList(t)

		require.NoError(t, CheckRequiredEnvVars())
	})

	t.Run("empty REQUIRED_ENV_VARS disables the check", func(t *testing.T) {
		t.Setenv("REQUIRED_ENV_VARS", "")
		t.Setenv("SQL_DSN", "")
		t.Setenv("LOG_SQL_DSN", "")

		require.NoError(t, CheckRequiredEnvVars())
	})

	t.Run("custom REQUIRED_ENV_VARS list", func(t *testing.T) {
		t.Setenv("REQUIRED_ENV_VARS", "FOO, BAR")
		t.Setenv("FOO", "")
		t.Setenv("BAR", "set")

		err := CheckRequiredEnvVars()
		require.Error(t, err)
		require.Contains(t, err.Error(), "FOO")
		require.NotContains(t, err.Error(), "BAR")
	})
}
