package other

import (
	nexusSchema "github.com/datadrivers/go-nexus-client/nexus3/schema"
	"github.com/datadrivers/terraform-provider-nexus/internal/schema/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceCleanUpPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to create a Nexus Cleanup Policy Rule.",

		Create: resourceCleanUpPolicyCreate,
		Read:   resourceCleanUpPolicyRead,
		Update: resourceCleanUpPolicyUpdate,
		Delete: resourceCleanUpPolicyDelete,
		Exists: resourceCleanUpPolicyExists,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": common.ResourceID,
			"name": {
				Description: "The name of the cleanup policy rule",
				ForceNew:    true,
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}

}
func getCleanUpPolicyFromResourceData(d *schema.ResourceData) nexusSchema.Script {
}

func resourceCleanUpPolicyCreate(d *schema.ResourceData, m interface{}) error {

}

func resourceCleanUpPolicyRead(d *schema.ResourceData, m interface{}) error {

}

func resourceCleanUpPolicyUpdate(d *schema.ResourceData, m interface{}) error {

}

func resourceCleanUpPolicyDelete(d *schema.ResourceData, m interface{}) error {

}

func resourceCleanUpPolicyExists(d *schema.ResourceData, m interface{}) (bool, error) {

}
