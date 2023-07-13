package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the PipelineParam resource defined in the Terraform configuration
func PipelineParamSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default": {
			Type:     schema.TypeString,
			Optional: true,
		},

		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"image_tag_config": {
			Type: schema.TypeSet, //GoType: ImageTagConfig
			Elem: &schema.Resource{
				Schema: ImageTagConfigSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"required": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},

		"resource_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"type": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"type_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"value": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying PipelineParam resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetPipelineParamResourceData(d *schema.ResourceData, m *models.PipelineParam) {
	d.Set("default", m.Default)
	d.Set("description", m.Description)
	d.Set("image_tag_config", SetImageTagConfigSubResourceData([]*models.ImageTagConfig{m.ImageTagConfig}))
	d.Set("name", m.Name)
	d.Set("required", m.Required)
	d.Set("resource_id", m.ResourceID)
	d.Set("type", m.Type)
	d.Set("type_name", m.TypeName)
	d.Set("value", m.Value)
}

// Iterate throught and update the PipelineParam resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetPipelineParamSubResourceData(m []*models.PipelineParam) (d []interface{}) {
	for _, pipelineParam := range m {
		if pipelineParam != nil {
			properties := make(map[string]interface{})
			properties["default"] = pipelineParam.Default
			properties["description"] = pipelineParam.Description
			properties["image_tag_config"] = SetImageTagConfigSubResourceData([]*models.ImageTagConfig{pipelineParam.ImageTagConfig})
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

// Function to perform the following actions:
// (1) Translate PipelineParam resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func PipelineParamModel(d *schema.ResourceData) *models.PipelineParam {
	defaultVar := d.Get("default").(string)
	description := d.Get("description").(string)
	var imageTagConfig *models.ImageTagConfig = nil //hit complex
	//imageTagConfigInterface, imageTagConfigIsSet := d.GetOk("image_tag_config")
	//if imageTagConfigIsSet {
	//	imageTagConfigMap := imageTagConfigInterface.([]interface{})[0].(map[string]interface{})
	imageTagConfig = ImageTagConfigModel(d)
	//}
	name := d.Get("name").(string)
	required := d.Get("required").(bool)
	resourceID := d.Get("resource_id").(string)
	typeVar := int64(d.Get("type").(int))
	typeName := d.Get("type_name").(string)
	value := d.Get("value").(string)

	return &models.PipelineParam{
		Default:        defaultVar,
		Description:    description,
		ImageTagConfig: imageTagConfig,
		Name:           name,
		Required:       required,
		ResourceID:     resourceID,
		Type:           typeVar,
		TypeName:       typeName,
		Value:          value,
	}
}

// Retrieve property field names for updating the PipelineParam resource
func GetPipelineParamPropertyFields() (t []string) {
	return []string{
		"default",
		"description",
		"image_tag_config",
		"name",
		"required",
		"resource_id",
		"type",
		"type_name",
		"value",
	}
}
