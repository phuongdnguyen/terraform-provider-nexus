package deprecated_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/nduyphuong/go-nexus-client/nexus3/schema/repository"
	"github.com/nduyphuong/terraform-provider-nexus/internal/acceptance"
)

func TestAccDataSourceRepositoryHosted(t *testing.T) {
	repoName := "maven-releases"
	resourceName := "data.nexus_repository.acceptance"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.AccPreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceRepositoryConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "id", repoName),
						resource.TestCheckResourceAttr(resourceName, "name", repoName),
						resource.TestCheckResourceAttr(resourceName, "format", repository.RepositoryFormatMaven2),
						resource.TestCheckResourceAttr(resourceName, "type", repository.RepositoryTypeHosted),
					),
				),
			},
		},
	})
}
