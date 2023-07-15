package repository

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nexus "github.com/nduyphuong/go-nexus-client/nexus3"
	"github.com/nduyphuong/go-nexus-client/nexus3/schema/repository"
	"github.com/nduyphuong/terraform-provider-nexus/internal/schema/common"
	repositorySchema "github.com/nduyphuong/terraform-provider-nexus/internal/schema/repository"
)

func ResourceRepositoryPypiGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to create a group pypi repository.",

		Create: resourcePypiGroupRepositoryCreate,
		Delete: resourcePypiGroupRepositoryDelete,
		Exists: resourcePypiGroupRepositoryExists,
		Read:   resourcePypiGroupRepositoryRead,
		Update: resourcePypiGroupRepositoryUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			// Common schemas
			"id":     common.ResourceID,
			"name":   repositorySchema.ResourceName,
			"online": repositorySchema.ResourceOnline,
			// Group schemas
			"group":   repositorySchema.ResourceGroup,
			"storage": repositorySchema.ResourceStorage,
		},
	}
}

func getPypiGroupRepositoryFromResourceData(resourceData *schema.ResourceData) repository.PypiGroupRepository {
	storageConfig := resourceData.Get("storage").([]interface{})[0].(map[string]interface{})
	groupConfig := resourceData.Get("group").([]interface{})[0].(map[string]interface{})
	groupMemberNamesInterface := groupConfig["member_names"].([]interface{})
	groupMemberNames := make([]string, 0)
	for _, v := range groupMemberNamesInterface {
		groupMemberNames = append(groupMemberNames, v.(string))
	}

	repo := repository.PypiGroupRepository{
		Name:   resourceData.Get("name").(string),
		Online: resourceData.Get("online").(bool),
		Storage: repository.Storage{
			BlobStoreName:               storageConfig["blob_store_name"].(string),
			StrictContentTypeValidation: storageConfig["strict_content_type_validation"].(bool),
		},
		Group: repository.Group{
			MemberNames: groupMemberNames,
		},
	}

	return repo
}

func setPypiGroupRepositoryToResourceData(repo *repository.PypiGroupRepository, resourceData *schema.ResourceData) error {
	resourceData.SetId(repo.Name)
	resourceData.Set("name", repo.Name)
	resourceData.Set("online", repo.Online)

	if err := resourceData.Set("storage", flattenStorage(&repo.Storage)); err != nil {
		return err
	}

	if err := resourceData.Set("group", flattenGroup(&repo.Group)); err != nil {
		return err
	}

	return nil
}

func resourcePypiGroupRepositoryCreate(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)

	repo := getPypiGroupRepositoryFromResourceData(resourceData)

	if err := client.Repository.Pypi.Group.Create(repo); err != nil {
		return err
	}
	resourceData.SetId(repo.Name)

	return resourcePypiGroupRepositoryRead(resourceData, m)
}

func resourcePypiGroupRepositoryRead(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)

	repo, err := client.Repository.Pypi.Group.Get(resourceData.Id())
	if err != nil {
		return err
	}

	if repo == nil {
		resourceData.SetId("")
		return nil
	}

	return setPypiGroupRepositoryToResourceData(repo, resourceData)
}

func resourcePypiGroupRepositoryUpdate(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)

	repoName := resourceData.Id()
	repo := getPypiGroupRepositoryFromResourceData(resourceData)
	repo1, err := client.Repository.Pypi.Group.Get(resourceData.Id())
	if err != nil {
		return err
	}
	if reflect.DeepEqual(repo1.Group.MemberNames, repo.Group.MemberNames) {
		return nil
	}
	if err := client.Repository.Pypi.Group.Update(repoName, repo); err != nil {
		return err
	}

	return resourcePypiGroupRepositoryRead(resourceData, m)
}

func resourcePypiGroupRepositoryDelete(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)
	return client.Repository.Pypi.Group.Delete(resourceData.Id())
}

func resourcePypiGroupRepositoryExists(resourceData *schema.ResourceData, m interface{}) (bool, error) {
	client := m.(*nexus.NexusClient)

	repo, err := client.Repository.Pypi.Group.Get(resourceData.Id())
	return repo != nil, err
}
