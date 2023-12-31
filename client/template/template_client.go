// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new template API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for template API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ApplicationCreateTemplate creates template set for custom app

create Template Set for Custom App
*/
func (a *Client) ApplicationCreateTemplate(params *ApplicationCreateTemplateParams) (*ApplicationCreateTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationCreateTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applicationCreateTemplate",
		Method:             "POST",
		PathPattern:        "/admin/app/template",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplicationCreateTemplateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationCreateTemplateOK), nil

}

/*
ApplicationsDeleteTemplate deletes app template

delete app template
*/
func (a *Client) ApplicationsDeleteTemplate(params *ApplicationsDeleteTemplateParams) (*ApplicationsDeleteTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationsDeleteTemplateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applicationsDeleteTemplate",
		Method:             "DELETE",
		PathPattern:        "/admin/app/template",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplicationsDeleteTemplateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationsDeleteTemplateOK), nil

}

/*
ApplicationsListTemplates gives all application templates

gives all application templates
*/
func (a *Client) ApplicationsListTemplates(params *ApplicationsListTemplatesParams) (*ApplicationsListTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationsListTemplatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "applicationsListTemplates",
		Method:             "GET",
		PathPattern:        "/admin/app/template",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplicationsListTemplatesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ApplicationsListTemplatesOK), nil

}

/*
GetTemplates gives all task templates

gives all task templates
*/
func (a *Client) GetTemplates(params *GetTemplatesParams) (*GetTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTemplatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTemplates",
		Method:             "GET",
		PathPattern:        "/admin/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTemplatesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTemplatesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
