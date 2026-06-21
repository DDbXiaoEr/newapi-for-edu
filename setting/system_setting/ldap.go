package system_setting

import "github.com/QuantumNous/new-api/setting/config"

type LDAPSettings struct {
	Enabled              bool   `json:"enabled"`
	ServerURL            string `json:"server_url"`
	BindDN               string `json:"bind_dn"`
	BindPassword         string `json:"bind_password"`
	BaseDN               string `json:"base_dn"`
	UserFilter           string `json:"user_filter"`
	UsernameAttribute    string `json:"username_attribute"`
	DisplayNameAttribute string `json:"display_name_attribute"`
	MailAttribute        string `json:"mail_attribute"`
	StartTLS             bool   `json:"start_tls"`
	SkipTLSVerify        bool   `json:"skip_tls_verify"`
	TimeoutSeconds       int    `json:"timeout_seconds"`
}

var defaultLDAPSettings = LDAPSettings{
	Enabled:              false,
	UserFilter:           "(|(uid={username})(sAMAccountName={username})(mail={username}))",
	UsernameAttribute:    "uid",
	DisplayNameAttribute: "cn",
	MailAttribute:        "mail",
	TimeoutSeconds:       5,
}

func init() {
	config.GlobalConfig.Register("ldap", &defaultLDAPSettings)
}

func GetLDAPSettings() *LDAPSettings {
	return &defaultLDAPSettings
}
