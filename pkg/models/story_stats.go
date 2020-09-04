// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StoryStats The stats object for Stories
//
// swagger:model StoryStats
type StoryStats struct {

	// The number of documents related to this Story.
	// Required: true
	NumRelatedDocuments *int64 `json:"num_related_documents"`
}

// Validate validates this story stats
func (m *StoryStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumRelatedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoryStats) validateNumRelatedDocuments(formats strfmt.Registry) error {

	if err := validate.Required("num_related_documents", "body", m.NumRelatedDocuments); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoryStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoryStats) UnmarshalBinary(b []byte) error {
	var res StoryStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}