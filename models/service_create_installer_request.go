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

// ServiceCreateInstallerRequest service create installer request
//
// swagger:model service.CreateInstallerRequest
type ServiceCreateInstallerRequest struct {

	// app ids
	// Required: true
	AppIds []string `json:"app_ids"`

	// metadata
	Metadata *ServiceCreateInstallerRequestMetadata `json:"metadata,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this service create installer request
func (m *ServiceCreateInstallerRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateInstallerRequest) validateAppIds(formats strfmt.Registry) error {

	if err := validate.Required("app_ids", "body", m.AppIds); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateInstallerRequest) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceCreateInstallerRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service create installer request based on the context it is used
func (m *ServiceCreateInstallerRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateInstallerRequest) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateInstallerRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateInstallerRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateInstallerRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceCreateInstallerRequestMetadata service create installer request metadata
//
// swagger:model ServiceCreateInstallerRequestMetadata
type ServiceCreateInstallerRequestMetadata struct {

	// community url
	// Required: true
	CommunityURL *string `json:"community_url"`

	// copyright markdown
	CopyrightMarkdown string `json:"copyright_markdown,omitempty"`

	// demo url
	DemoURL string `json:"demo_url,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// documentation url
	// Required: true
	DocumentationURL *string `json:"documentation_url"`

	// favicon url
	// Required: true
	FaviconURL *string `json:"favicon_url"`

	// footer markdown
	FooterMarkdown string `json:"footer_markdown,omitempty"`

	// github url
	// Required: true
	GithubURL *string `json:"github_url"`

	// homepage url
	// Required: true
	HomepageURL *string `json:"homepage_url"`

	// logo url
	// Required: true
	LogoURL *string `json:"logo_url"`

	// post install markdown
	PostInstallMarkdown string `json:"post_install_markdown,omitempty"`
}

// Validate validates this service create installer request metadata
func (m *ServiceCreateInstallerRequestMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunityURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentationURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaviconURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGithubURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHomepageURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogoURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceCreateInstallerRequestMetadata) validateCommunityURL(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"community_url", "body", m.CommunityURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateInstallerRequestMetadata) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateInstallerRequestMetadata) validateDocumentationURL(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"documentation_url", "body", m.DocumentationURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateInstallerRequestMetadata) validateFaviconURL(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"favicon_url", "body", m.FaviconURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateInstallerRequestMetadata) validateGithubURL(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"github_url", "body", m.GithubURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateInstallerRequestMetadata) validateHomepageURL(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"homepage_url", "body", m.HomepageURL); err != nil {
		return err
	}

	return nil
}

func (m *ServiceCreateInstallerRequestMetadata) validateLogoURL(formats strfmt.Registry) error {

	if err := validate.Required("metadata"+"."+"logo_url", "body", m.LogoURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service create installer request metadata based on context it is used
func (m *ServiceCreateInstallerRequestMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateInstallerRequestMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateInstallerRequestMetadata) UnmarshalBinary(b []byte) error {
	var res ServiceCreateInstallerRequestMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
