// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkServicesEdgeCacheKeysetDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetBasicExample(context),
			},
			{
				ResourceName:            "google_network_services_edge_cache_keyset.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name"},
			},
		},
	})
}

func testAccNetworkServicesEdgeCacheKeyset_networkServicesEdgeCacheKeysetBasicExample(context map[string]interface{}) string {
	return Nprintf(`

resource "google_network_services_edge_cache_keyset" "default" {
  name                 = "default%{random_suffix}"
  description          = "The default keyset"
  public_key {
    id = "my-public-key"
    value = "FHsTyFHNmvNpw4o7-rp-M1yqMyBF8vXSBRkZtkQ0RKY"
  }
  public_key {
    id = "my-public-key-2"
    value = "hzd03llxB1u5FOLKFkZ6_wCJqC7jtN0bg7xlBqS6WVM"
  }
}
`, context)
}

func testAccCheckNetworkServicesEdgeCacheKeysetDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_network_services_edge_cache_keyset" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{NetworkServicesBasePath}}projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("NetworkServicesEdgeCacheKeyset still exists at %s", url)
			}
		}

		return nil
	}
}