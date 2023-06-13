package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the AppReleaseStepStatus resource defined in the Terraform configuration
func AppReleaseStepStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"completion_time": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"created_time": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"message": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"status": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"updated_time": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying AppReleaseStepStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppReleaseStepStatusResourceData(d *schema.ResourceData, m *models.AppReleaseStepStatus) {
	d.Set("completion_time", m.CompletionTime)
	d.Set("created_time", m.CreatedTime)
	d.Set("message", m.Message)
	d.Set("status", m.Status)
	d.Set("updated_time", m.UpdatedTime)
}

// Iterate throught and update the AppReleaseStepStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppReleaseStepStatusSubResourceData(m []*models.AppReleaseStepStatus) (d []*map[string]interface{}) {
	for _, appReleaseStepStatus := range m {
		if appReleaseStepStatus != nil {
			properties := make(map[string]interface{})
			properties["completion_time"] = appReleaseStepStatus.CompletionTime
			properties["created_time"] = appReleaseStepStatus.CreatedTime
			properties["message"] = appReleaseStepStatus.Message
			properties["status"] = appReleaseStepStatus.Status
			properties["updated_time"] = appReleaseStepStatus.UpdatedTime
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AppReleaseStepStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppReleaseStepStatusModel(d *schema.ResourceData) *models.AppReleaseStepStatus {
	completionTime := d.Get("completion_time").(string)
	createdTime := d.Get("created_time").(string)
	message := d.Get("message").(string)
	status := d.Get("status").(string)
	updatedTime := d.Get("updated_time").(string)

	return &models.AppReleaseStepStatus{
		CompletionTime: completionTime,
		CreatedTime:    createdTime,
		Message:        message,
		Status:         status,
		UpdatedTime:    updatedTime,
	}
}

// Retrieve property field names for updating the AppReleaseStepStatus resource
func GetAppReleaseStepStatusPropertyFields() (t []string) {
	return []string{
		"completion_time",
		"created_time",
		"message",
		"status",
		"updated_time",
	}
}
