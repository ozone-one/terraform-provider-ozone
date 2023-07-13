package schemata

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the CreateReleaseRunRequest resource defined in the Terraform configuration
func CreateReleaseRunRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"release_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// "steps": {
		// 	Type: schema.TypeList, //GoType: []*ReleaseRunStep
		// 	Elem: &schema.Resource{
		// 		Schema: ReleaseRunStepSchema(),
		// 	},
		// 	ConfigMode: schema.SchemaConfigModeAttr,
		// 	Optional:   true,
		// 	ForceNew:   true,
		// },

		"params": {
			Type: schema.TypeList, //GoType: []*ReleaseParam
			Elem: &schema.Resource{
				Schema: ReleaseParamSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"trigger_resource_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"trigger_resource_kind": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

// Update the underlying CreateReleaseRunRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCreateReleaseRunRequestResourceData(d *schema.ResourceData, m *models.CreateReleaseRunRequest) {
	d.Set("name", m.Name)
	d.Set("release_id", m.ReleaseID)
	//d.Set("steps", SetReleaseRunStepSubResourceData(m.Steps))
	// d.Set("trigger_resource_id", m.TriggerResourceID)
	// d.Set("trigger_resource_kind", m.TriggerResourceKind)
	d.Set("params", SetReleaseParamSubResourceData(m.Params))
}

// Iterate throught and update the CreateReleaseRunRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCreateReleaseRunRequestSubResourceData(m []*models.CreateReleaseRunRequest) (d []*map[string]interface{}) {
	for _, createReleaseRunRequest := range m {
		if createReleaseRunRequest != nil {
			properties := make(map[string]interface{})
			properties["name"] = createReleaseRunRequest.Name
			properties["release_id"] = createReleaseRunRequest.ReleaseID
			//properties["steps"] = SetReleaseRunStepSubResourceData(createReleaseRunRequest.Steps)
			properties["params"] = SetReleaseParamSubResourceData(createReleaseRunRequest.Params)
			// properties["trigger_resource_id"] = createReleaseRunRequest.TriggerResourceID
			// properties["trigger_resource_kind"] = createReleaseRunRequest.TriggerResourceKind
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CreateReleaseRunRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CreateReleaseRunRequestModel(d *schema.ResourceData) *models.CreateReleaseRunRequest {
	name := d.Get("name").(string)
	releaseID := d.Get("release_id").(string)
	//steps := d.Get("steps").([]*models.ReleaseRunStep)
	// get release run steps
	releaseParams := make([]*models.ReleaseParam, 0)

	if v, ok := d.GetOk("params"); ok {
		for _, releaseParam := range v.([]interface{}) {
			params := releaseParam.(map[string]interface{})["params"]
			pipelineID := releaseParam.(map[string]interface{})["pipeline_id"].(string)
			var item = models.ReleaseParam{
				PipelineID: pipelineID,
				Params:     make([]*models.PipelineParam, 0),
			}
			for _, v := range params.([]interface{}) {
				item.Params = append(item.Params, pipelineParams(v))
			}
			releaseParams = append(releaseParams, &item)
		}
	}

	return &models.CreateReleaseRunRequest{
		Name:      name,
		ReleaseID: releaseID,
		Params:    releaseParams,
	}
}

// Retrieve property field names for updating the CreateReleaseRunRequest resource
func GetCreateReleaseRunRequestPropertyFields() (t []string) {
	return []string{
		"name",
		"release_id",
		"steps",
		"trigger_resource_id",
		"trigger_resource_kind",
	}
}

func stepParams(v interface{}) models.ReleaseRunStep {
	step := v.(map[string]interface{})

	log.Printf("stepParams %s %T", step, step)

	var item models.ReleaseRunStep
	item.Name = step["name"].(string)
	item.TypeID = int64(step["type_id"].(int))
	item.Description = step["description"].(string)
	if v := step["run_after"]; v != nil {
		item.RunAfter = make([]string, len(v.([]interface{})))
		for i, v := range v.([]interface{}) {
			item.RunAfter[i] = v.(string)
		}
	}
	// if v := step["run_after"].([]interface{}); v != nil {
	// 	item.RunAfter = make([]string, len(v))
	// 	for i, v := range v {
	// 		item.RunAfter[i] = v.(string)
	// 	}
	// }
	item.JiraApproval = jiraApprovalParam(step["jira_approval"])
	item.Pipeline = pipelineParam(step["pipeline"])
	item.SlackApproval = slackParams(step["slack_approval"])
	return item
}

func jiraApprovalParam(v interface{}) *models.ReleaseRunJiraApprovalStep {
	step, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	var item models.ReleaseRunJiraApprovalStep
	item.IssueKey = step["issue_key"].(string)
	item.ProjectID = step["project_id"].(string)
	item.ProviderID = step["provider_id"].(string)
	item.WebhookID = step["webhook_id"].(string)
	return &item
}

func pipelineParam(v interface{}) *models.ReleaseRunPipelineStep {
	steps, ok := v.([]interface{})
	if !ok {
		log.Printf("failed to convert pipeline step %s %T", v, v)
		return nil
	}
	step, ok := steps[0].(map[string]interface{})
	if !ok {
		log.Printf("failed to convert pipeline step %s %T", steps[0], steps[0])
		return nil
	}
	var item models.ReleaseRunPipelineStep
	item.PipelineID = step["pipeline_id"].(string)
	item.ClusterID = step["cluster_id"].(string)
	item.ApplicationID = step["application_id"].(string)
	item.EnvID = step["env_id"].(string)
	item.Timeout = int64(step["timeout"].(int))
	for _, v := range step["trigger_params"].([]interface{}) {
		item.TriggerParams = append(item.TriggerParams, pipelineParams(v))
	}

	return &item
}

func slackParams(v interface{}) *models.ReleaseRunSlackApprovalStep {
	slackApproval, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}
	var item models.ReleaseRunSlackApprovalStep
	item.CallbackID = slackApproval["callback_id"].(string)
	item.Ts = slackApproval["ts"].(string)
	item.WebhookID = slackApproval["webhook_id"].(string)
	return &item
}
