// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListLabels list labels
//
// swagger:model ListLabels
type ListLabels struct {

	// A true/false boolean indicating if the slim versions of the Label should be returned.
	Slim bool `json:"slim,omitempty"`
}

// Validate validates this list labels
func (m *ListLabels) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListLabels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListLabels) UnmarshalBinary(b []byte) error {
	var res ListLabels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
