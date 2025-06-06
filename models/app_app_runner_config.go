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

// AppAppRunnerConfig app app runner config
//
// swagger:model app.AppRunnerConfig
type AppAppRunnerConfig struct {

	// app config id
	AppConfigID string `json:"app_config_id,omitempty"`

	// app id
	AppID string `json:"app_id,omitempty"`

	// app runner type
	AppRunnerType AppAppRunnerType `json:"app_runner_type,omitempty"`

	// cloud platform
	CloudPlatform AppCloudPlatform `json:"cloud_platform,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// env vars
	EnvVars map[string]string `json:"env_vars,omitempty"`

	// helm driver
	HelmDriver string `json:"helm_driver,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app app runner config
func (m *AppAppRunnerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppRunnerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppRunnerConfig) validateAppRunnerType(formats strfmt.Registry) error {
	if swag.IsZero(m.AppRunnerType) { // not required
		return nil
	}

	if err := m.AppRunnerType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("app_runner_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("app_runner_type")
		}
		return err
	}

	return nil
}

func (m *AppAppRunnerConfig) validateCloudPlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	if err := m.CloudPlatform.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloud_platform")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloud_platform")
		}
		return err
	}

	return nil
}

// ContextValidate validate this app app runner config based on the context it is used
func (m *AppAppRunnerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppRunnerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudPlatform(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppRunnerConfig) contextValidateAppRunnerType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AppRunnerType) { // not required
		return nil
	}

	if err := m.AppRunnerType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("app_runner_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("app_runner_type")
		}
		return err
	}

	return nil
}

func (m *AppAppRunnerConfig) contextValidateCloudPlatform(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	if err := m.CloudPlatform.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloud_platform")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloud_platform")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppAppRunnerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAppRunnerConfig) UnmarshalBinary(b []byte) error {
	var res AppAppRunnerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
