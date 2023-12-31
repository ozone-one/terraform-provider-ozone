// Code generated by go-swagger; DO NOT EDIT.

package pipelinerun

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pipelinerun API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for pipelinerun API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
PipelineRunCreate runs application pipeline

runs application pipeline
*/
func (a *Client) PipelineRunCreate(params *PipelineRunCreateParams) (*PipelineRunCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPipelineRunCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pipelineRunCreate",
		Method:             "POST",
		PathPattern:        "/admin/pipelinerun",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PipelineRunCreateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PipelineRunCreateOK), nil

}

/*
PipelineRunGet gives all pipeline runs

gives all pipeline runs
*/
func (a *Client) PipelineRunGet(params *PipelineRunGetParams) (*PipelineRunGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPipelineRunGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pipelineRunGet",
		Method:             "GET",
		PathPattern:        "/admin/pipelinerun",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PipelineRunGetReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PipelineRunGetOK), nil

}

/*
PipelineRunGetByID gives all pipeline runs

gives all pipeline runs
*/
func (a *Client) PipelineRunGetByID(params *PipelineRunGetByIDParams) (*PipelineRunGetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPipelineRunGetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pipelineRunGetByID",
		Method:             "GET",
		PathPattern:        "/admin/pipelinerun/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PipelineRunGetByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PipelineRunGetByIDOK), nil

}

/*
ReRunPipelineRun retrieves pipeline run logs

retrieves pipeline run logs
*/
func (a *Client) ReRunPipelineRun(params *ReRunPipelineRunParams) (*ReRunPipelineRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReRunPipelineRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "re-run-pipeline-run",
		Method:             "POST",
		PathPattern:        "/admin/pipelinerun/{pipeline_run_id}/rerun",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReRunPipelineRunReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReRunPipelineRunOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
