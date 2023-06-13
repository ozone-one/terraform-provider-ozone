package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRunDetails resource defined in the Terraform configuration
func ReleaseRunDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"last_release_run_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"last_run": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ReleaseRunDetails resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseRunDetailsResourceData(d *schema.ResourceData, m *models.ReleaseRunDetails) {
	d.Set("last_release_run_id", m.LastReleaseRunID)
	d.Set("last_run", m.LastRun)
}

// Iterate throught and update the ReleaseRunDetails resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseRunDetailsSubResourceData(m []*models.ReleaseRunDetails) (d []*map[string]interface{}) {
	for _, releaseRunDetails := range m {
		if releaseRunDetails != nil {
			properties := make(map[string]interface{})
			properties["last_release_run_id"] = releaseRunDetails.LastReleaseRunID
			properties["last_run"] = releaseRunDetails.LastRun
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRunDetails resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseRunDetailsModel(d *schema.ResourceData) *models.ReleaseRunDetails {
	lastReleaseRunID := d.Get("last_release_run_id").(string)
	lastRun := d.Get("last_run").(string)

	return &models.ReleaseRunDetails{
		LastReleaseRunID: lastReleaseRunID,
		LastRun:          lastRun,
	}
}

// Retrieve property field names for updating the ReleaseRunDetails resource
func GetReleaseRunDetailsPropertyFields() (t []string) {
	return []string{
		"last_release_run_id",
		"last_run",
	}
}
