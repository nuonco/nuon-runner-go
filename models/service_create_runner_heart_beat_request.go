// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceCreateRunnerHeartBeatRequest service create runner heart beat request
//
// swagger:model service.CreateRunnerHeartBeatRequest
type ServiceCreateRunnerHeartBeatRequest struct {

	// alive time
	// Required: true
	AliveTime *TimeDuration `json:"alive_time"`

	// jobs in flight
	// Required: true
	JobsInFlight *int64 `json:"jobs_in_flight"`
}

// Validate validates this service create runner heart beat request
func (m *ServiceCreateRunnerHeartBeatRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAliveTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobsInFlight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateRunnerHeartBeatRequest) validateAliveTime(formats strfmt.Registry) error {

	if err := validate.Required("alive_time", "body", m.AliveTime); err != nil {
		return err
	}

	if err := validate.Required("alive_time", "body", m.AliveTime); err != nil {
		return err
	}

	if m.AliveTime != nil {
		if err := m.AliveTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alive_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alive_time")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCreateRunnerHeartBeatRequest) validateJobsInFlight(formats strfmt.Registry) error {

	if err := validate.Required("jobs_in_flight", "body", m.JobsInFlight); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service create runner heart beat request based on the context it is used
func (m *ServiceCreateRunnerHeartBeatRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAliveTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateRunnerHeartBeatRequest) contextValidateAliveTime(ctx context.Context, formats strfmt.Registry) error {

	if m.AliveTime != nil {

		if err := m.AliveTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alive_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alive_time")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateRunnerHeartBeatRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateRunnerHeartBeatRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateRunnerHeartBeatRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
