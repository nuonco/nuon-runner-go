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

// ReleaseHookEvent release hook event
//
// swagger:model release.HookEvent
type ReleaseHookEvent string

func NewReleaseHookEvent(value ReleaseHookEvent) *ReleaseHookEvent {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ReleaseHookEvent.
func (m ReleaseHookEvent) Pointer() *ReleaseHookEvent {
	return &m
}

const (

	// ReleaseHookEventPreDashInstall captures enum value "pre-install"
	ReleaseHookEventPreDashInstall ReleaseHookEvent = "pre-install"

	// ReleaseHookEventPostDashInstall captures enum value "post-install"
	ReleaseHookEventPostDashInstall ReleaseHookEvent = "post-install"

	// ReleaseHookEventPreDashDelete captures enum value "pre-delete"
	ReleaseHookEventPreDashDelete ReleaseHookEvent = "pre-delete"

	// ReleaseHookEventPostDashDelete captures enum value "post-delete"
	ReleaseHookEventPostDashDelete ReleaseHookEvent = "post-delete"

	// ReleaseHookEventPreDashUpgrade captures enum value "pre-upgrade"
	ReleaseHookEventPreDashUpgrade ReleaseHookEvent = "pre-upgrade"

	// ReleaseHookEventPostDashUpgrade captures enum value "post-upgrade"
	ReleaseHookEventPostDashUpgrade ReleaseHookEvent = "post-upgrade"

	// ReleaseHookEventPreDashRollback captures enum value "pre-rollback"
	ReleaseHookEventPreDashRollback ReleaseHookEvent = "pre-rollback"

	// ReleaseHookEventPostDashRollback captures enum value "post-rollback"
	ReleaseHookEventPostDashRollback ReleaseHookEvent = "post-rollback"

	// ReleaseHookEventTest captures enum value "test"
	ReleaseHookEventTest ReleaseHookEvent = "test"
)

// for schema
var releaseHookEventEnum []interface{}

func init() {
	var res []ReleaseHookEvent
	if err := json.Unmarshal([]byte(`["pre-install","post-install","pre-delete","post-delete","pre-upgrade","post-upgrade","pre-rollback","post-rollback","test"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		releaseHookEventEnum = append(releaseHookEventEnum, v)
	}
}

func (m ReleaseHookEvent) validateReleaseHookEventEnum(path, location string, value ReleaseHookEvent) error {
	if err := validate.EnumCase(path, location, value, releaseHookEventEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this release hook event
func (m ReleaseHookEvent) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReleaseHookEventEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this release hook event based on context it is used
func (m ReleaseHookEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
