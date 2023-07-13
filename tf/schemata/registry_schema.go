package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the Registry resource defined in the Terraform configuration
func RegistrySchema() map[string]*schema.Schema {
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

		"first_tag": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"latest_tag": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"number_of_images": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"project_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"provider_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"providers": {
			Type: schema.TypeList, //GoType: ResourceProviders
			Elem: &schema.Resource{
				Schema: ResourceProvidersSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"public": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},

		"registry_type_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"registry_type_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"resource_group_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"tags": {
			Type: schema.TypeList, //GoType: []string
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{},
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

		"url": {
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
		"workspace_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// Update the underlying Registry resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetRegistryResourceData(d *schema.ResourceData, m *models.Registry) {
	d.Set("account_id", m.AccountID)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("first_tag", m.FirstTag)
	d.Set("id", m.ID)
	d.Set("latest_tag", m.LatestTag)
	d.Set("name", m.Name)
	d.Set("number_of_images", m.NumberOfImages)
	d.Set("project_id", m.ProjectID)
	d.Set("provider_id", m.ProviderID)
	d.Set("providers", SetResourceProvidersSubResourceData([]*models.ResourceProviders{m.Providers}))
	d.Set("public", m.Public)
	d.Set("registry_type_id", m.RegistryTypeID)
	d.Set("registry_type_name", m.RegistryTypeName)
	d.Set("resource_group_name", m.ResourceGroupName)
	d.Set("tags", m.Tags)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("url", m.URL)
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the Registry resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetRegistrySubResourceData(m []*models.Registry) (d []*map[string]interface{}) {
	for _, registry := range m {
		if registry != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = registry.AccountID
			properties["created_at"] = registry.CreatedAt
			properties["created_by_id"] = registry.CreatedByID
			properties["created_by_name"] = registry.CreatedByName
			properties["deleted_at"] = registry.DeletedAt
			properties["description"] = registry.Description
			properties["first_tag"] = registry.FirstTag
			properties["id"] = registry.ID
			properties["latest_tag"] = registry.LatestTag
			properties["name"] = registry.Name
			properties["number_of_images"] = registry.NumberOfImages
			properties["project_id"] = registry.ProjectID
			properties["provider_id"] = registry.ProviderID
			properties["providers"] = SetResourceProvidersSubResourceData([]*models.ResourceProviders{registry.Providers})
			properties["public"] = registry.Public
			properties["registry_type_id"] = registry.RegistryTypeID
			properties["registry_type_name"] = registry.RegistryTypeName
			properties["resource_group_name"] = registry.ResourceGroupName
			properties["tags"] = registry.Tags
			properties["updated_at"] = registry.UpdatedAt
			properties["updated_by_id"] = registry.UpdatedByID
			properties["updated_by_name"] = registry.UpdatedByName
			properties["url"] = registry.URL
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{registry.Workspaces})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Registry resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func RegistryModel(d *schema.ResourceData) *models.Registry {
	accountID := d.Get("account_id").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	deletedAt := d.Get("deleted_at").(string)
	description := d.Get("description").(string)
	firstTag := d.Get("first_tag").(string)
	id, _ := d.Get("id").(string)
	latestTag := d.Get("latest_tag").(string)
	name := d.Get("name").(string)
	numberOfImages := int64(d.Get("number_of_images").(int))
	projectID := int64(d.Get("project_id").(int))
	providerID := d.Get("provider_id").(string)
	var providers *models.ResourceProviders = nil //hit complex
	//providersInterface, providersIsSet := d.GetOk("providers")
	//if providersIsSet {
	//	providersMap := providersInterface.([]interface{})[0].(map[string]interface{})
	providers = ResourceProvidersModel(d)
	//}
	public := d.Get("public").(bool)
	registryTypeID := int64(d.Get("registry_type_id").(int))
	registryTypeName := d.Get("registry_type_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	tags := d.Get("tags").([]string)
	updatedAt := d.Get("updated_at").(string)
	updatedByID := d.Get("updated_by_id").(string)
	updatedByName := d.Get("updated_by_name").(string)
	url := d.Get("url").(string)
	var workspaces *models.MongoResourceWorkspaces = nil //hit complex
	//workspacesInterface, workspacesIsSet := d.GetOk("workspaces")
	//if workspacesIsSet {
	//	workspacesMap := workspacesInterface.([]interface{})[0].(map[string]interface{})
	workspaces = MongoResourceWorkspacesModel(d)
	//}

	return &models.Registry{
		AccountID:         accountID,
		CreatedAt:         createdAt,
		CreatedByID:       createdByID,
		CreatedByName:     createdByName,
		DeletedAt:         deletedAt,
		Description:       description,
		FirstTag:          firstTag,
		ID:                id,
		LatestTag:         latestTag,
		Name:              name,
		NumberOfImages:    numberOfImages,
		ProjectID:         projectID,
		ProviderID:        providerID,
		Providers:         providers,
		Public:            public,
		RegistryTypeID:    registryTypeID,
		RegistryTypeName:  registryTypeName,
		ResourceGroupName: resourceGroupName,
		Tags:              tags,
		UpdatedAt:         updatedAt,
		UpdatedByID:       updatedByID,
		UpdatedByName:     updatedByName,
		URL:               url,
		Workspaces:        workspaces,
	}
}

// Retrieve property field names for updating the Registry resource
func GetRegistryPropertyFields() (t []string) {
	return []string{
		"account_id",
		"created_at",
		"created_by_id",
		"created_by_name",
		"deleted_at",
		"description",
		"first_tag",
		"id",
		"latest_tag",
		"name",
		"number_of_images",
		"project_id",
		"provider_id",
		"providers",
		"public",
		"registry_type_id",
		"registry_type_name",
		"resource_group_name",
		"tags",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
		"url",
		"workspaces",
	}
}
