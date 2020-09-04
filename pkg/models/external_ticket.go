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

// ExternalTicket A External Ticket allows you to create a link between an external system, like Zendesk, and a Clubhouse story.
//
// swagger:model ExternalTicket
type ExternalTicket struct {

	// The ID used in the external system.
	// Required: true
	ExternalID *string `json:"external_id"`

	// The full qualified url of the external ticket.
	ExternalURL string `json:"external_url,omitempty"`

	// A unique ID internal to Clubhouse.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// The Clubhouse Story ids associated with this External Ticket.
	// Required: true
	StoryIds []float64 `json:"story_ids"`
}

// Validate validates this external ticket
func (m *ExternalTicket) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalTicket) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *ExternalTicket) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExternalTicket) validateStoryIds(formats strfmt.Registry) error {

	if err := validate.Required("story_ids", "body", m.StoryIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalTicket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalTicket) UnmarshalBinary(b []byte) error {
	var res ExternalTicket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
