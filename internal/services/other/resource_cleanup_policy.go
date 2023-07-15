package other

import (
	"fmt"
	"github.com/datadrivers/terraform-provider-nexus/internal/schema/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	nexus "github.com/nduyphuong/go-nexus-client/nexus3"
	nexusSchema "github.com/nduyphuong/go-nexus-client/nexus3/schema"
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
			"format": {
				Description: "The format that this cleanup policy can be applied to",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"notes": {
				Description: "Notes for this policy",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"criteria": {
				Description: "Cleanup criteria",
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"last_downloaded_days": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Remove components that were published over this amount of time",
						},
						"last_blob_updated_days": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Remove components that haven't been downloaded in this amount of time",
						},
						"regex": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Remove components that have at least one asset name matching the following" +
								" regular expression pattern",
						},
					},
				},
			},
		},
	}
}

type CleanUpPolicy struct {
	Name     string
	Format   string
	Notes    string
	Criteria Criteria
}

type Criteria struct {
	LastDownloaded  int
	LastBlobUpdated int
	Regex           string
}

func getCleanUpPolicyFromResourceData(d *schema.ResourceData) CleanUpPolicy {
	c := CleanUpPolicy{
		Name:   d.Get("name").(string),
		Format: d.Get("format").(string),
		Notes:  d.Get("notes").(string),
	}
	if _, ok := d.GetOk("criteria"); ok {
		rList := d.Get("criteria").([]interface{})
		rConfig := rList[0].(map[string]interface{})
		c.Criteria = Criteria{
			LastDownloaded:  rConfig["last_downloaded_days"].(int),
			LastBlobUpdated: rConfig["last_blob_updated_days"].(int),
			Regex:           rConfig["regex"].(string),
		}
	}
	return c
}

func GetPayload(name, format, notes string, lastDownloaded, lastBlobUpdated int) string {
	return fmt.Sprintf(`
	{     "name":"%s",
		  "format":"%s",
		  "notes":"%s",
	      "criteria": {
          "lastBlobUpdated": %v,
          "lastDownloaded": %v
          }
	}
`, name, format, notes, lastDownloaded, lastBlobUpdated)
}
func NewCleanUpScript(name, format, notes string, criteria Criteria) nexusSchema.Script {
	var content = `// Original from:
// https://github.com/idealista/nexus-role/blob/master/files/scripts/cleanup_policy.groovy
import com.google.common.collect.Maps
import groovy.json.JsonSlurper
import groovy.json.JsonBuilder
import java.util.concurrent.TimeUnit

import org.sonatype.nexus.cleanup.storage.CleanupPolicyStorage
import static org.sonatype.nexus.repository.search.DefaultComponentMetadataProducer.IS_PRERELEASE_KEY
import static org.sonatype.nexus.repository.search.DefaultComponentMetadataProducer.LAST_BLOB_UPDATED_KEY
import static org.sonatype.nexus.repository.search.DefaultComponentMetadataProducer.LAST_DOWNLOADED_KEY


def cleanupPolicyStorage = container.lookup(CleanupPolicyStorage.class.getName())

try {
    parsed_args = new JsonSlurper().parseText(args)
} catch(Exception ex) {
    log.debug("list")
    def policies = []
    cleanupPolicyStorage.getAll().each {
        policies << toJsonString(it)
    }
    return policies
}

parsed_args.each {
    log.debug("Received arguments: <${it.key}=${it.value}> (${it.value.getClass()})")
}

if (parsed_args.name == null) {
    throw new Exception("Missing mandatory argument: name")
}

// "get" operation
if (parsed_args.size() == 1) {
    log.debug("get")
    existingPolicy = cleanupPolicyStorage.get(parsed_args.name)
    return toJsonString(existingPolicy)
}

// create and update use this
Map<String, String> criteriaMap = createCriteria(parsed_args)

// "update" operation
if (cleanupPolicyStorage.exists(parsed_args.name)) {
    log.debug("Updating Cleanup Policy <name=${parsed_args.name}>")
    existingPolicy = cleanupPolicyStorage.get(parsed_args.name)
    existingPolicy.setNotes(parsed_args.notes)
    existingPolicy.setCriteria(criteriaMap)
    cleanupPolicyStorage.update(existingPolicy)
    return toJsonString(existingPolicy)
}

// "create" operation
format = parsed_args.format == "all" ? "ALL_FORMATS" : parsed_args.format

log.debug("Creating Cleanup Policy <name=${parsed_args.name}>")
cleanupPolicy = cleanupPolicyStorage.newCleanupPolicy()

log.debug("Configuring Cleanup Policy <policy=${cleanupPolicy}>")
cleanupPolicy.setName(parsed_args.name)
cleanupPolicy.setNotes(parsed_args.notes)
cleanupPolicy.setFormat(format)
cleanupPolicy.setMode('delete')
cleanupPolicy.setCriteria(criteriaMap)

log.debug("Adding Cleanup Policy <policy=${cleanupPolicy}>")
cleanupPolicyStorage.add(cleanupPolicy)
return toJsonString(cleanupPolicy)


def Map<String, String> createCriteria(parsed_args) {
    Map<String, String> criteriaMap = Maps.newHashMap()
    if (parsed_args.criteria.lastBlobUpdated == null) {
        criteriaMap.remove(LAST_BLOB_UPDATED_KEY)
    } else {
        criteriaMap.put(LAST_BLOB_UPDATED_KEY, asStringSeconds(parsed_args.criteria.lastBlobUpdated))
    }
    if (parsed_args.criteria.lastDownloaded == null) {
        criteriaMap.remove(LAST_DOWNLOADED_KEY)
    } else {
        criteriaMap.put(LAST_DOWNLOADED_KEY, asStringSeconds(parsed_args.criteria.lastDownloaded))
    }
    log.debug("Using criteriaMap: ${criteriaMap}")

    return criteriaMap
}

def Integer asSeconds(days) {
    return days * TimeUnit.DAYS.toSeconds(1)
}

def String asStringSeconds(daysInt) {
    return String.valueOf(asSeconds(daysInt))
}

// There's got to be a better way to do this.
// using JsonOutput directly on the object causes a stack overflow
def String toJsonString(cleanup_policy) {
    def policyString = new JsonBuilder()
    policyString {
        name cleanup_policy.getName()
        notes cleanup_policy.getNotes()
        format cleanup_policy.getFormat()
        mode cleanup_policy.getMode()
        criteria cleanup_policy.getCriteria()
    }
    return policyString.toPrettyString()
}
`
	return nexusSchema.Script{
		Name:    name,
		Content: content,
		Type:    "groovy",
	}
}

