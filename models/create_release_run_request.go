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

// CreateReleaseRunRequest create release run request
//
// swagger:model createReleaseRunRequest
type CreateReleaseRunRequest struct {

	// name
	Name string `json:"name,omitempty"`

	// release ID
	ReleaseID string `json:"releaseID,omitempty"`

	// steps
	Steps []*ReleaseRunStep `json:"steps"`

	// trigger resource ID
	TriggerResourceID string `json:"triggerResourceID,omitempty"`

	// trigger resource kind
	TriggerResourceKind string `json:"triggerResourceKind,omitempty"`
}

// Validate validates this create release run request
func (m *CreateReleaseRunRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateReleaseRunRequest) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create release run request based on the context it is used
func (m *CreateReleaseRunRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateReleaseRunRequest) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {
			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateReleaseRunRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateReleaseRunRequest) UnmarshalBinary(b []byte) error {
	var res CreateReleaseRunRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
