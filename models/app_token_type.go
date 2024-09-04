// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AppTokenType app token type
//
// swagger:model app.TokenType
type AppTokenType string

func NewAppTokenType(value AppTokenType) *AppTokenType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AppTokenType.
func (m AppTokenType) Pointer() *AppTokenType {
	return &m
}

const (

	// AppTokenTypeAuth0 captures enum value "auth0"
	AppTokenTypeAuth0 AppTokenType = "auth0"

	// AppTokenTypeAdmin captures enum value "admin"
	AppTokenTypeAdmin AppTokenType = "admin"

	// AppTokenTypeStatic captures enum value "static"
	AppTokenTypeStatic AppTokenType = "static"

	// AppTokenTypeIntegration captures enum value "integration"
	AppTokenTypeIntegration AppTokenType = "integration"

	// AppTokenTypeCanary captures enum value "canary"
	AppTokenTypeCanary AppTokenType = "canary"
)

// for schema
var appTokenTypeEnum []interface{}

func init() {
	var res []AppTokenType
	if err := json.Unmarshal([]byte(`["auth0","admin","static","integration","canary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appTokenTypeEnum = append(appTokenTypeEnum, v)
	}
}

func (m AppTokenType) validateAppTokenTypeEnum(path, location string, value AppTokenType) error {
	if err := validate.EnumCase(path, location, value, appTokenTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app token type
func (m AppTokenType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppTokenTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this app token type based on context it is used
func (m AppTokenType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
