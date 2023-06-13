package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the VariableView resource defined in the Terraform configuration
func VariableViewSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"created_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"created_by_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"created_by_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"deleted_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"env_values": {
			Type: schema.TypeList, //GoType: map[string]string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			Optional: true,
			ForceNew: true,
		},

		"global": {
			Type:     schema.TypeBool,
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

		"updated_by_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"updated_by_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"value": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying VariableView resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetVariableViewResourceData(d *schema.ResourceData, m *models.VariableView) {
	d.Set("account_id", m.AccountID)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("env_values", m.EnvValues)
	d.Set("global", m.Global)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("value", m.Value)
}

// Iterate throught and update the VariableView resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetVariableViewSubResourceData(m []*models.VariableView) (d []*map[string]interface{}) {
	for _, variableView := range m {
		if variableView != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = variableView.AccountID
			properties["created_at"] = variableView.CreatedAt
			properties["created_by_id"] = variableView.CreatedByID
			properties["created_by_name"] = variableView.CreatedByName
			properties["deleted_at"] = variableView.DeletedAt
			properties["description"] = variableView.Description
			properties["env_values"] = variableView.EnvValues
			properties["global"] = variableView.Global
			properties["id"] = variableView.ID
			properties["name"] = variableView.Name
			properties["updated_at"] = variableView.UpdatedAt
			properties["updated_by_id"] = variableView.UpdatedByID
			properties["updated_by_name"] = variableView.UpdatedByName
			properties["value"] = variableView.Value
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate VariableView resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func VariableViewModel(d *schema.ResourceData) *models.VariableView {
	accountID := d.Get("account_id").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	description := d.Get("description").(string)
	envValues := d.Get("env_values").(map[string]string)
	global := d.Get("global").(bool)
	id, _ := d.Get("id").(string)
	name := d.Get("name").(string)
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)
	value := d.Get("value").(string)

	return &models.VariableView{
		AccountID:     accountID,
		CreatedAt:     createdAt,
		CreatedByID:   createdByID,
		CreatedByName: createdByName,
		DeletedAt:     deletedAt,
		Description:   description,
		EnvValues:     envValues,
		Global:        global,
		ID:            id,
		Name:          name,
		UpdatedAt:     updatedAt,
		UpdatedByID:   updatedByID,
		UpdatedByName: updatedByName,
		Value:         value,
	}
}

// Retrieve property field names for updating the VariableView resource
func GetVariableViewPropertyFields() (t []string) {
	return []string{
		"account_id",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"description",
		"env_values",
		"global",
		"id",
		"name",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
		"value",
	}
}
