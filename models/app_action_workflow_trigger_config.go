// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppActionWorkflowTriggerConfig app action workflow trigger config
//
// swagger:model app.ActionWorkflowTriggerConfig
type AppActionWorkflowTriggerConfig struct {

	// action workflow config id
	ActionWorkflowConfigID string `json:"action_workflow_config_id,omitempty"`

	// this belongs to an app config id
	AppConfigID string `json:"app_config_id,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// cron schedule
	CronSchedule string `json:"cron_schedule,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// individual fields for different types
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app action workflow trigger config
func (m *AppActionWorkflowTriggerConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this app action workflow trigger config based on context it is used
func (m *AppActionWorkflowTriggerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppActionWorkflowTriggerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppActionWorkflowTriggerConfig) UnmarshalBinary(b []byte) error {
	var res AppActionWorkflowTriggerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
