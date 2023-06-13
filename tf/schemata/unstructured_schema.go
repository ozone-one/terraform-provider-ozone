package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the Unstructured resource defined in the Terraform configuration
func UnstructuredSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"object": {
			Type: schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying Unstructured resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetUnstructuredResourceData(d *schema.ResourceData, m *models.Unstructured) {
	d.Set("object", m.Object)
}

// Iterate throught and update the Unstructured resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetUnstructuredSubResourceData(m []*models.Unstructured) (d []*map[string]interface{}) {
	for _, unstructured := range m {
		if unstructured != nil {
			properties := make(map[string]interface{})
			properties["object"] = unstructured.Object
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Unstructured resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func UnstructuredModel(d *schema.ResourceData) *models.Unstructured {
	object := d.Get("object")

	return &models.Unstructured{
		Object: object,
	}
}

// Retrieve property field names for updating the Unstructured resource
func GetUnstructuredPropertyFields() (t []string) {
	return []string{
		"object",
	}
}
