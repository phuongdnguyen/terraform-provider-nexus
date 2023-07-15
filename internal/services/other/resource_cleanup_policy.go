package other

import (
	nexus "github.com/datadrivers/go-nexus-client/nexus3"
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
						"last_downloaded": {
							Type: schema.TypeInt,
						},
						"last_blob_updated": {
							Type: schema.TypeInt,
						},
						"regex": {
							Type: schema.TypeString,
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
	Last_downloaded   int
	Last_blob_updated int
	Regex             string
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
			Last_downloaded:   rConfig["last_downloaded"].(int),
			Last_blob_updated: rConfig["last_blob_updated"].(int),
			Regex:             rConfig["regex"].(string),
		}
	}
	return c
}

var script = `// Original from:
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

func resourceCleanUpPolicyCreate(resourceData *schema.ResourceData, m interface{}) error {
	nexusClient := m.(*nexus.NexusClient)
	cu := getCleanUpPolicyFromResourceData(resourceData)

}

func resourceCleanUpPolicyRead(resourceData *schema.ResourceData, m interface{}) error {

}

func resourceCleanUpPolicyUpdate(resourceData *schema.ResourceData, m interface{}) error {

}

func resourceCleanUpPolicyDelete(resourceData *schema.ResourceData, m interface{}) error {

}

func resourceCleanUpPolicyExists(resourceData *schema.ResourceData, m interface{}) (bool, error) {

}
