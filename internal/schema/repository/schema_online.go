package repository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ResourceOnline = &schema.Schema{
		Default:     true,
		Description: "Whether this repository accepts incoming requests, defaults to `true` if unset",
		Optional:    true,
		Type:        schema.TypeBool,
	}
	DataSourceOnline = &schema.Schema{
		Description: "Whether this repository accepts incoming requests",
		Type:        schema.TypeBool,
		Computed:    true,
	}
)
