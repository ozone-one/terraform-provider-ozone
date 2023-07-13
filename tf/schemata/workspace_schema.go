package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the Workspace resource defined in the Terraform configuration
func WorkspaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

// Update the underlying Workspace resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetWorkspaceResourceData(d *schema.ResourceData, m *models.Workspace) {
	d.Set("created_at", m.CreatedAt)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

// Iterate throught and update the Workspace resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetWorkspaceSubResourceData(m []*models.Workspace) (d []*map[string]interface{}) {
	for _, workspace := range m {
		if workspace != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = workspace.CreatedAt
			properties["description"] = workspace.Description
			properties["id"] = workspace.ID
			properties["name"] = workspace.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Workspace resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func WorkspaceModel(d *schema.ResourceData) *models.Workspace {
	createdAt := d.Get("created_at").(string)
	description := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name := d.Get("name").(string)

	return &models.Workspace{
		CreatedAt:   createdAt,
		Description: description,
		ID:          id,
		Name:        name,
	}
}

// Retrieve property field names for updating the Workspace resource
func GetWorkspacePropertyFields() (t []string) {
	return []string{
		"created_at",
		"description",
		"id",
		"name",
	}
}
