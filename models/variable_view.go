// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VariableView variable view
//
// swagger:model variableView
type VariableView struct {

	// account id
	AccountID string `json:"account_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// created by name
	CreatedByName string `json:"created_by_name,omitempty"`

	// deleted at
	DeletedAt string `json:"deleted_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// env values
	EnvValues map[string]string `json:"envValues,omitempty"`

	// global
	Global bool `json:"global,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by id
	UpdatedByID string `json:"updated_by_id,omitempty"`

	// updated by name
	UpdatedByName string `json:"updated_by_name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this variable view
func (m *VariableView) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this variable view based on context it is used
func (m *VariableView) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VariableView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableView) UnmarshalBinary(b []byte) error {
	var res VariableView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
