package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

// Schema mapping representing the CustomApplication resource defined in the Terraform configuration
func CustomApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bootstrap_value": {
			Type: schema.TypeList, //GoType: TemplateBootstrapValues
			Elem: &schema.Resource{
				Schema: TemplateBootstrapValuesSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"build_type": {
			Type: schema.TypeList, //GoType: BuildType
			Elem: &schema.Resource{
				Schema: BuildTypeSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"build_type_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"deploy_type": {
			Type: schema.TypeList, //GoType: DeploymentType
			Elem: &schema.Resource{
				Schema: DeploymentTypeSchema(),
			},
			Optional: true,
			ForceNew: true,
		},

		"deploy_type_id": {
			Type:     schema.TypeInt,
			Optional: true,
			ForceNew: true,
		},

		"dns": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"dont_need_c_i": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},

		"file_location": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"registry_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"registry_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"repository_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"repository_name": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"template_branch": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"template_repo_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},

		"template_repository_id": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},

		"template_sub_folder": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}
}

// Update the underlying CustomApplication resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetCustomApplicationResourceData(d *schema.ResourceData, m *models.CustomApplication) {
	d.Set("bootstrap_value", SetTemplateBootstrapValuesSubResourceData([]*models.TemplateBootstrapValues{m.BootstrapValue}))
	d.Set("build_type", SetBuildTypeSubResourceData([]*models.BuildType{m.BuildType}))
	d.Set("build_type_id", m.BuildTypeID)
	d.Set("deploy_type", SetDeploymentTypeSubResourceData([]*models.DeploymentType{m.DeployType}))
	d.Set("deploy_type_id", m.DeployTypeID)
	d.Set("dns", m.DNS)
	d.Set("dont_need_c_i", m.DontNeedCI)
	d.Set("file_location", m.FileLocation)
	d.Set("registry_id", m.RegistryID)
	d.Set("registry_name", m.RegistryName)
	d.Set("repository_id", m.RepositoryID)
	d.Set("repository_name", m.RepositoryName)
	d.Set("template_branch", m.TemplateBranch)
	d.Set("template_repo_enabled", m.TemplateRepoEnabled)
	d.Set("template_repository_id", m.TemplateRepositoryID)
	d.Set("template_sub_folder", m.TemplateSubFolder)
}

// Iterate throught and update the CustomApplication resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetCustomApplicationSubResourceData(m []*models.CustomApplication) (d []*map[string]interface{}) {
	for _, customApplication := range m {
		if customApplication != nil {
			properties := make(map[string]interface{})
			properties["bootstrap_value"] = SetTemplateBootstrapValuesSubResourceData([]*models.TemplateBootstrapValues{customApplication.BootstrapValue})
			properties["build_type"] = SetBuildTypeSubResourceData([]*models.BuildType{customApplication.BuildType})
			properties["build_type_id"] = customApplication.BuildTypeID
			properties["deploy_type"] = SetDeploymentTypeSubResourceData([]*models.DeploymentType{customApplication.DeployType})
			properties["deploy_type_id"] = customApplication.DeployTypeID
			properties["dns"] = customApplication.DNS
			properties["dont_need_c_i"] = customApplication.DontNeedCI
			properties["file_location"] = customApplication.FileLocation
			properties["registry_id"] = customApplication.RegistryID
			properties["registry_name"] = customApplication.RegistryName
			properties["repository_id"] = customApplication.RepositoryID
			properties["repository_name"] = customApplication.RepositoryName
			properties["template_branch"] = customApplication.TemplateBranch
			properties["template_repo_enabled"] = customApplication.TemplateRepoEnabled
			properties["template_repository_id"] = customApplication.TemplateRepositoryID
			properties["template_sub_folder"] = customApplication.TemplateSubFolder
			d = append(d, &properties)
		}
	}
	return
}

// Function to perform the following actions:
// (1) Translate CustomApplication resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func CustomApplicationModel(d *schema.ResourceData) *models.CustomApplication {
	var bootstrapValue *models.TemplateBootstrapValues = nil //hit complex
	//bootstrapValueInterface, bootstrapValueIsSet := d.GetOk("bootstrap_value")
	//if bootstrapValueIsSet {
	//	bootstrapValueMap := bootstrapValueInterface.([]interface{})[0].(map[string]interface{})
	bootstrapValue = TemplateBootstrapValuesModel(d)
	//}
	var buildType *models.BuildType = nil //hit complex
	//build_typeInterface, build_typeIsSet := d.GetOk("build_type")
	//if build_typeIsSet {
	//	build_typeMap := build_typeInterface.([]interface{})[0].(map[string]interface{})
	buildType = BuildTypeModel(d)
	//}
	buildTypeID := int64(d.Get("build_type_id").(int))
	var deployType *models.DeploymentType = nil //hit complex
	//deploy_typeInterface, deploy_typeIsSet := d.GetOk("deploy_type")
	//if deploy_typeIsSet {
	//	deploy_typeMap := deploy_typeInterface.([]interface{})[0].(map[string]interface{})
	deployType = DeploymentTypeModel(d)
	//}
	deployTypeID := int64(d.Get("deploy_type_id").(int))
	dns := d.Get("dns").(string)
	dontNeedCI := d.Get("dont_need_c_i").(bool)
	fileLocation := d.Get("file_location").(string)
	registryID := d.Get("registry_id").(string)
	registryName := d.Get("registry_name").(string)
	repositoryID := d.Get("repository_id").(string)
	repositoryName := d.Get("repository_name").(string)
	templateBranch := d.Get("template_branch").(string)
	templateRepoEnabled := d.Get("template_repo_enabled").(bool)
	templateRepositoryID := d.Get("template_repository_id").(string)
	templateSubFolder := d.Get("template_sub_folder").(string)

	return &models.CustomApplication{
		BootstrapValue:       bootstrapValue,
		BuildType:            buildType,
		BuildTypeID:          buildTypeID,
		DeployType:           deployType,
		DeployTypeID:         deployTypeID,
		DNS:                  dns,
		DontNeedCI:           dontNeedCI,
		FileLocation:         fileLocation,
		RegistryID:           registryID,
		RegistryName:         registryName,
		RepositoryID:         repositoryID,
		RepositoryName:       repositoryName,
		TemplateBranch:       templateBranch,
		TemplateRepoEnabled:  templateRepoEnabled,
		TemplateRepositoryID: templateRepositoryID,
		TemplateSubFolder:    templateSubFolder,
	}
}

// Retrieve property field names for updating the CustomApplication resource
func GetCustomApplicationPropertyFields() (t []string) {
	return []string{
		"bootstrap_value",
		"build_type",
		"build_type_id",
		"deploy_type",
		"deploy_type_id",
		"dns",
		"dont_need_c_i",
		"file_location",
		"registry_id",
		"registry_name",
		"repository_id",
		"repository_name",
		"template_branch",
		"template_repo_enabled",
		"template_repository_id",
		"template_sub_folder",
	}
}
