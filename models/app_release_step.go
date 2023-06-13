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

// AppReleaseStep app release step
//
// swagger:model appReleaseStep
type AppReleaseStep struct {

	// description
	Description string `json:"description,omitempty"`

	// jira approval
	JiraApproval *ReleaseRunJiraApprovalStep `json:"jiraApproval,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// run after
	RunAfter []string `json:"runAfter"`

	// slack approval
	SlackApproval *ReleaseRunSlackApprovalStep `json:"slackApproval,omitempty"`

	// type ID
	TypeID int64 `json:"typeID,omitempty"`

	// workspaces
	Workspaces *MongoResourceWorkspaces `json:"workspaces,omitempty"`
}

// Validate validates this app release step
func (m *AppReleaseStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJiraApproval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackApproval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppReleaseStep) validateJiraApproval(formats strfmt.Registry) error {
	if swag.IsZero(m.JiraApproval) { // not required
		return nil
	}

	if m.JiraApproval != nil {
		if err := m.JiraApproval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jiraApproval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jiraApproval")
			}
			return err
		}
	}

	return nil
}

func (m *AppReleaseStep) validateSlackApproval(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackApproval) { // not required
		return nil
	}

	if m.SlackApproval != nil {
		if err := m.SlackApproval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slackApproval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slackApproval")
			}
			return err
		}
	}

	return nil
}

func (m *AppReleaseStep) validateWorkspaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspaces) { // not required
		return nil
	}

	if m.Workspaces != nil {
		if err := m.Workspaces.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspaces")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspaces")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app release step based on the context it is used
func (m *AppReleaseStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJiraApproval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackApproval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppReleaseStep) contextValidateJiraApproval(ctx context.Context, formats strfmt.Registry) error {

	if m.JiraApproval != nil {
		if err := m.JiraApproval.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jiraApproval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jiraApproval")
			}
			return err
		}
	}

	return nil
}

func (m *AppReleaseStep) contextValidateSlackApproval(ctx context.Context, formats strfmt.Registry) error {

	if m.SlackApproval != nil {
		if err := m.SlackApproval.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slackApproval")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slackApproval")
			}
			return err
		}
	}

	return nil
}

func (m *AppReleaseStep) contextValidateWorkspaces(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspaces != nil {
		if err := m.Workspaces.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspaces")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspaces")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppReleaseStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppReleaseStep) UnmarshalBinary(b []byte) error {
	var res AppReleaseStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}