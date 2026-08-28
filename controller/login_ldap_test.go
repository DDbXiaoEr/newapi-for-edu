package controller

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/i18n"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting/system_setting"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupLoginControllerTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	gin.SetMode(gin.TestMode)
	common.SetMainDatabaseType(common.DatabaseTypeSQLite)
	common.RedisEnabled = false
	common.PasswordLoginEnabled = true
	common.RegisterEnabled = true
	common.QuotaForNewUser = 0
	common.QuotaForInvitee = 0
	common.QuotaForInviter = 0
	if err := i18n.Init(); err != nil {
		t.Fatalf("failed to init i18n: %v", err)
	}

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", strings.ReplaceAll(t.Name(), "/", "_"))
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}
	model.DB = db
	model.LOG_DB = db

	if err := db.AutoMigrate(&model.User{}, &model.Log{}, &model.TwoFA{}, &model.TwoFABackupCode{}, &model.UserSession{}, &model.AuthFlow{}); err != nil {
		t.Fatalf("failed to migrate login test tables: %v", err)
	}

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
	})

	return db
}

func seedLoginUser(t *testing.T, db *gorm.DB, username string, password string, status int) *model.User {
	t.Helper()

	hashedPassword, err := common.Password2Hash(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	user := &model.User{
		Username:    username,
		Password:    hashedPassword,
		DisplayName: username,
		Role:        common.RoleCommonUser,
		Status:      status,
		Group:       "default",
		Quota:       0,
		AffCode:     common.GetRandomString(4),
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	return user
}

func performLoginRequest(t *testing.T, body string) (*httptest.ResponseRecorder, *http.Request) {
	t.Helper()

	recorder := httptest.NewRecorder()
	router := gin.New()
	store := cookie.NewStore([]byte("login-test-secret"))
	router.Use(sessions.Sessions("new-api-session", store))
	router.POST("/api/user/login", Login)

	request := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, request)
	return recorder, request
}

func fetchUserByID(t *testing.T, id int) *model.User {
	t.Helper()

	user, err := model.GetUserById(id, true)
	if err != nil {
		t.Fatalf("failed to fetch user by id %d: %v", id, err)
	}
	return user
}

func getSessionCookie(t *testing.T, recorder *httptest.ResponseRecorder, cookieName string) *http.Cookie {
	t.Helper()

	for _, cookie := range recorder.Result().Cookies() {
		if cookie.Name == cookieName {
			return cookie
		}
	}
	t.Fatalf("session cookie %q not found in response", cookieName)
	return nil
}

func TestLoginLocalPasswordSuccessSkipsLDAP(t *testing.T) {
	db := setupLoginControllerTestDB(t)
	user := seedLoginUser(t, db, "alice", "old-password", common.UserStatusEnabled)

	originalAuthenticateLDAP := authenticateLDAP
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
	})

	ldapCalled := false
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		ldapCalled = true
		return nil, nil
	}

	recorder, _ := performLoginRequest(t, `{"username":"alice","password":"old-password"}`)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected HTTP 200, got %d", recorder.Code)
	}

	var response struct {
		Success bool `json:"success"`
		Data    struct {
			User struct {
				ID int `json:"id"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := common.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if !response.Success {
		t.Fatalf("expected success response, got body: %s", recorder.Body.String())
	}
	if response.Data.User.ID != user.Id {
		t.Fatalf("expected logged in user id %d, got %d", user.Id, response.Data.User.ID)
	}
	if ldapCalled {
		t.Fatalf("expected LDAP authenticator not to be called")
	}

	storedUser := fetchUserByID(t, user.Id)
	if !common.ValidatePasswordAndHash("old-password", storedUser.Password) {
		t.Fatalf("expected stored password hash to remain unchanged")
	}
}

func TestLoginLocalPasswordMismatchDoesNotCallLDAPWhenDisabled(t *testing.T) {
	db := setupLoginControllerTestDB(t)
	user := seedLoginUser(t, db, "alice", "old-password", common.UserStatusEnabled)

	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	system_setting.GetLDAPSettings().Enabled = false
	ldapCalled := false
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		ldapCalled = true
		return nil, nil
	}

	_, err := authenticateLoginUser(context.Background(), "alice", "wrong-password")
	if !errors.Is(err, model.ErrInvalidCredentials) {
		t.Fatalf("expected invalid credentials, got %v", err)
	}
	if ldapCalled {
		t.Fatalf("expected LDAP authenticator not to be called when LDAP is disabled")
	}

	storedUser := fetchUserByID(t, user.Id)
	if !common.ValidatePasswordAndHash("old-password", storedUser.Password) {
		t.Fatalf("expected local password hash to remain unchanged")
	}
}

func TestLoginFallsBackToLDAPAndUpdatesLocalPassword(t *testing.T) {
	db := setupLoginControllerTestDB(t)
	user := seedLoginUser(t, db, "alice", "old-password", common.UserStatusEnabled)

	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	system_setting.GetLDAPSettings().Enabled = true
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		if identifier != "alice" {
			t.Fatalf("unexpected identifier %q", identifier)
		}
		if password != "new-password" {
			t.Fatalf("unexpected password %q", password)
		}
		return &service.LDAPUser{
			Username:    "alice",
			DisplayName: "Alice LDAP",
			Email:       "alice@example.com",
		}, nil
	}

	recorder, _ := performLoginRequest(t, `{"username":"alice","password":"new-password"}`)
	if recorder.Code != http.StatusOK {
		t.Fatalf("expected HTTP 200, got %d", recorder.Code)
	}

	var response struct {
		Success bool `json:"success"`
		Data    struct {
			User struct {
				ID int `json:"id"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := common.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if !response.Success {
		t.Fatalf("expected success response, got body: %s", recorder.Body.String())
	}
	if response.Data.User.ID != user.Id {
		t.Fatalf("expected logged in user id %d, got %d", user.Id, response.Data.User.ID)
	}

	storedUser := fetchUserByID(t, user.Id)
	if !common.ValidatePasswordAndHash("new-password", storedUser.Password) {
		t.Fatalf("expected local password to be updated after successful LDAP login")
	}
}

func TestLoginDisabledUserRejectsLDAPFallback(t *testing.T) {
	db := setupLoginControllerTestDB(t)
	user := seedLoginUser(t, db, "alice", "old-password", common.UserStatusDisabled)

	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	system_setting.GetLDAPSettings().Enabled = true
	ldapCalled := false
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		ldapCalled = true
		return &service.LDAPUser{Username: "alice"}, nil
	}

	recorder, _ := performLoginRequest(t, `{"username":"alice","password":"new-password"}`)

	var response struct {
		Success bool `json:"success"`
	}
	if err := common.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if response.Success {
		t.Fatalf("expected disabled user login to fail")
	}
	if ldapCalled {
		t.Fatalf("expected LDAP fallback to be skipped for disabled users")
	}

	storedUser := fetchUserByID(t, user.Id)
	if !common.ValidatePasswordAndHash("old-password", storedUser.Password) {
		t.Fatalf("expected disabled user password hash to remain unchanged")
	}
}

func TestLoginCreatesUserFromLDAPWhenLocalUserDoesNotExist(t *testing.T) {
	_ = setupLoginControllerTestDB(t)

	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	system_setting.GetLDAPSettings().Enabled = true
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		return &service.LDAPUser{
			Username:    "ldap-user",
			DisplayName: "LDAP User",
			Email:       "ldap-user@example.com",
		}, nil
	}

	recorder, _ := performLoginRequest(t, `{"username":"ldap-user","password":"directory-password"}`)

	var response struct {
		Success bool `json:"success"`
		Data    struct {
			User struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := common.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if !response.Success {
		t.Fatalf("expected successful LDAP auto-create login, got body: %s", recorder.Body.String())
	}
	if response.Data.User.Username != "ldap-user" {
		t.Fatalf("expected username ldap-user, got %q", response.Data.User.Username)
	}

	user := fetchUserByID(t, response.Data.User.ID)
	if user.Email != "ldap-user@example.com" {
		t.Fatalf("expected created user email to be saved, got %q", user.Email)
	}
	if !common.ValidatePasswordAndHash("directory-password", user.Password) {
		t.Fatalf("expected created user password to match directory password")
	}
}

func TestLoginCreatesUserFromLDAPWhenRegisterDisabled(t *testing.T) {
	_ = setupLoginControllerTestDB(t)

	originalRegisterEnabled := common.RegisterEnabled
	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		common.RegisterEnabled = originalRegisterEnabled
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	common.RegisterEnabled = false
	system_setting.GetLDAPSettings().Enabled = true
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		return &service.LDAPUser{
			Username:    "ldap-user-no-register",
			DisplayName: "LDAP User",
			Email:       "ldap-user-no-register@example.com",
		}, nil
	}

	recorder, _ := performLoginRequest(t, `{"username":"ldap-user-no-register","password":"directory-password"}`)

	var response struct {
		Success bool `json:"success"`
		Data    struct {
			User struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := common.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if !response.Success {
		t.Fatalf("expected successful LDAP auto-create login, got body: %s", recorder.Body.String())
	}
	if response.Data.User.Username == "" {
		t.Fatalf("expected non-empty username for auto-created LDAP user")
	}
}

func TestAuthenticateLoginUserLDAPFailuresDoNotUpdatePassword(t *testing.T) {
	db := setupLoginControllerTestDB(t)
	user := seedLoginUser(t, db, "alice", "old-password", common.UserStatusEnabled)

	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	system_setting.GetLDAPSettings().Enabled = true
	cases := []struct {
		name    string
		ldapErr error
		wantErr error
	}{
		{name: "search no result", ldapErr: service.ErrLDAPNoResult, wantErr: model.ErrInvalidCredentials},
		{name: "bind failed", ldapErr: service.ErrLDAPInvalidCredentials, wantErr: model.ErrInvalidCredentials},
		{name: "search multiple", ldapErr: service.ErrLDAPMultipleResults, wantErr: model.ErrDatabase},
		{name: "connect failed", ldapErr: service.ErrLDAPConfig, wantErr: model.ErrDatabase},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
				return nil, tc.ldapErr
			}

			_, err := authenticateLoginUser(context.Background(), "alice", "new-password")
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("expected error %v, got %v", tc.wantErr, err)
			}

			storedUser := fetchUserByID(t, user.Id)
			if !common.ValidatePasswordAndHash("old-password", storedUser.Password) {
				t.Fatalf("expected password hash unchanged when LDAP auth fails")
			}
		})
	}
}

