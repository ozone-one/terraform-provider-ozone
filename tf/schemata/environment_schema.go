package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the Environment resource defined in the Terraform configuration
func EnvironmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"workspace_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying Environment resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetEnvironmentResourceData(d *schema.ResourceData, m *models.Environment) {
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

// Iterate throught and update the Environment resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetEnvironmentSubResourceData(m []*models.Environment) (d []*map[string]interface{}) {
	for _, environment := range m {
		if environment != nil {
			properties := make(map[string]interface{})
			properties["id"] = environment.ID
			properties["name"] = environment.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Environment resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func EnvironmentModel(d *schema.ResourceData) *models.Environment {
	id, _ := d.Get("id").(string)
	name := d.Get("name").(string)

	return &models.Environment{
		ID:   id,
		Name: name,
	}
}

// Retrieve property field names for updating the Environment resource
func GetEnvironmentPropertyFields() (t []string) {
	return []string{
		"id",
		"name",
	}
}
