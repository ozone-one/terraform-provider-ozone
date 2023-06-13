package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the CreateWorkspaceRequest resource defined in the Terraform configuration
func CreateWorkspaceRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_member_ids": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			Optional: true,
			ForceNew: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying CreateWorkspaceRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCreateWorkspaceRequestResourceData(d *schema.ResourceData, m *models.CreateWorkspaceRequest) {
	d.Set("admin_member_ids", m.AdminMemberIds)
	d.Set("description", m.Description)
	d.Set("name", m.Name)
}

// Iterate throught and update the CreateWorkspaceRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCreateWorkspaceRequestSubResourceData(m []*models.CreateWorkspaceRequest) (d []*map[string]interface{}) {
	for _, createWorkspaceRequest := range m {
		if createWorkspaceRequest != nil {
			properties := make(map[string]interface{})
			properties["admin_member_ids"] = createWorkspaceRequest.AdminMemberIds
			properties["description"] = createWorkspaceRequest.Description
			properties["name"] = createWorkspaceRequest.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CreateWorkspaceRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CreateWorkspaceRequestModel(d *schema.ResourceData) *models.CreateWorkspaceRequest {
	adminMemberIds := d.Get("admin_member_ids").([]string)
	description := d.Get("description").(string)
	name := d.Get("name").(string)

	return &models.CreateWorkspaceRequest{
		AdminMemberIds: adminMemberIds,
		Description:    description,
		Name:           name,
	}
}

// Retrieve property field names for updating the CreateWorkspaceRequest resource
func GetCreateWorkspaceRequestPropertyFields() (t []string) {
	return []string{
		"admin_member_ids",
		"description",
		"name",
	}
}
