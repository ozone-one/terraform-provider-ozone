package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the MongoResourceWorkspaces resource defined in the Terraform configuration
func MongoResourceWorkspacesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"items": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},
	}
}

// Update the underlying MongoResourceWorkspaces resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetMongoResourceWorkspacesResourceData(d *schema.ResourceData, m *models.MongoResourceWorkspaces) {
	d.Set("items", m.Items)
}

// Iterate throught and update the MongoResourceWorkspaces resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetMongoResourceWorkspacesSubResourceData(m []*models.MongoResourceWorkspaces) (d []*map[string]interface{}) {
	for _, mongoResourceWorkspaces := range m {
		if mongoResourceWorkspaces != nil {
			properties := make(map[string]interface{})
			properties["items"] = mongoResourceWorkspaces.Items
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate MongoResourceWorkspaces resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func MongoResourceWorkspacesModel(d *schema.ResourceData) *models.MongoResourceWorkspaces {
	items := d.Get("items").([]string)

	return &models.MongoResourceWorkspaces{
		Items: items,
	}
}

// Retrieve property field names for updating the MongoResourceWorkspaces resource
func GetMongoResourceWorkspacesPropertyFields() (t []string) {
	return []string{
		"items",
	}
}
