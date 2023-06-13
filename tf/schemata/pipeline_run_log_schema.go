package schemata

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the PipelineRunLog resource defined in the Terraform configuration
func PipelineRunLogSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"message": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"timestamp": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying PipelineRunLog resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPipelineRunLogResourceData(d *schema.ResourceData, m *models.PipelineRunLog) {
	d.Set("id", m.ID)
	d.Set("message", m.Message)
	d.Set("timestamp", m.Timestamp)
}

// Iterate throught and update the PipelineRunLog resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPipelineRunLogSubResourceData(m []*models.PipelineRunLog) (d []*map[string]interface{}) {
	for _, pipelineRunLog := range m {
		if pipelineRunLog != nil {
			properties := make(map[string]interface{})
			properties["id"] = pipelineRunLog.ID
			properties["message"] = pipelineRunLog.Message
			properties["timestamp"] = pipelineRunLog.Timestamp
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PipelineRunLog resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PipelineRunLogModel(d *schema.ResourceData) *models.PipelineRunLog {
	id, _ := d.Get("id").(string)
	message := d.Get("message").(string)
	timestamp := d.Get("timestamp").(strfmt.DateTime)

	return &models.PipelineRunLog{
		ID:        id,
		Message:   message,
		Timestamp: timestamp,
	}
}

// Retrieve property field names for updating the PipelineRunLog resource
func GetPipelineRunLogPropertyFields() (t []string) {
	return []string{
		"id",
		"message",
		"timestamp",
	}
}
