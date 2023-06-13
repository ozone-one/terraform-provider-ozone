package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseStepType resource defined in the Terraform configuration
func ReleaseStepTypeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"deleted_at": {
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

		"updated_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ReleaseStepType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseStepTypeResourceData(d *schema.ResourceData, m *models.ReleaseStepType) {
	d.Set("created_at", m.CreatedAt)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("updated_at", m.UpdatedAt)
}

// Iterate throught and update the ReleaseStepType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseStepTypeSubResourceData(m []*models.ReleaseStepType) (d []*map[string]interface{}) {
	for _, releaseStepType := range m {
		if releaseStepType != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = releaseStepType.CreatedAt
			properties["deleted_at"] = releaseStepType.DeletedAt
			properties["id"] = releaseStepType.ID
			properties["name"] = releaseStepType.Name
			properties["updated_at"] = releaseStepType.UpdatedAt
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseStepType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseStepTypeModel(d *schema.ResourceData) *models.ReleaseStepType {
	createdAt := d.Get("created_at").(string)
	deletedAt := d.Get("deleted_at").(string)
	id, _ := d.Get("id").(int64)
	name := d.Get("name").(string)
	updatedAt := d.Get("updated_at").(string)

	return &models.ReleaseStepType{
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
		ID:        id,
		Name:      name,
		UpdatedAt: updatedAt,
	}
}

// Retrieve property field names for updating the ReleaseStepType resource
func GetReleaseStepTypePropertyFields() (t []string) {
	return []string{
		"created_at",
		"deleted_at",
		"id",
		"name",
		"updated_at",
	}
}
