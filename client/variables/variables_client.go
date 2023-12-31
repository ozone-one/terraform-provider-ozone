// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new variables API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for variables API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateVariable creates a new variable

create a new variable
*/
func (a *Client) CreateVariable(params *CreateVariableParams) (*CreateVariableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVariableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createVariable",
		Method:             "POST",
		PathPattern:        "/admin/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateVariableReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateVariableOK), nil

}

/*
GetVariablesList gives all variables

gives all variables
*/
func (a *Client) GetVariablesList(params *GetVariablesListParams) (*GetVariablesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariablesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVariablesList",
		Method:             "GET",
		PathPattern:        "/admin/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVariablesListReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVariablesListOK), nil

}

/*
VariableGetByID gets variable by id

get variable by id
*/
func (a *Client) VariableGetByID(params *VariableGetByIDParams) (*VariableGetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVariableGetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "variableGetByID",
		Method:             "GET",
		PathPattern:        "/admin/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VariableGetByIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VariableGetByIDOK), nil

}

/*
VariableUpdate updates existing variable

updates existing variable
*/
func (a *Client) VariableUpdate(params *VariableUpdateParams) (*VariableUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVariableUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "variableUpdate",
		Method:             "PUT",
		PathPattern:        "/admin/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VariableUpdateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VariableUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
