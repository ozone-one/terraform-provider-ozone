package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the DynatraceApplicationAPM resource defined in the Terraform configuration
func DynatraceApplicationAPMSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"provider_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying DynatraceApplicationAPM resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDynatraceApplicationAPMResourceData(d *schema.ResourceData, m *models.DynatraceApplicationAPM) {
	d.Set("provider_id", m.ProviderID)
}

// Iterate throught and update the DynatraceApplicationAPM resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDynatraceApplicationAPMSubResourceData(m []*models.DynatraceApplicationAPM) (d []*map[string]interface{}) {
	for _, dynatraceApplicationAPM := range m {
		if dynatraceApplicationAPM != nil {
			properties := make(map[string]interface{})
			properties["provider_id"] = dynatraceApplicationAPM.ProviderID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DynatraceApplicationAPM resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DynatraceApplicationAPMModel(d *schema.ResourceData) *models.DynatraceApplicationAPM {
	providerID := d.Get("provider_id").(string)

	return &models.DynatraceApplicationAPM{
		ProviderID: providerID,
	}
}

// Retrieve property field names for updating the DynatraceApplicationAPM resource
func GetDynatraceApplicationAPMPropertyFields() (t []string) {
	return []string{
		"provider_id",
	}
}
