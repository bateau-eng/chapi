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

// Team Group of Projects with the same Workflow.
//
// swagger:model Team
type Team struct {

	// The time/date the Team was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The description of the Team.
	// Required: true
	Description *string `json:"description"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The unique identifier of the Team.
	// Required: true
	ID *int64 `json:"id"`

	// The name of the Team.
	// Required: true
	Name *string `json:"name"`

	// A number representing the position of the Team in relation to every other Team within the Organization.
	// Required: true
	Position *float64 `json:"position"`

	// An array of IDs of projects within the Team.
	// Required: true
	ProjectIds []float64 `json:"project_ids"`

	// The time/date the Team was last updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// workflow
	// Required: true
	Workflow *Workflow `json:"workflow"`
}

// Validate validates this team
func (m *Team) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Team) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Team) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateProjectIds(formats strfmt.Registry) error {

	if err := validate.Required("project_ids", "body", m.ProjectIds); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Team) validateWorkflow(formats strfmt.Registry) error {

	if err := validate.Required("workflow", "body", m.Workflow); err != nil {
		return err
	}

	if m.Workflow != nil {
		if err := m.Workflow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workflow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Team) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Team) UnmarshalBinary(b []byte) error {
	var res Team
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
