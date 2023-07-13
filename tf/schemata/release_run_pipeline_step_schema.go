package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRunPipelineStep resource defined in the Terraform configuration
func ReleaseRunPipelineStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"cluster_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"env_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"pipeline_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"trigger_params": {
			Type: schema.TypeList, //GoType: []*PipelineParam
			Elem: &schema.Resource{
				Schema: PipelineParamSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},
	}
}

// Update the underlying ReleaseRunPipelineStep resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseRunPipelineStepResourceData(d *schema.ResourceData, m *models.ReleaseRunPipelineStep) {
	d.Set("application_id", m.ApplicationID)
	d.Set("cluster_id", m.ClusterID)
	d.Set("env_id", m.EnvID)
	d.Set("pipeline_id", m.PipelineID)
	d.Set("timeout", m.Timeout)
	d.Set("trigger_params", SetPipelineParamSubResourceData(m.TriggerParams))
}

// Iterate throught and update the ReleaseRunPipelineStep resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseRunPipelineStepSubResourceData(m []*models.ReleaseRunPipelineStep) (d []interface{}) {
	for _, releaseRunPipelineStep := range m {
		if releaseRunPipelineStep != nil {
			properties := make(map[string]interface{})
			properties["application_id"] = releaseRunPipelineStep.ApplicationID
			properties["cluster_id"] = releaseRunPipelineStep.ClusterID
			properties["env_id"] = releaseRunPipelineStep.EnvID
			properties["pipeline_id"] = releaseRunPipelineStep.PipelineID
			properties["timeout"] = releaseRunPipelineStep.Timeout
			properties["trigger_params"] = SetPipelineParamSubResourceData(releaseRunPipelineStep.TriggerParams)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRunPipelineStep resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseRunPipelineStepModel(d *schema.ResourceData) *models.ReleaseRunPipelineStep {
	applicationID := d.Get("application_id").(string)
	clusterID := d.Get("cluster_id").(string)
	envID := d.Get("env_id").(string)
	pipelineID := d.Get("pipeline_id").(string)
	timeout := int64(d.Get("timeout").(int))
	triggerParams := d.Get("trigger_params").([]*models.PipelineParam)

	return &models.ReleaseRunPipelineStep{
		ApplicationID: applicationID,
		ClusterID:     clusterID,
		EnvID:         envID,
		PipelineID:    pipelineID,
		Timeout:       timeout,
		TriggerParams: triggerParams,
	}
}

// Retrieve property field names for updating the ReleaseRunPipelineStep resource
func GetReleaseRunPipelineStepPropertyFields() (t []string) {
	return []string{
		"application_id",
		"cluster_id",
		"env_id",
		"pipeline_id",
		"timeout",
		"trigger_params",
	}
}
