package openstack

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccNetworkingV2FloatingIPAssociate_importBasic(t *testing.T) {
	resourceName := "openstack_networking_floatingip_associate_v2.fip_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkingV2FloatingIPAssociateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkingV2FloatingIPAssociateBasic(),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
