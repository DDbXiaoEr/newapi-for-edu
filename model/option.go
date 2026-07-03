package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/setting"
	"github.com/QuantumNous/new-api/setting/config"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/QuantumNous/new-api/setting/performance_setting"
	"github.com/QuantumNous/new-api/setting/ratio_setting"
	"github.com/QuantumNous/new-api/setting/system_setting"
)

type Option struct {
	Key   string `json:"key" gorm:"primaryKey"`
	Value string `json:"value"`
}

func AllOption() ([]*Option, error) {
	var options []*Option
	var err error
	err = DB.Find(&options).Error
	return options, err
}

func InitOptionMap() {
	common.OptionMapRWMutex.Lock()
	common.OptionMap = make(map[string]string)

	// 添加原有的系统配置
	common.OptionMap["FileUploadPermission"] = strconv.Itoa(common.FileUploadPermission)
	common.OptionMap["FileDownloadPermission"] = strconv.Itoa(common.FileDownloadPermission)
	common.OptionMap["ImageUploadPermission"] = strconv.Itoa(common.ImageUploadPermission)
	common.OptionMap["ImageDownloadPermission"] = strconv.Itoa(common.ImageDownloadPermission)
	common.OptionMap["PasswordLoginEnabled"] = strconv.FormatBool(common.PasswordLoginEnabled)
	common.OptionMap["PasswordRegisterEnabled"] = strconv.FormatBool(common.PasswordRegisterEnabled)
	common.OptionMap["EmailVerificationEnabled"] = strconv.FormatBool(common.EmailVerificationEnabled)
	common.OptionMap["GitHubOAuthEnabled"] = strconv.FormatBool(common.GitHubOAuthEnabled)
	common.OptionMap["LinuxDOOAuthEnabled"] = strconv.FormatBool(common.LinuxDOOAuthEnabled)
	common.OptionMap["TelegramOAuthEnabled"] = strconv.FormatBool(common.TelegramOAuthEnabled)
	common.OptionMap["WeChatAuthEnabled"] = strconv.FormatBool(common.WeChatAuthEnabled)
	common.OptionMap["TurnstileCheckEnabled"] = strconv.FormatBool(common.TurnstileCheckEnabled)
	common.OptionMap["RegisterEnabled"] = strconv.FormatBool(common.RegisterEnabled)
	common.OptionMap["AutomaticDisableChannelEnabled"] = strconv.FormatBool(common.AutomaticDisableChannelEnabled)
	common.OptionMap["AutomaticEnableChannelEnabled"] = strconv.FormatBool(common.AutomaticEnableChannelEnabled)
	common.OptionMap["LogConsumeEnabled"] = strconv.FormatBool(common.LogConsumeEnabled)
	common.OptionMap["SyslogEnabled"] = strconv.FormatBool(common.SyslogEnabled)
	common.OptionMap["SyslogNetwork"] = common.SyslogNetwork
	common.OptionMap["SyslogAddr"] = common.SyslogAddr
	common.OptionMap["SyslogTag"] = common.SyslogTag
	common.OptionMap["DisplayInCurrencyEnabled"] = strconv.FormatBool(common.DisplayInCurrencyEnabled)
	common.OptionMap["DisplayTokenStatEnabled"] = strconv.FormatBool(common.DisplayTokenStatEnabled)
	common.OptionMap["DrawingEnabled"] = strconv.FormatBool(common.DrawingEnabled)
	common.OptionMap["TaskEnabled"] = strconv.FormatBool(common.TaskEnabled)
	common.OptionMap["DataExportEnabled"] = strconv.FormatBool(common.DataExportEnabled)
	common.OptionMap["ChannelDisableThreshold"] = strconv.FormatFloat(common.ChannelDisableThreshold, 'f', -1, 64)
	common.OptionMap["EmailDomainRestrictionEnabled"] = strconv.FormatBool(common.EmailDomainRestrictionEnabled)
	common.OptionMap["EmailAliasRestrictionEnabled"] = strconv.FormatBool(common.EmailAliasRestrictionEnabled)
	common.OptionMap["EmailDomainWhitelist"] = strings.Join(common.EmailDomainWhitelist, ",")
	common.OptionMap["SMTPServer"] = ""
	common.OptionMap["SMTPFrom"] = ""
	common.OptionMap["SMTPPort"] = strconv.Itoa(common.SMTPPort)
	common.OptionMap["SMTPAccount"] = ""
	common.OptionMap["SMTPToken"] = ""
	common.OptionMap["SMTPSSLEnabled"] = strconv.FormatBool(common.SMTPSSLEnabled)
	common.OptionMap["SMTPForceAuthLogin"] = strconv.FormatBool(common.SMTPForceAuthLogin)
	common.OptionMap["Notice"] = ""
	common.OptionMap["About"] = ""
	common.OptionMap["HomePageContent"] = ""
	common.OptionMap["Footer"] = common.Footer
	common.OptionMap["SystemName"] = common.SystemName
	common.OptionMap["Logo"] = common.Logo
	common.OptionMap["ServerAddress"] = ""
	common.OptionMap["WorkerUrl"] = system_setting.WorkerUrl
	common.OptionMap["WorkerValidKey"] = system_setting.WorkerValidKey
	common.OptionMap["WorkerAllowHttpImageRequestEnabled"] = strconv.FormatBool(system_setting.WorkerAllowHttpImageRequestEnabled)
	common.OptionMap["PayAddress"] = ""
	common.OptionMap["CustomCallbackAddress"] = ""
	common.OptionMap["EpayId"] = ""
	common.OptionMap["EpayKey"] = ""
	common.OptionMap["Price"] = strconv.FormatFloat(operation_setting.Price, 'f', -1, 64)
	common.OptionMap["USDExchangeRate"] = strconv.FormatFloat(operation_setting.USDExchangeRate, 'f', -1, 64)
	common.OptionMap["MinTopUp"] = strconv.Itoa(operation_setting.MinTopUp)
	common.OptionMap["StripeMinTopUp"] = strconv.Itoa(setting.StripeMinTopUp)
	common.OptionMap["StripeApiSecret"] = setting.StripeApiSecret
	common.OptionMap["StripeWebhookSecret"] = setting.StripeWebhookSecret
	common.OptionMap["StripePriceId"] = setting.StripePriceId
	common.OptionMap["StripeUnitPrice"] = strconv.FormatFloat(setting.StripeUnitPrice, 'f', -1, 64)
	common.OptionMap["StripePromotionCodesEnabled"] = strconv.FormatBool(setting.StripePromotionCodesEnabled)
	common.OptionMap["CreemApiKey"] = setting.CreemApiKey
	common.OptionMap["CreemProducts"] = setting.CreemProducts
	common.OptionMap["CreemTestMode"] = strconv.FormatBool(setting.CreemTestMode)
	common.OptionMap["CreemWebhookSecret"] = setting.CreemWebhookSecret
	common.OptionMap["WaffoEnabled"] = strconv.FormatBool(setting.WaffoEnabled)
	common.OptionMap["WaffoApiKey"] = setting.WaffoApiKey
	common.OptionMap["WaffoPrivateKey"] = setting.WaffoPrivateKey
	common.OptionMap["WaffoPublicCert"] = setting.WaffoPublicCert
	common.OptionMap["WaffoSandboxPublicCert"] = setting.WaffoSandboxPublicCert
	common.OptionMap["WaffoSandboxApiKey"] = setting.WaffoSandboxApiKey
	common.OptionMap["WaffoSandboxPrivateKey"] = setting.WaffoSandboxPrivateKey
	common.OptionMap["WaffoSandbox"] = strconv.FormatBool(setting.WaffoSandbox)
	common.OptionMap["WaffoMerchantId"] = setting.WaffoMerchantId
	common.OptionMap["WaffoNotifyUrl"] = setting.WaffoNotifyUrl
	common.OptionMap["WaffoReturnUrl"] = setting.WaffoReturnUrl
	common.OptionMap["WaffoSubscriptionReturnUrl"] = setting.WaffoSubscriptionReturnUrl
	common.OptionMap["WaffoCurrency"] = setting.WaffoCurrency
	common.OptionMap["WaffoUnitPrice"] = strconv.FormatFloat(setting.WaffoUnitPrice, 'f', -1, 64)
	common.OptionMap["WaffoMinTopUp"] = strconv.Itoa(setting.WaffoMinTopUp)
	common.OptionMap["WaffoPayMethods"] = setting.WaffoPayMethods2JsonString()
	common.OptionMap["WaffoPancakeEnabled"] = strconv.FormatBool(setting.WaffoPancakeEnabled)
	common.OptionMap["WaffoPancakeSandbox"] = strconv.FormatBool(setting.WaffoPancakeSandbox)
	common.OptionMap["WaffoPancakeMerchantID"] = setting.WaffoPancakeMerchantID
	common.OptionMap["WaffoPancakePrivateKey"] = setting.WaffoPancakePrivateKey
	common.OptionMap["WaffoPancakeWebhookPublicKey"] = setting.WaffoPancakeWebhookPublicKey
	common.OptionMap["WaffoPancakeWebhookTestKey"] = setting.WaffoPancakeWebhookTestKey
	common.OptionMap["WaffoPancakeStoreID"] = setting.WaffoPancakeStoreID
	common.OptionMap["WaffoPancakeProductID"] = setting.WaffoPancakeProductID
	common.OptionMap["WaffoPancakeReturnURL"] = setting.WaffoPancakeReturnURL
	common.OptionMap["WaffoPancakeCurrency"] = setting.WaffoPancakeCurrency
	common.OptionMap["WaffoPancakeUnitPrice"] = strconv.FormatFloat(setting.WaffoPancakeUnitPrice, 'f', -1, 64)
	common.OptionMap["WaffoPancakeMinTopUp"] = strconv.Itoa(setting.WaffoPancakeMinTopUp)
	common.OptionMap["TopupGroupRatio"] = common.TopupGroupRatio2JSONString()
	common.OptionMap["Chats"] = setting.Chats2JsonString()
	common.OptionMap["AutoGroups"] = setting.AutoGroups2JsonString()
	common.OptionMap["DefaultUseAutoGroup"] = strconv.FormatBool(setting.DefaultUseAutoGroup)
	common.OptionMap["PayMethods"] = operation_setting.PayMethods2JsonString()
	common.OptionMap["GitHubClientId"] = ""
	common.OptionMap["GitHubClientSecret"] = ""
	common.OptionMap["TelegramBotToken"] = ""
	common.OptionMap["TelegramBotName"] = ""
	common.OptionMap["WeChatServerAddress"] = ""
	common.OptionMap["WeChatServerToken"] = ""
	common.OptionMap["WeChatAccountQRCodeImageURL"] = ""
	common.OptionMap["TurnstileSiteKey"] = ""
	common.OptionMap["TurnstileSecretKey"] = ""
	common.OptionMap["QuotaForNewUser"] = strconv.Itoa(common.QuotaForNewUser)
	common.OptionMap["QuotaForInviter"] = strconv.Itoa(common.QuotaForInviter)
	common.OptionMap["QuotaForInvitee"] = strconv.Itoa(common.QuotaForInvitee)
	common.OptionMap["QuotaRemindThreshold"] = strconv.Itoa(common.QuotaRemindThreshold)
	common.OptionMap["PreConsumedQuota"] = strconv.Itoa(common.PreConsumedQuota)
	common.OptionMap["ModelRequestRateLimitCount"] = strconv.Itoa(setting.ModelRequestRateLimitCount)
	common.OptionMap["ModelRequestRateLimitDurationMinutes"] = strconv.Itoa(setting.ModelRequestRateLimitDurationMinutes)
	common.OptionMap["ModelRequestRateLimitSuccessCount"] = strconv.Itoa(setting.ModelRequestRateLimitSuccessCount)
	common.OptionMap["ModelRequestRateLimitGroup"] = setting.ModelRequestRateLimitGroup2JSONString()
	common.OptionMap["ModelRatio"] = ratio_setting.ModelRatio2JSONString()
	common.OptionMap["ModelPrice"] = ratio_setting.ModelPrice2JSONString()
	common.OptionMap["CacheRatio"] = ratio_setting.CacheRatio2JSONString()
	common.OptionMap["CreateCacheRatio"] = ratio_setting.CreateCacheRatio2JSONString()
	common.OptionMap["GroupRatio"] = ratio_setting.GroupRatio2JSONString()
	common.OptionMap["GroupGroupRatio"] = ratio_setting.GroupGroupRatio2JSONString()
	common.OptionMap["UserUsableGroups"] = setting.UserUsableGroups2JSONString()
	common.OptionMap["CompletionRatio"] = ratio_setting.CompletionRatio2JSONString()
	common.OptionMap["ImageRatio"] = ratio_setting.ImageRatio2JSONString()
	common.OptionMap["AudioRatio"] = ratio_setting.AudioRatio2JSONString()
	common.OptionMap["AudioCompletionRatio"] = ratio_setting.AudioCompletionRatio2JSONString()
	common.OptionMap["TopUpLink"] = common.TopUpLink
	//common.OptionMap["ChatLink"] = common.ChatLink
	//common.OptionMap["ChatLink2"] = common.ChatLink2
	common.OptionMap["QuotaPerUnit"] = strconv.FormatFloat(common.QuotaPerUnit, 'f', -1, 64)
	common.OptionMap["RetryTimes"] = strconv.Itoa(common.RetryTimes)
	common.OptionMap["DataExportInterval"] = strconv.Itoa(common.DataExportInterval)
	common.OptionMap["DataExportDefaultTime"] = common.DataExportDefaultTime
	common.OptionMap["DefaultCollapseSidebar"] = strconv.FormatBool(common.DefaultCollapseSidebar)
	common.OptionMap["MjNotifyEnabled"] = strconv.FormatBool(setting.MjNotifyEnabled)
	common.OptionMap["MjAccountFilterEnabled"] = strconv.FormatBool(setting.MjAccountFilterEnabled)
	common.OptionMap["MjModeClearEnabled"] = strconv.FormatBool(setting.MjModeClearEnabled)
	common.OptionMap["MjForwardUrlEnabled"] = strconv.FormatBool(setting.MjForwardUrlEnabled)
	common.OptionMap["MjActionCheckSuccessEnabled"] = strconv.FormatBool(setting.MjActionCheckSuccessEnabled)
	common.OptionMap["CheckSensitiveEnabled"] = strconv.FormatBool(setting.CheckSensitiveEnabled)
	common.OptionMap["DemoSiteEnabled"] = strconv.FormatBool(operation_setting.DemoSiteEnabled)
	common.OptionMap["SelfUseModeEnabled"] = strconv.FormatBool(operation_setting.SelfUseModeEnabled)
	common.OptionMap["ModelRequestRateLimitEnabled"] = strconv.FormatBool(setting.ModelRequestRateLimitEnabled)
	common.OptionMap["CheckSensitiveOnPromptEnabled"] = strconv.FormatBool(setting.CheckSensitiveOnPromptEnabled)
	common.OptionMap["StopOnSensitiveEnabled"] = strconv.FormatBool(setting.StopOnSensitiveEnabled)
	common.OptionMap["SensitiveWords"] = setting.SensitiveWordsToString()
	common.OptionMap["StreamCacheQueueLength"] = strconv.Itoa(setting.StreamCacheQueueLength)
	common.OptionMap["AutomaticDisableKeywords"] = operation_setting.AutomaticDisableKeywordsToString()
	common.OptionMap["AutomaticDisableStatusCodes"] = operation_setting.AutomaticDisableStatusCodesToString()
	common.OptionMap["AutomaticRetryStatusCodes"] = operation_setting.AutomaticRetryStatusCodesToString()
	common.OptionMap["ExposeRatioEnabled"] = strconv.FormatBool(ratio_setting.IsExposeRatioEnabled())

	// 自动添加所有注册的模型配置
	modelConfigs := config.GlobalConfig.ExportAllConfigs()
	for k, v := range modelConfigs {
		common.OptionMap[k] = v
	}

	common.OptionMapRWMutex.Unlock()
	loadOptionsFromDatabase()
}

