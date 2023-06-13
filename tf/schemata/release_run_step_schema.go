package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRunStep resource defined in the Terraform configuration
func ReleaseRunStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		// "id": {
		// 	Type:     schema.TypeString,
		// 	Computed: true,
		// },

		"jira_approval": {
			Type: schema.TypeList, //GoType: ReleaseRunJiraApprovalStep
			Elem: &schema.Resource{
				Schema: ReleaseRunJiraApprovalStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"pipeline": {
			Type: schema.TypeList, //GoType: ReleaseRunPipelineStep
			Elem: &schema.Resource{
				Schema: ReleaseRunPipelineStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"run_after": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"slack_approval": {
			Type: schema.TypeList, //GoType: ReleaseRunSlackApprovalStep
			Elem: &schema.Resource{
				Schema: ReleaseRunSlackApprovalStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"status": {
			Type: schema.TypeList, //GoType: AppReleaseStepStatus
			Elem: &schema.Resource{
				Schema: AppReleaseStepStatusSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"type_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"workspaces": {
			Type: schema.TypeList, //GoType: MongoResourceWorkspaces
			Elem: &schema.Resource{
				Schema: MongoResourceWorkspacesSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},
	}
}

// Update the underlying ReleaseRunStep resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseRunStepResourceData(d *schema.ResourceData, m *models.ReleaseRunStep) {
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("jira_approval", SetReleaseRunJiraApprovalStepSubResourceData([]*models.ReleaseRunJiraApprovalStep{m.JiraApproval}))
	d.Set("name", m.Name)
	d.Set("pipeline", SetReleaseRunPipelineStepSubResourceData([]*models.ReleaseRunPipelineStep{m.Pipeline}))
	d.Set("run_after", m.RunAfter)
	d.Set("slack_approval", SetReleaseRunSlackApprovalStepSubResourceData([]*models.ReleaseRunSlackApprovalStep{m.SlackApproval}))
	d.Set("status", SetAppReleaseStepStatusSubResourceData([]*models.AppReleaseStepStatus{m.Status}))
	d.Set("timeout", m.Timeout)
	d.Set("type_id", m.TypeID)
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the ReleaseRunStep resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseRunStepSubResourceData(m []*models.ReleaseRunStep) (d []*map[string]interface{}) {
	for _, releaseRunStep := range m {
		if releaseRunStep != nil {
			properties := make(map[string]interface{})
			properties["description"] = releaseRunStep.Description
			properties["id"] = releaseRunStep.ID
			properties["jira_approval"] = SetReleaseRunJiraApprovalStepSubResourceData([]*models.ReleaseRunJiraApprovalStep{releaseRunStep.JiraApproval})
			properties["name"] = releaseRunStep.Name
			properties["pipeline"] = SetReleaseRunPipelineStepSubResourceData([]*models.ReleaseRunPipelineStep{releaseRunStep.Pipeline})
			properties["run_after"] = releaseRunStep.RunAfter
			properties["slack_approval"] = SetReleaseRunSlackApprovalStepSubResourceData([]*models.ReleaseRunSlackApprovalStep{releaseRunStep.SlackApproval})
			properties["status"] = SetAppReleaseStepStatusSubResourceData([]*models.AppReleaseStepStatus{releaseRunStep.Status})
			properties["timeout"] = releaseRunStep.Timeout
			properties["type_id"] = releaseRunStep.TypeID
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{releaseRunStep.Workspaces})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRunStep resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseRunStepModel(d *schema.ResourceData) *models.ReleaseRunStep {
	description := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	var jiraApproval *models.ReleaseRunJiraApprovalStep = nil //hit complex
	//jiraApprovalInterface, jiraApprovalIsSet := d.GetOk("jira_approval")
	//if jiraApprovalIsSet {
	//	jiraApprovalMap := jiraApprovalInterface.([]interface{})[0].(map[string]interface{})
	jiraApproval = ReleaseRunJiraApprovalStepModel(d)
	//}
	name := d.Get("name").(string)
	var pipeline *models.ReleaseRunPipelineStep = nil //hit complex
	//pipelineInterface, pipelineIsSet := d.GetOk("pipeline")
	//if pipelineIsSet {
	//	pipelineMap := pipelineInterface.([]interface{})[0].(map[string]interface{})
	pipeline = ReleaseRunPipelineStepModel(d)
	//}
	runAfter := d.Get("run_after").([]string)
	var slackApproval *models.ReleaseRunSlackApprovalStep = nil //hit complex
	//slackApprovalInterface, slackApprovalIsSet := d.GetOk("slack_approval")
	//if slackApprovalIsSet {
	//	slackApprovalMap := slackApprovalInterface.([]interface{})[0].(map[string]interface{})
	slackApproval = ReleaseRunSlackApprovalStepModel(d)
	//}
	var status *models.AppReleaseStepStatus = nil //hit complex
	//statusInterface, statusIsSet := d.GetOk("status")
	//if statusIsSet {
	//	statusMap := statusInterface.([]interface{})[0].(map[string]interface{})
	status = AppReleaseStepStatusModel(d)
	//}
	timeout := int64(d.Get("timeout").(int))
	typeID := int64(d.Get("type_id").(int))
	var workspaces *models.MongoResourceWorkspaces = nil //hit complex
	//workspacesInterface, workspacesIsSet := d.GetOk("workspaces")
	//if workspacesIsSet {
	//	workspacesMap := workspacesInterface.([]interface{})[0].(map[string]interface{})
	workspaces = MongoResourceWorkspacesModel(d)
	//}

	return &models.ReleaseRunStep{
		Description:   description,
		ID:            id,
		JiraApproval:  jiraApproval,
		Name:          name,
		Pipeline:      pipeline,
		RunAfter:      runAfter,
		SlackApproval: slackApproval,
		Status:        status,
		Timeout:       timeout,
		TypeID:        typeID,
		Workspaces:    workspaces,
	}
}

// Retrieve property field names for updating the ReleaseRunStep resource
func GetReleaseRunStepPropertyFields() (t []string) {
	return []string{
		"description",
		"id",
		"jira_approval",
		"name",
		"pipeline",
		"run_after",
		"slack_approval",
		"status",
		"timeout",
		"type_id",
		"workspaces",
	}
}
