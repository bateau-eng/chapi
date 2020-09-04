// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateCategory update category
//
// swagger:model UpdateCategory
type UpdateCategory struct {

	// A true/false boolean indicating if the Category has been archived.
	Archived bool `json:"archived,omitempty"`

	// The hex color to be displayed with the Category (for example, "#ff0000").
	Color *string `json:"color,omitempty"`

	// The new name of the Category.
	Name string `json:"name,omitempty"`
}

// Validate validates this update category
func (m *UpdateCategory) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCategory) UnmarshalBinary(b []byte) error {
	var res UpdateCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}