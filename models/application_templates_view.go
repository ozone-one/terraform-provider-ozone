// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationTemplatesView application templates view
//
// swagger:model applicationTemplatesView
type ApplicationTemplatesView struct {

	// account id
	AccountID string `json:"account_id,omitempty"`

	// application ID
	ApplicationID string `json:"applicationID,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// created by name
	CreatedByName string `json:"created_by_name,omitempty"`

	// deleted at
	DeletedAt string `json:"deleted_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// items
	Items []*ApplicationTemplateItem `json:"items"`

	// name
	Name string `json:"name,omitempty"`

	// params
	Params *TemplateBootstrapValues `json:"params,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// updated by id
	UpdatedByID string `json:"updated_by_id,omitempty"`

	// updated by name
	UpdatedByName string `json:"updated_by_name,omitempty"`
}

// Validate validates this application templates view
func (m *ApplicationTemplatesView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationTemplatesView) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationTemplatesView) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(m.Params) { // not required
		return nil
	}

	if m.Params != nil {
		if err := m.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application templates view based on the context it is used
func (m *ApplicationTemplatesView) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationTemplatesView) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationTemplatesView) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	if m.Params != nil {
		if err := m.Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationTemplatesView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationTemplatesView) UnmarshalBinary(b []byte) error {
	var res ApplicationTemplatesView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
