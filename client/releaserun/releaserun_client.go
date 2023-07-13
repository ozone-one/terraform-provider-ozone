// Code generated by go-swagger; DO NOT EDIT.

package releaserun

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new releaserun API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for releaserun API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetReleaseRunByID gives all applications

gives release run by id
*/
func (a *Client) GetReleaseRunByID(params *GetReleaseRunByIDParams) (*GetReleaseRunByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReleaseRunByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReleaseRunByID",
		Method:             "GET",
		PathPattern:        "/admin/releaserun/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReleaseRunByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReleaseRunByIDOK), nil

}

/*
ListReleaseRuns gives all applications

gives release runs
*/
func (a *Client) ListReleaseRuns(params *ListReleaseRunsParams) (*ListReleaseRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReleaseRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReleaseRuns",
		Method:             "GET",
		PathPattern:        "/admin/releaserun",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReleaseRunsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReleaseRunsOK), nil

}

/*
ReleaseRunCreate creates a new releaserun

gives release run by id
*/
func (a *Client) ReleaseRunCreate(params *ReleaseRunCreateParams) (*ReleaseRunCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReleaseRunCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "releaseRunCreate",
		Method:             "POST",
		PathPattern:        "/admin/releaserun",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReleaseRunCreateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReleaseRunCreateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
