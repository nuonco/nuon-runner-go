// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChartFile chart file
//
// swagger:model chart.File
type ChartFile struct {

	// Data is the template as byte data.
	Data []int64 `json:"data"`

	// Name is the path-like name of the template.
	Name string `json:"name,omitempty"`
}

// Validate validates this chart file
func (m *ChartFile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chart file based on context it is used
func (m *ChartFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChartFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChartFile) UnmarshalBinary(b []byte) error {
	var res ChartFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
