package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ApplicationTemplates resource defined in the Terraform configuration
func ApplicationTemplatesSchema() map[string]*schema.Schema {
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

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"items": {
			Type: schema.TypeList, //GoType: []*ApplicationTemplateItem
			Elem: &schema.Resource{
				Schema: ApplicationTemplateItemSchema(),
			},
			ConfigMode: schema.SchemaConfigModeAttr,
			Optional:   true,
			ForceNew:   true,
		},

		"name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"params": {
			Type: schema.TypeList, //GoType: TemplateBootstrapValues
			Elem: &schema.Resource{
				Schema: TemplateBootstrapValuesSchema(),
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
	}
}

// Update the underlying ApplicationTemplates resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetApplicationTemplatesResourceData(d *schema.ResourceData, m *models.ApplicationTemplates) {
	d.Set("account_id", m.AccountID)
	d.Set("application_id", m.ApplicationID)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("id", m.ID)
	d.Set("items", SetApplicationTemplateItemSubResourceData(m.Items))
	d.Set("name", m.Name)
	d.Set("params", SetTemplateBootstrapValuesSubResourceData([]*models.TemplateBootstrapValues{m.Params}))
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
}

// Iterate throught and update the ApplicationTemplates resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetApplicationTemplatesSubResourceData(m []*models.ApplicationTemplates) (d []*map[string]interface{}) {
	for _, applicationTemplates := range m {
		if applicationTemplates != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = applicationTemplates.AccountID
			properties["application_id"] = applicationTemplates.ApplicationID
			properties["created_at"] = applicationTemplates.CreatedAt
			properties["created_by_id"] = applicationTemplates.CreatedByID
			properties["created_by_name"] = applicationTemplates.CreatedByName
			properties["deleted_at"] = applicationTemplates.DeletedAt
			properties["id"] = applicationTemplates.ID
			properties["items"] = SetApplicationTemplateItemSubResourceData(applicationTemplates.Items)
			properties["name"] = applicationTemplates.Name
			properties["params"] = SetTemplateBootstrapValuesSubResourceData([]*models.TemplateBootstrapValues{applicationTemplates.Params})
			properties["updated_at"] = applicationTemplates.UpdatedAt
			properties["updated_by_id"] = applicationTemplates.UpdatedByID
			properties["updated_by_name"] = applicationTemplates.UpdatedByName
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ApplicationTemplates resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ApplicationTemplatesModel(d *schema.ResourceData) *models.ApplicationTemplates {
	accountID := d.Get("account_id").(string)
	applicationID := d.Get("application_id").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	id, _ := d.Get("id").(string)
	items := d.Get("items").([]*models.ApplicationTemplateItem)
	name := d.Get("name").(string)
	var params *models.TemplateBootstrapValues = nil //hit complex
	//paramsInterface, paramsIsSet := d.GetOk("params")
	//if paramsIsSet {
	//	paramsMap := paramsInterface.([]interface{})[0].(map[string]interface{})
	params = TemplateBootstrapValuesModel(d)
	//}
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)

	return &models.ApplicationTemplates{
		AccountID:     accountID,
		ApplicationID: applicationID,
		CreatedAt:     createdAt,
		CreatedByID:   createdByID,
		CreatedByName: createdByName,
		DeletedAt:     deletedAt,
		ID:            id,
		Items:         items,
		Name:          name,
		Params:        params,
		UpdatedAt:     updatedAt,
		UpdatedByID:   updatedByID,
		UpdatedByName: updatedByName,
	}
}

// Retrieve property field names for updating the ApplicationTemplates resource
func GetApplicationTemplatesPropertyFields() (t []string) {
	return []string{
		"account_id",
		"application_id",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"id",
		"items",
		"name",
		"params",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
	}
}
