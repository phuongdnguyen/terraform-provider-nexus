package repository_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/nduyphuong/go-nexus-client/nexus3/schema/repository"
	"github.com/nduyphuong/terraform-provider-nexus/internal/acceptance"
)

func testAccDataSourceRepositoryBowerHostedConfig() string {
	return `
data "nexus_repository_bower_hosted" "acceptance" {
	name   = nexus_repository_bower_hosted.acceptance.id
}`
}

func TestAccDataSourceRepositoryBowerHosted(t *testing.T) {
	repo := repository.BowerHostedRepository{
		Name:   fmt.Sprintf("acceptance-%s", acctest.RandString(10)),
		Online: true,
		Storage: repository.HostedStorage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: false,
		},
	}
	dataSourceName := "data.nexus_repository_bower_hosted.acceptance"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.AccPreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRepositoryBowerHostedConfig(repo) + testAccDataSourceRepositoryBowerHostedConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(dataSourceName, "id", repo.Name),
						resource.TestCheckResourceAttr(dataSourceName, "name", repo.Name),
						resource.TestCheckResourceAttr(dataSourceName, "online", strconv.FormatBool(repo.Online)),
						resource.TestCheckResourceAttr(dataSourceName, "storage.0.blob_store_name", repo.Storage.BlobStoreName),
						resource.TestCheckResourceAttr(dataSourceName, "storage.0.strict_content_type_validation", strconv.FormatBool(repo.Storage.StrictContentTypeValidation)),
					),
				),
			},
		},
	})
}
