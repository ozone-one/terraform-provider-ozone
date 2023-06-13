package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ApplicationTemplateItem resource defined in the Terraform configuration
func ApplicationTemplateItemSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"show_to_user": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},

		"yaml_string": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ApplicationTemplateItem resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetApplicationTemplateItemResourceData(d *schema.ResourceData, m *models.ApplicationTemplateItem) {
	d.Set("name", m.Name)
	d.Set("show_to_user", m.ShowToUser)
	d.Set("yaml_string", m.YamlString)
}

// Iterate throught and update the ApplicationTemplateItem resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetApplicationTemplateItemSubResourceData(m []*models.ApplicationTemplateItem) (d []*map[string]interface{}) {
	for _, applicationTemplateItem := range m {
		if applicationTemplateItem != nil {
			properties := make(map[string]interface{})
			properties["name"] = applicationTemplateItem.Name
			properties["show_to_user"] = applicationTemplateItem.ShowToUser
			properties["yaml_string"] = applicationTemplateItem.YamlString
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ApplicationTemplateItem resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ApplicationTemplateItemModel(d *schema.ResourceData) *models.ApplicationTemplateItem {
	name := d.Get("name").(string)
	showToUser := d.Get("show_to_user").(bool)
	yamlString := d.Get("yaml_string").(string)

	return &models.ApplicationTemplateItem{
		Name:       name,
		ShowToUser: showToUser,
		YamlString: yamlString,
	}
}

// Retrieve property field names for updating the ApplicationTemplateItem resource
func GetApplicationTemplateItemPropertyFields() (t []string) {
	return []string{
		"name",
		"show_to_user",
		"yaml_string",
	}
}
