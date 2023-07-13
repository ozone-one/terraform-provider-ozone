package schemata

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the CreatePipelineRunRequest resource defined in the Terraform configuration
func CreatePipelineRunRequestSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"application_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"cluster_id": {
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

		"env_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"kind": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"params": {
			Type: schema.TypeList, //GoType: []*PipelineParam
			Elem: &schema.Resource{
				Schema: PipelineParamSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"pipeline_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"pipeline_run_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying CreatePipelineRunRequest resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCreatePipelineRunRequestResourceData(d *schema.ResourceData, m *models.CreatePipelineRunRequest) {
	d.Set("account_id", m.AccountID)
	d.Set("application_id", m.ApplicationID)
	d.Set("cluster_id", m.ClusterID)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("env_id", m.EnvID)
	d.Set("kind", m.Kind)
	d.Set("name", m.Name)
	d.Set("params", SetPipelineParamSubResourceData(m.Params))
	d.Set("pipeline_id", m.PipelineID)
	d.Set("pipeline_run_id", m.PipelineRunID)
	d.Set("timeout", m.Timeout)
}

// Iterate throught and update the CreatePipelineRunRequest resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCreatePipelineRunRequestSubResourceData(m []*models.CreatePipelineRunRequest) (d []*map[string]interface{}) {
	for _, createPipelineRunRequest := range m {
		if createPipelineRunRequest != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = createPipelineRunRequest.AccountID
			properties["application_id"] = createPipelineRunRequest.ApplicationID
			properties["cluster_id"] = createPipelineRunRequest.ClusterID
			properties["created_at"] = createPipelineRunRequest.CreatedAt
			properties["created_by_id"] = createPipelineRunRequest.CreatedByID
			properties["created_by_name"] = createPipelineRunRequest.CreatedByName
			properties["deleted_at"] = createPipelineRunRequest.DeletedAt
			properties["description"] = createPipelineRunRequest.Description
			properties["env_id"] = createPipelineRunRequest.EnvID
			properties["kind"] = createPipelineRunRequest.Kind
			properties["name"] = createPipelineRunRequest.Name
			properties["params"] = SetPipelineParamSubResourceData(createPipelineRunRequest.Params)
			properties["pipeline_id"] = createPipelineRunRequest.PipelineID
			properties["pipeline_run_id"] = createPipelineRunRequest.PipelineRunID
			properties["timeout"] = createPipelineRunRequest.Timeout
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CreatePipelineRunRequest resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CreatePipelineRunRequestModel(d *schema.ResourceData) *models.CreatePipelineRunRequest {
	accountID := d.Get("account_id").(string)
	applicationID := d.Get("application_id").(string)
	clusterID := d.Get("cluster_id").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	//description := d.Get("description").(string)
	envID := d.Get("env_id").(string)
	kind := d.Get("kind").(string)
	name := d.Get("name").(string)
	pipelineID := d.Get("pipeline_id").(string)
	//pipelineRunID := d.Get("pipeline_run_id").(string)
	var timeout int64
	if d.Get("timeout") != nil {
		timeout = int64(d.Get("timeout").(int))
	}
	var paramsArray []*models.PipelineParam
	for _, v := range d.Get("trigger_params").([]interface{}) {
		item := pipelineParams(v)
		paramsArray = append(paramsArray, item)
	}

	return &models.CreatePipelineRunRequest{
		AccountID:     accountID,
		ApplicationID: applicationID,
		ClusterID:     clusterID,
		CreatedAt:     createdAt,
		CreatedByID:   createdByID,
		CreatedByName: createdByName,
		DeletedAt:     deletedAt,
		//Description:   description,
		EnvID:      envID,
		Kind:       kind,
		Name:       name,
		Params:     paramsArray,
		PipelineID: pipelineID,
		//PipelineRunID: pipelineRunID,
		Timeout: timeout,
	}
}

// Retrieve property field names for updating the CreatePipelineRunRequest resource
func GetCreatePipelineRunRequestPropertyFields() (t []string) {
	return []string{
		"account_id",
		"application_id",
		"cluster_id",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"description",
		"env_id",
		"kind",
		"name",
		"params",
		"pipeline_id",
		"pipeline_run_id",
		"timeout",
	}
}

func pipelineParams(v interface{}) *models.PipelineParam {
	var item models.PipelineParam
	if v.(map[string]interface{})["name"] != nil {
		item.Name = v.(map[string]interface{})["name"].(string)
	}
	if v.(map[string]interface{})["value"] != nil {
		item.Value = v.(map[string]interface{})["value"].(string)
	}
	if v.(map[string]interface{})["type"] != nil {
		item.Type = int64(v.(map[string]interface{})["type"].(int))
	}
	if v.(map[string]interface{})["description"] != nil {
		item.Description = v.(map[string]interface{})["description"].(string)
	}
	if v.(map[string]interface{})["required"] != nil {
		item.Required = v.(map[string]interface{})["required"].(bool)
	}
	if val := v.(map[string]interface{})["image_tag_config"]; val != nil {
		if plbc, ok := val.([]interface{}); ok {
			for _, v := range plbc {
				if v.(map[string]interface{})["pick_latest_branch_commit"] != nil {
					ok, _ := strconv.ParseBool(v.(map[string]interface{})["pick_latest_branch_commit"].(string))
					item.ImageTagConfig = &models.ImageTagConfig{
						PickLatestBranchCommit: ok,
					}
				}
			}
		} else {
			if plbc := val.(*schema.Set); plbc != nil {
				items := plbc.List()
				if len(items) > 0 {
					item.ImageTagConfig = &models.ImageTagConfig{
						PickLatestBranchCommit: items[0].(map[string]interface{})["pick_latest_branch_commit"].(bool),
					}
				}
			}
		}
		// if plbc := val.(map[string]interface{})["pick_latest_branch_commit"]; plbc != nil {
		// 	ok, _ := strconv.ParseBool(plbc.(string))
		// 	item.ImageTagConfig = &models.ImageTagConfig{
		// 		PickLatestBranchCommit: ok,
		// 	}
		// }
	}
	return &item
}
