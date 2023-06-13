package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the PipelineRunRollback resource defined in the Terraform configuration
func PipelineRunRollbackSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"pipeline_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"pipeline_run_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying PipelineRunRollback resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPipelineRunRollbackResourceData(d *schema.ResourceData, m *models.PipelineRunRollback) {
	d.Set("name", m.Name)
	d.Set("pipeline_id", m.PipelineID)
	d.Set("pipeline_run_id", m.PipelineRunID)
}

// Iterate throught and update the PipelineRunRollback resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPipelineRunRollbackSubResourceData(m []*models.PipelineRunRollback) (d []*map[string]interface{}) {
	for _, pipelineRunRollback := range m {
		if pipelineRunRollback != nil {
			properties := make(map[string]interface{})
			properties["name"] = pipelineRunRollback.Name
			properties["pipeline_id"] = pipelineRunRollback.PipelineID
			properties["pipeline_run_id"] = pipelineRunRollback.PipelineRunID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PipelineRunRollback resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PipelineRunRollbackModel(d *schema.ResourceData) *models.PipelineRunRollback {
	name := d.Get("name").(string)
	pipelineID := d.Get("pipeline_id").(string)
	pipelineRunID := d.Get("pipeline_run_id").(string)

	return &models.PipelineRunRollback{
		Name:          name,
		PipelineID:    pipelineID,
		PipelineRunID: pipelineRunID,
	}
}

// Retrieve property field names for updating the PipelineRunRollback resource
func GetPipelineRunRollbackPropertyFields() (t []string) {
	return []string{
		"name",
		"pipeline_id",
		"pipeline_run_id",
	}
}
