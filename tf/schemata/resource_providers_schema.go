package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ResourceProviders resource defined in the Terraform configuration
func ResourceProvidersSchema() map[string]*schema.Schema {
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

// Update the underlying ResourceProviders resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetResourceProvidersResourceData(d *schema.ResourceData, m *models.ResourceProviders) {
	d.Set("items", m.Items)
}

// Iterate throught and update the ResourceProviders resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetResourceProvidersSubResourceData(m []*models.ResourceProviders) (d []*map[string]interface{}) {
	for _, resourceProviders := range m {
		if resourceProviders != nil {
			properties := make(map[string]interface{})
			properties["items"] = resourceProviders.Items
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ResourceProviders resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ResourceProvidersModel(d *schema.ResourceData) *models.ResourceProviders {
	items := d.Get("items").([]string)

	return &models.ResourceProviders{
		Items: items,
	}
}

// Retrieve property field names for updating the ResourceProviders resource
func GetResourceProvidersPropertyFields() (t []string) {
	return []string{
		"items",
	}
}
