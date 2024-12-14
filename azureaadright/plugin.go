package azureaadright

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-azure"

// Plugin creates this (azure) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: isNotFoundError([]string{"ResourceGroupNotFound"}),
			},
		},
		// Default ignore config for the plugin
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrorPluginDefault(),
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "subscription_id",
				Hydrate: getSubscriptionIdForConnection,
			},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{			
			"azure_log_analytics_workspace":                                tableAzureLogAnalyticsWorkspace(ctx),
			"azure_monitor_activity_log_event":                             tableAzureMonitorActivityLogEvent(ctx),
			"azure_resource_group":                                         tableAzureResourceGroup(ctx),			
			"azure_role_assignment":                                        tableAzureIamRoleAssignment(ctx),
			"azure_role_definition":                                        tableAzureIamRoleDefinition(ctx),			
			"azure_security_center_auto_provisioning":                      tableAzureSecurityCenterAutoProvisioning(ctx),
			"azure_security_center_automation":                             tableAzureSecurityCenterAutomation(ctx),
			"azure_security_center_contact":                                tableAzureSecurityCenterContact(ctx),
			"azure_security_center_jit_network_access_policy":              tableAzureSecurityCenterJITNetworkAccessPolicy(ctx),
			"azure_security_center_setting":                                tableAzureSecurityCenterSetting(ctx),
			"azure_security_center_sub_assessment":                         tableAzureSecurityCenterSubAssessment(ctx),
			"azure_security_center_subscription_pricing":                   tableAzureSecurityCenterPricing(ctx),			
		},
	}

	return p
}

func getSubscriptionIdForConnection(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (any, error) {
	subscriptionID, err := getSubscriptionIDMemoized(ctx, d, h)
	if err != nil {
		return nil, err
	}
	return subscriptionID, nil
}
