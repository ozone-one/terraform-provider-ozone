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
)

// NewListApplicationReleaseRunParams creates a new ListApplicationReleaseRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListApplicationReleaseRunParams() *ListApplicationReleaseRunParams {
	return &ListApplicationReleaseRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListApplicationReleaseRunParamsWithTimeout creates a new ListApplicationReleaseRunParams object
// with the ability to set a timeout on a request.
func NewListApplicationReleaseRunParamsWithTimeout(timeout time.Duration) *ListApplicationReleaseRunParams {
	return &ListApplicationReleaseRunParams{
		timeout: timeout,
	}
}

// NewListApplicationReleaseRunParamsWithContext creates a new ListApplicationReleaseRunParams object
// with the ability to set a context for a request.
func NewListApplicationReleaseRunParamsWithContext(ctx context.Context) *ListApplicationReleaseRunParams {
	return &ListApplicationReleaseRunParams{
		Context: ctx,
	}
}

// NewListApplicationReleaseRunParamsWithHTTPClient creates a new ListApplicationReleaseRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewListApplicationReleaseRunParamsWithHTTPClient(client *http.Client) *ListApplicationReleaseRunParams {
	return &ListApplicationReleaseRunParams{
		HTTPClient: client,
	}
}

/*
ListApplicationReleaseRunParams contains all the parameters to send to the API endpoint

	for the list application release run operation.

	Typically these are written to a http.Request.
*/
type ListApplicationReleaseRunParams struct {

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

// WithDefaults hydrates default values in the list application release run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListApplicationReleaseRunParams) WithDefaults() *ListApplicationReleaseRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list application release run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListApplicationReleaseRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list application release run params
func (o *ListApplicationReleaseRunParams) WithTimeout(timeout time.Duration) *ListApplicationReleaseRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list application release run params
func (o *ListApplicationReleaseRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list application release run params
func (o *ListApplicationReleaseRunParams) WithContext(ctx context.Context) *ListApplicationReleaseRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list application release run params
func (o *ListApplicationReleaseRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list application release run params
func (o *ListApplicationReleaseRunParams) WithHTTPClient(client *http.Client) *ListApplicationReleaseRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list application release run params
func (o *ListApplicationReleaseRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXWorkspaceID adds the xWorkspaceID to the list application release run params
func (o *ListApplicationReleaseRunParams) WithXWorkspaceID(xWorkspaceID string) *ListApplicationReleaseRunParams {
	o.SetXWorkspaceID(xWorkspaceID)
	return o
}

// SetXWorkspaceID adds the xWorkspaceId to the list application release run params
func (o *ListApplicationReleaseRunParams) SetXWorkspaceID(xWorkspaceID string) {
	o.XWorkspaceID = xWorkspaceID
}

// WithID adds the id to the list application release run params
func (o *ListApplicationReleaseRunParams) WithID(id string) *ListApplicationReleaseRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list application release run params
func (o *ListApplicationReleaseRunParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ListApplicationReleaseRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
