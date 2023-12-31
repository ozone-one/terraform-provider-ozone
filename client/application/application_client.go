// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new application API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for application API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ApplicationCreate creates new application

creates new application
*/
func (a *Client) ApplicationCreate(params *ApplicationCreateParams) (*ApplicationCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applicationCreate",
		Method:             "POST",
		PathPattern:        "/admin/app",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplicationCreateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationCreateOK), nil

}

/*
ApplicationGetByID gets application by id

get application by id
*/
func (a *Client) ApplicationGetByID(params *ApplicationGetByIDParams) (*ApplicationGetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationGetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applicationGetByID",
		Method:             "GET",
		PathPattern:        "/admin/app/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplicationGetByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationGetByIDOK), nil

}

/*
ApplicationsList gives all applications

gives all applications
*/
func (a *Client) ApplicationsList(params *ApplicationsListParams) (*ApplicationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applicationsList",
		Method:             "GET",
		PathPattern:        "/admin/app",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplicationsListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationsListOK), nil

}

/*
ApplicationsUpdate updates existing application

updates existing application
*/
func (a *Client) ApplicationsUpdate(params *ApplicationsUpdateParams) (*ApplicationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applicationsUpdate",
		Method:             "PUT",
		PathPattern:        "/admin/app/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplicationsUpdateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationsUpdateOK), nil

}

/*
ListApplicationReleaseRun gets release runs of specific application

get release runs of specific application
*/
func (a *Client) ListApplicationReleaseRun(params *ListApplicationReleaseRunParams) (*ListApplicationReleaseRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListApplicationReleaseRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listApplicationReleaseRun",
		Method:             "GET",
		PathPattern:        "/admin/app/{id}/releaserun",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListApplicationReleaseRunReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListApplicationReleaseRunOK), nil

}

/*
ListApplicationReleases gets releases of specific application

get releases of specific application
*/
func (a *Client) ListApplicationReleases(params *ListApplicationReleasesParams) (*ListApplicationReleasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListApplicationReleasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listApplicationReleases",
		Method:             "GET",
		PathPattern:        "/admin/app/{id}/releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListApplicationReleasesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListApplicationReleasesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
