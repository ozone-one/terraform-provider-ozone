package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the GithubComOzoneCloudCommonEntitiesApplicationApplication resource defined in the Terraform configuration
func GithubComOzoneCloudCommonEntitiesApplicationApplicationSchema() map[string]*schema.Schema {
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

		"status": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"type_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"type_name": {
			Type:     schema.TypeString,
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

		"workspaces": {
			Type: schema.TypeList, //GoType: MongoResourceWorkspaces
			Elem: &schema.Resource{
				Schema: MongoResourceWorkspacesSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},
	}
}

// Update the underlying GithubComOzoneCloudCommonEntitiesApplicationApplication resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetGithubComOzoneCloudCommonEntitiesApplicationApplicationResourceData(d *schema.ResourceData, m *models.GithubComOzoneCloudCommonEntitiesApplicationApplication) {
	d.Set("account_id", m.AccountID)
	d.Set("apm", SetApplicationAPMSubResourceData([]*models.ApplicationAPM{m.Apm}))
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("custom", SetCustomApplicationSubResourceData([]*models.CustomApplication{m.Custom}))
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("environments", SetEnvironmentSubResourceData(m.Environments))
	d.Set("id", m.ID)
	d.Set("ml_verification", SetApplicationMLVerificationSubResourceData([]*models.ApplicationMLVerification{m.MlVerification}))
	d.Set("name", m.Name)
	d.Set("status", m.Status)
	d.Set("type_id", m.TypeID)
	d.Set("type_name", m.TypeName)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the GithubComOzoneCloudCommonEntitiesApplicationApplication resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetGithubComOzoneCloudCommonEntitiesApplicationApplicationSubResourceData(m []*models.GithubComOzoneCloudCommonEntitiesApplicationApplication) (d []*map[string]interface{}) {
	for _, githubComOzoneCloudCommonEntitiesApplicationApplication := range m {
		if githubComOzoneCloudCommonEntitiesApplicationApplication != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = githubComOzoneCloudCommonEntitiesApplicationApplication.AccountID
			properties["apm"] = SetApplicationAPMSubResourceData([]*models.ApplicationAPM{githubComOzoneCloudCommonEntitiesApplicationApplication.Apm})
			properties["created_at"] = githubComOzoneCloudCommonEntitiesApplicationApplication.CreatedAt
			properties["created_by_id"] = githubComOzoneCloudCommonEntitiesApplicationApplication.CreatedByID
			properties["created_by_name"] = githubComOzoneCloudCommonEntitiesApplicationApplication.CreatedByName
			properties["custom"] = SetCustomApplicationSubResourceData([]*models.CustomApplication{githubComOzoneCloudCommonEntitiesApplicationApplication.Custom})
			properties["deleted_at"] = githubComOzoneCloudCommonEntitiesApplicationApplication.DeletedAt
			properties["description"] = githubComOzoneCloudCommonEntitiesApplicationApplication.Description
			properties["environments"] = SetEnvironmentSubResourceData(githubComOzoneCloudCommonEntitiesApplicationApplication.Environments)
			properties["id"] = githubComOzoneCloudCommonEntitiesApplicationApplication.ID
			properties["ml_verification"] = SetApplicationMLVerificationSubResourceData([]*models.ApplicationMLVerification{githubComOzoneCloudCommonEntitiesApplicationApplication.MlVerification})
			properties["name"] = githubComOzoneCloudCommonEntitiesApplicationApplication.Name
			properties["status"] = githubComOzoneCloudCommonEntitiesApplicationApplication.Status
			properties["type_id"] = githubComOzoneCloudCommonEntitiesApplicationApplication.TypeID
			properties["type_name"] = githubComOzoneCloudCommonEntitiesApplicationApplication.TypeName
			properties["updated_at"] = githubComOzoneCloudCommonEntitiesApplicationApplication.UpdatedAt
			properties["updated_by_id"] = githubComOzoneCloudCommonEntitiesApplicationApplication.UpdatedByID
			properties["updated_by_name"] = githubComOzoneCloudCommonEntitiesApplicationApplication.UpdatedByName
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{githubComOzoneCloudCommonEntitiesApplicationApplication.Workspaces})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate GithubComOzoneCloudCommonEntitiesApplicationApplication resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func GithubComOzoneCloudCommonEntitiesApplicationApplicationModel(d *schema.ResourceData) *models.GithubComOzoneCloudCommonEntitiesApplicationApplication {
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
	id, _ := d.Get("id").(string)
	var mlVerification *models.ApplicationMLVerification = nil //hit complex
	//mlVerificationInterface, mlVerificationIsSet := d.GetOk("ml_verification")
	//if mlVerificationIsSet {
	//	mlVerificationMap := mlVerificationInterface.([]interface{})[0].(map[string]interface{})
	mlVerification = ApplicationMLVerificationModel(d)
	//}
	name := d.Get("name").(string)
	status := d.Get("status").(string)
	typeID := int64(d.Get("type_id").(int))
	typeName := d.Get("type_name").(string)
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)
	var workspaces *models.MongoResourceWorkspaces = nil //hit complex
	//workspacesInterface, workspacesIsSet := d.GetOk("workspaces")
	//if workspacesIsSet {
	//	workspacesMap := workspacesInterface.([]interface{})[0].(map[string]interface{})
	workspaces = MongoResourceWorkspacesModel(d)
	//}

	return &models.GithubComOzoneCloudCommonEntitiesApplicationApplication{
		AccountID:      accountID,
		Apm:            apm,
		CreatedAt:      createdAt,
		CreatedByID:    createdByID,
		CreatedByName:  createdByName,
		Custom:         custom,
		DeletedAt:      deletedAt,
		Description:    description,
		Environments:   environments,
		ID:             id,
		MlVerification: mlVerification,
		Name:           name,
		Status:         status,
		TypeID:         typeID,
		TypeName:       typeName,
		UpdatedAt:      updatedAt,
		UpdatedByID:    updatedByID,
		UpdatedByName:  updatedByName,
		Workspaces:     workspaces,
	}
}

// Retrieve property field names for updating the GithubComOzoneCloudCommonEntitiesApplicationApplication resource
func GetGithubComOzoneCloudCommonEntitiesApplicationApplicationPropertyFields() (t []string) {
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
		"id",
		"ml_verification",
		"name",
		"status",
		"type_id",
		"type_name",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
		"workspaces",
	}
}
