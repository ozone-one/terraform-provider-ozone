package schemata

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ReleaseRun resource defined in the Terraform configuration
func ReleaseRunSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"app_release_id": {
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

		"uid": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"release": {
			Type: schema.TypeList, //GoType: ReleaseDetails
			Elem: &schema.Resource{
				Schema: ReleaseDetailsSchema(),
			},
			Optional:   true,
			ForceNew:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
		},

		"status": {
			Type: schema.TypeList, //GoType: ReleaseRunStatus
			Elem: &schema.Resource{
				Schema: ReleaseRunStatusSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Computed:   true,
			ForceNew:   true,
		},

		"params": {
			Type: schema.TypeList, //GoType: []*ReleaseRunStep
			Elem: &schema.Resource{
				Schema: ReleaseParamSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Required:   true,
			ForceNew:   true,
		},

		"trigger_resource_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"trigger_resource_kind": {
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

		"release_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"workspace_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

// Update the underlying ReleaseRun resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetReleaseRunResourceData(d *schema.ResourceData, m *models.ReleaseRun) {
	d.Set("account_id", m.AccountID)
	d.Set("app_release_id", m.AppReleaseID)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("uid", m.ID)
	d.Set("id", m.ID)
	if err := d.Set("name", m.Name); err != nil {
		log.Printf("the name field has not been set err:%s", err.Error())
	} else {
		log.Printf("the name field has been set to %s", m.Name)
	}
	d.Set("release", SetReleaseDetailsSubResourceData([]*models.ReleaseDetails{m.Release}))
	d.Set("status", SetReleaseRunStatusSubResourceData([]*models.ReleaseRunStatus{m.Status}))
	//d.Set("steps", SetReleaseRunStepSubResourceData(m.Steps))
	//d.Set("params", SetReleaseParamSubResourceData(m.Params))
	d.Set("trigger_resource_id", m.TriggerResourceID)
	d.Set("trigger_resource_kind", m.TriggerResourceKind)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the ReleaseRun resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetReleaseRunSubResourceData(m []*models.ReleaseRun) (d []*map[string]interface{}) {
	for _, releaseRun := range m {
		if releaseRun != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = releaseRun.AccountID
			properties["app_release_id"] = releaseRun.AppReleaseID
			properties["created_at"] = releaseRun.CreatedAt
			properties["created_by_id"] = releaseRun.CreatedByID
			properties["created_by_name"] = releaseRun.CreatedByName
			properties["deleted_at"] = releaseRun.DeletedAt
			properties["id"] = releaseRun.ID
			properties["name"] = releaseRun.Name
			properties["release"] = SetReleaseDetailsSubResourceData([]*models.ReleaseDetails{releaseRun.Release})
			properties["status"] = SetReleaseRunStatusSubResourceData([]*models.ReleaseRunStatus{releaseRun.Status})
			properties["steps"] = SetReleaseRunStepSubResourceData(releaseRun.Steps)
			properties["trigger_resource_id"] = releaseRun.TriggerResourceID
			properties["trigger_resource_kind"] = releaseRun.TriggerResourceKind
			properties["updated_at"] = releaseRun.UpdatedAt
			properties["updated_by_id"] = releaseRun.UpdatedByID
			properties["updated_by_name"] = releaseRun.UpdatedByName
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{releaseRun.Workspaces})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ReleaseRun resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ReleaseRunModel(d *schema.ResourceData) *models.ReleaseRun {
	accountID := d.Get("account_id").(string)
	appReleaseID := d.Get("app_release_id").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	id, _ := d.Get("id").(string)
	name := d.Get("name").(string)
	var release *models.ReleaseDetails = nil //hit complex
	//releaseInterface, releaseIsSet := d.GetOk("release")
	//if releaseIsSet {
	//	releaseMap := releaseInterface.([]interface{})[0].(map[string]interface{})
	release = ReleaseDetailsModel(d)
	//}
	var status *models.ReleaseRunStatus = nil //hit complex
	//statusInterface, statusIsSet := d.GetOk("status")
	//if statusIsSet {
	//	statusMap := statusInterface.([]interface{})[0].(map[string]interface{})
	status = ReleaseRunStatusModel(d)
	//}
	steps := d.Get("steps").([]*models.ReleaseRunStep)
	triggerResourceID := d.Get("trigger_resource_id").(string)
	triggerResourceKind := d.Get("trigger_resource_kind").(string)
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)
	var workspaces *models.MongoResourceWorkspaces = nil //hit complex
	//workspacesInterface, workspacesIsSet := d.GetOk("workspaces")
	//if workspacesIsSet {
	//	workspacesMap := workspacesInterface.([]interface{})[0].(map[string]interface{})
	workspaces = MongoResourceWorkspacesModel(d)
	//}

	return &models.ReleaseRun{
		AccountID:           accountID,
		AppReleaseID:        appReleaseID,
		CreatedAt:           createdAt,
		CreatedByID:         createdByID,
		CreatedByName:       createdByName,
		DeletedAt:           deletedAt,
		ID:                  id,
		Name:                name,
		Release:             release,
		Status:              status,
		Steps:               steps,
		TriggerResourceID:   triggerResourceID,
		TriggerResourceKind: triggerResourceKind,
		UpdatedAt:           updatedAt,
		UpdatedByID:         updatedByID,
		UpdatedByName:       updatedByName,
		Workspaces:          workspaces,
	}
}

// Retrieve property field names for updating the ReleaseRun resource
func GetReleaseRunPropertyFields() (t []string) {
	return []string{
		"account_id",
		"app_release_id",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"id",
		"name",
		"release",
		"status",
		"steps",
		"trigger_resource_id",
		"trigger_resource_kind",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
		"workspaces",
	}
}
