package repository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	ResourceGroup = &schema.Schema{
		Description: "Configuration for repository group",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"member_names": {
					Description: "Member repositories names",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"order": &schema.Schema{
								Type:     schema.TypeInt,
								Required: true,
							},
							"name": &schema.Schema{
								Type:     schema.TypeString,
								Required: true,
							},
						},
					},
					MinItems: 1,
					Required: true,
					// Set: func(v interface{}) int {
					// 	return schema.HashString(strings.ToLower(v.(string)))
					// },
					Type: schema.TypeSet,
				},
			},
		},
		MaxItems: 1,
		Required: true,
		Type:     schema.TypeList,
	}
	ResourceGroupDeploy = &schema.Schema{
		Description: "Configuration for repository group",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"member_names": {
					Description: "Member repositories names",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"order": &schema.Schema{
								Type:     schema.TypeInt,
								Required: true,
							},
							"name": &schema.Schema{
								Type:     schema.TypeString,
								Required: true,
							},
						},
					},
					MinItems: 1,
					Required: true,
					Type:     schema.TypeSet,
				},
				"writable_member": {
					Description: "Pro-only: This field is for the Group Deployment feature available in NXRM Pro.",
					Optional:    true,
					Type:        schema.TypeString,
				},
			},
		},
		MaxItems: 1,
		Required: true,
		Type:     schema.TypeList,
	}
	DataSourceGroup = &schema.Schema{
		Description: "Configuration for repository group",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"member_names": {
					Description: "Member repositories names",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"order": {
								Type:     schema.TypeInt,
								Required: true,
							},
							"name": {
								Type:     schema.TypeString,
								Required: true,
							},
						},
					},
					Computed: true,
					Type:     schema.TypeSet,
				},
			},
		},
		Computed: true,
		Type:     schema.TypeList,
	}
	DataSourceGroupDeploy = &schema.Schema{
		Description: "Configuration for repository group",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"member_names": {
					Description: "Member repositories names",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"order": {
								Type:     schema.TypeInt,
								Required: true,
							},
							"name": {
								Type:     schema.TypeString,
								Required: true,
							},
						},
					},
					Computed: true,
					Type:     schema.TypeSet,
				},
				"writable_member": {
					Description: "Pro-only: This field is for the Group Deployment feature available in NXRM Pro.",
					Computed:    true,
					Type:        schema.TypeString,
				},
			},
		},
		Computed: true,
		Type:     schema.TypeList,
	}
)
