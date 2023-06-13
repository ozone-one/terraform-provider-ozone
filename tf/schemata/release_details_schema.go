package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseDetails resource defined in the Terraform configuration
func ReleaseDetailsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ReleaseDetails resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseDetailsResourceData(d *schema.ResourceData, m *models.ReleaseDetails) {
	d.Set("name", m.Name)
}

// Iterate throught and update the ReleaseDetails resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseDetailsSubResourceData(m []*models.ReleaseDetails) (d []*map[string]interface{}) {
	for _, releaseDetails := range m {
		if releaseDetails != nil {
			properties := make(map[string]interface{})
			properties["name"] = releaseDetails.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseDetails resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseDetailsModel(d *schema.ResourceData) *models.ReleaseDetails {
	name := d.Get("name").(string)

	return &models.ReleaseDetails{
		Name: name,
	}
}

// Retrieve property field names for updating the ReleaseDetails resource
func GetReleaseDetailsPropertyFields() (t []string) {
	return []string{
		"name",
	}
}
