package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the AppRelease resource defined in the Terraform configuration
func AppReleaseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
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

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"last_release_run": {
			Type: schema.TypeList, //GoType: ReleaseRunDetails
			Elem: &schema.Resource{
				Schema: ReleaseRunDetailsSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"steps": {
			Type: schema.TypeList, //GoType: []*AppReleaseStep
			Elem: &schema.Resource{
				Schema: AppReleaseStepSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
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

		"workspace_id": {
			Type:     schema.TypeString,
			Required: true,
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

// Update the underlying AppRelease resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetAppReleaseResourceData(d *schema.ResourceData, m *models.AppRelease) {
	d.Set("account_id", m.AccountID)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("last_release_run", SetReleaseRunDetailsSubResourceData([]*models.ReleaseRunDetails{m.LastReleaseRun}))
	d.Set("name", m.Name)
	d.Set("steps", SetAppReleaseStepSubResourceData(m.Steps))
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the AppRelease resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetAppReleaseSubResourceData(m []*models.AppRelease) (d []*map[string]interface{}) {
	for _, appRelease := range m {
		if appRelease != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = appRelease.AccountID
			properties["created_at"] = appRelease.CreatedAt
			properties["created_by_id"] = appRelease.CreatedByID
			properties["created_by_name"] = appRelease.CreatedByName
			properties["deleted_at"] = appRelease.DeletedAt
			properties["description"] = appRelease.Description
			properties["id"] = appRelease.ID
			properties["last_release_run"] = SetReleaseRunDetailsSubResourceData([]*models.ReleaseRunDetails{appRelease.LastReleaseRun})
			properties["name"] = appRelease.Name
			properties["steps"] = SetAppReleaseStepSubResourceData(appRelease.Steps)
			properties["updated_at"] = appRelease.UpdatedAt
			properties["updated_by_id"] = appRelease.UpdatedByID
			properties["updated_by_name"] = appRelease.UpdatedByName
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{appRelease.Workspaces})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate AppRelease resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func AppReleaseModel(d *schema.ResourceData) *models.AppRelease {
	accountID := d.Get("account_id").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	description := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	var lastReleaseRun *models.ReleaseRunDetails = nil //hit complex
	//lastReleaseRunInterface, lastReleaseRunIsSet := d.GetOk("last_release_run")
	//if lastReleaseRunIsSet {
	//	lastReleaseRunMap := lastReleaseRunInterface.([]interface{})[0].(map[string]interface{})
	lastReleaseRun = ReleaseRunDetailsModel(d)
	//}
	name := d.Get("name").(string)
	steps := d.Get("steps").([]*models.AppReleaseStep)
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)
	var workspaces *models.MongoResourceWorkspaces = nil //hit complex
	//workspacesInterface, workspacesIsSet := d.GetOk("workspaces")
	//if workspacesIsSet {
	//	workspacesMap := workspacesInterface.([]interface{})[0].(map[string]interface{})
	workspaces = MongoResourceWorkspacesModel(d)
	//}

	return &models.AppRelease{
		AccountID:      accountID,
		CreatedAt:      createdAt,
		CreatedByID:    createdByID,
		CreatedByName:  createdByName,
		DeletedAt:      deletedAt,
		Description:    description,
		ID:             id,
		LastReleaseRun: lastReleaseRun,
		Name:           name,
		Steps:          steps,
		UpdatedAt:      updatedAt,
		UpdatedByID:    updatedByID,
		UpdatedByName:  updatedByName,
		Workspaces:     workspaces,
	}
}

// Retrieve property field names for updating the AppRelease resource
func GetAppReleasePropertyFields() (t []string) {
	return []string{
		"account_id",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"description",
		"id",
		"last_release_run",
		"name",
		"steps",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
		"workspaces",
	}
}
