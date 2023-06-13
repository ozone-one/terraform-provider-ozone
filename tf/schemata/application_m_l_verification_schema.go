package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ApplicationMLVerification resource defined in the Terraform configuration
func ApplicationMLVerificationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},

		"enabled_envs": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			Optional: true,
			ForceNew: true,
		},

		"sensitivity": {
			Type:     schema.TypeFloat,
			Optional: true,
			ForceNew: true,
		},

		"wait_time_min": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ApplicationMLVerification resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetApplicationMLVerificationResourceData(d *schema.ResourceData, m *models.ApplicationMLVerification) {
	d.Set("enabled", m.Enabled)
	d.Set("enabled_envs", m.EnabledEnvs)
	d.Set("sensitivity", m.Sensitivity)
	d.Set("wait_time_min", m.WaitTimeMin)
}

// Iterate throught and update the ApplicationMLVerification resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetApplicationMLVerificationSubResourceData(m []*models.ApplicationMLVerification) (d []*map[string]interface{}) {
	for _, applicationMLVerification := range m {
		if applicationMLVerification != nil {
			properties := make(map[string]interface{})
			properties["enabled"] = applicationMLVerification.Enabled
			properties["enabled_envs"] = applicationMLVerification.EnabledEnvs
			properties["sensitivity"] = applicationMLVerification.Sensitivity
			properties["wait_time_min"] = applicationMLVerification.WaitTimeMin
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ApplicationMLVerification resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ApplicationMLVerificationModel(d *schema.ResourceData) *models.ApplicationMLVerification {
	enabled := d.Get("enabled").(bool)
	enabledEnvs := d.Get("enabled_envs").([]string)
	sensitivity := float64(d.Get("sensitivity").(float64))
	waitTimeMin := int64(d.Get("wait_time_min").(int))

	return &models.ApplicationMLVerification{
		Enabled:     enabled,
		EnabledEnvs: enabledEnvs,
		Sensitivity: sensitivity,
		WaitTimeMin: waitTimeMin,
	}
}

// Retrieve property field names for updating the ApplicationMLVerification resource
func GetApplicationMLVerificationPropertyFields() (t []string) {
	return []string{
		"enabled",
		"enabled_envs",
		"sensitivity",
		"wait_time_min",
	}
}
