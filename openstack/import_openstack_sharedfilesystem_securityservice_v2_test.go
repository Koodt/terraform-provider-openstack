package openstack

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccSFSV2SecurityService_importBasic(t *testing.T) {
	resourceName := "openstack_sharedfilesystem_securityservice_v2.securityservice_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckSFS(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSFSV2SecurityServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSFSV2SecurityServiceConfigBasic,
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
				},
			},
		},
	})
}
