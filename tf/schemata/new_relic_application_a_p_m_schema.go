package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the NewRelicApplicationAPM resource defined in the Terraform configuration
func NewRelicApplicationAPMSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"env_app_name": {
			Type: schema.TypeList, //GoType: map[string]string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			Optional: true,
			ForceNew: true,
		},

		"provider_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying NewRelicApplicationAPM resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetNewRelicApplicationAPMResourceData(d *schema.ResourceData, m *models.NewRelicApplicationAPM) {
	d.Set("env_app_name", m.EnvAppName)
	d.Set("provider_id", m.ProviderID)
}

// Iterate throught and update the NewRelicApplicationAPM resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetNewRelicApplicationAPMSubResourceData(m []*models.NewRelicApplicationAPM) (d []*map[string]interface{}) {
	for _, newRelicApplicationAPM := range m {
		if newRelicApplicationAPM != nil {
			properties := make(map[string]interface{})
			properties["env_app_name"] = newRelicApplicationAPM.EnvAppName
			properties["provider_id"] = newRelicApplicationAPM.ProviderID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate NewRelicApplicationAPM resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func NewRelicApplicationAPMModel(d *schema.ResourceData) *models.NewRelicApplicationAPM {
	envAppName := d.Get("env_app_name").(map[string]string)
	providerID := d.Get("provider_id").(string)

	return &models.NewRelicApplicationAPM{
		EnvAppName: envAppName,
		ProviderID: providerID,
	}
}

// Retrieve property field names for updating the NewRelicApplicationAPM resource
func GetNewRelicApplicationAPMPropertyFields() (t []string) {
	return []string{
		"env_app_name",
		"provider_id",
	}
}
