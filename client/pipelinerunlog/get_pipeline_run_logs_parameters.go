// Code generated by go-swagger; DO NOT EDIT.

package pipelinerunlog

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
	"github.com/go-openapi/swag"
)

// NewGetPipelineRunLogsParams creates a new GetPipelineRunLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPipelineRunLogsParams() *GetPipelineRunLogsParams {
	return &GetPipelineRunLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineRunLogsParamsWithTimeout creates a new GetPipelineRunLogsParams object
// with the ability to set a timeout on a request.
func NewGetPipelineRunLogsParamsWithTimeout(timeout time.Duration) *GetPipelineRunLogsParams {
	return &GetPipelineRunLogsParams{
		timeout: timeout,
	}
}

// NewGetPipelineRunLogsParamsWithContext creates a new GetPipelineRunLogsParams object
// with the ability to set a context for a request.
func NewGetPipelineRunLogsParamsWithContext(ctx context.Context) *GetPipelineRunLogsParams {
	return &GetPipelineRunLogsParams{
		Context: ctx,
	}
}

// NewGetPipelineRunLogsParamsWithHTTPClient creates a new GetPipelineRunLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPipelineRunLogsParamsWithHTTPClient(client *http.Client) *GetPipelineRunLogsParams {
	return &GetPipelineRunLogsParams{
		HTTPClient: client,
	}
}

/*
GetPipelineRunLogsParams contains all the parameters to send to the API endpoint

	for the get pipeline run logs operation.

	Typically these are written to a http.Request.
*/
type GetPipelineRunLogsParams struct {

	/* End.

	   ends at
	*/
	End *int64

	/* Start.

	   starts at
	*/
	Start *int64

	/* PipelineRunID.

	   pipeline run id
	*/
	PipelineRunID string

	/* StepName.

	   step name in the respective task run
	*/
	StepName *string

	/* TaskRunName.

	   pipeline task run name
	*/
	TaskRunName *string

	/* Text.

	   search query text
	*/
	Text *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pipeline run logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineRunLogsParams) WithDefaults() *GetPipelineRunLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pipeline run logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineRunLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithTimeout(timeout time.Duration) *GetPipelineRunLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithContext(ctx context.Context) *GetPipelineRunLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithHTTPClient(client *http.Client) *GetPipelineRunLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnd adds the end to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithEnd(end *int64) *GetPipelineRunLogsParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetEnd(end *int64) {
	o.End = end
}

// WithStart adds the start to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithStart(start *int64) *GetPipelineRunLogsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetStart(start *int64) {
	o.Start = start
}

// WithPipelineRunID adds the pipelineRunID to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithPipelineRunID(pipelineRunID string) *GetPipelineRunLogsParams {
	o.SetPipelineRunID(pipelineRunID)
	return o
}

// SetPipelineRunID adds the pipelineRunId to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetPipelineRunID(pipelineRunID string) {
	o.PipelineRunID = pipelineRunID
}

// WithStepName adds the stepName to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithStepName(stepName *string) *GetPipelineRunLogsParams {
	o.SetStepName(stepName)
	return o
}

// SetStepName adds the stepName to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetStepName(stepName *string) {
	o.StepName = stepName
}

// WithTaskRunName adds the taskRunName to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithTaskRunName(taskRunName *string) *GetPipelineRunLogsParams {
	o.SetTaskRunName(taskRunName)
	return o
}

// SetTaskRunName adds the taskRunName to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetTaskRunName(taskRunName *string) {
	o.TaskRunName = taskRunName
}

// WithText adds the text to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) WithText(text *string) *GetPipelineRunLogsParams {
	o.SetText(text)
	return o
}

// SetText adds the text to the get pipeline run logs params
func (o *GetPipelineRunLogsParams) SetText(text *string) {
	o.Text = text
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineRunLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.End != nil {

		// query param _end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("_end", qEnd); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param _start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("_start", qStart); err != nil {
				return err
			}
		}
	}

	// path param pipeline_run_id
	if err := r.SetPathParam("pipeline_run_id", o.PipelineRunID); err != nil {
		return err
	}

	if o.StepName != nil {

		// query param stepName
		var qrStepName string

		if o.StepName != nil {
			qrStepName = *o.StepName
		}
		qStepName := qrStepName
		if qStepName != "" {

			if err := r.SetQueryParam("stepName", qStepName); err != nil {
				return err
			}
		}
	}

	if o.TaskRunName != nil {

		// query param taskRunName
		var qrTaskRunName string

		if o.TaskRunName != nil {
			qrTaskRunName = *o.TaskRunName
		}
		qTaskRunName := qrTaskRunName
		if qTaskRunName != "" {

			if err := r.SetQueryParam("taskRunName", qTaskRunName); err != nil {
				return err
			}
		}
	}

	if o.Text != nil {

		// query param text
		var qrText string

		if o.Text != nil {
			qrText = *o.Text
		}
		qText := qrText
		if qText != "" {

			if err := r.SetQueryParam("text", qText); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}