func TestLoginLDAPSuccessStillRequiresTwoFA(t *testing.T) {
	db := setupLoginControllerTestDB(t)
	user := seedLoginUser(t, db, "alice", "old-password", common.UserStatusEnabled)
	twoFA := &model.TwoFA{
		UserId:    user.Id,
		Secret:    "test-secret",
		IsEnabled: true,
	}
	if err := db.Create(twoFA).Error; err != nil {
		t.Fatalf("failed to create 2FA record: %v", err)
	}

	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	system_setting.GetLDAPSettings().Enabled = true
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		return &service.LDAPUser{Username: "alice"}, nil
	}

	recorder, _ := performLoginRequest(t, `{"username":"alice","password":"new-password"}`)

	var response struct {
		Success bool `json:"success"`
		Data    struct {
			RequireTwoFA bool `json:"require_2fa"`
		} `json:"data"`
	}
	if err := common.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if !response.Success {
		t.Fatalf("expected success response, got body: %s", recorder.Body.String())
	}
	if !response.Data.RequireTwoFA {
		t.Fatalf("expected LDAP login to still require 2FA")
	}

	storedUser := fetchUserByID(t, user.Id)
	if !common.ValidatePasswordAndHash("new-password", storedUser.Password) {
		t.Fatalf("expected password to be updated before 2FA completion")
	}
}

