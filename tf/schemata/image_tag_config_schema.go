package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ImageTagConfig resource defined in the Terraform configuration
func ImageTagConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pick_latest_branch_commit": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ImageTagConfig resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetImageTagConfigResourceData(d *schema.ResourceData, m *models.ImageTagConfig) {
	d.Set("pick_latest_branch_commit", m.PickLatestBranchCommit)
}

// Iterate throught and update the ImageTagConfig resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetImageTagConfigSubResourceData(m []*models.ImageTagConfig) (d []*map[string]interface{}) {
	for _, imageTagConfig := range m {
		if imageTagConfig != nil {
			properties := make(map[string]interface{})
			properties["pick_latest_branch_commit"] = imageTagConfig.PickLatestBranchCommit
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ImageTagConfig resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ImageTagConfigModel(d *schema.ResourceData) *models.ImageTagConfig {
	pickLatestBranchCommit := d.Get("pick_latest_branch_commit").(bool)

	return &models.ImageTagConfig{
		PickLatestBranchCommit: pickLatestBranchCommit,
	}
}

// Retrieve property field names for updating the ImageTagConfig resource
func GetImageTagConfigPropertyFields() (t []string) {
	return []string{
		"pick_latest_branch_commit",
	}
}
