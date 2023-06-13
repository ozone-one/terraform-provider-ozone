// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new releases API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for releases API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetReleaseByID gets releases of specific application

get releases of specific application
*/
func (a *Client) GetReleaseByID(params *GetReleaseByIDParams) (*GetReleaseByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReleaseByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getReleaseByID",
		Method:             "GET",
		PathPattern:        "/admin/release/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReleaseByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReleaseByIDOK), nil

}

/*
ListReleaseStepTypes gives all applications

gives all applications
*/
func (a *Client) ListReleaseStepTypes(params *ListReleaseStepTypesParams) (*ListReleaseStepTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReleaseStepTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReleaseStepTypes",
		Method:             "GET",
		PathPattern:        "/admin/releasesteptypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReleaseStepTypesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReleaseStepTypesOK), nil

}

/*
ListReleases gives all application releases

gives all application releases
*/
func (a *Client) ListReleases(params *ListReleasesParams) (*ListReleasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listReleases",
		Method:             "GET",
		PathPattern:        "/admin/release",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListReleasesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReleasesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}