package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ElasticApplicationAPM resource defined in the Terraform configuration
func ElasticApplicationAPMSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"provider_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"service_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ElasticApplicationAPM resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetElasticApplicationAPMResourceData(d *schema.ResourceData, m *models.ElasticApplicationAPM) {
	d.Set("provider_id", m.ProviderID)
	d.Set("service_name", m.ServiceName)
}

// Iterate throught and update the ElasticApplicationAPM resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetElasticApplicationAPMSubResourceData(m []*models.ElasticApplicationAPM) (d []*map[string]interface{}) {
	for _, elasticApplicationAPM := range m {
		if elasticApplicationAPM != nil {
			properties := make(map[string]interface{})
			properties["provider_id"] = elasticApplicationAPM.ProviderID
			properties["service_name"] = elasticApplicationAPM.ServiceName
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ElasticApplicationAPM resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ElasticApplicationAPMModel(d *schema.ResourceData) *models.ElasticApplicationAPM {
	providerID := d.Get("provider_id").(string)
	serviceName := d.Get("service_name").(string)

	return &models.ElasticApplicationAPM{
		ProviderID:  providerID,
		ServiceName: serviceName,
	}
}

// Retrieve property field names for updating the ElasticApplicationAPM resource
func GetElasticApplicationAPMPropertyFields() (t []string) {
	return []string{
		"provider_id",
		"service_name",
	}
}
