// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TemplateVisibility template visibility
//
// swagger:model templateVisibility
type TemplateVisibility struct {

	// show to user
	ShowToUser bool `json:"showToUser,omitempty"`
}

// Validate validates this template visibility
func (m *TemplateVisibility) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this template visibility based on context it is used
func (m *TemplateVisibility) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TemplateVisibility) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateVisibility) UnmarshalBinary(b []byte) error {
	var res TemplateVisibility
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}