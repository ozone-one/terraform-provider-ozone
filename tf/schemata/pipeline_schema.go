package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

func PipelineSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"deleted_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_by": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"workspace_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"account_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"params": {

			Type: schema.TypeList, //GoType: []*PipelineParam
			Elem: &schema.Resource{
				Schema: PipelineParamSchema(),
			},
			Optional:   true,
			ForceNew:   true,
			ConfigMode: schema.SchemaConfigModeAttr,
		},
	}
}

func SetPipelineResourceData(d *schema.ResourceData, m *models.Pipeline) {
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("name", m.Name)
	d.Set("id", m.ID)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by", m.UpdatedByName)
	d.Set("account_id", m.AccountID)
	d.Set("params", handleParamResource(m.Params))
}

func handleParamResource(m []*models.PipelineParam) (d []interface{}) {
	for _, pipelineParam := range m {
		if pipelineParam != nil {
			properties := make(map[string]interface{})
			properties["default"] = pipelineParam.Default
			properties["description"] = pipelineParam.Description
			properties["image_tag_config"] = handleImageTagConfigResource(pipelineParam.ImageTagConfig)
			properties["name"] = pipelineParam.Name
			properties["required"] = pipelineParam.Required
			properties["resource_id"] = pipelineParam.ResourceID
			properties["type"] = pipelineParam.Type
			properties["type_name"] = pipelineParam.TypeName
			properties["value"] = pipelineParam.Value
			d = append(d, &properties)
		}
	}
	return
}

func handleImageTagConfigResource(m *models.ImageTagConfig) (d []interface{}) {
	if m != nil {
		d = make([]interface{}, 0)
		properties := make(map[string]interface{})
		properties["pick_latest_branch_commit"] = m.PickLatestBranchCommit
	}
	return
}
