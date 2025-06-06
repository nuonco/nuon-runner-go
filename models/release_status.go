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

// ReleaseStatus release status
//
// swagger:model release.Status
type ReleaseStatus string

func NewReleaseStatus(value ReleaseStatus) *ReleaseStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ReleaseStatus.
func (m ReleaseStatus) Pointer() *ReleaseStatus {
	return &m
}

const (

	// ReleaseStatusUnknown captures enum value "unknown"
	ReleaseStatusUnknown ReleaseStatus = "unknown"

	// ReleaseStatusDeployed captures enum value "deployed"
	ReleaseStatusDeployed ReleaseStatus = "deployed"

	// ReleaseStatusUninstalled captures enum value "uninstalled"
	ReleaseStatusUninstalled ReleaseStatus = "uninstalled"

	// ReleaseStatusSuperseded captures enum value "superseded"
	ReleaseStatusSuperseded ReleaseStatus = "superseded"

	// ReleaseStatusFailed captures enum value "failed"
	ReleaseStatusFailed ReleaseStatus = "failed"

	// ReleaseStatusUninstalling captures enum value "uninstalling"
	ReleaseStatusUninstalling ReleaseStatus = "uninstalling"

	// ReleaseStatusPendingDashInstall captures enum value "pending-install"
	ReleaseStatusPendingDashInstall ReleaseStatus = "pending-install"

	// ReleaseStatusPendingDashUpgrade captures enum value "pending-upgrade"
	ReleaseStatusPendingDashUpgrade ReleaseStatus = "pending-upgrade"

	// ReleaseStatusPendingDashRollback captures enum value "pending-rollback"
	ReleaseStatusPendingDashRollback ReleaseStatus = "pending-rollback"
)

// for schema
var releaseStatusEnum []interface{}

func init() {
	var res []ReleaseStatus
	if err := json.Unmarshal([]byte(`["unknown","deployed","uninstalled","superseded","failed","uninstalling","pending-install","pending-upgrade","pending-rollback"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		releaseStatusEnum = append(releaseStatusEnum, v)
	}
}

func (m ReleaseStatus) validateReleaseStatusEnum(path, location string, value ReleaseStatus) error {
	if err := validate.EnumCase(path, location, value, releaseStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this release status
func (m ReleaseStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReleaseStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this release status based on context it is used
func (m ReleaseStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
