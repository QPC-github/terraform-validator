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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIapTunnelIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.tunnelResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapTunnelIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel roles/iap.tunnelResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccIapTunnelIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel roles/iap.tunnelResourceAccessor", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapTunnelIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.tunnelResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccIapTunnelIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel roles/iap.tunnelResourceAccessor user:admin@hashicorptest.com", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIapTunnelIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/iap.tunnelResourceAccessor",
		"org_id":        getTestOrgFromEnv(t),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {},
			"time":   {},
		},
		Steps: []resource.TestStep{
			{
				Config: testAccIapTunnelIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_iap_tunnel_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIapTunnelIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_iap_tunnel_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/iap_tunnel", fmt.Sprintf("tf-test%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIapTunnelIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_tunnel_iam_member" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccIapTunnelIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_iap_tunnel_iam_policy" "foo" {
  project = google_project_service.project_service.project
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapTunnelIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

data "google_iam_policy" "foo" {
}

resource "google_iap_tunnel_iam_policy" "foo" {
  project = google_project_service.project_service.project
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccIapTunnelIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_tunnel_iam_binding" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccIapTunnelIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "project" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
}

resource "time_sleep" "wait_60_seconds" {
  depends_on = [google_project.project]

  create_duration = "60s"
}

resource "google_project_service" "project_service" {
  project = google_project.project.project_id
  service = "iap.googleapis.com"

  # Needed for CI tests for permissions to propagate, should not be needed for actual usage
  depends_on = [time_sleep.wait_60_seconds]
}

resource "google_iap_tunnel_iam_binding" "foo" {
  project = google_project_service.project_service.project
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
