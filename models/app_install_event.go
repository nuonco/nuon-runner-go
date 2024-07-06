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

// AppInstallEvent app install event
//
// swagger:model app.InstallEvent
type AppInstallEvent struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// created by
	CreatedBy *AppAccount `json:"created_by,omitempty"`

	// created by id
	CreatedByID string `json:"created_by_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// install id
	InstallID string `json:"install_id,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// operation name
	OperationName string `json:"operation_name,omitempty"`

	// operation status
	OperationStatus AppOperationStatus `json:"operation_status,omitempty"`

	// org id
	OrgID string `json:"org_id,omitempty"`

	// payload
	Payload map[string]string `json:"payload,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this app install event
func (m *AppInstallEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallEvent) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *AppInstallEvent) validateOperationStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.OperationStatus) { // not required
		return nil
	}

	if err := m.OperationStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operation_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operation_status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this app install event based on the context it is used
func (m *AppInstallEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperationStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppInstallEvent) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AppInstallEvent) contextValidateOperationStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.OperationStatus) { // not required
		return nil
	}

	if err := m.OperationStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operation_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operation_status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppInstallEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppInstallEvent) UnmarshalBinary(b []byte) error {
	var res AppInstallEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
