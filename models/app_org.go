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

// AppOrg app org
//
// swagger:model app.Org
type AppOrg struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// health checks
	HealthChecks []*AppOrgHealthCheck `json:"health_checks"`

	// id
	ID string `json:"id,omitempty"`

	// Filled in at read time
	LatestHealthCheck struct {
		AppOrgHealthCheck
	} `json:"latest_health_check,omitempty"`

	// logo url
	LogoURL string `json:"logo_url,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// notifications config
	NotificationsConfig *AppNotificationsConfig `json:"notifications_config,omitempty"`

	// runner group
	RunnerGroup *AppRunnerGroup `json:"runner_group,omitempty"`

	// sandbox mode
	SandboxMode bool `json:"sandbox_mode,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// vcs connections
	VcsConnections []*AppVCSConnection `json:"vcs_connections"`
}

// Validate validates this app org
func (m *AppOrg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestHealthCheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunnerGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppOrg) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *AppOrg) validateHealthChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthChecks) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthChecks); i++ {
		if swag.IsZero(m.HealthChecks[i]) { // not required
			continue
		}

		if m.HealthChecks[i] != nil {
			if err := m.HealthChecks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("health_checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("health_checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppOrg) validateLatestHealthCheck(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestHealthCheck) { // not required
		return nil
	}

	return nil
}

func (m *AppOrg) validateNotificationsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NotificationsConfig) { // not required
		return nil
	}

	if m.NotificationsConfig != nil {
		if err := m.NotificationsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppOrg) validateRunnerGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.RunnerGroup) { // not required
		return nil
	}

	if m.RunnerGroup != nil {
		if err := m.RunnerGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runner_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runner_group")
			}
			return err
		}
	}

	return nil
}

func (m *AppOrg) validateVcsConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.VcsConnections); i++ {
		if swag.IsZero(m.VcsConnections[i]) { // not required
			continue
		}

		if m.VcsConnections[i] != nil {
			if err := m.VcsConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcs_connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcs_connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this app org based on the context it is used
func (m *AppOrg) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestHealthCheck(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotificationsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunnerGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVcsConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppOrg) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppOrg) contextValidateHealthChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HealthChecks); i++ {

		if m.HealthChecks[i] != nil {

			if swag.IsZero(m.HealthChecks[i]) { // not required
				return nil
			}

			if err := m.HealthChecks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("health_checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("health_checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppOrg) contextValidateLatestHealthCheck(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *AppOrg) contextValidateNotificationsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NotificationsConfig != nil {

		if swag.IsZero(m.NotificationsConfig) { // not required
			return nil
		}

		if err := m.NotificationsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications_config")
			}
			return err
		}
	}

	return nil
}

func (m *AppOrg) contextValidateRunnerGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.RunnerGroup != nil {

		if swag.IsZero(m.RunnerGroup) { // not required
			return nil
		}

		if err := m.RunnerGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runner_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runner_group")
			}
			return err
		}
	}

	return nil
}

func (m *AppOrg) contextValidateVcsConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VcsConnections); i++ {

		if m.VcsConnections[i] != nil {

			if swag.IsZero(m.VcsConnections[i]) { // not required
				return nil
			}

			if err := m.VcsConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vcs_connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vcs_connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppOrg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppOrg) UnmarshalBinary(b []byte) error {
	var res AppOrg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
