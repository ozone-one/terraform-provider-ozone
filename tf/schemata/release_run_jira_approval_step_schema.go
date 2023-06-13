package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRunJiraApprovalStep resource defined in the Terraform configuration
func ReleaseRunJiraApprovalStepSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"issue_key": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"project_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"provider_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"webhook_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ReleaseRunJiraApprovalStep resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseRunJiraApprovalStepResourceData(d *schema.ResourceData, m *models.ReleaseRunJiraApprovalStep) {
	d.Set("issue_key", m.IssueKey)
	d.Set("project_id", m.ProjectID)
	d.Set("provider_id", m.ProviderID)
	d.Set("webhook_id", m.WebhookID)
}

// Iterate throught and update the ReleaseRunJiraApprovalStep resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseRunJiraApprovalStepSubResourceData(m []*models.ReleaseRunJiraApprovalStep) (d []*map[string]interface{}) {
	for _, releaseRunJiraApprovalStep := range m {
		if releaseRunJiraApprovalStep != nil {
			properties := make(map[string]interface{})
			properties["issue_key"] = releaseRunJiraApprovalStep.IssueKey
			properties["project_id"] = releaseRunJiraApprovalStep.ProjectID
			properties["provider_id"] = releaseRunJiraApprovalStep.ProviderID
			properties["webhook_id"] = releaseRunJiraApprovalStep.WebhookID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRunJiraApprovalStep resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseRunJiraApprovalStepModel(d *schema.ResourceData) *models.ReleaseRunJiraApprovalStep {
	issueKey := d.Get("issue_key").(string)
	projectID := d.Get("project_id").(string)
	providerID := d.Get("provider_id").(string)
	webhookID := d.Get("webhook_id").(string)

	return &models.ReleaseRunJiraApprovalStep{
		IssueKey:   issueKey,
		ProjectID:  projectID,
		ProviderID: providerID,
		WebhookID:  webhookID,
	}
}

// Retrieve property field names for updating the ReleaseRunJiraApprovalStep resource
func GetReleaseRunJiraApprovalStepPropertyFields() (t []string) {
	return []string{
		"issue_key",
		"project_id",
		"provider_id",
		"webhook_id",
	}
}