func loadOptionsFromDatabase() {
	options, _ := AllOption()
	for _, option := range options {
		err := updateOptionMap(option.Key, option.Value)
		if err != nil {
			common.SysLog("failed to update option map: " + err.Error())
		}
	}
}

func SyncOptions(frequency int) {
	for {
		time.Sleep(time.Duration(frequency) * time.Second)
		common.SysLog("syncing options from database")
		loadOptionsFromDatabase()
	}
}

func UpdateOption(key string, value string) error {
	// Save to database first
	option := Option{
		Key: key,
	}
	// https://gorm.io/docs/update.html#Save-All-Fields
	DB.FirstOrCreate(&option, Option{Key: key})
	option.Value = value
	// Save is a combination function.
	// If save value does not contain primary key, it will execute Create,
	// otherwise it will execute Update (with all fields).
	DB.Save(&option)
	// Update OptionMap
	return updateOptionMap(key, value)
}

type optionHandler func(key, value string) error

var optionHandlers = map[string]optionHandler{
	// Permission (int)
	"FileUploadPermission":    func(_, v string) error { val, _ := strconv.Atoi(v); common.FileUploadPermission = val; return nil },
	"FileDownloadPermission":  func(_, v string) error { val, _ := strconv.Atoi(v); common.FileDownloadPermission = val; return nil },
	"ImageUploadPermission":   func(_, v string) error { val, _ := strconv.Atoi(v); common.ImageUploadPermission = val; return nil },
	"ImageDownloadPermission": func(_, v string) error { val, _ := strconv.Atoi(v); common.ImageDownloadPermission = val; return nil },

	// Bool / Enabled flags
	"PasswordRegisterEnabled":         func(_, v string) error { common.PasswordRegisterEnabled = v == "true"; return nil },
	"PasswordLoginEnabled":            func(_, v string) error { common.PasswordLoginEnabled = v == "true"; return nil },
	"EmailVerificationEnabled":        func(_, v string) error { common.EmailVerificationEnabled = v == "true"; return nil },
	"GitHubOAuthEnabled":              func(_, v string) error { common.GitHubOAuthEnabled = v == "true"; return nil },
	"LinuxDOOAuthEnabled":             func(_, v string) error { common.LinuxDOOAuthEnabled = v == "true"; return nil },
	"WeChatAuthEnabled":               func(_, v string) error { common.WeChatAuthEnabled = v == "true"; return nil },
	"TelegramOAuthEnabled":            func(_, v string) error { common.TelegramOAuthEnabled = v == "true"; return nil },
	"TurnstileCheckEnabled":           func(_, v string) error { common.TurnstileCheckEnabled = v == "true"; return nil },
	"RegisterEnabled":                 func(_, v string) error { common.RegisterEnabled = v == "true"; return nil },
	"EmailDomainRestrictionEnabled":   func(_, v string) error { common.EmailDomainRestrictionEnabled = v == "true"; return nil },
	"EmailAliasRestrictionEnabled":    func(_, v string) error { common.EmailAliasRestrictionEnabled = v == "true"; return nil },
	"AutomaticDisableChannelEnabled":  func(_, v string) error { common.AutomaticDisableChannelEnabled = v == "true"; return nil },
	"AutomaticEnableChannelEnabled":   func(_, v string) error { common.AutomaticEnableChannelEnabled = v == "true"; return nil },
	"LogConsumeEnabled":               func(_, v string) error { common.LogConsumeEnabled = v == "true"; return nil },
	"DisplayTokenStatEnabled":         func(_, v string) error { common.DisplayTokenStatEnabled = v == "true"; return nil },
	"DrawingEnabled":                  func(_, v string) error { common.DrawingEnabled = v == "true"; return nil },
	"TaskEnabled":                     func(_, v string) error { common.TaskEnabled = v == "true"; return nil },
	"DataExportEnabled":               func(_, v string) error { common.DataExportEnabled = v == "true"; return nil },
	"DefaultCollapseSidebar":          func(_, v string) error { common.DefaultCollapseSidebar = v == "true"; return nil },
	"MjNotifyEnabled":                 func(_, v string) error { setting.MjNotifyEnabled = v == "true"; return nil },
	"MjAccountFilterEnabled":          func(_, v string) error { setting.MjAccountFilterEnabled = v == "true"; return nil },
	"MjModeClearEnabled":              func(_, v string) error { setting.MjModeClearEnabled = v == "true"; return nil },
	"MjForwardUrlEnabled":             func(_, v string) error { setting.MjForwardUrlEnabled = v == "true"; return nil },
	"MjActionCheckSuccessEnabled":     func(_, v string) error { setting.MjActionCheckSuccessEnabled = v == "true"; return nil },
	"CheckSensitiveEnabled":           func(_, v string) error { setting.CheckSensitiveEnabled = v == "true"; return nil },
	"DemoSiteEnabled":                 func(_, v string) error { operation_setting.DemoSiteEnabled = v == "true"; return nil },
	"SelfUseModeEnabled":              func(_, v string) error { operation_setting.SelfUseModeEnabled = v == "true"; return nil },
	"CheckSensitiveOnPromptEnabled":   func(_, v string) error { setting.CheckSensitiveOnPromptEnabled = v == "true"; return nil },
	"ModelRequestRateLimitEnabled":    func(_, v string) error { setting.ModelRequestRateLimitEnabled = v == "true"; return nil },
	"StopOnSensitiveEnabled":          func(_, v string) error { setting.StopOnSensitiveEnabled = v == "true"; return nil },
	"SMTPSSLEnabled":                  func(_, v string) error { common.SMTPSSLEnabled = v == "true"; return nil },
	"SMTPForceAuthLogin":              func(_, v string) error { common.SMTPForceAuthLogin = v == "true"; return nil },
	"WorkerAllowHttpImageRequestEnabled": func(_, v string) error { system_setting.WorkerAllowHttpImageRequestEnabled = v == "true"; return nil },
	"DefaultUseAutoGroup":             func(_, v string) error { setting.DefaultUseAutoGroup = v == "true"; return nil },
	"StripePromotionCodesEnabled":     func(_, v string) error { setting.StripePromotionCodesEnabled = v == "true"; return nil },
	"CreemTestMode":                   func(_, v string) error { setting.CreemTestMode = v == "true"; return nil },
	"WaffoEnabled":                    func(_, v string) error { setting.WaffoEnabled = v == "true"; return nil },
	"WaffoSandbox":                    func(_, v string) error { setting.WaffoSandbox = v == "true"; return nil },
	"WaffoPancakeEnabled":             func(_, v string) error { setting.WaffoPancakeEnabled = v == "true"; return nil },
	"WaffoPancakeSandbox":             func(_, v string) error { setting.WaffoPancakeSandbox = v == "true"; return nil },

	// ExposeRatioEnabled (has setter function)
	"ExposeRatioEnabled": func(_, v string) error { ratio_setting.SetExposeRatioEnabled(v == "true"); return nil },

	// DisplayInCurrencyEnabled (syncs config to general_setting.quota_display_type)
	"DisplayInCurrencyEnabled": func(_, v string) error {
		newVal := "USD"
		if v != "true" {
			newVal = "TOKENS"
		}
		if cfg := config.GlobalConfig.Get("general_setting"); cfg != nil {
			_ = config.UpdateConfigFromMap(cfg, map[string]string{"quota_display_type": newVal})
		}
		return nil
	},

	// Syslog (triggers logger sync)
	"SyslogEnabled": func(_, v string) error {
		common.SyslogEnabled = v == "true"
		logger.SyncSyslogFromConfig()
		return nil
	},
	"SyslogNetwork": func(_, v string) error {
		common.SyslogNetwork = v
		logger.SyncSyslogFromConfig()
		return nil
	},
	"SyslogAddr": func(_, v string) error {
		common.SyslogAddr = v
		logger.SyncSyslogFromConfig()
		return nil
	},
	"SyslogTag": func(_, v string) error {
		common.SyslogTag = v
		logger.SyncSyslogFromConfig()
		return nil
	},

	// String assignments
	"EmailDomainWhitelist":         func(_, v string) error { common.EmailDomainWhitelist = strings.Split(v, ","); return nil },
	"SMTPServer":                   func(_, v string) error { common.SMTPServer = v; return nil },
	"SMTPAccount":                  func(_, v string) error { common.SMTPAccount = v; return nil },
	"SMTPFrom":                     func(_, v string) error { common.SMTPFrom = v; return nil },
	"SMTPToken":                    func(_, v string) error { common.SMTPToken = v; return nil },
	"ServerAddress":                func(_, v string) error { system_setting.ServerAddress = v; return nil },
	"WorkerUrl":                    func(_, v string) error { system_setting.WorkerUrl = v; return nil },
	"WorkerValidKey":               func(_, v string) error { system_setting.WorkerValidKey = v; return nil },
	"PayAddress":                   func(_, v string) error { operation_setting.PayAddress = v; return nil },
	"CustomCallbackAddress":        func(_, v string) error { operation_setting.CustomCallbackAddress = v; return nil },
	"EpayId":                       func(_, v string) error { operation_setting.EpayId = v; return nil },
	"EpayKey":                      func(_, v string) error { operation_setting.EpayKey = v; return nil },
	"StripeApiSecret":              func(_, v string) error { setting.StripeApiSecret = v; return nil },
	"StripeWebhookSecret":          func(_, v string) error { setting.StripeWebhookSecret = v; return nil },
	"StripePriceId":                func(_, v string) error { setting.StripePriceId = v; return nil },
	"CreemApiKey":                  func(_, v string) error { setting.CreemApiKey = v; return nil },
	"CreemProducts":                func(_, v string) error { setting.CreemProducts = v; return nil },
	"CreemWebhookSecret":           func(_, v string) error { setting.CreemWebhookSecret = v; return nil },
	"WaffoApiKey":                  func(_, v string) error { setting.WaffoApiKey = v; return nil },
	"WaffoPrivateKey":              func(_, v string) error { setting.WaffoPrivateKey = v; return nil },
	"WaffoPublicCert":              func(_, v string) error { setting.WaffoPublicCert = v; return nil },
	"WaffoSandboxPublicCert":       func(_, v string) error { setting.WaffoSandboxPublicCert = v; return nil },
	"WaffoSandboxApiKey":           func(_, v string) error { setting.WaffoSandboxApiKey = v; return nil },
	"WaffoSandboxPrivateKey":       func(_, v string) error { setting.WaffoSandboxPrivateKey = v; return nil },
	"WaffoMerchantId":              func(_, v string) error { setting.WaffoMerchantId = v; return nil },
	"WaffoNotifyUrl":               func(_, v string) error { setting.WaffoNotifyUrl = v; return nil },
	"WaffoReturnUrl":               func(_, v string) error { setting.WaffoReturnUrl = v; return nil },
	"WaffoSubscriptionReturnUrl":   func(_, v string) error { setting.WaffoSubscriptionReturnUrl = v; return nil },
	"WaffoCurrency":                func(_, v string) error { setting.WaffoCurrency = v; return nil },
	"WaffoPancakeMerchantID":       func(_, v string) error { setting.WaffoPancakeMerchantID = v; return nil },
	"WaffoPancakePrivateKey":       func(_, v string) error { setting.WaffoPancakePrivateKey = v; return nil },
	"WaffoPancakeWebhookPublicKey": func(_, v string) error { setting.WaffoPancakeWebhookPublicKey = v; return nil },
	"WaffoPancakeWebhookTestKey":   func(_, v string) error { setting.WaffoPancakeWebhookTestKey = v; return nil },
	"WaffoPancakeStoreID":          func(_, v string) error { setting.WaffoPancakeStoreID = v; return nil },
	"WaffoPancakeProductID":        func(_, v string) error { setting.WaffoPancakeProductID = v; return nil },
	"WaffoPancakeReturnURL":        func(_, v string) error { setting.WaffoPancakeReturnURL = v; return nil },
	"WaffoPancakeCurrency":         func(_, v string) error { setting.WaffoPancakeCurrency = v; return nil },
	"GitHubClientId":               func(_, v string) error { common.GitHubClientId = v; return nil },
	"GitHubClientSecret":           func(_, v string) error { common.GitHubClientSecret = v; return nil },
	"LinuxDOClientId":              func(_, v string) error { common.LinuxDOClientId = v; return nil },
	"LinuxDOClientSecret":          func(_, v string) error { common.LinuxDOClientSecret = v; return nil },
	"Footer":                       func(_, v string) error { common.Footer = v; return nil },
	"SystemName":                   func(_, v string) error { common.SystemName = v; return nil },
	"Logo":                         func(_, v string) error { common.Logo = v; return nil },
	"WeChatServerAddress":          func(_, v string) error { common.WeChatServerAddress = v; return nil },
	"WeChatServerToken":            func(_, v string) error { common.WeChatServerToken = v; return nil },
	"WeChatAccountQRCodeImageURL":  func(_, v string) error { common.WeChatAccountQRCodeImageURL = v; return nil },
	"TelegramBotToken":             func(_, v string) error { common.TelegramBotToken = v; return nil },
	"TelegramBotName":              func(_, v string) error { common.TelegramBotName = v; return nil },
	"TurnstileSiteKey":             func(_, v string) error { common.TurnstileSiteKey = v; return nil },
	"TurnstileSecretKey":           func(_, v string) error { common.TurnstileSecretKey = v; return nil },
	"DataExportDefaultTime":        func(_, v string) error { common.DataExportDefaultTime = v; return nil },
	"TopUpLink":                    func(_, v string) error { common.TopUpLink = v; return nil },

	// Int assignments
	"SMTPPort":                          func(_, v string) error { val, _ := strconv.Atoi(v); common.SMTPPort = val; return nil },
	"LinuxDOMinimumTrustLevel":          func(_, v string) error { val, _ := strconv.Atoi(v); common.LinuxDOMinimumTrustLevel = val; return nil },
	"MinTopUp":                          func(_, v string) error { val, _ := strconv.Atoi(v); operation_setting.MinTopUp = val; return nil },
	"StripeMinTopUp":                    func(_, v string) error { val, _ := strconv.Atoi(v); setting.StripeMinTopUp = val; return nil },
	"WaffoMinTopUp":                     func(_, v string) error { val, _ := strconv.Atoi(v); setting.WaffoMinTopUp = val; return nil },
	"WaffoPancakeMinTopUp":              func(_, v string) error { val, _ := strconv.Atoi(v); setting.WaffoPancakeMinTopUp = val; return nil },
	"QuotaForNewUser":                   func(_, v string) error { val, _ := strconv.Atoi(v); common.QuotaForNewUser = val; return nil },
	"QuotaForInviter":                   func(_, v string) error { val, _ := strconv.Atoi(v); common.QuotaForInviter = val; return nil },
	"QuotaForInvitee":                   func(_, v string) error { val, _ := strconv.Atoi(v); common.QuotaForInvitee = val; return nil },
	"QuotaRemindThreshold":              func(_, v string) error { val, _ := strconv.Atoi(v); common.QuotaRemindThreshold = val; return nil },
	"PreConsumedQuota":                  func(_, v string) error { val, _ := strconv.Atoi(v); common.PreConsumedQuota = val; return nil },
	"ModelRequestRateLimitCount":        func(_, v string) error { val, _ := strconv.Atoi(v); setting.ModelRequestRateLimitCount = val; return nil },
	"ModelRequestRateLimitDurationMinutes": func(_, v string) error { val, _ := strconv.Atoi(v); setting.ModelRequestRateLimitDurationMinutes = val; return nil },
	"ModelRequestRateLimitSuccessCount":    func(_, v string) error { val, _ := strconv.Atoi(v); setting.ModelRequestRateLimitSuccessCount = val; return nil },
	"RetryTimes":                        func(_, v string) error { val, _ := strconv.Atoi(v); common.RetryTimes = val; return nil },
	"DataExportInterval":                func(_, v string) error { val, _ := strconv.Atoi(v); common.DataExportInterval = val; return nil },
	"StreamCacheQueueLength":            func(_, v string) error { val, _ := strconv.Atoi(v); setting.StreamCacheQueueLength = val; return nil },

	// Float64 assignments
	"Price":                  func(_, v string) error { val, _ := strconv.ParseFloat(v, 64); operation_setting.Price = val; return nil },
	"USDExchangeRate":        func(_, v string) error { val, _ := strconv.ParseFloat(v, 64); operation_setting.USDExchangeRate = val; return nil },
	"StripeUnitPrice":        func(_, v string) error { val, _ := strconv.ParseFloat(v, 64); setting.StripeUnitPrice = val; return nil },
	"WaffoUnitPrice":         func(_, v string) error { val, _ := strconv.ParseFloat(v, 64); setting.WaffoUnitPrice = val; return nil },
	"WaffoPancakeUnitPrice":  func(_, v string) error { val, _ := strconv.ParseFloat(v, 64); setting.WaffoPancakeUnitPrice = val; return nil },
	"ChannelDisableThreshold": func(_, v string) error { val, _ := strconv.ParseFloat(v, 64); common.ChannelDisableThreshold = val; return nil },
	"QuotaPerUnit":           func(_, v string) error { val, _ := strconv.ParseFloat(v, 64); common.QuotaPerUnit = val; return nil },

	// Complex JSON/struct handlers
	"Chats":                      func(_, v string) error { return setting.UpdateChatsByJsonString(v) },
	"AutoGroups":                 func(_, v string) error { return setting.UpdateAutoGroupsByJsonString(v) },
	"TopupGroupRatio":            func(_, v string) error { return common.UpdateTopupGroupRatioByJSONString(v) },
	"ModelRequestRateLimitGroup": func(_, v string) error { return setting.UpdateModelRequestRateLimitGroupByJSONString(v) },
	"ModelRatio":                 func(_, v string) error { return ratio_setting.UpdateModelRatioByJSONString(v) },
	"GroupRatio":                 func(_, v string) error { return ratio_setting.UpdateGroupRatioByJSONString(v) },
	"GroupGroupRatio":            func(_, v string) error { return ratio_setting.UpdateGroupGroupRatioByJSONString(v) },
	"UserUsableGroups":           func(_, v string) error { return setting.UpdateUserUsableGroupsByJSONString(v) },
	"CompletionRatio":            func(_, v string) error { return ratio_setting.UpdateCompletionRatioByJSONString(v) },
	"ModelPrice":                 func(_, v string) error { return ratio_setting.UpdateModelPriceByJSONString(v) },
	"CacheRatio":                 func(_, v string) error { return ratio_setting.UpdateCacheRatioByJSONString(v) },
	"CreateCacheRatio":           func(_, v string) error { return ratio_setting.UpdateCreateCacheRatioByJSONString(v) },
	"ImageRatio":                 func(_, v string) error { return ratio_setting.UpdateImageRatioByJSONString(v) },
	"AudioRatio":                 func(_, v string) error { return ratio_setting.UpdateAudioRatioByJSONString(v) },
	"AudioCompletionRatio":       func(_, v string) error { return ratio_setting.UpdateAudioCompletionRatioByJSONString(v) },
	"PayMethods":                 func(_, v string) error { return operation_setting.UpdatePayMethodsByJsonString(v) },
	"SensitiveWords":             func(_, v string) error { setting.SensitiveWordsFromString(v); return nil },
	"AutomaticDisableKeywords":   func(_, v string) error { operation_setting.AutomaticDisableKeywordsFromString(v); return nil },
	"AutomaticDisableStatusCodes": func(_, v string) error { return operation_setting.AutomaticDisableStatusCodesFromString(v) },
	"AutomaticRetryStatusCodes":  func(_, v string) error { return operation_setting.AutomaticRetryStatusCodesFromString(v) },

	// WaffoPayMethods - value is read directly from OptionMap; no global variable to sync
	"WaffoPayMethods": func(_, v string) error { return nil },
}

func updateOptionMap(key string, value string) (err error) {
	common.OptionMapRWMutex.Lock()
	defer common.OptionMapRWMutex.Unlock()
	common.OptionMap[key] = value

	if handleConfigUpdate(key, value) {
		return nil
	}

	handler, ok := optionHandlers[key]
	if ok {
		return handler(key, value)
	}
	return nil
}

// handleConfigUpdate 处理分层配置更新，返回是否已处理
func handleConfigUpdate(key, value string) bool {
	parts := strings.SplitN(key, ".", 2)
	if len(parts) != 2 {
		return false // 不是分层配置
	}

	configName := parts[0]
	configKey := parts[1]

	// 获取配置对象
	cfg := config.GlobalConfig.Get(configName)
	if cfg == nil {
		return false // 未注册的配置
	}

	// 更新配置
	configMap := map[string]string{
		configKey: value,
	}
	config.UpdateConfigFromMap(cfg, configMap)

	// 特定配置的后处理
	if configName == "performance_setting" {
		// 同步磁盘缓存配置到 common 包
		performance_setting.UpdateAndSync()
	}

	return true // 已处理
}
