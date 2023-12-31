// Code generated by go-swagger; DO NOT EDIT.

package releaserun

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

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// NewReleaseRunCreateParams creates a new ReleaseRunCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReleaseRunCreateParams() *ReleaseRunCreateParams {
	return &ReleaseRunCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReleaseRunCreateParamsWithTimeout creates a new ReleaseRunCreateParams object
// with the ability to set a timeout on a request.
func NewReleaseRunCreateParamsWithTimeout(timeout time.Duration) *ReleaseRunCreateParams {
	return &ReleaseRunCreateParams{
		timeout: timeout,
	}
}

// NewReleaseRunCreateParamsWithContext creates a new ReleaseRunCreateParams object
// with the ability to set a context for a request.
func NewReleaseRunCreateParamsWithContext(ctx context.Context) *ReleaseRunCreateParams {
	return &ReleaseRunCreateParams{
		Context: ctx,
	}
}

// NewReleaseRunCreateParamsWithHTTPClient creates a new ReleaseRunCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewReleaseRunCreateParamsWithHTTPClient(client *http.Client) *ReleaseRunCreateParams {
	return &ReleaseRunCreateParams{
		HTTPClient: client,
	}
}

/*
ReleaseRunCreateParams contains all the parameters to send to the API endpoint

	for the release run create operation.

	Typically these are written to a http.Request.
*/
type ReleaseRunCreateParams struct {

	/* XWorkspaceID.

	   workspace id
	*/
	XWorkspaceID string

	/* Request.

	   release run request
	*/
	Request *models.CreateReleaseRunRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the release run create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseRunCreateParams) WithDefaults() *ReleaseRunCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the release run create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseRunCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the release run create params
func (o *ReleaseRunCreateParams) WithTimeout(timeout time.Duration) *ReleaseRunCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the release run create params
func (o *ReleaseRunCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the release run create params
func (o *ReleaseRunCreateParams) WithContext(ctx context.Context) *ReleaseRunCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the release run create params
func (o *ReleaseRunCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the release run create params
func (o *ReleaseRunCreateParams) WithHTTPClient(client *http.Client) *ReleaseRunCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the release run create params
func (o *ReleaseRunCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXWorkspaceID adds the xWorkspaceID to the release run create params
func (o *ReleaseRunCreateParams) WithXWorkspaceID(xWorkspaceID string) *ReleaseRunCreateParams {
	o.SetXWorkspaceID(xWorkspaceID)
	return o
}

// SetXWorkspaceID adds the xWorkspaceId to the release run create params
func (o *ReleaseRunCreateParams) SetXWorkspaceID(xWorkspaceID string) {
	o.XWorkspaceID = xWorkspaceID
}

// WithRequest adds the request to the release run create params
func (o *ReleaseRunCreateParams) WithRequest(request *models.CreateReleaseRunRequest) *ReleaseRunCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the release run create params
func (o *ReleaseRunCreateParams) SetRequest(request *models.CreateReleaseRunRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ReleaseRunCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Workspace-ID
	if err := r.SetHeaderParam("X-Workspace-ID", o.XWorkspaceID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
