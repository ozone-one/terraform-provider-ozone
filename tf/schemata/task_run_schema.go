package schemata

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

// Schema mapping representing the Unstructured resource defined in the Terraform configuration
func TaskRunSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"steps": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: StepSchema(),
			},
			Optional:   true,
			ForceNew:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
		},
	}
}

func SetTaskRunResourceData(m interface{}) (result []interface{}) {
	data, err := json.Marshal(m)
	if err != nil {
		log.Printf("failed to marshal object err: %s", err.Error())
		return
	}
	var pipelineRun v1beta1.PipelineRun
	err = json.Unmarshal(data, &pipelineRun)
	if err != nil {
		log.Printf("failed to unmarhsal object err: %s", err.Error())
		return
	}
	for _, taskRun := range pipelineRun.Status.TaskRuns {
		properties := make(map[string]interface{})
		properties["name"] = taskRun.PipelineTaskName
		properties["steps"] = SetStepResourceDataArray(taskRun.Status.Steps)
		if len(taskRun.Status.Conditions) > 0 {
			properties["status"] = taskRun.Status.Conditions[0].Reason
		}

		result = append(result, properties)
	}
	return
}

/*
func SetPipelineRunSecretInjectionSubResourceData(m []*models.PipelineRunSecretInjection) (d []*map[string]interface{}) {
	for _, pipelineRunSecretInjection := range m {
		if pipelineRunSecretInjection != nil {
			properties := make(map[string]interface{})
			properties["secret_name"] = pipelineRunSecretInjection.SecretName
			d = append(d, &properties)
		}
	}
	return
}

*/
