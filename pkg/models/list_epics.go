// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListEpics list epics
//
// swagger:model ListEpics
type ListEpics struct {

	// A true/false boolean indicating whether to return Epics with their descriptions.
	IncludesDescription bool `json:"includes_description,omitempty"`
}

// Validate validates this list epics
func (m *ListEpics) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListEpics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListEpics) UnmarshalBinary(b []byte) error {
	var res ListEpics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
