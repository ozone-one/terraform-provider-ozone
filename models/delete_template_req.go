// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteTemplateReq delete template req
//
// swagger:model deleteTemplateReq
type DeleteTemplateReq struct {

	// application ID
	ApplicationID string `json:"applicationID,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this delete template req
func (m *DeleteTemplateReq) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete template req based on context it is used
func (m *DeleteTemplateReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteTemplateReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteTemplateReq) UnmarshalBinary(b []byte) error {
	var res DeleteTemplateReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
