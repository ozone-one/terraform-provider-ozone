// Code generated by go-swagger; DO NOT EDIT.

package pipeline

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

// NewPipelineGetByIDParams creates a new PipelineGetByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPipelineGetByIDParams() *PipelineGetByIDParams {
	return &PipelineGetByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPipelineGetByIDParamsWithTimeout creates a new PipelineGetByIDParams object
// with the ability to set a timeout on a request.
func NewPipelineGetByIDParamsWithTimeout(timeout time.Duration) *PipelineGetByIDParams {
	return &PipelineGetByIDParams{
		timeout: timeout,
	}
}

// NewPipelineGetByIDParamsWithContext creates a new PipelineGetByIDParams object
// with the ability to set a context for a request.
func NewPipelineGetByIDParamsWithContext(ctx context.Context) *PipelineGetByIDParams {
	return &PipelineGetByIDParams{
		Context: ctx,
	}
}

// NewPipelineGetByIDParamsWithHTTPClient creates a new PipelineGetByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPipelineGetByIDParamsWithHTTPClient(client *http.Client) *PipelineGetByIDParams {
	return &PipelineGetByIDParams{
		HTTPClient: client,
	}
}

/*
PipelineGetByIDParams contains all the parameters to send to the API endpoint

	for the pipeline get by ID operation.

	Typically these are written to a http.Request.
*/
type PipelineGetByIDParams struct {

	/* ID.

	   pipeline id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pipeline get by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineGetByIDParams) WithDefaults() *PipelineGetByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pipeline get by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PipelineGetByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pipeline get by ID params
func (o *PipelineGetByIDParams) WithTimeout(timeout time.Duration) *PipelineGetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pipeline get by ID params
func (o *PipelineGetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pipeline get by ID params
func (o *PipelineGetByIDParams) WithContext(ctx context.Context) *PipelineGetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pipeline get by ID params
func (o *PipelineGetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pipeline get by ID params
func (o *PipelineGetByIDParams) WithHTTPClient(client *http.Client) *PipelineGetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pipeline get by ID params
func (o *PipelineGetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the pipeline get by ID params
func (o *PipelineGetByIDParams) WithID(id string) *PipelineGetByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the pipeline get by ID params
func (o *PipelineGetByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PipelineGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}