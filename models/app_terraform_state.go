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

// AppTerraformState app terraform state
//
// swagger:model app.TerraformState
type AppTerraformState struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// data
	Data []int64 `json:"data"`

	// id
	ID string `json:"id,omitempty"`

	// lock
	Lock []int64 `json:"lock"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// revision
	Revision int64 `json:"revision,omitempty"`

	// terraform workspace
	TerraformWorkspace *AppTerraformWorkspace `json:"terraformWorkspace,omitempty"`

	// terraform workspace ID
	TerraformWorkspaceID string `json:"terraformWorkspaceID,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app terraform state
func (m *AppTerraformState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTerraformWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppTerraformState) validateTerraformWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.TerraformWorkspace) { // not required
		return nil
	}

	if m.TerraformWorkspace != nil {
		if err := m.TerraformWorkspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraformWorkspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraformWorkspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this app terraform state based on the context it is used
func (m *AppTerraformState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTerraformWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppTerraformState) contextValidateTerraformWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.TerraformWorkspace != nil {

		if swag.IsZero(m.TerraformWorkspace) { // not required
			return nil
		}

		if err := m.TerraformWorkspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraformWorkspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraformWorkspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppTerraformState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppTerraformState) UnmarshalBinary(b []byte) error {
	var res AppTerraformState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
