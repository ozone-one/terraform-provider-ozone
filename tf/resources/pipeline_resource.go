package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	client "github.com/ozone-one/terraform-provider-ozone/client"
	"github.com/ozone-one/terraform-provider-ozone/client/pipeline"
	"github.com/ozone-one/terraform-provider-ozone/tf/schemata"
	"golang.org/x/net/context"
)

func DataResourcePipeline() *schema.Resource {
	return &schema.Resource{
		ReadContext: pipelineReadDataSource,
		Schema:      schemata.PipelineSchema(),
	}
}

func pipelineReadDataSource(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var (
		start int64 = 0
		end   int64 = 1
	)
	nameSlug := d.Get("name").(string)
	workspaceID := d.Get("workspace_id").(string)
	params := pipeline.NewPipelineListParams().
		WithContext(ctx).
		WithStart(&start).
		WithEnd(&end).WithNameSlug(&nameSlug).WithXWorkspaceID(workspaceID)
	client := m.(*client.AppBeMasterAPI)
	resp, err := client.Pipeline.PipelineList(params)
	if err != nil {
		return diag.Errorf("unexpected: %s", err)
	}
	respModel := resp.GetPayload()
	if len(respModel.Items) == 0 {
		return diag.Errorf("no pipeline found with name %s", d.Get("name").(string))
	}
	d.SetId(respModel.Items[0].ID)
	schemata.SetPipelineResourceData(d, respModel.Items[0])

	return nil
}
