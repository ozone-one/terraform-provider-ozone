package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the CreateAppTemplateReq resource defined in the Terraform configuration
func CreateAppTemplateReqSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"bootstrap_value": {
			Type: schema.TypeList, //GoType: TemplateBootstrapValues
			Elem: &schema.Resource{
				Schema: TemplateBootstrapValuesSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying CreateAppTemplateReq resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCreateAppTemplateReqResourceData(d *schema.ResourceData, m *models.CreateAppTemplateReq) {
	d.Set("application_id", m.ApplicationID)
	d.Set("bootstrap_value", SetTemplateBootstrapValuesSubResourceData([]*models.TemplateBootstrapValues{m.BootstrapValue}))
	d.Set("name", m.Name)
}

// Iterate throught and update the CreateAppTemplateReq resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCreateAppTemplateReqSubResourceData(m []*models.CreateAppTemplateReq) (d []*map[string]interface{}) {
	for _, createAppTemplateReq := range m {
		if createAppTemplateReq != nil {
			properties := make(map[string]interface{})
			properties["application_id"] = createAppTemplateReq.ApplicationID
			properties["bootstrap_value"] = SetTemplateBootstrapValuesSubResourceData([]*models.TemplateBootstrapValues{createAppTemplateReq.BootstrapValue})
			properties["name"] = createAppTemplateReq.Name
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CreateAppTemplateReq resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CreateAppTemplateReqModel(d *schema.ResourceData) *models.CreateAppTemplateReq {
	applicationID := d.Get("application_id").(string)
	var bootstrapValue *models.TemplateBootstrapValues = nil //hit complex
	//bootstrapValueInterface, bootstrapValueIsSet := d.GetOk("bootstrap_value")
	//if bootstrapValueIsSet {
	//	bootstrapValueMap := bootstrapValueInterface.([]interface{})[0].(map[string]interface{})
	bootstrapValue = TemplateBootstrapValuesModel(d)
	//}
	name := d.Get("name").(string)

	return &models.CreateAppTemplateReq{
		ApplicationID:  applicationID,
		BootstrapValue: bootstrapValue,
		Name:           name,
	}
}

// Retrieve property field names for updating the CreateAppTemplateReq resource
func GetCreateAppTemplateReqPropertyFields() (t []string) {
	return []string{
		"application_id",
		"bootstrap_value",
		"name",
	}
}
