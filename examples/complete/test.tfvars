// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

product_family  = "dso"
product_service = "test"

tags = {
  Purpose = "Terratest"
}

enable_rbac_authorization = true

# These are not really secrets, they are just dummies for the test
secret_name  = "test-secret" #pragma: allowlist secret
secret_value = "test-value"  #pragma: allowlist secret
