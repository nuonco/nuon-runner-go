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

// AppAppInputConfig app app input config
//
// swagger:model app.AppInputConfig
type AppAppInputConfig struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// input groups
	InputGroups []*AppAppInputGroup `json:"input_groups"`

	// inputs
	Inputs []*AppAppInput `json:"inputs"`

	// install inputs
	InstallInputs []*AppInstallInputs `json:"install_inputs"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app app input config
func (m *AppAppInputConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstallInputs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppInputConfig) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *AppAppInputConfig) validateInputGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.InputGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InputGroups); i++ {
		if swag.IsZero(m.InputGroups[i]) { // not required
			continue
		}

		if m.InputGroups[i] != nil {
			if err := m.InputGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppAppInputConfig) validateInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppAppInputConfig) validateInstallInputs(formats strfmt.Registry) error {
	if swag.IsZero(m.InstallInputs) { // not required
		return nil
	}

	for i := 0; i < len(m.InstallInputs); i++ {
		if swag.IsZero(m.InstallInputs[i]) { // not required
			continue
		}

		if m.InstallInputs[i] != nil {
			if err := m.InstallInputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this app app input config based on the context it is used
func (m *AppAppInputConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstallInputs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppAppInputConfig) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppAppInputConfig) contextValidateInputGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InputGroups); i++ {

		if m.InputGroups[i] != nil {

			if swag.IsZero(m.InputGroups[i]) { // not required
				return nil
			}

			if err := m.InputGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("input_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("input_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppAppInputConfig) contextValidateInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Inputs); i++ {

		if m.Inputs[i] != nil {

			if swag.IsZero(m.Inputs[i]) { // not required
				return nil
			}

			if err := m.Inputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppAppInputConfig) contextValidateInstallInputs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstallInputs); i++ {

		if m.InstallInputs[i] != nil {

			if swag.IsZero(m.InstallInputs[i]) { // not required
				return nil
			}

			if err := m.InstallInputs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("install_inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppAppInputConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAppInputConfig) UnmarshalBinary(b []byte) error {
	var res AppAppInputConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
