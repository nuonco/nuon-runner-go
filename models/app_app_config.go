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

// AppAppConfig app app config
//
// swagger:model app.AppConfig
type AppAppConfig struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// format
	Format AppAppConfigFmt `json:"format,omitempty"`

	// generated terraform
	GeneratedTerraform string `json:"generated_terraform,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// status
	Status AppAppConfigStatus `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// fields that are filled in via after query or views
	Version int64 `json:"version,omitempty"`
}

// Validate validates this app app config
func (m *AppAppConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppConfig) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *AppAppConfig) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	if err := m.Format.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("format")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("format")
		}
		return err
	}

	return nil
}

func (m *AppAppConfig) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this app app config based on the context it is used
func (m *AppAppConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppConfig) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppAppConfig) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	if err := m.Format.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("format")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("format")
		}
		return err
	}

	return nil
}

func (m *AppAppConfig) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppAppConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAppConfig) UnmarshalBinary(b []byte) error {
	var res AppAppConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
