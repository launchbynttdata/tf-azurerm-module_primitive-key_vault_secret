package common

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableKeyVaultSecret(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionId) == 0 {
		t.Fatal("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	rgName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	keyVaultName := terraform.Output(t, ctx.TerratestTerraformOptions(), "key_vault_name")
	secretName := terraform.Output(t, ctx.TerratestTerraformOptions(), "secret_name")

	t.Run("KeyVaultExists", func(t *testing.T) {
		keyVault := azure.GetKeyVault(t, rgName, keyVaultName, subscriptionId)
		assert.Equal(t, keyVaultName, *keyVault.Name, "Expected Key Vault name to be %s, but got %s", keyVaultName, *keyVault.Name)
	})

	t.Run("SecretExists", func(t *testing.T) {
		secretExists := azure.KeyVaultSecretExists(t, keyVaultName, secretName)
		assert.True(t, secretExists, "Expected secret %s to exist in Key Vault %s, but it does not", secretName, keyVaultName)
	})
}
