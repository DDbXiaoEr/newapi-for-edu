package controller

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service"

	"github.com/QuantumNous/new-api/pkg/testdb"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func setupCodexOAuthTest(t *testing.T) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	previousRedis, previousSecret := common.RedisEnabled, common.SessionSecret
	previousExchange := exchangeCodexAuthorizationCode
	testdb.OpenBound(t, &model.DB, &model.LOG_DB, model.InitColumnNames, testdb.Options{Models: []any{&model.AuthFlow{}, &model.Channel{}}})
	common.RedisEnabled = false
	common.SessionSecret = "codex-oauth-test-secret"
	t.Cleanup(func() {
		common.RedisEnabled, common.SessionSecret = previousRedis, previousSecret
		exchangeCodexAuthorizationCode = previousExchange
	})
}

func newCodexOAuthJWT(t *testing.T, accountID, email string) string {
	t.Helper()
	payload, err := common.Marshal(map[string]any{
		"https://api.openai.com/auth": map[string]any{"chatgpt_account_id": accountID},
		"email":                       email,
	})
	require.NoError(t, err)
	return "header." + base64.RawURLEncoding.EncodeToString(payload) + ".sig"
}

func TestStartCodexOAuthStoresPKCEInAuthFlow(t *testing.T) {
	setupCodexOAuthTest(t)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/channel/codex/oauth/start", nil)
	c.Set("id", 7)

	StartCodexOAuth(c)

	require.Equal(t, http.StatusOK, recorder.Code)
	var response struct {
		Success bool `json:"success"`
		Data    struct {
			AuthorizeURL string `json:"authorize_url"`
			FlowToken    string `json:"flow_token"`
		} `json:"data"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &response))
	require.True(t, response.Success)
	require.NotEmpty(t, response.Data.FlowToken)
	require.Contains(t, response.Data.AuthorizeURL, "state="+response.Data.FlowToken)

	flow, err := model.GetAuthFlow(response.Data.FlowToken, model.AuthFlowMatch{
		Purpose: model.AuthFlowPurposeCodexOAuth,
		UserId:  7,
	})
	require.NoError(t, err)
	var payload codexOAuthFlowPayload
	require.NoError(t, common.UnmarshalJsonStr(flow.Payload, &payload))
	require.NotEmpty(t, payload.Verifier)
	require.Equal(t, 0, payload.ChannelID)
}

func TestCompleteCodexOAuthUsesAuthFlowInsteadOfCookie(t *testing.T) {
	setupCodexOAuthTest(t)
	payload, err := common.Marshal(codexOAuthFlowPayload{Verifier: "pkce-verifier"})
	require.NoError(t, err)
	token, _, err := model.CreateAuthFlow(model.AuthFlowCreate{
		Purpose:   model.AuthFlowPurposeCodexOAuth,
		UserId:    7,
		Payload:   string(payload),
		ExpiresAt: time.Now().Add(time.Minute),
	})
	require.NoError(t, err)

	accessToken := newCodexOAuthJWT(t, "acct-1", "user@example.com")
	exchangeCodexAuthorizationCode = func(ctx context.Context, code, verifier, proxyURL string) (*service.CodexOAuthTokenResult, error) {
		require.Equal(t, "auth-code", code)
		require.Equal(t, "pkce-verifier", verifier)
		return &service.CodexOAuthTokenResult{
			AccessToken:  accessToken,
			RefreshToken: "refresh-token",
			ExpiresAt:    time.Now().Add(time.Hour),
		}, nil
	}

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	body := fmt.Sprintf(`{"input":"https://localhost:1455/auth/callback?code=auth-code&state=%s"}`, token)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/channel/codex/oauth/complete", bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", 7)

	CompleteCodexOAuth(c)

	require.Equal(t, http.StatusOK, recorder.Code)
	var response struct {
		Success bool `json:"success"`
		Data    struct {
			Key       string `json:"key"`
			AccountID string `json:"account_id"`
			Email     string `json:"email"`
		} `json:"data"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &response))
	require.True(t, response.Success, recorder.Body.String())
	require.Equal(t, "acct-1", response.Data.AccountID)
	require.Equal(t, "user@example.com", response.Data.Email)
	require.Contains(t, response.Data.Key, "refresh-token")

	_, err = model.GetAuthFlow(token, model.AuthFlowMatch{Purpose: model.AuthFlowPurposeCodexOAuth, UserId: 7})
	require.ErrorIs(t, err, model.ErrAuthFlowConsumed)
}

func TestCompleteCodexOAuthRejectsOtherUsersFlow(t *testing.T) {
	setupCodexOAuthTest(t)
	payload, err := common.Marshal(codexOAuthFlowPayload{Verifier: "pkce-verifier"})
	require.NoError(t, err)
	token, _, err := model.CreateAuthFlow(model.AuthFlowCreate{
		Purpose:   model.AuthFlowPurposeCodexOAuth,
		UserId:    7,
		Payload:   string(payload),
		ExpiresAt: time.Now().Add(time.Minute),
	})
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	body := fmt.Sprintf(`{"input":"https://localhost:1455/auth/callback?code=auth-code&state=%s"}`, token)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/channel/codex/oauth/complete", bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("id", 8)

	CompleteCodexOAuth(c)

	require.Equal(t, http.StatusOK, recorder.Code)
	var response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &response))
	require.False(t, response.Success)
	require.Contains(t, response.Message, "oauth flow not started")
}

func TestStartCodexOAuthForChannelRejectsNonCodex(t *testing.T) {
	setupCodexOAuthTest(t)
	channel := model.Channel{Type: constant.ChannelTypeOpenAI, Name: "openai", Key: "sk-test", Status: common.ChannelStatusEnabled}
	require.NoError(t, model.DB.Create(&channel).Error)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = gin.Params{{Key: "id", Value: strconv.Itoa(channel.Id)}}
	c.Request = httptest.NewRequest(http.MethodPost, "/api/channel/"+strconv.Itoa(channel.Id)+"/codex/oauth/start", nil)
	c.Set("id", 7)

	StartCodexOAuthForChannel(c)

	require.Equal(t, http.StatusOK, recorder.Code)
	var response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	require.NoError(t, common.Unmarshal(recorder.Body.Bytes(), &response))
	require.False(t, response.Success)
	require.Contains(t, response.Message, "not Codex")
}
