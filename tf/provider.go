package tf

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "github.com/ozone-one/terraform-provider-ozone/client"
	"github.com/ozone-one/terraform-provider-ozone/tf/resources"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OZ_WORKSPACE_ID", nil),
			},
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
		},
		ResourcesMap: map[string]*schema.Resource{
			"oz_pipelinerun": resources.Pipelinerun(),
			"oz_releaserun":  resources.Releaserun(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	workspaceID := d.Get("workspace_id").(string)
	key := d.Get("api_key").(string)
	host := d.Get("host").(string)
	config := client.NewConfig()
	config.SetAccessKey(&key)
	config.SetWorkspaceID(&workspaceID)
	config.SetAccountDomain(&host)

	c := client.New(config)

	return c, nil
}
