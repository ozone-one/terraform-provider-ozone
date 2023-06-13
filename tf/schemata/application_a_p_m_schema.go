package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ApplicationAPM resource defined in the Terraform configuration
func ApplicationAPMSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apm_type": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"datadog": {
			Type: schema.TypeList, //GoType: DatadogApplicationAPM
			Elem: &schema.Resource{
				Schema: DatadogApplicationAPMSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"dynatrace": {
			Type: schema.TypeList, //GoType: DynatraceApplicationAPM
			Elem: &schema.Resource{
				Schema: DynatraceApplicationAPMSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"elastic": {
			Type: schema.TypeList, //GoType: ElasticApplicationAPM
			Elem: &schema.Resource{
				Schema: ElasticApplicationAPMSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},

		"newrelic": {
			Type: schema.TypeList, //GoType: NewRelicApplicationAPM
			Elem: &schema.Resource{
				Schema: NewRelicApplicationAPMSchema(),
			},
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ApplicationAPM resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetApplicationAPMResourceData(d *schema.ResourceData, m *models.ApplicationAPM) {
	d.Set("apm_type", m.ApmType)
	d.Set("datadog", SetDatadogApplicationAPMSubResourceData([]*models.DatadogApplicationAPM{m.Datadog}))
	d.Set("dynatrace", SetDynatraceApplicationAPMSubResourceData([]*models.DynatraceApplicationAPM{m.Dynatrace}))
	d.Set("elastic", SetElasticApplicationAPMSubResourceData([]*models.ElasticApplicationAPM{m.Elastic}))
	d.Set("enabled", m.Enabled)
	d.Set("newrelic", SetNewRelicApplicationAPMSubResourceData([]*models.NewRelicApplicationAPM{m.Newrelic}))
}

// Iterate throught and update the ApplicationAPM resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetApplicationAPMSubResourceData(m []*models.ApplicationAPM) (d []*map[string]interface{}) {
	for _, applicationAPM := range m {
		if applicationAPM != nil {
			properties := make(map[string]interface{})
			properties["apm_type"] = applicationAPM.ApmType
			properties["datadog"] = SetDatadogApplicationAPMSubResourceData([]*models.DatadogApplicationAPM{applicationAPM.Datadog})
			properties["dynatrace"] = SetDynatraceApplicationAPMSubResourceData([]*models.DynatraceApplicationAPM{applicationAPM.Dynatrace})
			properties["elastic"] = SetElasticApplicationAPMSubResourceData([]*models.ElasticApplicationAPM{applicationAPM.Elastic})
			properties["enabled"] = applicationAPM.Enabled
			properties["newrelic"] = SetNewRelicApplicationAPMSubResourceData([]*models.NewRelicApplicationAPM{applicationAPM.Newrelic})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ApplicationAPM resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ApplicationAPMModel(d *schema.ResourceData) *models.ApplicationAPM {
	apmType := int64(d.Get("apm_type").(int))
	var datadog *models.DatadogApplicationAPM = nil //hit complex
	//datadogInterface, datadogIsSet := d.GetOk("datadog")
	//if datadogIsSet {
	//	datadogMap := datadogInterface.([]interface{})[0].(map[string]interface{})
	datadog = DatadogApplicationAPMModel(d)
	//}
	var dynatrace *models.DynatraceApplicationAPM = nil //hit complex
	//dynatraceInterface, dynatraceIsSet := d.GetOk("dynatrace")
	//if dynatraceIsSet {
	//	dynatraceMap := dynatraceInterface.([]interface{})[0].(map[string]interface{})
	dynatrace = DynatraceApplicationAPMModel(d)
	//}
	var elastic *models.ElasticApplicationAPM = nil //hit complex
	//elasticInterface, elasticIsSet := d.GetOk("elastic")
	//if elasticIsSet {
	//	elasticMap := elasticInterface.([]interface{})[0].(map[string]interface{})
	elastic = ElasticApplicationAPMModel(d)
	//}
	enabled := d.Get("enabled").(bool)
	var newrelic *models.NewRelicApplicationAPM = nil //hit complex
	//newrelicInterface, newrelicIsSet := d.GetOk("newrelic")
	//if newrelicIsSet {
	//	newrelicMap := newrelicInterface.([]interface{})[0].(map[string]interface{})
	newrelic = NewRelicApplicationAPMModel(d)
	//}

	return &models.ApplicationAPM{
		ApmType:   apmType,
		Datadog:   datadog,
		Dynatrace: dynatrace,
		Elastic:   elastic,
		Enabled:   enabled,
		Newrelic:  newrelic,
	}
}

// Retrieve property field names for updating the ApplicationAPM resource
func GetApplicationAPMPropertyFields() (t []string) {
	return []string{
		"apm_type",
		"datadog",
		"dynatrace",
		"elastic",
		"enabled",
		"newrelic",
	}
}
