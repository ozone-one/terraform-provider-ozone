package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the UnstructuredUnstructured resource defined in the Terraform configuration
func UnstructuredUnstructuredSchema() map[string]*schema.Schema {
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

// Update the underlying UnstructuredUnstructured resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetUnstructuredUnstructuredResourceData(d *schema.ResourceData, m *models.UnstructuredUnstructured) {
	d.Set("object", m.Object)
}

// Iterate throught and update the UnstructuredUnstructured resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetUnstructuredUnstructuredSubResourceData(m []*models.UnstructuredUnstructured) (d []*map[string]interface{}) {
	for _, unstructuredUnstructured := range m {
		if unstructuredUnstructured != nil {
			properties := make(map[string]interface{})
			properties["object"] = unstructuredUnstructured.Object
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate UnstructuredUnstructured resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func UnstructuredUnstructuredModel(d *schema.ResourceData) *models.UnstructuredUnstructured {
	object := d.Get("object")

	return &models.UnstructuredUnstructured{
		Object: object,
	}
}

// Retrieve property field names for updating the UnstructuredUnstructured resource
func GetUnstructuredUnstructuredPropertyFields() (t []string) {
	return []string{
		"object",
	}
}
