// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationAPM application a p m
//
// swagger:model applicationAPM
type ApplicationAPM struct {

	// apm type
	ApmType int64 `json:"apmType,omitempty"`

	// datadog
	Datadog *DatadogApplicationAPM `json:"datadog,omitempty"`

	// dynatrace
	Dynatrace *DynatraceApplicationAPM `json:"dynatrace,omitempty"`

	// elastic
	Elastic *ElasticApplicationAPM `json:"elastic,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// newrelic
	Newrelic *NewRelicApplicationAPM `json:"newrelic,omitempty"`
}

// Validate validates this application a p m
func (m *ApplicationAPM) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatadog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynatrace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElastic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewrelic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationAPM) validateDatadog(formats strfmt.Registry) error {
	if swag.IsZero(m.Datadog) { // not required
		return nil
	}

	if m.Datadog != nil {
		if err := m.Datadog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datadog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datadog")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAPM) validateDynatrace(formats strfmt.Registry) error {
	if swag.IsZero(m.Dynatrace) { // not required
		return nil
	}

	if m.Dynatrace != nil {
		if err := m.Dynatrace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynatrace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynatrace")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAPM) validateElastic(formats strfmt.Registry) error {
	if swag.IsZero(m.Elastic) { // not required
		return nil
	}

	if m.Elastic != nil {
		if err := m.Elastic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elastic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elastic")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAPM) validateNewrelic(formats strfmt.Registry) error {
	if swag.IsZero(m.Newrelic) { // not required
		return nil
	}

	if m.Newrelic != nil {
		if err := m.Newrelic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newrelic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newrelic")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application a p m based on the context it is used
func (m *ApplicationAPM) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatadog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDynatrace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElastic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNewrelic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationAPM) contextValidateDatadog(ctx context.Context, formats strfmt.Registry) error {

	if m.Datadog != nil {
		if err := m.Datadog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datadog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datadog")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAPM) contextValidateDynatrace(ctx context.Context, formats strfmt.Registry) error {

	if m.Dynatrace != nil {
		if err := m.Dynatrace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynatrace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynatrace")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAPM) contextValidateElastic(ctx context.Context, formats strfmt.Registry) error {

	if m.Elastic != nil {
		if err := m.Elastic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elastic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elastic")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationAPM) contextValidateNewrelic(ctx context.Context, formats strfmt.Registry) error {

	if m.Newrelic != nil {
		if err := m.Newrelic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("newrelic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("newrelic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationAPM) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationAPM) UnmarshalBinary(b []byte) error {
	var res ApplicationAPM
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
