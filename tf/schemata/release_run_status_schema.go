package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRunStatus resource defined in the Terraform configuration
func ReleaseRunStatusSchema() map[string]*schema.Schema {
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

// Update the underlying ReleaseRunStatus resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseRunStatusResourceData(d *schema.ResourceData, m *models.ReleaseRunStatus) {
	d.Set("completion_time", m.CompletionTime)
	d.Set("created_time", m.CreatedTime)
	d.Set("message", m.Message)
	d.Set("status", m.Status)
	d.Set("updated_time", m.UpdatedTime)
}

// Iterate throught and update the ReleaseRunStatus resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseRunStatusSubResourceData(m []*models.ReleaseRunStatus) (d []*map[string]interface{}) {
	for _, releaseRunStatus := range m {
		if releaseRunStatus != nil {
			properties := make(map[string]interface{})
			properties["completion_time"] = releaseRunStatus.CompletionTime
			properties["created_time"] = releaseRunStatus.CreatedTime
			properties["message"] = releaseRunStatus.Message
			properties["status"] = releaseRunStatus.Status
			properties["updated_time"] = releaseRunStatus.UpdatedTime
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRunStatus resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseRunStatusModel(d *schema.ResourceData) *models.ReleaseRunStatus {
	completionTime := d.Get("completion_time").(string)
	createdTime := d.Get("created_time").(string)
	message := d.Get("message").(string)
	status := d.Get("status").(string)
	updatedTime := d.Get("updated_time").(string)

	return &models.ReleaseRunStatus{
		CompletionTime: completionTime,
		CreatedTime:    createdTime,
		Message:        message,
		Status:         status,
		UpdatedTime:    updatedTime,
	}
}

// Retrieve property field names for updating the ReleaseRunStatus resource
func GetReleaseRunStatusPropertyFields() (t []string) {
	return []string{
		"completion_time",
		"created_time",
		"message",
		"status",
		"updated_time",
	}
}
