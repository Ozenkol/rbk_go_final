package setting

type SettingKey string

const (
	SettingKeyEmail       SettingKey = "Email provider"
	SettingKeyTelephony   SettingKey = "Phone Number"
	SettingKeyPayment     SettingKey = "Address"
	SettingKeyStorage     SettingKey = "Storage"
	SettingKeyShipping    SettingKey = "Shipping"
	SettingKeyLLMProvider SettingKey = "LLM Provider"
)