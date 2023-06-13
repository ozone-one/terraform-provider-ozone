package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the CreateApplicationRequest resource defined in the Terraform configuration
func CreateApplicationRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apm": {
			Type: schema.TypeList, //GoType: ApplicationAPM
			Elem: &schema.Resource{
				Schema: ApplicationAPMSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"custom": {
			Type: schema.TypeList, //GoType: CustomApplication
			Elem: &schema.Resource{
				Schema: CustomApplicationSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"envs": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			Optional: true,
			ForceNew: true,
		},

		"ml_verification": {
			Type: schema.TypeList, //GoType: ApplicationMLVerification
			Elem: &schema.Resource{
				Schema: ApplicationMLVerificationSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"pipelines": {
			Type: schema.TypeList, //GoType: []*MongoObjectIDStruct
			Elem: &schema.Resource{
				Schema: MongoObjectIDStructSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"type_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

// Update the underlying CreateApplicationRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCreateApplicationRequestResourceData(d *schema.ResourceData, m *models.CreateApplicationRequest) {
	d.Set("apm", SetApplicationAPMSubResourceData([]*models.ApplicationAPM{m.Apm}))
	d.Set("custom", SetCustomApplicationSubResourceData([]*models.CustomApplication{m.Custom}))
	d.Set("description", m.Description)
	d.Set("envs", m.Envs)
	d.Set("ml_verification", SetApplicationMLVerificationSubResourceData([]*models.ApplicationMLVerification{m.MlVerification}))
	d.Set("name", m.Name)
	d.Set("pipelines", SetMongoObjectIDStructSubResourceData(m.Pipelines))
	d.Set("type_id", m.TypeID)
}

// Iterate throught and update the CreateApplicationRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCreateApplicationRequestSubResourceData(m []*models.CreateApplicationRequest) (d []*map[string]interface{}) {
	for _, createApplicationRequest := range m {
		if createApplicationRequest != nil {
			properties := make(map[string]interface{})
			properties["apm"] = SetApplicationAPMSubResourceData([]*models.ApplicationAPM{createApplicationRequest.Apm})
			properties["custom"] = SetCustomApplicationSubResourceData([]*models.CustomApplication{createApplicationRequest.Custom})
			properties["description"] = createApplicationRequest.Description
			properties["envs"] = createApplicationRequest.Envs
			properties["ml_verification"] = SetApplicationMLVerificationSubResourceData([]*models.ApplicationMLVerification{createApplicationRequest.MlVerification})
			properties["name"] = createApplicationRequest.Name
			properties["pipelines"] = SetMongoObjectIDStructSubResourceData(createApplicationRequest.Pipelines)
			properties["type_id"] = createApplicationRequest.TypeID
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CreateApplicationRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CreateApplicationRequestModel(d *schema.ResourceData) *models.CreateApplicationRequest {
	var apm *models.ApplicationAPM = nil //hit complex
	//apmInterface, apmIsSet := d.GetOk("apm")
	//if apmIsSet {
	//	apmMap := apmInterface.([]interface{})[0].(map[string]interface{})
	apm = ApplicationAPMModel(d)
	//}
	var custom *models.CustomApplication = nil //hit complex
	//customInterface, customIsSet := d.GetOk("custom")
	//if customIsSet {
	//	customMap := customInterface.([]interface{})[0].(map[string]interface{})
	custom = CustomApplicationModel(d)
	//}
	description := d.Get("description").(string)
	envs := d.Get("envs").([]string)
	var mlVerification *models.ApplicationMLVerification = nil //hit complex
	//mlVerificationInterface, mlVerificationIsSet := d.GetOk("ml_verification")
	//if mlVerificationIsSet {
	//	mlVerificationMap := mlVerificationInterface.([]interface{})[0].(map[string]interface{})
	mlVerification = ApplicationMLVerificationModel(d)
	//}
	name := d.Get("name").(string)
	pipelines := d.Get("pipelines").([]*models.MongoObjectIDStruct)
	typeID := int64(d.Get("type_id").(int))

	return &models.CreateApplicationRequest{
		Apm:            apm,
		Custom:         custom,
		Description:    description,
		Envs:           envs,
		MlVerification: mlVerification,
		Name:           &name,
		Pipelines:      pipelines,
		TypeID:         &typeID,
	}
}

// Retrieve property field names for updating the CreateApplicationRequest resource
func GetCreateApplicationRequestPropertyFields() (t []string) {
	return []string{
		"apm",
		"custom",
		"description",
		"envs",
		"ml_verification",
		"name",
		"pipelines",
		"type_id",
	}
}
