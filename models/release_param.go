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

// ReleaseParam release param
//
// swagger:model releaseParam
type ReleaseParam struct {

	// params
	Params []*PipelineParam `json:"params"`

	// pipeline ID
	PipelineID string `json:"pipelineID,omitempty"`
}

// Validate validates this release param
func (m *ReleaseParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseParam) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(m.Params) { // not required
		return nil
	}

	for i := 0; i < len(m.Params); i++ {
		if swag.IsZero(m.Params[i]) { // not required
			continue
		}

		if m.Params[i] != nil {
			if err := m.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this release param based on the context it is used
func (m *ReleaseParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseParam) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Params); i++ {

		if m.Params[i] != nil {
			if err := m.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseParam) UnmarshalBinary(b []byte) error {
	var res ReleaseParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
