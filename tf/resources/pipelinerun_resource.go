//// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"log"

	"github.com/ozone-one/terraform-provider-ozone/client/pipelinerun"
	client "github.com/ozone-one/terraform-provider-ozone/client"
	"github.com/ozone-one/terraform-provider-ozone/tf/schemata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

/*
Pipelinerun pipelinerun API
*/

func Pipelinerun() *schema.Resource {
	return &schema.Resource{
		ReadContext:   pipelineRunReadResource,
		CreateContext: pipelineRunCreate,
		DeleteContext:func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
			return nil
		},
		Schema: schemata.PipelineRunSchema(),
	}
}

func DataResourcePipelinerun() *schema.Resource {
	return &schema.Resource{
		ReadContext: pipelineRunReadData,
		Schema: schemata.PipelineRunSchema(),
	}
}

func pipelineRunCreate(ctx context.Context, d *schema.ResourceData,m interface{}) diag.Diagnostics {

	model := schemata.CreatePipelineRunRequestModel(d)
	workspaceID := d.Get("workspace_id").(string)
	params := pipelinerun.NewPipelineRunCreateParams().WithRequest(model).WithXWorkspaceID(workspaceID)

	client := m.(*client.AppBeMasterAPI)

	resp, err := client.Pipelinerun.PipelineRunCreate(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		//diags = append(diags, diag.Errorf("unexpected: %s", err)...)
		return diag.Errorf("unexpected: %s", err)
	}

	respModel := resp.GetPayload()
	fmt.Printf("[TRACE] response: %+v", respModel)
	schemata.SetPipelineRunResourceData(d, respModel)
	d.SetId(resp.Payload.ID)

	return nil
}


func pipelineRunReadResource(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[TRACE] setting using ReousrceID id: %s", d.Id())
	return pipelineRunGetByID(ctx,d,m,d.Id())
}
func pipelineRunReadData(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Get("uid").(string)
	log.Printf("[TRACE] setting using UID id: %s", id)
	return pipelineRunGetByID(ctx,d,m,id)
}
func pipelineRunGetByID(ctx context.Context,d *schema.ResourceData, m interface{},id string) diag.Diagnostics {


	params := pipelinerun.NewPipelineRunGetByIDParams()
	
	params.ID = id

	client := m.(*client.AppBeMasterAPI)
	
	resp, err := client.Pipelinerun.PipelineRunGetByID(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		return diag.Errorf("failed to get pipelune run by id: %s", err)
	}

	respModel := resp.GetPayload()
	d.SetId(respModel.ID)
	schemata.SetPipelineRunResourceData(d, respModel)

	return nil
}