func resourceCleanUpPolicyCreate(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)
	cu := getCleanUpPolicyFromResourceData(resourceData)
	script := NewCleanUpScript(cu.Name, cu.Format, cu.Notes, cu.Criteria)
	//create script
	if err := client.Script.Create(&script); err != nil {
		return err
	}
	//	run script

	payload := GetPayload(cu.Name, cu.Format, cu.Notes, cu.Criteria.LastDownloaded, cu.Criteria.LastBlobUpdated)
	fmt.Printf("payload: %v", payload)
	fmt.Printf("script: %v", script)
	if err := client.Script.Run(script.Name, payload); err != nil {
		return err
	}
	//cleanup policy name is equal to script name
	resourceData.SetId(cu.Name)
	return resourceCleanUpPolicyRead(resourceData, m)
}

func resourceCleanUpPolicyRead(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)

	script, err := client.Script.Get(resourceData.Id())
	if err != nil {
		return err
	}

	if script == nil {
		resourceData.SetId("")
		return nil
	}
	/**Instead of set data from what we get from nexus
	We set data to what are provided in resource data
	**/
	resourceData.Set("name", script.Name)
	cu := getCleanUpPolicyFromResourceData(resourceData)
	resourceData.Set("format", cu.Format)
	resourceData.Set("notes", cu.Notes)
	resourceData.Set("criteria", flattenCriteria(&cu.Criteria))
	return nil
}

func resourceCleanUpPolicyUpdate(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)
	// if name is changed, resource is force to recreate already
	if resourceData.HasChangeExcept("name") {
		cu := getCleanUpPolicyFromResourceData(resourceData)
		script := NewCleanUpScript(cu.Name, cu.Format, cu.Notes, cu.Criteria)
		if err := client.Script.Update(&script); err != nil {
			return err
		}
		payload := GetPayload(cu.Name, cu.Format, cu.Notes, cu.Criteria.LastDownloaded, cu.Criteria.LastBlobUpdated)
		if err := client.Script.Run(cu.Name, payload); err != nil {
			return err
		}
	}
	return resourceCleanUpPolicyRead(resourceData, m)
}

func resourceCleanUpPolicyDelete(resourceData *schema.ResourceData, m interface{}) error {
	client := m.(*nexus.NexusClient)

	if err := client.Script.Delete(resourceData.Id()); err != nil {
		return err
	}

	resourceData.SetId("")
	return nil
}

func resourceCleanUpPolicyExists(resourceData *schema.ResourceData, m interface{}) (bool, error) {
	client := m.(*nexus.NexusClient)

	script, err := client.Script.Get(resourceData.Id())
	return script != nil, err
}
