// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V2File v2 file
//
// swagger:model v2.File
type V2File struct {

	// Data is the template as byte data.
	Data []int64 `json:"data"`

	// Name is the path-like name of the template.
	Name string `json:"name,omitempty"`
}

// Validate validates this v2 file
func (m *V2File) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v2 file based on context it is used
func (m *V2File) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V2File) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V2File) UnmarshalBinary(b []byte) error {
	var res V2File
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
