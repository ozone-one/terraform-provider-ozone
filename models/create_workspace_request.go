// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateWorkspaceRequest create workspace request
//
// swagger:model createWorkspaceRequest
type CreateWorkspaceRequest struct {

	// admin member ids
	AdminMemberIds []string `json:"admin_member_ids"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create workspace request
func (m *CreateWorkspaceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create workspace request based on context it is used
func (m *CreateWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res CreateWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
