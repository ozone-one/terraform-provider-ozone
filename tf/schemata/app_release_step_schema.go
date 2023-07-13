package schemata

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the AppReleaseStep resource defined in the Terraform configuration
func AppReleaseStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"jira_approval": {
			Type: schema.TypeList, //GoType: ReleaseRunJiraApprovalStep
			Elem: &schema.Resource{
				Schema: ReleaseRunJiraApprovalStepSchema(),
			},
			Optional:   true,
			ForceNew:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"run_after": {
			Type:       schema.TypeList, //GoType: []string
			ConfigMode: schema.SchemaConfigModeAttr,
			Elem:       &schema.Schema{Type: schema.TypeString},
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
		"pipeline": {
			Type: schema.TypeSet, //GoType: ReleaseRunPipelineStep
			Elem: &schema.Resource{
				Schema: ReleaseRunPipelineStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},
	}
}

// Update the underlying AppReleaseStep resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppReleaseStepResourceData(d *schema.ResourceData, m *models.AppReleaseStep) {
	d.Set("description", m.Description)
	d.Set("jira_approval", SetReleaseRunJiraApprovalStepSubResourceData([]*models.ReleaseRunJiraApprovalStep{m.JiraApproval}))
	d.Set("name", m.Name)
	d.Set("run_after", m.RunAfter)
	d.Set("slack_approval", SetReleaseRunSlackApprovalStepSubResourceData([]*models.ReleaseRunSlackApprovalStep{m.SlackApproval}))
	d.Set("type_id", m.TypeID)
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the AppReleaseStep resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppReleaseStepSubResourceData(m []*models.AppReleaseStep) (result []interface{}) {
	for _, appReleaseStep := range m {
		if appReleaseStep != nil {
			log.Printf("[DEBUG] Updating app_release_step subresource data for: %s", appReleaseStep.Name)
			properties := make(map[string]interface{})
			properties["description"] = appReleaseStep.Description
			properties["jira_approval"] = SetReleaseRunJiraApprovalStepSubResourceData([]*models.ReleaseRunJiraApprovalStep{appReleaseStep.JiraApproval})
			properties["name"] = appReleaseStep.Name
			properties["run_after"] = appReleaseStep.RunAfter
			properties["slack_approval"] = SetReleaseRunSlackApprovalStepSubResourceData([]*models.ReleaseRunSlackApprovalStep{appReleaseStep.SlackApproval})
			properties["type_id"] = appReleaseStep.TypeID
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{appReleaseStep.Workspaces})
			if appReleaseStep.Pipeline != nil {
				properties["pipeline"] = SetReleaseRunPipelineStepSubResourceData([]*models.ReleaseRunPipelineStep{appReleaseStep.Pipeline})
			}
			result = append(result, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AppReleaseStep resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppReleaseStepModel(d *schema.ResourceData) *models.AppReleaseStep {
	description := d.Get("description").(string)
	var jiraApproval *models.ReleaseRunJiraApprovalStep = nil //hit complex
	//jiraApprovalInterface, jiraApprovalIsSet := d.GetOk("jira_approval")
	//if jiraApprovalIsSet {
	//	jiraApprovalMap := jiraApprovalInterface.([]interface{})[0].(map[string]interface{})
	jiraApproval = ReleaseRunJiraApprovalStepModel(d)
	//}
	name := d.Get("name").(string)
	runAfter := d.Get("run_after").([]string)
	var slackApproval *models.ReleaseRunSlackApprovalStep = nil //hit complex
	//slackApprovalInterface, slackApprovalIsSet := d.GetOk("slack_approval")
	//if slackApprovalIsSet {
	//	slackApprovalMap := slackApprovalInterface.([]interface{})[0].(map[string]interface{})
	slackApproval = ReleaseRunSlackApprovalStepModel(d)
	//}
	typeID := int64(d.Get("type_id").(int))
	var workspaces *models.MongoResourceWorkspaces = nil //hit complex
	//workspacesInterface, workspacesIsSet := d.GetOk("workspaces")
	//if workspacesIsSet {
	//	workspacesMap := workspacesInterface.([]interface{})[0].(map[string]interface{})
	workspaces = MongoResourceWorkspacesModel(d)
	//}

	return &models.AppReleaseStep{
		Description:   description,
		JiraApproval:  jiraApproval,
		Name:          name,
		RunAfter:      runAfter,
		SlackApproval: slackApproval,
		TypeID:        typeID,
		Workspaces:    workspaces,
	}
}

// Retrieve property field names for updating the AppReleaseStep resource
func GetAppReleaseStepPropertyFields() (t []string) {
	return []string{
		"description",
		"jira_approval",
		"name",
		"run_after",
		"slack_approval",
		"type_id",
		"workspaces",
	}
}
