package controller

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/i18n"
	"github.com/QuantumNous/new-api/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func performLDAPLoginTestRequest(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()

	gin.SetMode(gin.TestMode)
	if err := i18n.Init(); err != nil {
		t.Fatalf("failed to init i18n: %v", err)
	}

	recorder := httptest.NewRecorder()
	router := gin.New()
	router.POST("/api/option/ldap/test", TestLDAPLogin)

	request := httptest.NewRequest(http.MethodPost, "/api/option/ldap/test", bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, request)
	return recorder
}

func TestLDAPLoginRejectsMissingCredentials(t *testing.T) {
	original := testLDAPLogin
	t.Cleanup(func() {
		testLDAPLogin = original
	})
	called := false
	testLDAPLogin = func(ctx context.Context, identifier string, password string) (*service.LDAPLoginTestResult, error) {
		called = true
		return nil, nil
	}

	recorder := performLDAPLoginTestRequest(t, `{"username":"alice"}`)
	require.Equal(t, http.StatusOK, recorder.Code)

	var payload struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &payload))
	require.False(t, payload.Success)
	require.Equal(t, "LDAP username and password are required", payload.Message)
	require.False(t, called)
}

func TestLDAPLoginReturnsAttributeNamesWithoutValues(t *testing.T) {
	original := testLDAPLogin
	t.Cleanup(func() {
		testLDAPLogin = original
	})
	testLDAPLogin = func(ctx context.Context, identifier string, password string) (*service.LDAPLoginTestResult, error) {
		if identifier != "alice" || password != "secret" {
			t.Fatalf("unexpected credentials identifier=%q password=%q", identifier, password)
		}
		return &service.LDAPLoginTestResult{
			AttributeNames: []string{"dn", "mail", "memberOf", "uid"},
		}, nil
	}

	recorder := performLDAPLoginTestRequest(t, `{"username":"alice","password":"secret"}`)
	require.Equal(t, http.StatusOK, recorder.Code)

	var payload struct {
		Success bool `json:"success"`
		Data    struct {
			AttributeNames []string `json:"attribute_names"`
		} `json:"data"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &payload))
	require.True(t, payload.Success)
	require.Equal(t, []string{"dn", "mail", "memberOf", "uid"}, payload.Data.AttributeNames)

	raw := map[string]any{}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &raw))
	data, _ := raw["data"].(map[string]any)
	require.NotNil(t, data)
	_, hasValues := data["attribute_values"]
	require.False(t, hasValues)
}

func TestLDAPLoginMapsDirectoryErrors(t *testing.T) {
	original := testLDAPLogin
	t.Cleanup(func() {
		testLDAPLogin = original
	})
	testLDAPLogin = func(ctx context.Context, identifier string, password string) (*service.LDAPLoginTestResult, error) {
		return nil, service.ErrLDAPInvalidCredentials
	}

	recorder := performLDAPLoginTestRequest(t, `{"username":"alice","password":"wrong"}`)
	require.Equal(t, http.StatusOK, recorder.Code)

	var payload struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &payload))
	require.False(t, payload.Success)
	require.Equal(t, "LDAP username or password is incorrect", payload.Message)
}

func TestLDAPLoginRejectsInvalidJSON(t *testing.T) {
	original := testLDAPLogin
	t.Cleanup(func() {
		testLDAPLogin = original
	})
	testLDAPLogin = func(ctx context.Context, identifier string, password string) (*service.LDAPLoginTestResult, error) {
		return nil, errors.New("should not be called")
	}

	recorder := performLDAPLoginTestRequest(t, `{`)
	require.Equal(t, http.StatusOK, recorder.Code)

	var payload struct {
		Success bool `json:"success"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &payload))
	require.False(t, payload.Success)
}
