package oauth

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/i18n"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/system_setting"
	"github.com/gin-gonic/gin"
)

func init() {
	Register("cas", &CasProvider{})
}

type CasProvider struct{}

type casServiceResponse struct {
	XMLName               xml.Name          `xml:"http://www.yale.edu/tp/cas serviceResponse"`
	AuthenticationSuccess *casAuthSuccess   `xml:"http://www.yale.edu/tp/cas authenticationSuccess"`
	AuthenticationFailure *casAuthFailure   `xml:"http://www.yale.edu/tp/cas authenticationFailure"`
}

type casAuthSuccess struct {
	User       string        `xml:"user"`
	Attributes casAttributes `xml:"attributes"`
}

type casAttributes struct {
	Entries []casAttribute `xml:",any"`
}

type casAttribute struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (p *CasProvider) GetName() string {
	return "CAS"
}

func (p *CasProvider) IsEnabled() bool {
	return system_setting.GetCasSettings().Enabled
}

func (p *CasProvider) ExchangeToken(ctx context.Context, code string, c *gin.Context) (*OAuthToken, error) {
	if code == "" {
		return nil, NewOAuthError(i18n.MsgOAuthInvalidCode, nil)
	}

	settings := system_setting.GetCasSettings()
	serviceUrl := fmt.Sprintf("%s/oauth/cas", system_setting.ServerAddress)

	logger.LogDebug(ctx, "[OAuth-CAS] ExchangeToken: ticket=%s..., service=%s", code[:min(len(code), 10)], serviceUrl)

	validateUrl, err := url.Parse(strings.TrimRight(settings.ServerUrl, "/") + "/serviceValidate")
	if err != nil {
		return nil, NewOAuthErrorWithRaw(i18n.MsgOAuthConnectFailed, map[string]any{"Provider": "CAS"}, err.Error())
	}

	q := validateUrl.Query()
	q.Set("service", serviceUrl)
	q.Set("ticket", code)
	validateUrl.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", validateUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("[OAuth-CAS] ExchangeToken error: %s", err.Error()))
		return nil, NewOAuthErrorWithRaw(i18n.MsgOAuthConnectFailed, map[string]any{"Provider": "CAS"}, err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("[OAuth-CAS] ExchangeToken read error: %s", err.Error()))
		return nil, err
	}

	logger.LogDebug(ctx, "[OAuth-CAS] ExchangeToken response: %s", string(body))

	var casResp casServiceResponse
	err = xml.Unmarshal(body, &casResp)
	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("[OAuth-CAS] ExchangeToken XML parse error: %s", err.Error()))
		return nil, NewOAuthError(i18n.MsgOAuthTokenFailed, map[string]any{"Provider": "CAS"})
	}

	if casResp.AuthenticationFailure != nil {
		codeStr := casResp.AuthenticationFailure.Code
		descStr := casResp.AuthenticationFailure.Description
		logger.LogError(ctx, fmt.Sprintf("[OAuth-CAS] ExchangeToken failed: code=%s, desc=%s", codeStr, descStr))
		return nil, NewOAuthErrorWithRaw(i18n.MsgOAuthTokenFailed, map[string]any{"Provider": "CAS"}, descStr)
	}

	if casResp.AuthenticationSuccess == nil || casResp.AuthenticationSuccess.User == "" {
		logger.LogError(ctx, "[OAuth-CAS] ExchangeToken failed: no user in response")
		return nil, NewOAuthError(i18n.MsgOAuthUserInfoEmpty, map[string]any{"Provider": "CAS"})
	}

	logger.LogDebug(ctx, "[OAuth-CAS] ExchangeToken success: user=%s", casResp.AuthenticationSuccess.User)

	return &OAuthToken{
		AccessToken: code,
		TokenType:   "cas",
		IDToken:     string(body),
	}, nil
}

func (p *CasProvider) GetUserInfo(ctx context.Context, token *OAuthToken) (*OAuthUser, error) {
	settings := system_setting.GetCasSettings()

	logger.LogDebug(ctx, "[OAuth-CAS] GetUserInfo: parsing from cached CAS response")

	if token.IDToken == "" {
		return nil, NewOAuthError(i18n.MsgOAuthGetUserErr, map[string]any{"Provider": "CAS"})
	}

	var casResp casServiceResponse
	err := xml.Unmarshal([]byte(token.IDToken), &casResp)
	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("[OAuth-CAS] GetUserInfo XML parse error: %s", err.Error()))
		return nil, err
	}

	if casResp.AuthenticationSuccess == nil {
		return nil, NewOAuthError(i18n.MsgOAuthGetUserErr, map[string]any{"Provider": "CAS"})
	}

	username := casResp.AuthenticationSuccess.User
	displayName := username
	email := ""
	providerUserID := username

	attrMap := make(map[string]string)
	for _, attr := range casResp.AuthenticationSuccess.Attributes.Entries {
		attrMap[strings.ToLower(attr.XMLName.Local)] = attr.Value
	}

	if settings.UsernameAttribute != "" {
		if v, ok := attrMap[strings.ToLower(settings.UsernameAttribute)]; ok && v != "" {
			username = v
			providerUserID = v
		}
	}

	if settings.DisplayNameAttribute != "" {
		if v, ok := attrMap[strings.ToLower(settings.DisplayNameAttribute)]; ok && v != "" {
			displayName = v
		}
	} else {
		if v, ok := attrMap["displayname"]; ok && v != "" {
			displayName = v
		} else if v, ok := attrMap["name"]; ok && v != "" {
			displayName = v
		}
	}

	if settings.EmailAttribute != "" {
		if v, ok := attrMap[strings.ToLower(settings.EmailAttribute)]; ok && v != "" {
			email = v
		}
	} else {
		if v, ok := attrMap["email"]; ok && v != "" {
			email = v
		} else if v, ok := attrMap["mail"]; ok && v != "" {
			email = v
		}
	}

	if settings.AccessAttribute != "" {
		attrValue, ok := attrMap[strings.ToLower(settings.AccessAttribute)]
		if !ok || (settings.AccessAttributeValue != "" && attrValue != settings.AccessAttributeValue) {
			logger.LogWarn(ctx, fmt.Sprintf("[OAuth-CAS] GetUserInfo: access denied by attribute restriction (attr=%s, value=%s, required=%s)",
				settings.AccessAttribute, attrValue, settings.AccessAttributeValue))
			return nil, &CasAccessDeniedError{}
		}
	}

	logger.LogDebug(ctx, "[OAuth-CAS] GetUserInfo success: user=%s, displayName=%s, email=%s", username, displayName, email)

	return &OAuthUser{
		ProviderUserID: providerUserID,
		Username:       username,
		DisplayName:    displayName,
		Email:          email,
	}, nil
}

func (p *CasProvider) IsUserIDTaken(providerUserID string) bool {
	return model.IsCasIdAlreadyTaken(providerUserID)
}

func (p *CasProvider) FillUserByProviderID(user *model.User, providerUserID string) error {
	user.CasId = providerUserID
	return user.FillUserByCasId()
}

func (p *CasProvider) SetProviderUserID(user *model.User, providerUserID string) {
	user.CasId = providerUserID
}

func (p *CasProvider) GetProviderPrefix() string {
	return "cas_"
}

type casAuthFailure struct {
	Code        string `xml:"code,attr"`
	Description string `xml:",chardata"`
}

type CasAccessDeniedError struct{}

func (e *CasAccessDeniedError) Error() string {
	return "CAS access denied: attribute restriction"
}
