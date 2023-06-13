// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetReleaseByIDParams creates a new GetReleaseByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReleaseByIDParams() *GetReleaseByIDParams {
	return &GetReleaseByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleaseByIDParamsWithTimeout creates a new GetReleaseByIDParams object
// with the ability to set a timeout on a request.
func NewGetReleaseByIDParamsWithTimeout(timeout time.Duration) *GetReleaseByIDParams {
	return &GetReleaseByIDParams{
		timeout: timeout,
	}
}

// NewGetReleaseByIDParamsWithContext creates a new GetReleaseByIDParams object
// with the ability to set a context for a request.
func NewGetReleaseByIDParamsWithContext(ctx context.Context) *GetReleaseByIDParams {
	return &GetReleaseByIDParams{
		Context: ctx,
	}
}

// NewGetReleaseByIDParamsWithHTTPClient creates a new GetReleaseByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReleaseByIDParamsWithHTTPClient(client *http.Client) *GetReleaseByIDParams {
	return &GetReleaseByIDParams{
		HTTPClient: client,
	}
}

/*
GetReleaseByIDParams contains all the parameters to send to the API endpoint

	for the get release by ID operation.

	Typically these are written to a http.Request.
*/
type GetReleaseByIDParams struct {

	/* XWorkspaceID.

	   workspace id
	*/
	XWorkspaceID string

	/* ID.

	   mongo id of the application
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get release by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReleaseByIDParams) WithDefaults() *GetReleaseByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get release by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReleaseByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get release by ID params
func (o *GetReleaseByIDParams) WithTimeout(timeout time.Duration) *GetReleaseByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get release by ID params
func (o *GetReleaseByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get release by ID params
func (o *GetReleaseByIDParams) WithContext(ctx context.Context) *GetReleaseByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get release by ID params
func (o *GetReleaseByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get release by ID params
func (o *GetReleaseByIDParams) WithHTTPClient(client *http.Client) *GetReleaseByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get release by ID params
func (o *GetReleaseByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXWorkspaceID adds the xWorkspaceID to the get release by ID params
func (o *GetReleaseByIDParams) WithXWorkspaceID(xWorkspaceID string) *GetReleaseByIDParams {
	o.SetXWorkspaceID(xWorkspaceID)
	return o
}

// SetXWorkspaceID adds the xWorkspaceId to the get release by ID params
func (o *GetReleaseByIDParams) SetXWorkspaceID(xWorkspaceID string) {
	o.XWorkspaceID = xWorkspaceID
}

// WithID adds the id to the get release by ID params
func (o *GetReleaseByIDParams) WithID(id string) *GetReleaseByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get release by ID params
func (o *GetReleaseByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleaseByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Workspace-ID
	if err := r.SetHeaderParam("X-Workspace-ID", o.XWorkspaceID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}