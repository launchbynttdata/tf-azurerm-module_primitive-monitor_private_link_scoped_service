package testimpl

import (
	"os"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	resourceId := terraform.Output(t, ctx.TerratestTerraformOptions(), "id")

	t.Run("TfOutputsNotEmpty", func(t *testing.T) {
		assert.NotEmpty(t, resourceId, "Scoped service resource ID must not be empty")
	})

	t.Run("CheckMonitorPrivateLinkScope", func(t *testing.T) {
		logAnalyticsName := terraform.Output(t, ctx.TerratestTerraformOptions(), "log_analytics_workspace_name")
		rgName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")

		workspace := azure.GetLogAnalyticsWorkspace(t, logAnalyticsName, rgName, subscriptionId)
		if assert.NotEmpty(t, *workspace.PrivateLinkScopedResources) {
			assert.Equal(t, strings.ToLower(resourceId), strings.ToLower(*(*workspace.PrivateLinkScopedResources)[0].ResourceID))
		}
	})
}
