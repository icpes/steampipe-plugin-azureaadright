package azureaadright

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type azureaadrightConfig struct {
	TenantID            *string  `hcl:"tenant_id"`
	SubscriptionID      *string  `hcl:"subscription_id"`
	ClientID            *string  `hcl:"client_id"`
	ClientSecret        *string  `hcl:"client_secret"`
	CertificatePath     *string  `hcl:"certificate_path"`
	CertificatePassword *string  `hcl:"certificate_password"`
	Username            *string  `hcl:"username"`
	Password            *string  `hcl:"password"`
	Environment         *string  `hcl:"environment"`
	IgnoreErrorCodes    []string `hcl:"ignore_error_codes,optional"`
}

func ConfigInstance() interface{} {
	return &azureaadrightConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) azureaadrightConfig {
	if connection == nil || connection.Config == nil {
		return azureaadrightConfig{}
	}
	config, _ := connection.Config.(azureaadrightConfig)
	return config
}
