package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the ApplicationTemplatesView resource defined in the Terraform configuration
func ApplicationTemplatesViewSchema() map[string]*schema.Schema {
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

// Update the underlying ApplicationTemplatesView resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetApplicationTemplatesViewResourceData(d *schema.ResourceData, m *models.ApplicationTemplatesView) {
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

// Iterate throught and update the ApplicationTemplatesView resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetApplicationTemplatesViewSubResourceData(m []*models.ApplicationTemplatesView) (d []*map[string]interface{}) {
	for _, applicationTemplatesView := range m {
		if applicationTemplatesView != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = applicationTemplatesView.AccountID
			properties["application_id"] = applicationTemplatesView.ApplicationID
			properties["created_at"] = applicationTemplatesView.CreatedAt
			properties["created_by_id"] = applicationTemplatesView.CreatedByID
			properties["created_by_name"] = applicationTemplatesView.CreatedByName
			properties["deleted_at"] = applicationTemplatesView.DeletedAt
			properties["id"] = applicationTemplatesView.ID
			properties["items"] = SetApplicationTemplateItemSubResourceData(applicationTemplatesView.Items)
			properties["name"] = applicationTemplatesView.Name
			properties["params"] = SetTemplateBootstrapValuesSubResourceData([]*models.TemplateBootstrapValues{applicationTemplatesView.Params})
			properties["updated_at"] = applicationTemplatesView.UpdatedAt
			properties["updated_by_id"] = applicationTemplatesView.UpdatedByID
			properties["updated_by_name"] = applicationTemplatesView.UpdatedByName
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate ApplicationTemplatesView resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func ApplicationTemplatesViewModel(d *schema.ResourceData) *models.ApplicationTemplatesView {
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

	return &models.ApplicationTemplatesView{
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

// Retrieve property field names for updating the ApplicationTemplatesView resource
func GetApplicationTemplatesViewPropertyFields() (t []string) {
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
