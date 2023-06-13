// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewApplicationCreateParams creates a new ApplicationCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationCreateParams() *ApplicationCreateParams {
	return &ApplicationCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationCreateParamsWithTimeout creates a new ApplicationCreateParams object
// with the ability to set a timeout on a request.
func NewApplicationCreateParamsWithTimeout(timeout time.Duration) *ApplicationCreateParams {
	return &ApplicationCreateParams{
		timeout: timeout,
	}
}

// NewApplicationCreateParamsWithContext creates a new ApplicationCreateParams object
// with the ability to set a context for a request.
func NewApplicationCreateParamsWithContext(ctx context.Context) *ApplicationCreateParams {
	return &ApplicationCreateParams{
		Context: ctx,
	}
}

// NewApplicationCreateParamsWithHTTPClient creates a new ApplicationCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationCreateParamsWithHTTPClient(client *http.Client) *ApplicationCreateParams {
	return &ApplicationCreateParams{
		HTTPClient: client,
	}
}

/*
ApplicationCreateParams contains all the parameters to send to the API endpoint

	for the application create operation.

	Typically these are written to a http.Request.
*/
type ApplicationCreateParams struct {

	/* XWorkspaceID.

	   workspace id
	*/
	XWorkspaceID string

	/* Request.

	   create application request
	*/
	Request *models.CreateApplicationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationCreateParams) WithDefaults() *ApplicationCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the application create params
func (o *ApplicationCreateParams) WithTimeout(timeout time.Duration) *ApplicationCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application create params
func (o *ApplicationCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application create params
func (o *ApplicationCreateParams) WithContext(ctx context.Context) *ApplicationCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application create params
func (o *ApplicationCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application create params
func (o *ApplicationCreateParams) WithHTTPClient(client *http.Client) *ApplicationCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application create params
func (o *ApplicationCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXWorkspaceID adds the xWorkspaceID to the application create params
func (o *ApplicationCreateParams) WithXWorkspaceID(xWorkspaceID string) *ApplicationCreateParams {
	o.SetXWorkspaceID(xWorkspaceID)
	return o
}

// SetXWorkspaceID adds the xWorkspaceId to the application create params
func (o *ApplicationCreateParams) SetXWorkspaceID(xWorkspaceID string) {
	o.XWorkspaceID = xWorkspaceID
}

// WithRequest adds the request to the application create params
func (o *ApplicationCreateParams) WithRequest(request *models.CreateApplicationRequest) *ApplicationCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the application create params
func (o *ApplicationCreateParams) SetRequest(request *models.CreateApplicationRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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