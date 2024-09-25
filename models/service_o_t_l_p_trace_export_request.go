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

// ServiceOTLPTraceExportRequest service o t l p trace export request
//
// swagger:model service.OTLPTraceExportRequest
type ServiceOTLPTraceExportRequest struct {

	// resource spans
	ResourceSpans []*ServiceOTLPTraceExportRequestResourceSpansItems0 `json:"resourceSpans"`
}

// Validate validates this service o t l p trace export request
func (m *ServiceOTLPTraceExportRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceSpans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequest) validateResourceSpans(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceSpans) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceSpans); i++ {
		if swag.IsZero(m.ResourceSpans[i]) { // not required
			continue
		}

		if m.ResourceSpans[i] != nil {
			if err := m.ResourceSpans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceSpans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceSpans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service o t l p trace export request based on the context it is used
func (m *ServiceOTLPTraceExportRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceSpans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequest) contextValidateResourceSpans(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceSpans); i++ {

		if m.ResourceSpans[i] != nil {

			if swag.IsZero(m.ResourceSpans[i]) { // not required
				return nil
			}

			if err := m.ResourceSpans[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourceSpans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resourceSpans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequest) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0 service o t l p trace export request resource spans items0
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0
type ServiceOTLPTraceExportRequestResourceSpansItems0 struct {

	// resource
	Resource *ServiceOTLPTraceExportRequestResourceSpansItems0Resource `json:"resource,omitempty"`

	// scope spans
	ScopeSpans []*ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0 `json:"scopeSpans"`
}

// Validate validates this service o t l p trace export request resource spans items0
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeSpans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) validateResource(formats strfmt.Registry) error {
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

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) validateScopeSpans(formats strfmt.Registry) error {
	if swag.IsZero(m.ScopeSpans) { // not required
		return nil
	}

	for i := 0; i < len(m.ScopeSpans); i++ {
		if swag.IsZero(m.ScopeSpans[i]) { // not required
			continue
		}

		if m.ScopeSpans[i] != nil {
			if err := m.ScopeSpans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scopeSpans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scopeSpans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service o t l p trace export request resource spans items0 based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScopeSpans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) contextValidateScopeSpans(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ScopeSpans); i++ {

		if m.ScopeSpans[i] != nil {

			if swag.IsZero(m.ScopeSpans[i]) { // not required
				return nil
			}

			if err := m.ScopeSpans[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scopeSpans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scopeSpans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0Resource service o t l p trace export request resource spans items0 resource
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0Resource
type ServiceOTLPTraceExportRequestResourceSpansItems0Resource struct {

	// attributes
	Attributes []*ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0 `json:"attributes"`
}

// Validate validates this service o t l p trace export request resource spans items0 resource
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0Resource) validateAttributes(formats strfmt.Registry) error {
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
					return ve.ValidateName("resource" + "." + "attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resource" + "." + "attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service o t l p trace export request resource spans items0 resource based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0Resource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0Resource) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {

			if swag.IsZero(m.Attributes[i]) { // not required
				return nil
			}

			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resource" + "." + "attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resource" + "." + "attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0Resource) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0 service o t l p trace export request resource spans items0 resource attributes items0
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0
type ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value `json:"value,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 resource attributes items0
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service o t l p trace export request resource spans items0 resource attributes items0 based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value service o t l p trace export request resource spans items0 resource attributes items0 value
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value
type ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value struct {

	// string value
	StringValue string `json:"stringValue,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 resource attributes items0 value
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service o t l p trace export request resource spans items0 resource attributes items0 value based on context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ResourceAttributesItems0Value
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0 service o t l p trace export request resource spans items0 scope spans items0
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0
type ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0 struct {

	// scope
	Scope *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope `json:"scope,omitempty"`

	// spans
	Spans []*ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0 `json:"spans"`
}

// Validate validates this service o t l p trace export request resource spans items0 scope spans items0
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) validateScope(formats strfmt.Registry) error {
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

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) validateSpans(formats strfmt.Registry) error {
	if swag.IsZero(m.Spans) { // not required
		return nil
	}

	for i := 0; i < len(m.Spans); i++ {
		if swag.IsZero(m.Spans[i]) { // not required
			continue
		}

		if m.Spans[i] != nil {
			if err := m.Spans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service o t l p trace export request resource spans items0 scope spans items0 based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) contextValidateSpans(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Spans); i++ {

		if m.Spans[i] != nil {

			if swag.IsZero(m.Spans[i]) { // not required
				return nil
			}

			if err := m.Spans[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("spans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope service o t l p trace export request resource spans items0 scope spans items0 scope
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope
type ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope struct {

	// attributes
	Attributes []*ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0 `json:"attributes"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 scope spans items0 scope
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope) validateAttributes(formats strfmt.Registry) error {
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
					return ve.ValidateName("scope" + "." + "attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scope" + "." + "attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service o t l p trace export request resource spans items0 scope spans items0 scope based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {

			if swag.IsZero(m.Attributes[i]) { // not required
				return nil
			}

			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scope" + "." + "attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scope" + "." + "attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0Scope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0 service o t l p trace export request resource spans items0 scope spans items0 scope attributes items0
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0
type ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value `json:"value,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 scope spans items0 scope attributes items0
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service o t l p trace export request resource spans items0 scope spans items0 scope attributes items0 based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value service o t l p trace export request resource spans items0 scope spans items0 scope attributes items0 value
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value
type ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value struct {

	// string value
	StringValue string `json:"stringValue,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 scope spans items0 scope attributes items0 value
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service o t l p trace export request resource spans items0 scope spans items0 scope attributes items0 value based on context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0ScopeAttributesItems0Value
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0 service o t l p trace export request resource spans items0 scope spans items0 spans items0
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0
type ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0 struct {

	// attributes
	Attributes []*ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0 `json:"attributes"`

	// end time unix nano
	EndTimeUnixNano string `json:"endTimeUnixNano,omitempty"`

	// kind
	Kind int64 `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parent span Id
	ParentSpanID string `json:"parentSpanId,omitempty"`

	// span Id
	SpanID string `json:"spanId,omitempty"`

	// start time unix nano
	StartTimeUnixNano string `json:"startTimeUnixNano,omitempty"`

	// trace Id
	TraceID string `json:"traceId,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 scope spans items0 spans items0
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0) validateAttributes(formats strfmt.Registry) error {
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

// ContextValidate validate this service o t l p trace export request resource spans items0 scope spans items0 spans items0 based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0 service o t l p trace export request resource spans items0 scope spans items0 spans items0 attributes items0
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0
type ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value `json:"value,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 scope spans items0 spans items0 attributes items0
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service o t l p trace export request resource spans items0 scope spans items0 spans items0 attributes items0 based on the context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {

		if swag.IsZero(m.Value) { // not required
			return nil
		}

		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value service o t l p trace export request resource spans items0 scope spans items0 spans items0 attributes items0 value
//
// swagger:model ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value
type ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value struct {

	// string value
	StringValue string `json:"stringValue,omitempty"`
}

// Validate validates this service o t l p trace export request resource spans items0 scope spans items0 spans items0 attributes items0 value
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service o t l p trace export request resource spans items0 scope spans items0 spans items0 attributes items0 value based on context it is used
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value) UnmarshalBinary(b []byte) error {
	var res ServiceOTLPTraceExportRequestResourceSpansItems0ScopeSpansItems0SpansItems0AttributesItems0Value
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}