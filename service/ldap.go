package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/setting/system_setting"

	ldap "github.com/go-ldap/ldap/v3"
)

var (
	ErrLDAPDisabled           = errors.New("ldap disabled")
	ErrLDAPInvalidCredentials = errors.New("ldap invalid credentials")
	ErrLDAPNoResult           = errors.New("ldap no result")
	ErrLDAPMultipleResults    = errors.New("ldap multiple results")
	ErrLDAPConfig             = errors.New("ldap config error")
)

type LDAPUser struct {
	DN          string
	Username    string
	DisplayName string
	Email       string
	Attributes  map[string][]string
}

var AuthenticateLDAP = authenticateLDAP

var (
	ldapDial      = dialLDAP
	ldapCloseConn = func(conn *ldap.Conn) {
		if conn != nil {
			conn.Close()
		}
	}
	ldapBindSearch = bindLDAPConnection
	ldapFindUser   = findLDAPEntry
	ldapBindUser   = func(conn *ldap.Conn, dn string, password string) error {
		return conn.Bind(dn, password)
	}
)

func authenticateLDAP(ctx context.Context, identifier string, password string) (*LDAPUser, error) {
	_ = ctx

	settings := system_setting.GetLDAPSettings()
	if settings == nil || !settings.Enabled {
		return nil, ErrLDAPDisabled
	}

	identifier = strings.TrimSpace(identifier)
	if identifier == "" || password == "" {
		return nil, ErrLDAPInvalidCredentials
	}
	if strings.TrimSpace(settings.ServerURL) == "" || strings.TrimSpace(settings.BaseDN) == "" {
		return nil, ErrLDAPConfig
	}

	conn, err := ldapDial(settings)
	if err != nil {
		return nil, err
	}
	defer ldapCloseConn(conn)

	if err := ldapBindSearch(conn, settings); err != nil {
		return nil, err
	}

	entry, err := ldapFindUser(conn, settings, identifier)
	if err != nil {
		return nil, err
	}

	if err := ldapBindUser(conn, entry.DN, password); err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials) {
			return nil, ErrLDAPInvalidCredentials
		}
		return nil, fmt.Errorf("%w: %v", ErrLDAPConfig, err)
	}

	return entryToLDAPUser(entry, settings, identifier), nil
}

func dialLDAP(settings *system_setting.LDAPSettings) (*ldap.Conn, error) {
	timeout := time.Duration(settings.TimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 5 * time.Second
	}

	parsedURL, err := url.Parse(strings.TrimSpace(settings.ServerURL))
	if err != nil {
		return nil, fmt.Errorf("%w: invalid server url: %v", ErrLDAPConfig, err)
	}

	dialer := &net.Dialer{Timeout: timeout}
	opts := []ldap.DialOpt{ldap.DialWithDialer(dialer)}

	if parsedURL.Scheme == "ldaps" {
		tlsConfig := &tls.Config{InsecureSkipVerify: settings.SkipTLSVerify}
		if host := parsedURL.Hostname(); host != "" {
			tlsConfig.ServerName = host
		}
		opts = append(opts, ldap.DialWithTLSConfig(tlsConfig))
	}

	conn, err := ldap.DialURL(strings.TrimSpace(settings.ServerURL), opts...)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrLDAPConfig, err)
	}

	if parsedURL.Scheme == "ldap" && settings.StartTLS {
		tlsConfig := &tls.Config{InsecureSkipVerify: settings.SkipTLSVerify}
		if host := parsedURL.Hostname(); host != "" {
			tlsConfig.ServerName = host
		}
		if err := conn.StartTLS(tlsConfig); err != nil {
			conn.Close()
			return nil, fmt.Errorf("%w: %v", ErrLDAPConfig, err)
		}
	}

	return conn, nil
}

func bindLDAPConnection(conn *ldap.Conn, settings *system_setting.LDAPSettings) error {
	if strings.TrimSpace(settings.BindDN) == "" {
		return nil
	}
	if err := conn.Bind(settings.BindDN, settings.BindPassword); err != nil {
		if ldap.IsErrorWithCode(err, ldap.LDAPResultInvalidCredentials) {
			return ErrLDAPInvalidCredentials
		}
		return fmt.Errorf("%w: %v", ErrLDAPConfig, err)
	}
	return nil
}

func findLDAPEntry(conn *ldap.Conn, settings *system_setting.LDAPSettings, identifier string) (*ldap.Entry, error) {
	filter := strings.TrimSpace(settings.UserFilter)
	if filter == "" {
		filter = "(|(uid={username})(sAMAccountName={username})(mail={username}))"
	}
	filter = strings.ReplaceAll(filter, "{username}", ldap.EscapeFilter(identifier))

	attributes := uniqueStrings([]string{
		strings.TrimSpace(settings.UsernameAttribute),
		strings.TrimSpace(settings.DisplayNameAttribute),
		strings.TrimSpace(settings.MailAttribute),
		"uid",
		"sAMAccountName",
		"cn",
		"mail",
	})
	for _, rule := range settings.GroupAssignmentRules {
		attributes = uniqueStrings(append(attributes, rule.Attribute))
	}

	timeout := time.Duration(settings.TimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 5 * time.Second
	}

	searchRequest := ldap.NewSearchRequest(
		settings.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		2,
		int(timeout.Seconds()),
		false,
		filter,
		attributes,
		nil,
	)

	result, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrLDAPConfig, err)
	}
	if len(result.Entries) == 0 {
		return nil, ErrLDAPNoResult
	}
	if len(result.Entries) > 1 {
		return nil, ErrLDAPMultipleResults
	}

	return result.Entries[0], nil
}

func entryToLDAPUser(entry *ldap.Entry, settings *system_setting.LDAPSettings, fallback string) *LDAPUser {
	username := firstNonEmpty(
		entry.GetAttributeValue(settings.UsernameAttribute),
		entry.GetAttributeValue("uid"),
		entry.GetAttributeValue("sAMAccountName"),
		entry.GetAttributeValue("mail"),
		fallback,
	)
	displayName := firstNonEmpty(
		entry.GetAttributeValue(settings.DisplayNameAttribute),
		entry.GetAttributeValue("cn"),
		username,
	)
	email := firstNonEmpty(
		entry.GetAttributeValue(settings.MailAttribute),
		entry.GetAttributeValue("mail"),
	)

	return &LDAPUser{
		DN:          entry.DN,
		Username:    strings.TrimSpace(username),
		DisplayName: strings.TrimSpace(displayName),
		Email:       strings.TrimSpace(email),
		Attributes:  ldapEntryAttributes(entry),
	}
}

func ldapEntryAttributes(entry *ldap.Entry) map[string][]string {
	if entry == nil {
		return nil
	}
	attributes := make(map[string][]string, len(entry.Attributes)+1)
	if dn := strings.TrimSpace(entry.DN); dn != "" {
		attributes["dn"] = []string{dn}
	}
	for _, attr := range entry.Attributes {
		name := strings.TrimSpace(attr.Name)
		if name == "" {
			continue
		}
		key := strings.ToLower(name)
		for _, value := range attr.Values {
			value = strings.TrimSpace(value)
			if value == "" {
				continue
			}
			attributes[key] = append(attributes[key], value)
		}
	}
	return attributes
}

func uniqueStrings(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}
