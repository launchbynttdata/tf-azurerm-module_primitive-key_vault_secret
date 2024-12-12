product_family  = "dso"
product_service = "test"

tags = {
  Purpose = "Terratest"
}

enable_rbac_authorization = true

# These are not really secrets, they are just dummies for the test
secret_name  = "test-secret" #pragma: allowlist secret
secret_value = "test-value"  #pragma: allowlist secret
