package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ErrorMessage resource defined in the Terraform configuration
func ErrorMessageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"message": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ErrorMessage resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetErrorMessageResourceData(d *schema.ResourceData, m *models.ErrorMessage) {
	d.Set("message", m.Message)
}

// Iterate throught and update the ErrorMessage resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetErrorMessageSubResourceData(m []*models.ErrorMessage) (d []*map[string]interface{}) {
	for _, errorMessage := range m {
		if errorMessage != nil {
			properties := make(map[string]interface{})
			properties["message"] = errorMessage.Message
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ErrorMessage resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ErrorMessageModel(d *schema.ResourceData) *models.ErrorMessage {
	message := d.Get("message").(string)

	return &models.ErrorMessage{
		Message: message,
	}
}

// Retrieve property field names for updating the ErrorMessage resource
func GetErrorMessagePropertyFields() (t []string) {
	return []string{
		"message",
	}
}
