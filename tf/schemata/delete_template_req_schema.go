package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the DeleteTemplateReq resource defined in the Terraform configuration
func DeleteTemplateReqSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
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

// Update the underlying DeleteTemplateReq resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetDeleteTemplateReqResourceData(d *schema.ResourceData, m *models.DeleteTemplateReq) {
	d.Set("application_id", m.ApplicationID)
	d.Set("id", m.ID)
}

// Iterate throught and update the DeleteTemplateReq resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetDeleteTemplateReqSubResourceData(m []*models.DeleteTemplateReq) (d []*map[string]interface{}) {
	for _, deleteTemplateReq := range m {
		if deleteTemplateReq != nil {
			properties := make(map[string]interface{})
			properties["application_id"] = deleteTemplateReq.ApplicationID
			properties["id"] = deleteTemplateReq.ID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate DeleteTemplateReq resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func DeleteTemplateReqModel(d *schema.ResourceData) *models.DeleteTemplateReq {
	applicationID := d.Get("application_id").(string)
	id, _ := d.Get("id").(string)

	return &models.DeleteTemplateReq{
		ApplicationID: applicationID,
		ID:            id,
	}
}

// Retrieve property field names for updating the DeleteTemplateReq resource
func GetDeleteTemplateReqPropertyFields() (t []string) {
	return []string{
		"application_id",
		"id",
	}
}
