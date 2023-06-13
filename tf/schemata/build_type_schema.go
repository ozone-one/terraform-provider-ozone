package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the BuildType resource defined in the Terraform configuration
func BuildTypeSchema() map[string]*schema.Schema {
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

// Update the underlying BuildType resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetBuildTypeResourceData(d *schema.ResourceData, m *models.BuildType) {
	d.Set("created_at", m.CreatedAt)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("updated_at", m.UpdatedAt)
}

// Iterate throught and update the BuildType resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetBuildTypeSubResourceData(m []*models.BuildType) (d []*map[string]interface{}) {
	for _, buildType := range m {
		if buildType != nil {
			properties := make(map[string]interface{})
			properties["created_at"] = buildType.CreatedAt
			properties["deleted_at"] = buildType.DeletedAt
			properties["id"] = buildType.ID
			properties["name"] = buildType.Name
			properties["updated_at"] = buildType.UpdatedAt
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate BuildType resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func BuildTypeModel(d *schema.ResourceData) *models.BuildType {
	createdAt := d.Get("created_at").(string)
	deletedAt := d.Get("deleted_at").(string)
	id, _ := d.Get("id").(int64)
	name := d.Get("name").(string)
	updatedAt := d.Get("updated_at").(string)

	return &models.BuildType{
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
		ID:        id,
		Name:      name,
		UpdatedAt: updatedAt,
	}
}

// Retrieve property field names for updating the BuildType resource
func GetBuildTypePropertyFields() (t []string) {
	return []string{
		"created_at",
		"deleted_at",
		"id",
		"name",
		"updated_at",
	}
}
