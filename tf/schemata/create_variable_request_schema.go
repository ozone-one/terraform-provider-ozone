package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the CreateVariableRequest resource defined in the Terraform configuration
func CreateVariableRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

		"name": {
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

// Update the underlying CreateVariableRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCreateVariableRequestResourceData(d *schema.ResourceData, m *models.CreateVariableRequest) {
	d.Set("description", m.Description)
	d.Set("env_values", m.EnvValues)
	d.Set("global", m.Global)
	d.Set("name", m.Name)
	d.Set("value", m.Value)
}

// Iterate throught and update the CreateVariableRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCreateVariableRequestSubResourceData(m []*models.CreateVariableRequest) (d []*map[string]interface{}) {
	for _, createVariableRequest := range m {
		if createVariableRequest != nil {
			properties := make(map[string]interface{})
			properties["description"] = createVariableRequest.Description
			properties["env_values"] = createVariableRequest.EnvValues
			properties["global"] = createVariableRequest.Global
			properties["name"] = createVariableRequest.Name
			properties["value"] = createVariableRequest.Value
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CreateVariableRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CreateVariableRequestModel(d *schema.ResourceData) *models.CreateVariableRequest {
	description := d.Get("description").(string)
	envValues := d.Get("env_values").(map[string]string)
	global := d.Get("global").(bool)
	name := d.Get("name").(string)
	value := d.Get("value").(string)

	return &models.CreateVariableRequest{
		Description: description,
		EnvValues:   envValues,
		Global:      global,
		Name:        name,
		Value:       value,
	}
}

// Retrieve property field names for updating the CreateVariableRequest resource
func GetCreateVariableRequestPropertyFields() (t []string) {
	return []string{
		"description",
		"env_values",
		"global",
		"name",
		"value",
	}
}
