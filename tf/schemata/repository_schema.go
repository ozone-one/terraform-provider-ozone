package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the Repository resource defined in the Terraform configuration
func RepositorySchema() map[string]*schema.Schema {
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

		"default_branch": {
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

		"repo_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"repository_type_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"repository_type_name": {
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
	}
}

// Update the underlying Repository resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetRepositoryResourceData(d *schema.ResourceData, m *models.Repository) {
	d.Set("account_id", m.AccountID)
	d.Set("created_at", m.CreatedAt)
	d.Set("created_by_id", m.CreatedByID)
	d.Set("created_by_name", m.CreatedByName)
	d.Set("default_branch", m.DefaultBranch)
	d.Set("deleted_at", m.DeletedAt)
	d.Set("description", m.Description)
	d.Set("id", m.ID)
	d.Set("name", m.Name)
	d.Set("provider_id", m.ProviderID)
	d.Set("providers", SetResourceProvidersSubResourceData([]*models.ResourceProviders{m.Providers}))
	d.Set("repo_name", m.RepoName)
	d.Set("repository_type_id", m.RepositoryTypeID)
	d.Set("repository_type_name", m.RepositoryTypeName)
	d.Set("updated_at", m.UpdatedAt)
	d.Set("updated_by_id", m.UpdatedByID)
	d.Set("updated_by_name", m.UpdatedByName)
	d.Set("url", m.URL)
	d.Set("workspaces", SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{m.Workspaces}))
}

// Iterate throught and update the Repository resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetRepositorySubResourceData(m []*models.Repository) (d []*map[string]interface{}) {
	for _, repository := range m {
		if repository != nil {
			properties := make(map[string]interface{})
			properties["account_id"] = repository.AccountID
			properties["created_at"] = repository.CreatedAt
			properties["created_by_id"] = repository.CreatedByID
			properties["created_by_name"] = repository.CreatedByName
			properties["default_branch"] = repository.DefaultBranch
			properties["deleted_at"] = repository.DeletedAt
			properties["description"] = repository.Description
			properties["id"] = repository.ID
			properties["name"] = repository.Name
			properties["provider_id"] = repository.ProviderID
			properties["providers"] = SetResourceProvidersSubResourceData([]*models.ResourceProviders{repository.Providers})
			properties["repo_name"] = repository.RepoName
			properties["repository_type_id"] = repository.RepositoryTypeID
			properties["repository_type_name"] = repository.RepositoryTypeName
			properties["updated_at"] = repository.UpdatedAt
			properties["updated_by_id"] = repository.UpdatedByID
			properties["updated_by_name"] = repository.UpdatedByName
			properties["url"] = repository.URL
			properties["workspaces"] = SetMongoResourceWorkspacesSubResourceData([]*models.MongoResourceWorkspaces{repository.Workspaces})
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate Repository resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func RepositoryModel(d *schema.ResourceData) *models.Repository {
	accountID := d.Get("account_id").(string)
	createdAt := d.Get("created_at").(string)
	createdByID := d.Get("created_by_id").(string)
	createdByName := d.Get("created_by_name").(string)
	defaultBranch := d.Get("default_branch").(string)
	deletedAt := d.Get("deleted_at").(string)
	description := d.Get("description").(string)
	id, _ := d.Get("id").(string)
	name := d.Get("name").(string)
	providerID := d.Get("provider_id").(string)
	var providers *models.ResourceProviders = nil //hit complex
	//providersInterface, providersIsSet := d.GetOk("providers")
	//if providersIsSet {
	//	providersMap := providersInterface.([]interface{})[0].(map[string]interface{})
	providers = ResourceProvidersModel(d)
	//}
	repoName := d.Get("repo_name").(string)
	repositoryTypeID := int64(d.Get("repository_type_id").(int))
	repositoryTypeName := d.Get("repository_type_name").(string)
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

	return &models.Repository{
		AccountID:          accountID,
		CreatedAt:          createdAt,
		CreatedByID:        createdByID,
		CreatedByName:      createdByName,
		DefaultBranch:      defaultBranch,
		DeletedAt:          deletedAt,
		Description:        description,
		ID:                 id,
		Name:               name,
		ProviderID:         providerID,
		Providers:          providers,
		RepoName:           repoName,
		RepositoryTypeID:   repositoryTypeID,
		RepositoryTypeName: repositoryTypeName,
		UpdatedAt:          updatedAt,
		UpdatedByID:        updatedByID,
		UpdatedByName:      updatedByName,
		URL:                url,
		Workspaces:         workspaces,
	}
}

// Retrieve property field names for updating the Repository resource
func GetRepositoryPropertyFields() (t []string) {
	return []string{
		"account_id",
		"created_at",
		"created_by_id",
		"created_by_name",
		"default_branch",
		"deleted_at",
		"description",
		"id",
		"name",
		"provider_id",
		"providers",
		"repo_name",
		"repository_type_id",
		"repository_type_name",
		"updated_at",
		"updated_by_id",
		"updated_by_name",
		"url",
		"workspaces",
	}
}
