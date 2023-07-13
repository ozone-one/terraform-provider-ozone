package tf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "github.com/ozone-one/terraform-provider-ozone/client"
	"github.com/ozone-one/terraform-provider-ozone/tf/resources"
)

const (
	serviceAccountBasePath = "/api/service_account/api"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OZ_API_KEY", nil),
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OZ_HOST", nil),
			},
			"base_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OZ_BASE_PATH", serviceAccountBasePath),
			},
			"scheme": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OZ_SCHEME", "https"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"oz_pipelinerun": resources.Pipelinerun(),
			"oz_releaserun":  resources.Releaserun(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"oz_microservice": resources.DataResourceMicroservice(),
			"oz_repository":   resources.DataResourceRepository(),
			"oz_registry":     resources.DataResourceRegistry(),
			"oz_environment":  resources.DataResourceEnvironment(),
			"oz_pipeline":     resources.DataResourcePipeline(),
			"oz_variable":     resources.DataResourceVariable(),
			"oz_cluster":      resources.DataResourceCluster(),
			"oz_project":      resources.DataResourceWorkspace(),
			"oz_release":      resources.DataResourceRelease(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	//workspaceID := d.Get("workspace_id").(string)
	key := d.Get("api_key").(string)
	host := d.Get("host").(string)
	basepath := d.Get("base_path").(string)
	scheme := d.Get("scheme").(string)
	config := client.NewConfig()
	config.SetAccessKey(&key)
	//config.SetWorkspaceID(&workspaceID)
	config.SetAccountDomain(&host)
	config.SetBasePath(&basepath)
	config.SetSchemes([]string{scheme})

	c := client.New(config)

	return c, nil
}
