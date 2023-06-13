package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the PipelineRun resource defined in the Terraform configuration
func PipelineRunSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"application_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"application_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"cluster_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"cluster_name": {
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

		"env_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"kind": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"ml_verification": {
			Type: schema.TypeList, //GoType: PipelineRunMLVerification
			Elem: &schema.Resource{
				Schema: PipelineRunMLVerificationSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"pipeline_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"rollback": {
			Type: schema.TypeList, //GoType: PipelineRunRollback
			Elem: &schema.Resource{
				Schema: PipelineRunRollbackSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"secret_injection": {
			Type: schema.TypeList, //GoType: PipelineRunSecretInjection
			Elem: &schema.Resource{
				Schema: PipelineRunSecretInjectionSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		// "tekton_pipeline_run": {
		// 	Type: schema.TypeList, //GoType: DynamicResource
		// 	Elem: &schema.Resource{
		// 		Schema: DynamicResourceSchema(),
		// 	},
		// 	Optional: true,
		// 	ForceNew: true,
		// },

		"trigger_params": {
			Type: schema.TypeList, //GoType: []*PipelineParam
			Elem: &schema.Resource{
				Schema: PipelineParamSchema(),
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

		"uid": {
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
		"task_run": {
			Type: schema.TypeList, //GoType: []*TaskRun
			Elem: &schema.Resource{
				Schema: TaskRunSchema(),
			},
			Optional:   true,
			ForceNew:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
		},
	}
}

// Update the underlying PipelineRun resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPipelineRunResourceData(d *schema.ResourceData, m *models.PipelineRun) {
	d.Set("account_id", m.AccountID)
	d.Set("application_id", m.ApplicationID)
	d.Set("application_name", m.ApplicationName)
	d.Set("cluster_id", m.ClusterID)
	d.Set("cluster_name", m.ClusterName)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("env_id", m.EnvID)
	d.Set("id", m.ID)
	d.Set("kind", m.Kind)
	d.Set("ml_verification", SetPipelineRunMLVerificationSubResourceData([]*models.PipelineRunMLVerification{m.MlVerification}))
	d.Set("name", m.Name)
	d.Set("pipeline_id", m.PipelineID)
	d.Set("rollback", SetPipelineRunRollbackSubResourceData([]*models.PipelineRunRollback{m.Rollback}))
	d.Set("secret_injection", SetPipelineRunSecretInjectionSubResourceData([]*models.PipelineRunSecretInjection{m.SecretInjection}))
	//d.Set("tekton_pipeline_run", SetDynamicResourceSubResourceData([]*models.DynamicResource{m.TektonPipelineRun}))
	d.Set("trigger_params", SetPipelineParamSubResourceData(m.TriggerParams))
	d.Set("trigger_resource_id", m.TriggerResourceID)
	d.Set("trigger_resource_kind", m.TriggerResourceKind)
	d.Set("uid", m.UID)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("task_run", SetTaskRunResourceData(m.TektonPipelineRun.Object))
}

// Iterate throught and update the PipelineRun resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPipelineRunSubResourceData(m []*models.PipelineRun) (d []*map[string]interface{}) {
	for _, pipelineRun := range m {
		if pipelineRun != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = pipelineRun.AccountID
			properties["application_id"] = pipelineRun.ApplicationID
			properties["application_name"] = pipelineRun.ApplicationName
			properties["cluster_id"] = pipelineRun.ClusterID
			properties["cluster_name"] = pipelineRun.ClusterName
			properties["created_at"] = pipelineRun.CreatedAt
			properties["created_by_id"] = pipelineRun.CreatedByID
			properties["created_by_name"] = pipelineRun.CreatedByName
			properties["deleted_at"] = pipelineRun.DeletedAt
			properties["env_id"] = pipelineRun.EnvID
			properties["id"] = pipelineRun.ID
			properties["kind"] = pipelineRun.Kind
			properties["ml_verification"] = SetPipelineRunMLVerificationSubResourceData([]*models.PipelineRunMLVerification{pipelineRun.MlVerification})
			properties["name"] = pipelineRun.Name
			properties["pipeline_id"] = pipelineRun.PipelineID
			properties["rollback"] = SetPipelineRunRollbackSubResourceData([]*models.PipelineRunRollback{pipelineRun.Rollback})
			properties["secret_injection"] = SetPipelineRunSecretInjectionSubResourceData([]*models.PipelineRunSecretInjection{pipelineRun.SecretInjection})
			//properties["tekton_pipeline_run"] = SetDynamicResourceSubResourceData([]*models.DynamicResource{pipelineRun.TektonPipelineRun})
			properties["trigger_params"] = SetPipelineParamSubResourceData(pipelineRun.TriggerParams)
			properties["trigger_resource_id"] = pipelineRun.TriggerResourceID
			properties["trigger_resource_kind"] = pipelineRun.TriggerResourceKind
			properties["uid"] = pipelineRun.UID
			properties["updated_at"] = pipelineRun.UpdatedAt
			properties["updated_by_id"] = pipelineRun.UpdatedByID
			properties["updated_by_name"] = pipelineRun.UpdatedByName
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate PipelineRun resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PipelineRunModel(d *schema.ResourceData) *models.PipelineRun {
	accountID := d.Get("account_id").(string)
	applicationID := d.Get("application_id").(string)
	applicationName := d.Get("application_name").(string)
	clusterID := d.Get("cluster_id").(string)
	clusterName := d.Get("cluster_name").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	envID := d.Get("env_id").(string)
	id, _ := d.Get("id").(string)
	kind := d.Get("kind").(string)
	var mlVerification *models.PipelineRunMLVerification = nil //hit complex
	//mlVerificationInterface, mlVerificationIsSet := d.GetOk("ml_verification")
	//if mlVerificationIsSet {
	//	mlVerificationMap := mlVerificationInterface.([]interface{})[0].(map[string]interface{})
	mlVerification = PipelineRunMLVerificationModel(d)
	//}
	name := d.Get("name").(string)
	pipelineID := d.Get("pipeline_id").(string)
	var rollback *models.PipelineRunRollback = nil //hit complex
	//rollbackInterface, rollbackIsSet := d.GetOk("rollback")
	//if rollbackIsSet {
	//	rollbackMap := rollbackInterface.([]interface{})[0].(map[string]interface{})
	rollback = PipelineRunRollbackModel(d)
	//}
	var secretInjection *models.PipelineRunSecretInjection = nil //hit complex
	//secretInjectionInterface, secretInjectionIsSet := d.GetOk("secret_injection")
	//if secretInjectionIsSet {
	//	secretInjectionMap := secretInjectionInterface.([]interface{})[0].(map[string]interface{})
	secretInjection = PipelineRunSecretInjectionModel(d)
	//}
	var tektonPipelineRun *models.DynamicResource = nil //hit complex
	//tektonPipelineRunInterface, tektonPipelineRunIsSet := d.GetOk("tekton_pipeline_run")
	//if tektonPipelineRunIsSet {
	//	tektonPipelineRunMap := tektonPipelineRunInterface.([]interface{})[0].(map[string]interface{})
	tektonPipelineRun = DynamicResourceModel(d)
	//}
	triggerParams := d.Get("trigger_params").([]*models.PipelineParam)
	triggerResourceID := d.Get("trigger_resource_id").(string)
	triggerResourceKind := d.Get("trigger_resource_kind").(string)
	uid := d.Get("uid").(string)
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)

	return &models.PipelineRun{
		AccountID:           accountID,
		ApplicationID:       applicationID,
		ApplicationName:     applicationName,
		ClusterID:           clusterID,
		ClusterName:         clusterName,
		CreatedAt:           createdAt,
		CreatedByID:         createdByID,
		CreatedByName:       createdByName,
		DeletedAt:           deletedAt,
		EnvID:               envID,
		ID:                  id,
		Kind:                kind,
		MlVerification:      mlVerification,
		Name:                name,
		PipelineID:          pipelineID,
		Rollback:            rollback,
		SecretInjection:     secretInjection,
		TektonPipelineRun:   tektonPipelineRun,
		TriggerParams:       triggerParams,
		TriggerResourceID:   triggerResourceID,
		TriggerResourceKind: triggerResourceKind,
		UID:                 uid,
		UpdatedAt:           updatedAt,
		UpdatedByID:         updatedByID,
		UpdatedByName:       updatedByName,
	}
}

// Retrieve property field names for updating the PipelineRun resource
func GetPipelineRunPropertyFields() (t []string) {
	return []string{
		"account_id",
		"application_id",
		"application_name",
		"cluster_id",
		"cluster_name",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"env_id",
		"id",
		"kind",
		"ml_verification",
		"name",
		"pipeline_id",
		"rollback",
		"secret_injection",
		"tekton_pipeline_run",
		"trigger_params",
		"trigger_resource_id",
		"trigger_resource_kind",
		"uid",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
	}
}
