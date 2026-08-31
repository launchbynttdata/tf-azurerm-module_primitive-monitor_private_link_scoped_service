package testimpl

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	subscriptionID := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	resourceID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "id")

	t.Run("TfOutputsNotEmpty", func(t *testing.T) {
		assert.NotEmpty(t, resourceID, "Scoped service resource ID must not be empty")
	})

	t.Run("CheckMonitorPrivateLinkScope", func(t *testing.T) {
		logAnalyticsName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "log_analytics_workspace_name")
		rgName := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "resource_group_name")
		scopeID := terraform.OutputContext(t, context.Background(), ctx.TerratestTerraformOptions(), "monitor_private_link_scope_id")
		scopeName := privateLinkScopeName(scopeID)
		require.NotEmpty(t, scopeName, "monitor_private_link_scope_id must contain a scope name")

		credential, err := azidentity.NewDefaultAzureCredential(nil)
		require.NoError(t, err, "failed to create Azure credential")

		client, err := armmonitor.NewPrivateLinkScopedResourcesClient(subscriptionID, credential, nil)
		require.NoError(t, err, "failed to create PrivateLinkScopedResourcesClient")

		scopedResource, err := client.Get(context.Background(), rgName, scopeName, logAnalyticsName, nil)
		require.NoError(t, err, "failed to get private link scoped resource")
		require.NotNil(t, scopedResource.Properties)
		require.NotNil(t, scopedResource.Properties.LinkedResourceID)

		assert.Equal(t, strings.ToLower(resourceID), strings.ToLower(*scopedResource.Properties.LinkedResourceID))
	})
}

func privateLinkScopeName(scopeID string) string {
	const marker = "/privatelinkscopes/"
	idx := strings.Index(strings.ToLower(scopeID), marker)
	if idx < 0 {
		return ""
	}
	rest := scopeID[idx+len(marker):]
	if slash := strings.Index(rest, "/"); slash >= 0 {
		rest = rest[:slash]
	}
	return rest
}
