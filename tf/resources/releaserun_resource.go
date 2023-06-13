//// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (

	"log"
	"github.com/ozone-one/terraform-provider-ozone/client/releaserun"
	client "github.com/ozone-one/terraform-provider-ozone/client"
	"github.com/ozone-one/terraform-provider-ozone/tf/schemata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

/*
Releaserun releaserun API
*/

func Releaserun() *schema.Resource {
	return &schema.Resource{
		ReadContext:   releaseRunReadResource,
		CreateContext: releaseRunCreate,
		DeleteContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
			return nil
		},
		Schema: schemata.ReleaseRunSchema(),
	}
}


func DataResourceReleaserun() *schema.Resource {
	return &schema.Resource{
		ReadContext:   releaseRunReadData,
		Schema: schemata.ReleaseRunSchema(),
	}
}

func releaseRunGetByID(ctx context.Context, d *schema.ResourceData,m interface{},id string) diag.Diagnostics {

	params := releaserun.NewGetReleaseRunByIDParams()
	params.SetID(id)

	client := m.(*client.AppBeMasterAPI)

	resp, err := client.Releaserun.GetReleaseRunByID(params)
	if err != nil {
		log.Printf("failed to make api request error: %v", err)
		return diag.Errorf("unexpected error: %s", err)
	}

	respModel := resp.GetPayload()
	//schemata.SetReleaseRunResourceData(d, respModel)
	d.SetId(respModel.ID)
	schemata.SetReleaseRunResourceData(d, respModel)
	log.Printf("[TRACE] response: %+v", respModel.ID)
	return nil
}


func releaseRunCreate(ctx context.Context, d *schema.ResourceData,m interface{}) diag.Diagnostics {

	model := schemata.CreateReleaseRunRequestModel(d)
	params := releaserun.NewReleaseRunCreateParams().WithRequest(model)

	client := m.(*client.AppBeMasterAPI)

	resp, err := client.Releaserun.ReleaseRunCreate(params)
	log.Printf("[TRACE] response: %v", resp)
	if err != nil {
		log.Printf("failed to make api request error: %+v,params:= %+v", err.Error(),model.Steps[0])
		return diag.Errorf("unexpected: %s", err)
	}

	respModel := resp.GetPayload()
	d.SetId(respModel.ID)
	schemata.SetReleaseRunResourceData(d, respModel)

	return nil
}


func releaseRunReadResource(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[TRACE] setting using ReousrceID id: %s", d.Id())
	return releaseRunGetByID(ctx,d,m,d.Id())
}
func releaseRunReadData(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	id := d.Get("uid").(string)
	log.Printf("[TRACE] setting using UID id: %s", id)
	return releaseRunGetByID(ctx,d,m,id)
}