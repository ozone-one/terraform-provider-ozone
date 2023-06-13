package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRunSlackApprovalStep resource defined in the Terraform configuration
func ReleaseRunSlackApprovalStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"callback_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"ts": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"webhook_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ReleaseRunSlackApprovalStep resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseRunSlackApprovalStepResourceData(d *schema.ResourceData, m *models.ReleaseRunSlackApprovalStep) {
	d.Set("callback_id", m.CallbackID)
	d.Set("ts", m.Ts)
	d.Set("webhook_id", m.WebhookID)
}

// Iterate throught and update the ReleaseRunSlackApprovalStep resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseRunSlackApprovalStepSubResourceData(m []*models.ReleaseRunSlackApprovalStep) (d []*map[string]interface{}) {
	for _, releaseRunSlackApprovalStep := range m {
		if releaseRunSlackApprovalStep != nil {
			properties := make(map[string]interface{})
			properties["callback_id"] = releaseRunSlackApprovalStep.CallbackID
			properties["ts"] = releaseRunSlackApprovalStep.Ts
			properties["webhook_id"] = releaseRunSlackApprovalStep.WebhookID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRunSlackApprovalStep resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseRunSlackApprovalStepModel(d *schema.ResourceData) *models.ReleaseRunSlackApprovalStep {
	callbackID := d.Get("callback_id").(string)
	ts := d.Get("ts").(string)
	webhookID := d.Get("webhook_id").(string)

	return &models.ReleaseRunSlackApprovalStep{
		CallbackID: callbackID,
		Ts:         ts,
		WebhookID:  webhookID,
	}
}

// Retrieve property field names for updating the ReleaseRunSlackApprovalStep resource
func GetReleaseRunSlackApprovalStepPropertyFields() (t []string) {
	return []string{
		"callback_id",
		"ts",
		"webhook_id",
	}
}
