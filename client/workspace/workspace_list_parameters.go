// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

// NewWorkspaceListParams creates a new WorkspaceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkspaceListParams() *WorkspaceListParams {
	return &WorkspaceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkspaceListParamsWithTimeout creates a new WorkspaceListParams object
// with the ability to set a timeout on a request.
func NewWorkspaceListParamsWithTimeout(timeout time.Duration) *WorkspaceListParams {
	return &WorkspaceListParams{
		timeout: timeout,
	}
}

// NewWorkspaceListParamsWithContext creates a new WorkspaceListParams object
// with the ability to set a context for a request.
func NewWorkspaceListParamsWithContext(ctx context.Context) *WorkspaceListParams {
	return &WorkspaceListParams{
		Context: ctx,
	}
}

// NewWorkspaceListParamsWithHTTPClient creates a new WorkspaceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkspaceListParamsWithHTTPClient(client *http.Client) *WorkspaceListParams {
	return &WorkspaceListParams{
		HTTPClient: client,
	}
}

/*
WorkspaceListParams contains all the parameters to send to the API endpoint

	for the workspace list operation.

	Typically these are written to a http.Request.
*/
type WorkspaceListParams struct {

	/* NameSlug.

	   find by workspace name
	*/
	NameSlug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workspace list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceListParams) WithDefaults() *WorkspaceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workspace list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkspaceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workspace list params
func (o *WorkspaceListParams) WithTimeout(timeout time.Duration) *WorkspaceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workspace list params
func (o *WorkspaceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workspace list params
func (o *WorkspaceListParams) WithContext(ctx context.Context) *WorkspaceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workspace list params
func (o *WorkspaceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workspace list params
func (o *WorkspaceListParams) WithHTTPClient(client *http.Client) *WorkspaceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workspace list params
func (o *WorkspaceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNameSlug adds the nameSlug to the workspace list params
func (o *WorkspaceListParams) WithNameSlug(nameSlug *string) *WorkspaceListParams {
	o.SetNameSlug(nameSlug)
	return o
}

// SetNameSlug adds the nameSlug to the workspace list params
func (o *WorkspaceListParams) SetNameSlug(nameSlug *string) {
	o.NameSlug = nameSlug
}

// WriteToRequest writes these params to a swagger request
func (o *WorkspaceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NameSlug != nil {

		// query param name_slug
		var qrNameSlug string

		if o.NameSlug != nil {
			qrNameSlug = *o.NameSlug
		}
		qNameSlug := qrNameSlug
		if qNameSlug != "" {

			if err := r.SetQueryParam("name_slug", qNameSlug); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
