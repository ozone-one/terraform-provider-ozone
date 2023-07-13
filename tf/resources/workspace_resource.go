package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "github.com/ozone-one/terraform-provider-ozone/client"
	"github.com/ozone-one/terraform-provider-ozone/client/workspace"
	"github.com/ozone-one/terraform-provider-ozone/tf/schemata"
	"golang.org/x/net/context"
)

func DataResourceWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: workspaceReadDataSource,
		Schema:      schemata.WorkspaceSchema(),
	}
}

func workspaceReadDataSource(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	nameSlug := d.Get("name").(string)
	params := workspace.NewWorkspaceListParams().
		WithContext(ctx).WithNameSlug(&nameSlug)
	client := m.(*client.AppBeMasterAPI)
	resp, err := client.Workspace.WorkspaceList(params)
	if err != nil {
		return diag.Errorf("unexpected: %s", err)
	}
	respModel := resp.GetPayload()
	if len(respModel) == 0 {
		return diag.Errorf("no workspace found with name %s", d.Get("name").(string))
	}
	d.SetId(respModel[0].ID)
	schemata.SetWorkspaceResourceData(d, respModel[0])

	return nil
}
