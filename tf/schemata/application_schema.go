package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the Application resource defined in the Terraform configuration
func ApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"apm": {
			Type: schema.TypeList, //GoType: ApplicationAPM
			Elem: &schema.Resource{
				Schema: ApplicationAPMSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"created_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"created_by_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"created_by_name": {
			Type:     schema.TypeString,
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

		"deleted_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"environments": {
			Type: schema.TypeList, //GoType: []*Environment
			Elem: &schema.Resource{
				Schema: EnvironmentSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"envs": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
			},
			Optional: true,
			ForceNew: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
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
			Optional: true,
			ForceNew: true,
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
			Optional: true,
			ForceNew: true,
		},

		"updated_at": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"updated_by_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"updated_by_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying Application resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetApplicationResourceData(d *schema.ResourceData, m *models.Application) {
	d.Set("account_id", m.AccountID)
	d.Set("apm", SetApplicationAPMSubResourceData([]*models.ApplicationAPM{m.Apm}))
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("custom", SetCustomApplicationSubResourceData([]*models.CustomApplication{m.Custom}))
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("environments", SetEnvironmentSubResourceData(m.Environments))
	d.Set("envs", m.Envs)
	d.Set("id", m.ID)
	d.Set("ml_verification", SetApplicationMLVerificationSubResourceData([]*models.ApplicationMLVerification{m.MlVerification}))
	d.Set("name", m.Name)
	d.Set("pipelines", SetMongoObjectIDStructSubResourceData(m.Pipelines))
	d.Set("type_id", m.TypeID)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
}

// Iterate throught and update the Application resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetApplicationSubResourceData(m []*models.Application) (d []*map[string]interface{}) {
	for _, application := range m {
		if application != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = application.AccountID
			properties["apm"] = SetApplicationAPMSubResourceData([]*models.ApplicationAPM{application.Apm})
			properties["created_at"] = application.CreatedAt
			properties["created_by_id"] = application.CreatedByID
			properties["created_by_name"] = application.CreatedByName
			properties["custom"] = SetCustomApplicationSubResourceData([]*models.CustomApplication{application.Custom})
			properties["deleted_at"] = application.DeletedAt
			properties["description"] = application.Description
			properties["environments"] = SetEnvironmentSubResourceData(application.Environments)
			properties["envs"] = application.Envs
			properties["id"] = application.ID
			properties["ml_verification"] = SetApplicationMLVerificationSubResourceData([]*models.ApplicationMLVerification{application.MlVerification})
			properties["name"] = application.Name
			properties["pipelines"] = SetMongoObjectIDStructSubResourceData(application.Pipelines)
			properties["type_id"] = application.TypeID
			properties["updated_at"] = application.UpdatedAt
			properties["updated_by_id"] = application.UpdatedByID
			properties["updated_by_name"] = application.UpdatedByName
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Application resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ApplicationModel(d *schema.ResourceData) *models.Application {
	accountID := d.Get("account_id").(string)
	var apm *models.ApplicationAPM = nil //hit complex
	//apmInterface, apmIsSet := d.GetOk("apm")
	//if apmIsSet {
	//	apmMap := apmInterface.([]interface{})[0].(map[string]interface{})
	apm = ApplicationAPMModel(d)
	//}
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	var custom *models.CustomApplication = nil //hit complex
	//customInterface, customIsSet := d.GetOk("custom")
	//if customIsSet {
	//	customMap := customInterface.([]interface{})[0].(map[string]interface{})
	custom = CustomApplicationModel(d)
	//}
	deletedAt := d.Get("deleted_at").(string)
	description := d.Get("description").(string)
	environments := d.Get("environments").([]*models.Environment)
	envs := d.Get("envs").([]string)
	id, _ := d.Get("id").(string)
	var mlVerification *models.ApplicationMLVerification = nil //hit complex
	//mlVerificationInterface, mlVerificationIsSet := d.GetOk("ml_verification")
	//if mlVerificationIsSet {
	//	mlVerificationMap := mlVerificationInterface.([]interface{})[0].(map[string]interface{})
	mlVerification = ApplicationMLVerificationModel(d)
	//}
	name := d.Get("name").(string)
	pipelines := d.Get("pipelines").([]*models.MongoObjectIDStruct)
	typeID := int64(d.Get("type_id").(int))
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)

	return &models.Application{
		AccountID:      accountID,
		Apm:            apm,
		CreatedAt:      createdAt,
		CreatedByID:    createdByID,
		CreatedByName:  createdByName,
		Custom:         custom,
		DeletedAt:      deletedAt,
		Description:    description,
		Environments:   environments,
		Envs:           envs,
		ID:             id,
		MlVerification: mlVerification,
		Name:           name,
		Pipelines:      pipelines,
		TypeID:         typeID,
		UpdatedAt:      updatedAt,
		UpdatedByID:    updatedByID,
		UpdatedByName:  updatedByName,
	}
}

// Retrieve property field names for updating the Application resource
func GetApplicationPropertyFields() (t []string) {
	return []string{
		"account_id",
		"apm",
		"created_at",
		"created_by_id",
		"created_by_name",
		"custom",
		"deleted_at",
		"description",
		"environments",
		"envs",
		"id",
		"ml_verification",
		"name",
		"pipelines",
		"type_id",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
	}
}
