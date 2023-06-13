// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeploymentType deployment type
//
// swagger:model deploymentType
type DeploymentType struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// deleted at
	DeletedAt string `json:"deleted_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this deployment type
func (m *DeploymentType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this deployment type based on context it is used
func (m *DeploymentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentType) UnmarshalBinary(b []byte) error {
	var res DeploymentType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
