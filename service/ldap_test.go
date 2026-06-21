package service

import (
	"context"
	"errors"
	"testing"

	"github.com/QuantumNous/new-api/setting/system_setting"

	ldap "github.com/go-ldap/ldap/v3"
)

func withMockLDAPHooks(t *testing.T) {
	t.Helper()

	originalDial := ldapDial
	originalCloseConn := ldapCloseConn
	originalBindSearch := ldapBindSearch
	originalFindUser := ldapFindUser
	originalBindUser := ldapBindUser
	originalSettings := *system_setting.GetLDAPSettings()

	t.Cleanup(func() {
		ldapDial = originalDial
		ldapCloseConn = originalCloseConn
		ldapBindSearch = originalBindSearch
		ldapFindUser = originalFindUser
		ldapBindUser = originalBindUser
		*system_setting.GetLDAPSettings() = originalSettings
	})
}

func TestAuthenticateLDAPDisabledSkipsDial(t *testing.T) {
	withMockLDAPHooks(t)

	settings := system_setting.GetLDAPSettings()
	settings.Enabled = false

	dialCalled := false
	ldapDial = func(settings *system_setting.LDAPSettings) (*ldap.Conn, error) {
		dialCalled = true
		return nil, nil
	}

	_, err := authenticateLDAP(context.Background(), "alice", "password")
	if !errors.Is(err, ErrLDAPDisabled) {
		t.Fatalf("expected ErrLDAPDisabled, got %v", err)
	}
	if dialCalled {
		t.Fatalf("expected LDAP dial not to be called when LDAP is disabled")
	}
}

func TestAuthenticateLDAPValidatesRequiredInputAndConfig(t *testing.T) {
	withMockLDAPHooks(t)

	settings := system_setting.GetLDAPSettings()
	settings.Enabled = true
	settings.ServerURL = ""
	settings.BaseDN = ""

	_, err := authenticateLDAP(context.Background(), "alice", "password")
	if !errors.Is(err, ErrLDAPConfig) {
		t.Fatalf("expected ErrLDAPConfig for missing config, got %v", err)
	}

	settings.ServerURL = "ldap://example.com:389"
	settings.BaseDN = "dc=example,dc=com"
	_, err = authenticateLDAP(context.Background(), "   ", "password")
	if !errors.Is(err, ErrLDAPInvalidCredentials) {
		t.Fatalf("expected ErrLDAPInvalidCredentials for empty identifier, got %v", err)
	}

	_, err = authenticateLDAP(context.Background(), "alice", "")
	if !errors.Is(err, ErrLDAPInvalidCredentials) {
		t.Fatalf("expected ErrLDAPInvalidCredentials for empty password, got %v", err)
	}
}

func TestAuthenticateLDAPSuccessfulFlowPreservesRawPassword(t *testing.T) {
	withMockLDAPHooks(t)

	settings := system_setting.GetLDAPSettings()
	settings.Enabled = true
	settings.ServerURL = "ldap://example.com:389"
	settings.BaseDN = "dc=example,dc=com"
	settings.UsernameAttribute = "uid"
	settings.DisplayNameAttribute = "cn"
	settings.MailAttribute = "mail"

	ldapDial = func(settings *system_setting.LDAPSettings) (*ldap.Conn, error) {
		return nil, nil
	}
	ldapCloseConn = func(conn *ldap.Conn) {}
	ldapBindSearch = func(conn *ldap.Conn, settings *system_setting.LDAPSettings) error {
		return nil
	}
	ldapFindUser = func(conn *ldap.Conn, settings *system_setting.LDAPSettings, identifier string) (*ldap.Entry, error) {
		return &ldap.Entry{
			DN: "uid=alice,dc=example,dc=com",
			Attributes: []*ldap.EntryAttribute{
				{Name: "uid", Values: []string{"alice"}},
				{Name: "cn", Values: []string{"Alice LDAP"}},
				{Name: "mail", Values: []string{"alice@example.com"}},
			},
		}, nil
	}

	var gotPassword string
	ldapBindUser = func(conn *ldap.Conn, dn string, password string) error {
		gotPassword = password
		return nil
	}

	user, err := authenticateLDAP(context.Background(), "alice", " password-with-space ")
	if err != nil {
		t.Fatalf("expected LDAP auth success, got %v", err)
	}
	if gotPassword != " password-with-space " {
		t.Fatalf("expected raw password to be preserved, got %q", gotPassword)
	}
	if user.Username != "alice" || user.DisplayName != "Alice LDAP" || user.Email != "alice@example.com" {
		t.Fatalf("unexpected LDAP user mapping result: %+v", user)
	}
}

