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

func TestAccOSLoginSSHPublicKey_osLoginSshKeyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOSLoginSSHPublicKeyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccOSLoginSSHPublicKey_osLoginSshKeyBasicExample(context),
			},
			{
				ResourceName:            "google_os_login_ssh_public_key.cache",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"user", "project"},
			},
		},
	})
}

func testAccOSLoginSSHPublicKey_osLoginSshKeyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_client_openid_userinfo" "me" {
}

resource "google_os_login_ssh_public_key" "cache" {
  user =  data.google_client_openid_userinfo.me.email
  key = file("test-fixtures/ssh_rsa.pub")
}
`, context)
}

func testAccCheckOSLoginSSHPublicKeyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_os_login_ssh_public_key" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{OSLoginBasePath}}users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("OSLoginSSHPublicKey still exists at %s", url)
			}
		}

		return nil
	}
}
