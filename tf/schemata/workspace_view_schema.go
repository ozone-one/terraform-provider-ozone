package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the WorkspaceView resource defined in the Terraform configuration
func WorkspaceViewSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by_id": {
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
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying WorkspaceView resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetWorkspaceViewResourceData(d *schema.ResourceData, m *models.WorkspaceView) {
	d.Set("created_by_id", m.CreatedByID)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
}

// Iterate throught and update the WorkspaceView resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetWorkspaceViewSubResourceData(m []*models.WorkspaceView) (d []*map[string]interface{}) {
	for _, workspaceView := range m {
		if workspaceView != nil {
			properties := make(map[string]interface{})
			properties["created_by_id"] = workspaceView.CreatedByID
			properties["description"] = workspaceView.Description
			properties["id"] = workspaceView.ID
			properties["name"] = workspaceView.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate WorkspaceView resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func WorkspaceViewModel(d *schema.ResourceData) *models.WorkspaceView {
	createdByID := d.Get("created_by_id").(string)
	description := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name := d.Get("name").(string)

	return &models.WorkspaceView{
		CreatedByID: createdByID,
		Description: description,
		ID:          id,
		Name:        name,
	}
}

// Retrieve property field names for updating the WorkspaceView resource
func GetWorkspaceViewPropertyFields() (t []string) {
	return []string{
		"created_by_id",
		"description",
		"id",
		"name",
	}
}
