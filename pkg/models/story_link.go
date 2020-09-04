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

// StoryLink Story links allow you create semantic relationships between two stories. Relationship types are relates to, blocks / blocked by, and duplicates / is duplicated by. The format is `subject -> link -> object`, or for example "story 5 blocks story 6".
//
// swagger:model StoryLink
type StoryLink struct {

	// The time/date when the Story Link was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The unique identifier of the Story Link.
	// Required: true
	ID *int64 `json:"id"`

	// The ID of the object Story.
	// Required: true
	ObjectID *int64 `json:"object_id"`

	// The ID of the subject Story.
	// Required: true
	SubjectID *int64 `json:"subject_id"`

	// The time/date when the Story Link was last updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// How the subject Story acts on the object Story. This can be "blocks", "duplicates", or "relates to".
	// Required: true
	Verb *string `json:"verb"`
}

// Validate validates this story link
func (m *StoryLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoryLink) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StoryLink) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *StoryLink) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *StoryLink) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("object_id", "body", m.ObjectID); err != nil {
		return err
	}

	return nil
}

func (m *StoryLink) validateSubjectID(formats strfmt.Registry) error {

	if err := validate.Required("subject_id", "body", m.SubjectID); err != nil {
		return err
	}

	return nil
}

func (m *StoryLink) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StoryLink) validateVerb(formats strfmt.Registry) error {

	if err := validate.Required("verb", "body", m.Verb); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoryLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoryLink) UnmarshalBinary(b []byte) error {
	var res StoryLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}