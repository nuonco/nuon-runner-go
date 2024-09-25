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

// ServiceOTLPLogExportRequest service o t l p log export request
//
// swagger:model service.OTLPLogExportRequest
type ServiceOTLPLogExportRequest struct {

	// resource logs
	ResourceLogs []*ServiceOTLPLogExportRequestResourceLogsItems0 `json:"resourceLogs"`
}

// Validate validates this service o t l p log export request
func (m *ServiceOTLPLogExportRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequest) validateResourceLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceLogs); i++ {
		if swag.IsZero(m.ResourceLogs[i]) { // not required
			continue
		}

		if m.ResourceLogs[i] != nil {
			if err := m.ResourceLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service o t l p log export request based on the context it is used
func (m *ServiceOTLPLogExportRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequest) contextValidateResourceLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceLogs); i++ {

		if m.ResourceLogs[i] != nil {

			if swag.IsZero(m.ResourceLogs[i]) { // not required
				return nil
			}

			if err := m.ResourceLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequest) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPLogExportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPLogExportRequestResourceLogsItems0 service o t l p log export request resource logs items0
//
// swagger:model ServiceOTLPLogExportRequestResourceLogsItems0
type ServiceOTLPLogExportRequestResourceLogsItems0 struct {

	// resource
	Resource *ServiceResource `json:"resource,omitempty"`

	// scope logs
	ScopeLogs []*ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0 `json:"scopeLogs"`
}

// Validate validates this service o t l p log export request resource logs items0
func (m *ServiceOTLPLogExportRequestResourceLogsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0) validateScopeLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.ScopeLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.ScopeLogs); i++ {
		if swag.IsZero(m.ScopeLogs[i]) { // not required
			continue
		}

		if m.ScopeLogs[i] != nil {
			if err := m.ScopeLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scopeLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scopeLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service o t l p log export request resource logs items0 based on the context it is used
func (m *ServiceOTLPLogExportRequestResourceLogsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScopeLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {

		if swag.IsZero(m.Resource) { // not required
			return nil
		}

		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0) contextValidateScopeLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ScopeLogs); i++ {

		if m.ScopeLogs[i] != nil {

			if swag.IsZero(m.ScopeLogs[i]) { // not required
				return nil
			}

			if err := m.ScopeLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scopeLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scopeLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequestResourceLogsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequestResourceLogsItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPLogExportRequestResourceLogsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0 service o t l p log export request resource logs items0 scope logs items0
//
// swagger:model ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0
type ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0 struct {

	// log records
	LogRecords []*ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0 `json:"logRecords"`

	// schema Url
	SchemaURL string `json:"schemaUrl,omitempty"`

	// scope
	Scope *ServiceScope `json:"scope,omitempty"`
}

// Validate validates this service o t l p log export request resource logs items0 scope logs items0
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogRecords(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) validateLogRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.LogRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.LogRecords); i++ {
		if swag.IsZero(m.LogRecords[i]) { // not required
			continue
		}

		if m.LogRecords[i] != nil {
			if err := m.LogRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logRecords" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logRecords" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service o t l p log export request resource logs items0 scope logs items0 based on the context it is used
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) contextValidateLogRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LogRecords); i++ {

		if m.LogRecords[i] != nil {

			if swag.IsZero(m.LogRecords[i]) { // not required
				return nil
			}

			if err := m.LogRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logRecords" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logRecords" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {

		if swag.IsZero(m.Scope) { // not required
			return nil
		}

		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0 service o t l p log export request resource logs items0 scope logs items0 log records items0
//
// swagger:model ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0
type ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0 struct {

	// attributes
	Attributes []*ServiceAttribute `json:"attributes"`

	// body
	Body *ServiceBody `json:"body,omitempty"`

	// dropped attributes count
	DroppedAttributesCount int64 `json:"droppedAttributesCount,omitempty"`

	// flags
	Flags int64 `json:"flags,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`

	// severity number
	SeverityNumber int64 `json:"severityNumber,omitempty"`

	// severity text
	SeverityText string `json:"severityText,omitempty"`

	// span Id
	SpanID string `json:"spanId,omitempty"`

	// time unix nano
	TimeUnixNano string `json:"timeUnixNano,omitempty"`

	// trace Id
	TraceID string `json:"traceId,omitempty"`
}

// Validate validates this service o t l p log export request resource logs items0 scope logs items0 log records items0
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) validateBody(formats strfmt.Registry) error {
	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service o t l p log export request resource logs items0 scope logs items0 log records items0 based on the context it is used
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBody(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {

			if swag.IsZero(m.Attributes[i]) { // not required
				return nil
			}

			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) contextValidateBody(ctx context.Context, formats strfmt.Registry) error {

	if m.Body != nil {

		if swag.IsZero(m.Body) { // not required
			return nil
		}

		if err := m.Body.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPLogExportRequestResourceLogsItems0ScopeLogsItems0LogRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