func TestAuthenticateLDAPPropagatesSearchAndConnectionFailures(t *testing.T) {
	withMockLDAPHooks(t)

	settings := system_setting.GetLDAPSettings()
	settings.Enabled = true
	settings.ServerURL = "ldap://example.com:389"
	settings.BaseDN = "dc=example,dc=com"

	t.Run("search no result", func(t *testing.T) {
		ldapDial = func(settings *system_setting.LDAPSettings) (*ldap.Conn, error) {
			return nil, nil
		}
		ldapCloseConn = func(conn *ldap.Conn) {}
		ldapBindSearch = func(conn *ldap.Conn, settings *system_setting.LDAPSettings) error {
			return nil
		}
		ldapFindUser = func(conn *ldap.Conn, settings *system_setting.LDAPSettings, identifier string) (*ldap.Entry, error) {
			return nil, ErrLDAPNoResult
		}

		_, err := authenticateLDAP(context.Background(), "alice", "password")
		if !errors.Is(err, ErrLDAPNoResult) {
			t.Fatalf("expected ErrLDAPNoResult, got %v", err)
		}
	})

	t.Run("search multiple result", func(t *testing.T) {
		ldapDial = func(settings *system_setting.LDAPSettings) (*ldap.Conn, error) {
			return nil, nil
		}
		ldapCloseConn = func(conn *ldap.Conn) {}
		ldapBindSearch = func(conn *ldap.Conn, settings *system_setting.LDAPSettings) error {
			return nil
		}
		ldapFindUser = func(conn *ldap.Conn, settings *system_setting.LDAPSettings, identifier string) (*ldap.Entry, error) {
			return nil, ErrLDAPMultipleResults
		}

		_, err := authenticateLDAP(context.Background(), "alice", "password")
		if !errors.Is(err, ErrLDAPMultipleResults) {
			t.Fatalf("expected ErrLDAPMultipleResults, got %v", err)
		}
	})

	t.Run("connection failed", func(t *testing.T) {
		ldapDial = func(settings *system_setting.LDAPSettings) (*ldap.Conn, error) {
			return nil, ErrLDAPConfig
		}

		_, err := authenticateLDAP(context.Background(), "alice", "password")
		if !errors.Is(err, ErrLDAPConfig) {
			t.Fatalf("expected ErrLDAPConfig, got %v", err)
		}
	})
}

func TestAuthenticateLDAPMapsBindInvalidCredentials(t *testing.T) {
	withMockLDAPHooks(t)

	settings := system_setting.GetLDAPSettings()
	settings.Enabled = true
	settings.ServerURL = "ldap://example.com:389"
	settings.BaseDN = "dc=example,dc=com"

	ldapDial = func(settings *system_setting.LDAPSettings) (*ldap.Conn, error) {
		return nil, nil
	}
	ldapCloseConn = func(conn *ldap.Conn) {}
	ldapBindSearch = func(conn *ldap.Conn, settings *system_setting.LDAPSettings) error {
		return nil
	}
	ldapFindUser = func(conn *ldap.Conn, settings *system_setting.LDAPSettings, identifier string) (*ldap.Entry, error) {
		return &ldap.Entry{DN: "uid=alice,dc=example,dc=com"}, nil
	}
	ldapBindUser = func(conn *ldap.Conn, dn string, password string) error {
		return &ldap.Error{ResultCode: ldap.LDAPResultInvalidCredentials, Err: errors.New("invalid credentials")}
	}

	_, err := authenticateLDAP(context.Background(), "alice", "password")
	if !errors.Is(err, ErrLDAPInvalidCredentials) {
		t.Fatalf("expected ErrLDAPInvalidCredentials, got %v", err)
	}
}
