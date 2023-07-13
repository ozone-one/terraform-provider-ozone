package schemata

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ozone-one/terraform-provider-ozone/models"
)

func ClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"workspace_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"account_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func SetupClusterResourceData(d *schema.ResourceData, data *models.Cluster) {
	d.Set("name", data.Name)
	d.Set("account_id", data.AccountID)
	d.Set("id", data.ID)
}
