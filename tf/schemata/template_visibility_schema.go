package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the TemplateVisibility resource defined in the Terraform configuration
func TemplateVisibilitySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"show_to_user": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying TemplateVisibility resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTemplateVisibilityResourceData(d *schema.ResourceData, m *models.TemplateVisibility) {
	d.Set("show_to_user", m.ShowToUser)
}

// Iterate throught and update the TemplateVisibility resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTemplateVisibilitySubResourceData(m []*models.TemplateVisibility) (d []*map[string]interface{}) {
	for _, templateVisibility := range m {
		if templateVisibility != nil {
			properties := make(map[string]interface{})
			properties["show_to_user"] = templateVisibility.ShowToUser
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate TemplateVisibility resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TemplateVisibilityModel(d *schema.ResourceData) *models.TemplateVisibility {
	showToUser := d.Get("show_to_user").(bool)

	return &models.TemplateVisibility{
		ShowToUser: showToUser,
	}
}

// Retrieve property field names for updating the TemplateVisibility resource
func GetTemplateVisibilityPropertyFields() (t []string) {
	return []string{
		"show_to_user",
	}
}
