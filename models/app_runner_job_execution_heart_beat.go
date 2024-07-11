// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppRunnerJobExecutionHeartBeat app runner job execution heart beat
//
// swagger:model app.RunnerJobExecutionHeartBeat
type AppRunnerJobExecutionHeartBeat struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// runner id
	RunnerID string `json:"runner_id,omitempty"`

	// runner job execution id
	RunnerJobExecutionID string `json:"runner_job_execution_id,omitempty"`

	// runner job id
	RunnerJobID string `json:"runner_job_id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app runner job execution heart beat
func (m *AppRunnerJobExecutionHeartBeat) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app runner job execution heart beat based on context it is used
func (m *AppRunnerJobExecutionHeartBeat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppRunnerJobExecutionHeartBeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppRunnerJobExecutionHeartBeat) UnmarshalBinary(b []byte) error {
	var res AppRunnerJobExecutionHeartBeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}