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

// ServiceCreateExternalImageComponentConfigRequest service create external image component config request
//
// swagger:model service.CreateExternalImageComponentConfigRequest
type ServiceCreateExternalImageComponentConfigRequest struct {

	// aws ecr image config
	AwsEcrImageConfig *ServiceAwsECRImageConfigRequest `json:"aws_ecr_image_config,omitempty"`

	// image url
	// Required: true
	ImageURL *string `json:"image_url"`

	// tag
	// Required: true
	Tag *string `json:"tag"`
}

// Validate validates this service create external image component config request
func (m *ServiceCreateExternalImageComponentConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsEcrImageConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateExternalImageComponentConfigRequest) validateAwsEcrImageConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsEcrImageConfig) { // not required
		return nil
	}

	if m.AwsEcrImageConfig != nil {
		if err := m.AwsEcrImageConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_ecr_image_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_ecr_image_config")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCreateExternalImageComponentConfigRequest) validateImageURL(formats strfmt.Registry) error {

	if err := validate.Required("image_url", "body", m.ImageURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateExternalImageComponentConfigRequest) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service create external image component config request based on the context it is used
func (m *ServiceCreateExternalImageComponentConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsEcrImageConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateExternalImageComponentConfigRequest) contextValidateAwsEcrImageConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsEcrImageConfig != nil {

		if swag.IsZero(m.AwsEcrImageConfig) { // not required
			return nil
		}

		if err := m.AwsEcrImageConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws_ecr_image_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws_ecr_image_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateExternalImageComponentConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateExternalImageComponentConfigRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateExternalImageComponentConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
