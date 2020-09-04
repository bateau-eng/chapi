// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EpicWorkflow Epic Workflow is the array of defined Epic States. Epic Workflow can be queried using the API but must be updated in the Clubhouse UI.
//
// swagger:model EpicWorkflow
type EpicWorkflow struct {

	// The date the Epic Workflow was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The unique ID of the default Epic State that new Epics are assigned by default.
	// Required: true
	DefaultEpicStateID *int64 `json:"default_epic_state_id"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// A map of the Epic States in this Epic Workflow.
	// Required: true
	EpicStates []*EpicState `json:"epic_states"`

	// The unique ID of the Epic Workflow.
	// Required: true
	ID *int64 `json:"id"`

	// The date the Epic Workflow was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this epic workflow
func (m *EpicWorkflow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultEpicStateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEpicStates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EpicWorkflow) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EpicWorkflow) validateDefaultEpicStateID(formats strfmt.Registry) error {

	if err := validate.Required("default_epic_state_id", "body", m.DefaultEpicStateID); err != nil {
		return err
	}

	return nil
}

func (m *EpicWorkflow) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *EpicWorkflow) validateEpicStates(formats strfmt.Registry) error {

	if err := validate.Required("epic_states", "body", m.EpicStates); err != nil {
		return err
	}

	for i := 0; i < len(m.EpicStates); i++ {
		if swag.IsZero(m.EpicStates[i]) { // not required
			continue
		}

		if m.EpicStates[i] != nil {
			if err := m.EpicStates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("epic_states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EpicWorkflow) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *EpicWorkflow) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EpicWorkflow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EpicWorkflow) UnmarshalBinary(b []byte) error {
	var res EpicWorkflow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
