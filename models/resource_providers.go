// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceProviders resource providers
//
// swagger:model resourceProviders
type ResourceProviders struct {

	// items
	Items []string `json:"items"`
}

// Validate validates this resource providers
func (m *ResourceProviders) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this resource providers based on context it is used
func (m *ResourceProviders) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceProviders) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceProviders) UnmarshalBinary(b []byte) error {
	var res ResourceProviders
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
