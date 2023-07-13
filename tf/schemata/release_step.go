package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRunStep resource defined in the Terraform configuration
func ReleaseParamSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pipeline_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"params": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: PipelineParamSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
		},
	}
}

// Update the underlying ReleaseRunStep resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseParamResourceData(d *schema.ResourceData, m *models.ReleaseParam) {
	d.Set("pipeline_id", m.PipelineID)
	d.Set("params", SetPipelineParamSubResourceData(m.Params))
}

// Iterate throught and update the ReleaseRunStep resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseParamSubResourceData(m []*models.ReleaseParam) (d []interface{}) {
	for _, releaseRunStep := range m {
		if releaseRunStep != nil {
			properties := make(map[string]interface{})
			properties["pipeline_id"] = releaseRunStep.PipelineID
			properties["params"] = SetPipelineParamSubResourceData(releaseRunStep.Params)
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRunStep resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
// func ReleaseParamModel(d *schema.ResourceData) *models.ReleaseParam {

// 	pipelineID := d.Get("pipeline_id").(string)
// 	params := pipelineParam(d)

// 	return &models.ReleaseParam{
// 		PipelineID: pipelineID,
// 		Params:     params,
// 	}
// }

// Retrieve property field names for updating the ReleaseRunStep resource
func GetReleaseParamPropertyFields() (t []string) {
	return []string{
		"pipeline_id",
		"params",
	}
}
