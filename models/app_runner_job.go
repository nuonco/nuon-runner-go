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

// AppRunnerJob app runner job
//
// swagger:model app.RunnerJob
type AppRunnerJob struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// execution timeout
	ExecutionTimeout int64 `json:"execution_timeout,omitempty"`

	// executions
	Executions []*AppRunnerJobExecution `json:"executions"`

	// id
	ID string `json:"id,omitempty"`

	// max executions
	MaxExecutions int64 `json:"max_executions,omitempty"`

	// org
	Org *AppOrg `json:"org,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// overall timeout
	OverallTimeout int64 `json:"overall_timeout,omitempty"`

	// queue timeout
	QueueTimeout int64 `json:"queue_timeout,omitempty"`

	// runner id
	RunnerID string `json:"runner_id,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app runner job
func (m *AppRunnerJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRunnerJob) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *AppRunnerJob) validateExecutions(formats strfmt.Registry) error {
	if swag.IsZero(m.Executions) { // not required
		return nil
	}

	for i := 0; i < len(m.Executions); i++ {
		if swag.IsZero(m.Executions[i]) { // not required
			continue
		}

		if m.Executions[i] != nil {
			if err := m.Executions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("executions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("executions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppRunnerJob) validateOrg(formats strfmt.Registry) error {
	if swag.IsZero(m.Org) { // not required
		return nil
	}

	if m.Org != nil {
		if err := m.Org.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("org")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("org")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app runner job based on the context it is used
func (m *AppRunnerJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExecutions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppRunnerJob) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {

		if swag.IsZero(m.CreatedBy) { // not required
			return nil
		}

		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *AppRunnerJob) contextValidateExecutions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Executions); i++ {

		if m.Executions[i] != nil {

			if swag.IsZero(m.Executions[i]) { // not required
				return nil
			}

			if err := m.Executions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("executions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("executions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppRunnerJob) contextValidateOrg(ctx context.Context, formats strfmt.Registry) error {

	if m.Org != nil {

		if swag.IsZero(m.Org) { // not required
			return nil
		}

		if err := m.Org.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("org")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("org")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppRunnerJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRunnerJob) UnmarshalBinary(b []byte) error {
	var res AppRunnerJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