func TestLoginLDAPSuccessThenVerify2FACompletesLogin(t *testing.T) {
	db := setupLoginControllerTestDB(t)
	user := seedLoginUser(t, db, "alice", "old-password", common.UserStatusEnabled)
	twoFA := &model.TwoFA{
		UserId:    user.Id,
		Secret:    "test-secret",
		IsEnabled: true,
	}
	if err := db.Create(twoFA).Error; err != nil {
		t.Fatalf("failed to create 2FA record: %v", err)
	}
	backupCode := "ABCD1234"
	if err := model.CreateBackupCodes(user.Id, []string{backupCode}); err != nil {
		t.Fatalf("failed to create backup code: %v", err)
	}

	originalAuthenticateLDAP := authenticateLDAP
	originalLDAPSettings := *system_setting.GetLDAPSettings()
	t.Cleanup(func() {
		authenticateLDAP = originalAuthenticateLDAP
		*system_setting.GetLDAPSettings() = originalLDAPSettings
	})

	system_setting.GetLDAPSettings().Enabled = true
	authenticateLDAP = func(ctx context.Context, identifier string, password string) (*service.LDAPUser, error) {
		return &service.LDAPUser{Username: "alice"}, nil
	}

	router := gin.New()
	store := cookie.NewStore([]byte("login-test-secret"))
	router.Use(sessions.Sessions("new-api-session", store))
	router.POST("/api/user/login", Login)
	router.POST("/api/user/2fa/login", Verify2FALogin)

	loginRecorder := httptest.NewRecorder()
	loginRequest := httptest.NewRequest(http.MethodPost, "/api/user/login", bytes.NewBufferString(`{"username":"alice","password":"new-password"}`))
	loginRequest.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(loginRecorder, loginRequest)

	var loginResponse struct {
		Success bool `json:"success"`
		Data    struct {
			RequireTwoFA bool   `json:"require_2fa"`
			FlowToken    string `json:"flow_token"`
		} `json:"data"`
	}
	if err := common.Unmarshal(loginRecorder.Body.Bytes(), &loginResponse); err != nil {
		t.Fatalf("failed to decode login response: %v", err)
	}
	if !loginResponse.Success || !loginResponse.Data.RequireTwoFA {
		t.Fatalf("expected require_2fa response, got body: %s", loginRecorder.Body.String())
	}
	if loginResponse.Data.FlowToken == "" {
		t.Fatalf("expected flow_token in login response")
	}

	verifyRecorder := httptest.NewRecorder()
	verifyBody := fmt.Sprintf(`{"code":"ABCD1234","flow_token":%q}`, loginResponse.Data.FlowToken)
	verifyRequest := httptest.NewRequest(http.MethodPost, "/api/user/2fa/login", bytes.NewBufferString(verifyBody))
	verifyRequest.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(verifyRecorder, verifyRequest)

	var verifyResponse struct {
		Success bool `json:"success"`
		Data    struct {
			User struct {
				ID int `json:"id"`
			} `json:"user"`
		} `json:"data"`
	}
	if err := common.Unmarshal(verifyRecorder.Body.Bytes(), &verifyResponse); err != nil {
		t.Fatalf("failed to decode verify response: %v", err)
	}
	if !verifyResponse.Success {
		t.Fatalf("expected successful 2FA verification, got body: %s", verifyRecorder.Body.String())
	}
	if verifyResponse.Data.User.ID != user.Id {
		t.Fatalf("expected logged in user id %d after 2FA, got %d", user.Id, verifyResponse.Data.User.ID)
	}

	storedUser := fetchUserByID(t, user.Id)
	if !common.ValidatePasswordAndHash("new-password", storedUser.Password) {
		t.Fatalf("expected password to remain updated after 2FA completion")
	}
}
