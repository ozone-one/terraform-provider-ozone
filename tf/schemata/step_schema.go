package schemata

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

// Schema mapping representing the Unstructured resource defined in the Terraform configuration
func StepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"step_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

func SetStepResourceData(d *schema.ResourceData, m map[string]interface{}) {
	d.Set("step_name", m["name"])
	d.Set("status", m["status"])
}

func SetStepResourceDataArray(m []v1beta1.StepState) (d []map[string]interface{}) {
	for _, step := range m {
		var status string
		if step.Terminated != nil {
			status = step.Terminated.Reason
		} else if step.Running != nil {
			status = "running"
		} else if step.Waiting != nil {
			status = step.Waiting.Reason
		}
		log.Printf("SetStepResourceDataArray-step_name %+v--status:-%+v----", step.Name, status)
		var item = map[string]interface{}{
			"step_name": step.Name,
			"status":    status,
		}
		d = append(d, item)
	}
	return
}
