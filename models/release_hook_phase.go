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

// ReleaseHookPhase release hook phase
//
// swagger:model release.HookPhase
type ReleaseHookPhase string

func NewReleaseHookPhase(value ReleaseHookPhase) *ReleaseHookPhase {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ReleaseHookPhase.
func (m ReleaseHookPhase) Pointer() *ReleaseHookPhase {
	return &m
}

const (

	// ReleaseHookPhaseUnknown captures enum value "Unknown"
	ReleaseHookPhaseUnknown ReleaseHookPhase = "Unknown"

	// ReleaseHookPhaseRunning captures enum value "Running"
	ReleaseHookPhaseRunning ReleaseHookPhase = "Running"

	// ReleaseHookPhaseSucceeded captures enum value "Succeeded"
	ReleaseHookPhaseSucceeded ReleaseHookPhase = "Succeeded"

	// ReleaseHookPhaseFailed captures enum value "Failed"
	ReleaseHookPhaseFailed ReleaseHookPhase = "Failed"
)

// for schema
var releaseHookPhaseEnum []interface{}

func init() {
	var res []ReleaseHookPhase
	if err := json.Unmarshal([]byte(`["Unknown","Running","Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		releaseHookPhaseEnum = append(releaseHookPhaseEnum, v)
	}
}

func (m ReleaseHookPhase) validateReleaseHookPhaseEnum(path, location string, value ReleaseHookPhase) error {
	if err := validate.EnumCase(path, location, value, releaseHookPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this release hook phase
func (m ReleaseHookPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReleaseHookPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this release hook phase based on context it is used
func (m ReleaseHookPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
