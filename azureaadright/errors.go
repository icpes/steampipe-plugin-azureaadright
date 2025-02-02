package azureaadright

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// isNotFoundError:: function which returns an ErrorPredicate for Azure API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		azureaadrightConfig := GetConfig(d.Connection)

		// If the get or list hydrate functions have an overriding IgnoreConfig
		// defined using the isNotFoundError function, then it should
		// also check for errors in the "ignore_error_codes" config argument
		allErrors := append(notFoundErrors, azureaadrightConfig.IgnoreErrorCodes...)
		// Added to support regex in not found errors
		for _, pattern := range allErrors {
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}

// shouldIgnoreErrorPluginDefault:: Plugin level default function to ignore a set errors for hydrate functions based on "ignore_error_codes" config argument
func shouldIgnoreErrorPluginDefault() plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		if !hasIgnoredErrorCodes(d.Connection) {
			return false
		}

		azureaadrightConfig := GetConfig(d.Connection)
		// Added to support regex in ignoring errors
		for _, pattern := range azureaadrightConfig.IgnoreErrorCodes {
			if strings.Contains(err.Error(), pattern) {
				return true
			}
		}
		return false
	}
}

func hasIgnoredErrorCodes(connection *plugin.Connection) bool {
	azureaadrightConfig := GetConfig(connection)
	return len(azureaadrightConfig.IgnoreErrorCodes) > 0
}
