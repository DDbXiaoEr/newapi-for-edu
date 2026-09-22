package service

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/google/uuid"
)

var (
	ErrLoginSessionInvalid  = errors.New("login session is invalid")
	ErrLoginSessionRevoked  = errors.New("login session is revoked")
	ErrLoginSessionMismatch = errors.New("login session does not match the expected session")
)

type LoginSessionView struct {
	SID          string `json:"sid"`
	Current      bool   `json:"current"`
	LoginMethod  string `json:"login_method"`
	IP           string `json:"ip"`
	UserAgent    string `json:"user_agent"`
	CreatedAt    int64  `json:"created_at"`
	LastActiveAt int64  `json:"last_active_at"`
	ExpiresAt    int64  `json:"expires_at"`
}

type AuthBundle struct {
	AccessToken     string           `json:"access_token"`
	TokenType       string           `json:"token_type"`
	AccessExpiresAt int64            `json:"access_expires_at"`
	Session         LoginSessionView `json:"session"`
}

func CreateLoginSession(userID int, loginMethod, ip, userAgent string) (*AuthBundle, error) {
	return createLoginSession(userID, 0, loginMethod, ip, userAgent)
}

func CreateLoginSessionAtAuthVersion(userID int, expectedAuthVersion int64, loginMethod, ip, userAgent string) (*AuthBundle, error) {
	if expectedAuthVersion <= 0 {
		return nil, ErrLoginSessionInvalid
	}
	return createLoginSession(userID, expectedAuthVersion, loginMethod, ip, userAgent)
}

func createLoginSession(userID int, expectedAuthVersion int64, loginMethod, ip, userAgent string) (*AuthBundle, error) {
	user, err := model.GetUserCache(userID)
	if err != nil {
		return nil, err
	}
	if user.Status != common.UserStatusEnabled || user.AuthVersion <= 0 {
		return nil, ErrLoginSessionInvalid
	}
	if expectedAuthVersion > 0 && user.AuthVersion != expectedAuthVersion {
		return nil, ErrLoginSessionRevoked
	}
	now := time.Now().Unix()
	activeCount, err := model.CountActiveUserSessions(userID, now)
	if err != nil {
		return nil, err
	}
	if activeCount >= int64(common.UserSessionActiveLimit) {
		return nil, model.ErrUserSessionLimit
	}
	issuanceCount, err := model.CountUserSessionsCreatedSince(userID, now-common.UserSessionIssuanceWindowSeconds)
	if err != nil {
		return nil, err
	}
	if issuanceCount >= int64(common.UserSessionIssuanceLimit) {
		return nil, model.ErrUserSessionIssuanceLimit
	}
	session, err := newLoginSession(userID, user.AuthVersion, loginMethod, ip, userAgent)
	if err != nil {
		return nil, err
	}
	if err := model.CreateUserSession(session); err != nil {
		return nil, err
	}
	bundle, err := issueAuthBundle(session, true)
	if err != nil {
		_, _ = model.RevokeUserSession(userID, session.SID, "token_issue_failed")
		return nil, err
	}
	return bundle, nil
}

func newLoginSession(userID int, authVersion int64, loginMethod, ip, userAgent string) (*model.UserSession, error) {
	// The refresh flow was removed, but user_sessions still carries a
	// non-empty refresh_hash column; store an unusable random digest so an old
	// deleted code path can never accept a guessed secret.
	refreshSecret, err := common.GenerateRandomCharsKey(64)
	if err != nil {
		return nil, err
	}
	now := time.Now().Unix()
	session := &model.UserSession{
		SID:             uuid.NewString(),
		UserID:          userID,
		Version:         1,
		UserAuthVersion: authVersion,
		Status:          model.UserSessionStatusActive,
		RefreshHash:     hashRefreshSecret(refreshSecret),
		LoginMethod:     strings.TrimSpace(loginMethod),
		IP:              truncateAuthMetadata(ip, 64),
		UserAgent:       truncateAuthMetadata(userAgent, 512),
		CreatedAt:       now,
		LastActiveAt:    now,
		ExpiresAt:       time.Unix(now, 0).Add(LoginSessionTTL).Unix(),
	}
	if session.LoginMethod == "" {
		session.LoginMethod = "unknown"
	}
	return session, nil
}

func ValidateLoginSession(identity AuthIdentity) (*model.UserSession, *model.UserBase, error) {
	session, err := model.GetUserSessionCached(identity.SessionID)
	if err != nil {
		if errors.Is(err, model.ErrUserSessionInactive) {
			return nil, nil, ErrLoginSessionRevoked
		}
		return nil, nil, err
	}
	now := time.Now().Unix()
	if session.UserID != identity.UserID || session.Status != model.UserSessionStatusActive || session.RevokedAt != 0 || session.ExpiresAt <= now || session.Version != identity.SessionVersion || session.UserAuthVersion != identity.UserAuthVersion {
		return nil, nil, ErrLoginSessionRevoked
	}
	user, err := model.GetUserCache(identity.UserID)
	if err != nil {
		return nil, nil, err
	}
	if user.Status != common.UserStatusEnabled || user.AuthVersion != identity.UserAuthVersion {
		return nil, nil, ErrLoginSessionRevoked
	}
	return session, user, nil
}

