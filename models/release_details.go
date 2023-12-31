// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReleaseDetails release details
//
// swagger:model releaseDetails
type ReleaseDetails struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this release details
func (m *ReleaseDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this release details based on context it is used
func (m *ReleaseDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseDetails) UnmarshalBinary(b []byte) error {
	var res ReleaseDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
