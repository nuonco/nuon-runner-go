// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceCreateRunnerJobExecutionResultRequest service create runner job execution result request
//
// swagger:model service.CreateRunnerJobExecutionResultRequest
type ServiceCreateRunnerJobExecutionResultRequest struct {

	// error code
	ErrorCode int64 `json:"error_code,omitempty"`

	// error metadata
	ErrorMetadata map[string]string `json:"error_metadata,omitempty"`

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this service create runner job execution result request
func (m *ServiceCreateRunnerJobExecutionResultRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service create runner job execution result request based on context it is used
func (m *ServiceCreateRunnerJobExecutionResultRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateRunnerJobExecutionResultRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateRunnerJobExecutionResultRequest) UnmarshalBinary(b []byte) error {
	var res ServiceCreateRunnerJobExecutionResultRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}