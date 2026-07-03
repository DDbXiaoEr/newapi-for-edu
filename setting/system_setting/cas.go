package system_setting

import "github.com/QuantumNous/new-api/setting/config"

type CasSettings struct {
	Enabled               bool   `json:"enabled"`
	ServerUrl             string `json:"server_url"`
	ServiceId             string `json:"service_id"`
	UsernameAttribute     string `json:"username_attribute"`
	DisplayNameAttribute  string `json:"display_name_attribute"`
	EmailAttribute        string `json:"email_attribute"`
	AccessAttribute       string `json:"access_attribute"`
	AccessAttributeValue  string `json:"access_attribute_value"`
}

var defaultCasSettings = CasSettings{
	UsernameAttribute:    "uid",
	DisplayNameAttribute: "displayname",
	EmailAttribute:       "mail",
}

func init() {
	config.GlobalConfig.Register("cas", &defaultCasSettings)
}

func GetCasSettings() *CasSettings {
	return &defaultCasSettings
}
