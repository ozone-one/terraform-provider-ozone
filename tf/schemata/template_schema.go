package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the Template resource defined in the Terraform configuration
func TemplateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"categories": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
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

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"string_data": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"template_format_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"template_type_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"u": {
			Type: schema.TypeList, //GoType: UnstructuredUnstructured
			Elem: &schema.Resource{
				Schema: UnstructuredUnstructuredSchema(),
			},
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

		"visibility": {
			Type: schema.TypeList, //GoType: TemplateVisibility
			Elem: &schema.Resource{
				Schema: TemplateVisibilitySchema(),
			},
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

// Update the underlying Template resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetTemplateResourceData(d *schema.ResourceData, m *models.Template) {
	d.Set("account_id", m.AccountID)
	d.Set("categories", m.Categories)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("string_data", m.StringData)
	d.Set("template_format_id", m.TemplateFormatID)
	d.Set("template_type_id", m.TemplateTypeID)
	d.Set("u", SetUnstructuredUnstructuredSubResourceData([]*models.UnstructuredUnstructured{m.U}))
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("visibility", SetTemplateVisibilitySubResourceData([]*models.TemplateVisibility{m.Visibility}))
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the Template resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetTemplateSubResourceData(m []*models.Template) (d []*map[string]interface{}) {
	for _, template := range m {
		if template != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = template.AccountID
			properties["categories"] = template.Categories
			properties["created_at"] = template.CreatedAt
			properties["created_by_id"] = template.CreatedByID
			properties["created_by_name"] = template.CreatedByName
			properties["deleted_at"] = template.DeletedAt
			properties["description"] = template.Description
			properties["id"] = template.ID
			properties["name"] = template.Name
			properties["string_data"] = template.StringData
			properties["template_format_id"] = template.TemplateFormatID
			properties["template_type_id"] = template.TemplateTypeID
			properties["u"] = SetUnstructuredUnstructuredSubResourceData([]*models.UnstructuredUnstructured{template.U})
			properties["updated_at"] = template.UpdatedAt
			properties["updated_by_id"] = template.UpdatedByID
			properties["updated_by_name"] = template.UpdatedByName
			properties["visibility"] = SetTemplateVisibilitySubResourceData([]*models.TemplateVisibility{template.Visibility})
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{template.Workspaces})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Template resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func TemplateModel(d *schema.ResourceData) *models.Template {
	accountID := d.Get("account_id").(string)
	categories := d.Get("categories").([]string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	description := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name := d.Get("name").(string)
	stringData := d.Get("string_data").(string)
	templateFormatID := int64(d.Get("template_format_id").(int))
	templateTypeID := int64(d.Get("template_type_id").(int))
	var u *models.UnstructuredUnstructured = nil //hit complex
	//uInterface, uIsSet := d.GetOk("u")
	//if uIsSet {
	//	uMap := uInterface.([]interface{})[0].(map[string]interface{})
	u = UnstructuredUnstructuredModel(d)
	//}
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)
	var visibility *models.TemplateVisibility = nil //hit complex
	//visibilityInterface, visibilityIsSet := d.GetOk("visibility")
	//if visibilityIsSet {
	//	visibilityMap := visibilityInterface.([]interface{})[0].(map[string]interface{})
	visibility = TemplateVisibilityModel(d)
	//}
	var workspaces *models.MongoResourceWorkspaces = nil //hit complex
	//workspacesInterface, workspacesIsSet := d.GetOk("workspaces")
	//if workspacesIsSet {
	//	workspacesMap := workspacesInterface.([]interface{})[0].(map[string]interface{})
	workspaces = MongoResourceWorkspacesModel(d)
	//}

	return &models.Template{
		AccountID:        accountID,
		Categories:       categories,
		CreatedAt:        createdAt,
		CreatedByID:      createdByID,
		CreatedByName:    createdByName,
		DeletedAt:        deletedAt,
		Description:      description,
		ID:               id,
		Name:             name,
		StringData:       stringData,
		TemplateFormatID: templateFormatID,
		TemplateTypeID:   templateTypeID,
		U:                u,
		UpdatedAt:        updatedAt,
		UpdatedByID:      updatedByID,
		UpdatedByName:    updatedByName,
		Visibility:       visibility,
		Workspaces:       workspaces,
	}
}

// Retrieve property field names for updating the Template resource
func GetTemplatePropertyFields() (t []string) {
	return []string{
		"account_id",
		"categories",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"description",
		"id",
		"name",
		"string_data",
		"template_format_id",
		"template_type_id",
		"u",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
		"visibility",
		"workspaces",
	}
}
