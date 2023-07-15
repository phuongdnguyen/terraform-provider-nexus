package repository

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nduyphuong/terraform-provider-nexus/internal/schema/common"
	"github.com/nduyphuong/terraform-provider-nexus/internal/schema/repository"
)

func DataSourceRepositoryBowerGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get an existing bower group repository.",

		Read: dataSourceRepositoryBowerGroupRead,
		Schema: map[string]*schema.Schema{
			// Common schemas
			"id":     common.DataSourceID,
			"name":   repository.DataSourceName,
			"online": repository.DataSourceOnline,
			// Group schemas
			"group":   repository.DataSourceGroup,
			"storage": repository.DataSourceStorage,
		},
	}
}

func dataSourceRepositoryBowerGroupRead(resourceData *schema.ResourceData, m interface{}) error {
	resourceData.SetId(resourceData.Get("name").(string))

	return resourceBowerGroupRepositoryRead(resourceData, m)
}
