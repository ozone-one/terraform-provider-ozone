package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the PipelineRunSecretInjection resource defined in the Terraform configuration
func PipelineRunSecretInjectionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"secret_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying PipelineRunSecretInjection resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPipelineRunSecretInjectionResourceData(d *schema.ResourceData, m *models.PipelineRunSecretInjection) {
	d.Set("secret_name", m.SecretName)
}

// Iterate throught and update the PipelineRunSecretInjection resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
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

// Function to perform the following actions:
// (1) Translate PipelineRunSecretInjection resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PipelineRunSecretInjectionModel(d *schema.ResourceData) *models.PipelineRunSecretInjection {
	secretName := d.Get("secret_name").(string)

	return &models.PipelineRunSecretInjection{
		SecretName: secretName,
	}
}

// Retrieve property field names for updating the PipelineRunSecretInjection resource
func GetPipelineRunSecretInjectionPropertyFields() (t []string) {
	return []string{
		"secret_name",
	}
}
