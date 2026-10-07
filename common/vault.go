package common

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Environment variables that control optional configuration loading from a
// HashiCorp Vault KV v2 secret. When VAULT_ENABLED is unset (the default) the
// process keeps its existing behavior: values come from the process
// environment and the .env file only.
const (
	VaultEnabledEnv    = "VAULT_ENABLED"
	VaultAddrEnv       = "VAULT_ADDR"
	VaultTokenEnv      = "VAULT_TOKEN"
	VaultSecretPathEnv = "VAULT_SECRET_PATH"
	VaultNamespaceEnv  = "VAULT_NAMESPACE"
	VaultSkipVerifyEnv = "VAULT_SKIP_VERIFY"

	defaultVaultAddr       = "http://127.0.0.1:8200"
	defaultVaultSecretPath = "secret/data/newapi_edu"
	vaultRequestTimeout    = 10 * time.Second
)

// vaultKVv2Response is the subset of the KV v2 read response we consume.
type vaultKVv2Response struct {
	Data struct {
		Data map[string]any `json:"data"`
	} `json:"data"`
}

// VaultEnvEnabled reports whether configuration should be loaded from Vault.
// The default is false so existing deployments are unaffected by this switch.
func VaultEnvEnabled() bool {
	return GetEnvOrDefaultBool(VaultEnabledEnv, false)
}

// LoadEnvFromVault reads the configured KV v2 secret and injects every
// key/value pair into the process environment, overriding any value already
// set (including values loaded from .env). It is a no-op unless VAULT_ENABLED
// is true. It intentionally uses the standard library HTTP client so no new
// dependency is required.
func LoadEnvFromVault() error {
	if !VaultEnvEnabled() {
		return nil
	}

	addr := strings.TrimRight(strings.TrimSpace(GetEnvOrDefaultString(VaultAddrEnv, defaultVaultAddr)), "/")
	token := strings.TrimSpace(os.Getenv(VaultTokenEnv))
	path := strings.Trim(strings.TrimSpace(GetEnvOrDefaultString(VaultSecretPathEnv, defaultVaultSecretPath)), "/")

	if addr == "" {
		return fmt.Errorf("%s is enabled but %s is empty", VaultEnabledEnv, VaultAddrEnv)
	}
	if token == "" {
		return fmt.Errorf("%s is enabled but %s is empty", VaultEnabledEnv, VaultTokenEnv)
	}
	if path == "" {
		return fmt.Errorf("%s is enabled but %s is empty", VaultEnabledEnv, VaultSecretPathEnv)
	}

	req, err := http.NewRequest(http.MethodGet, addr+"/v1/"+path, nil)
	if err != nil {
		return fmt.Errorf("failed to build Vault request: %w", err)
	}
	req.Header.Set("X-Vault-Token", token)
	if namespace := strings.TrimSpace(os.Getenv(VaultNamespaceEnv)); namespace != "" {
		req.Header.Set("X-Vault-Namespace", namespace)
	}

	client := &http.Client{Timeout: vaultRequestTimeout}
	if GetEnvOrDefaultBool(VaultSkipVerifyEnv, false) {
		client.Transport = &http.Transport{TLSClientConfig: InsecureTLSConfig}
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to read Vault secret %q: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return fmt.Errorf("failed to read Vault secret %q: unexpected status %d: %s", path, resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var payload vaultKVv2Response
	if err := DecodeJson(resp.Body, &payload); err != nil {
		return fmt.Errorf("failed to decode Vault secret %q: %w", path, err)
	}
	if len(payload.Data.Data) == 0 {
		return fmt.Errorf("Vault secret %q contains no data", path)
	}

	keys := make([]string, 0, len(payload.Data.Data))
	for key, value := range payload.Data.Data {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}
		if err := os.Setenv(key, vaultValueToString(value)); err != nil {
			return fmt.Errorf("failed to set environment variable %q from Vault: %w", key, err)
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	SysLog(fmt.Sprintf("loaded %d environment variables from Vault path %q: %s", len(keys), path, strings.Join(keys, ", ")))
	return nil
}

func vaultValueToString(value any) string {
	switch v := value.(type) {
	case nil:
		return ""
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}