// ValidateSessionReference validates a server-side flow bound to an existing
// dashboard session without requiring an access token on the callback request.
func ValidateSessionReference(userID int, sid string) (AuthIdentity, error) {
	if userID <= 0 || strings.TrimSpace(sid) == "" {
		return AuthIdentity{}, ErrLoginSessionInvalid
	}
	session, err := model.GetUserSessionCached(sid)
	if err != nil {
		return AuthIdentity{}, err
	}
	identity := AuthIdentity{
		UserID:          userID,
		SessionID:       sid,
		UserAuthVersion: session.UserAuthVersion,
		SessionVersion:  session.Version,
	}
	if _, _, err := ValidateLoginSession(identity); err != nil {
		return AuthIdentity{}, err
	}
	return identity, nil
}

// AdvanceCurrentSessionSecurity increments the user's global auth version,
// preserves only the current browser session at a new session version and
// returns a replacement access token. Call after a successful 2FA/passkey
// security-setting mutation that did not already advance AuthVersion.
func AdvanceCurrentSessionSecurity(identity AuthIdentity, reason string) (*AuthBundle, error) {
	nextUserAuthVersion, err := model.BumpUserAuthVersion(identity.UserID)
	if err != nil {
		return nil, err
	}
	return advanceCurrentSessionToVersion(identity, nextUserAuthVersion, reason)
}

// AdvanceCurrentSessionToUserVersion is used when the security mutation and
// AuthVersion increment were committed in the same transaction (for example,
// a password change).
func AdvanceCurrentSessionToUserVersion(identity AuthIdentity, reason string) (*AuthBundle, error) {
	user, err := model.GetUserCache(identity.UserID)
	if err != nil {
		return nil, err
	}
	if user.Status != common.UserStatusEnabled || user.AuthVersion <= identity.UserAuthVersion {
		return nil, ErrLoginSessionRevoked
	}
	return advanceCurrentSessionToVersion(identity, user.AuthVersion, reason)
}

func advanceCurrentSessionToVersion(identity AuthIdentity, nextUserAuthVersion int64, reason string) (*AuthBundle, error) {
	session, err := model.AdvanceUserSessionAuthVersion(
		identity.UserID,
		identity.SessionID,
		identity.SessionVersion,
		identity.UserAuthVersion,
		nextUserAuthVersion,
	)
	if err != nil {
		return nil, err
	}
	if _, err := model.RevokeOtherUserSessions(identity.UserID, identity.SessionID, reason); err != nil {
		return nil, err
	}
	return issueAuthBundle(session, true)
}

func ListLoginSessions(userID int, currentSID string) ([]LoginSessionView, error) {
	sessions, err := model.ListActiveUserSessions(userID, currentSID, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	views := make([]LoginSessionView, 0, len(sessions))
	for i := range sessions {
		views = append(views, sessionView(&sessions[i], sessions[i].SID == currentSID))
	}
	return views, nil
}

func issueAuthBundle(session *model.UserSession, current bool) (*AuthBundle, error) {
	identity := AuthIdentity{
		UserID:          session.UserID,
		SessionID:       session.SID,
		UserAuthVersion: session.UserAuthVersion,
		SessionVersion:  session.Version,
	}
	accessToken, accessExpiresAt, err := IssueAccessToken(identity)
	if err != nil {
		return nil, err
	}
	return &AuthBundle{
		AccessToken:     accessToken,
		TokenType:       "Bearer",
		AccessExpiresAt: accessExpiresAt,
		Session:         sessionView(session, current),
	}, nil
}

func sessionView(session *model.UserSession, current bool) LoginSessionView {
	return LoginSessionView{
		SID:          session.SID,
		Current:      current,
		LoginMethod:  session.LoginMethod,
		IP:           session.IP,
		UserAgent:    session.UserAgent,
		CreatedAt:    session.CreatedAt,
		LastActiveAt: session.LastActiveAt,
		ExpiresAt:    session.ExpiresAt,
	}
}

func hashRefreshSecret(secret string) string {
	return common.GenerateHMACWithKey(authSigningKey("refresh"), secret)
}

func truncateAuthMetadata(value string, max int) string {
	value = strings.TrimSpace(value)
	if len(value) <= max {
		return value
	}
	return value[:max]
}

func authSessionErrorCode(err error) (int, string) {
	switch {
	case errors.Is(err, model.ErrUserSessionLimit):
		return http.StatusConflict, "AUTH_SESSION_LIMIT"
	case errors.Is(err, model.ErrUserSessionIssuanceLimit):
		return http.StatusTooManyRequests, "AUTH_SESSION_ISSUANCE_LIMIT"
	case errors.Is(err, ErrLoginSessionMismatch):
		return http.StatusConflict, "AUTH_SESSION_MISMATCH"
	case errors.Is(err, ErrAuthTokenExpired):
		return http.StatusUnauthorized, "AUTH_TOKEN_EXPIRED"
	case errors.Is(err, ErrLoginSessionRevoked):
		return http.StatusUnauthorized, "AUTH_SESSION_REVOKED"
	case errors.Is(err, ErrAuthTokenInvalid):
		return http.StatusUnauthorized, "AUTH_UNAUTHORIZED"
	default:
		return http.StatusInternalServerError, "AUTH_INTERNAL_ERROR"
	}
}

func AuthSessionErrorCode(err error) (int, string) {
	return authSessionErrorCode(err)
}

func FormatAuthError(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("authentication failed: %v", err)
}
