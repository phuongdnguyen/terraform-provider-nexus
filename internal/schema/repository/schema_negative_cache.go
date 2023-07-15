package repository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ResourceNegativeCache = &schema.Schema{
		Description: "Configuration of the negative cache handling",
		Optional:    true,
		Type:        schema.TypeList,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"enabled": {
					Default:     NegativeCacheDefaultEnabled,
					Description: "Whether to cache responses for content not present in the proxied repository, defaults to `false` if unset",
					Optional:    true,
					Type:        schema.TypeBool,
				},
				"ttl": {
					Default:     NegativeCacheDefaultTTL,
					Description: "How long to cache the fact that a file was not found in the repository (in minutes), defaults is `1440` if unset",
					Optional:    true,
					Type:        schema.TypeInt,
				},
			},
		},
	}
	ResourceNegativeCacheEnabled = &schema.Schema{
		Description: "Configuration of the negative cache handling, defaults to `false` if unset",
		Optional:    true,
		Type:        schema.TypeBool,
		Default:     NegativeCacheDefaultEnabled,
	}
	ResourceNegativeCacheTTL = &schema.Schema{
		Description: "Configuration of the negative cache handling, defaults is `1440` if unset",
		Optional:    true,
		Type:        schema.TypeInt,
		Default:     NegativeCacheDefaultTTL,
	}
	DataSourceNegativeCache = &schema.Schema{
		Description: "Configuration of the negative cache handling",
		Type:        schema.TypeList,
		Computed:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"enabled": {
					Description: "Whether to cache responses for content not present in the proxied repository",
					Type:        schema.TypeBool,
					Computed:    true,
				},
				"ttl": {
					Description: "How long to cache the fact that a file was not found in the repository (in minutes)",
					Type:        schema.TypeInt,
					Computed:    true,
				},
			},
		},
	}
)
