package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the PipelineRunMLVerification resource defined in the Terraform configuration
func PipelineRunMLVerificationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"message": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"started_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"updated_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying PipelineRunMLVerification resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPipelineRunMLVerificationResourceData(d *schema.ResourceData, m *models.PipelineRunMLVerification) {
	d.Set("message", m.Message)
	d.Set("started_at", m.StartedAt)
	d.Set("updated_at", m.UpdatedAt)
}

// Iterate throught and update the PipelineRunMLVerification resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPipelineRunMLVerificationSubResourceData(m []*models.PipelineRunMLVerification) (d []*map[string]interface{}) {
	for _, pipelineRunMLVerification := range m {
		if pipelineRunMLVerification != nil {
			properties := make(map[string]interface{})
			properties["message"] = pipelineRunMLVerification.Message
			properties["started_at"] = pipelineRunMLVerification.StartedAt
			properties["updated_at"] = pipelineRunMLVerification.UpdatedAt
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PipelineRunMLVerification resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PipelineRunMLVerificationModel(d *schema.ResourceData) *models.PipelineRunMLVerification {
	message := d.Get("message").(string)
	startedAt := d.Get("started_at").(string)
	updatedAt := d.Get("updated_at").(string)

	return &models.PipelineRunMLVerification{
		Message:   message,
		StartedAt: startedAt,
		UpdatedAt: updatedAt,
	}
}

// Retrieve property field names for updating the PipelineRunMLVerification resource
func GetPipelineRunMLVerificationPropertyFields() (t []string) {
	return []string{
		"message",
		"started_at",
		"updated_at",
	}
}
