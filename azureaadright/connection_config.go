package azureaadright

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type azureaadrightConfig struct {
	TenantID            *string `hcl:"tenant_id"`
	ClientID            *string `hcl:"client_id"`
	ClientSecret        *string `hcl:"client_secret"`
	CertificatePath     *string `hcl:"certificate_path"`
	CertificatePassword *string `hcl:"certificate_password"`
	EnableMsi           *bool   `hcl:"enable_msi"`
	MsiEndpoint         *string `hcl:"msi_endpoint"`
	Environment         *string `hcl:"environment"`
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
