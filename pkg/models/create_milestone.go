// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateMilestone create milestone
//
// swagger:model CreateMilestone
type CreateMilestone struct {

	// An array of IDs of Categories attached to the Milestone.
	Categories []*CreateCategoryParams `json:"categories"`

	// A manual override for the time/date the Milestone was completed.
	// Format: date-time
	CompletedAtOverride strfmt.DateTime `json:"completed_at_override,omitempty"`

	// The Milestone's description.
	Description string `json:"description,omitempty"`

	// The name of the Milestone.
	// Required: true
	Name *string `json:"name"`

	// A manual override for the time/date the Milestone was started.
	// Format: date-time
	StartedAtOverride strfmt.DateTime `json:"started_at_override,omitempty"`

	// The workflow state that the Milestone is in.
	// Enum: [done in progress to do]
	State string `json:"state,omitempty"`
}

// Validate validates this create milestone
func (m *CreateMilestone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMilestone) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {
		if swag.IsZero(m.Categories[i]) { // not required
			continue
		}

		if m.Categories[i] != nil {
			if err := m.Categories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateMilestone) validateCompletedAtOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateMilestone) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateMilestone) validateStartedAtOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

var createMilestoneTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["done","in progress","to do"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createMilestoneTypeStatePropEnum = append(createMilestoneTypeStatePropEnum, v)
	}
}

const (

	// CreateMilestoneStateDone captures enum value "done"
	CreateMilestoneStateDone string = "done"

	// CreateMilestoneStateInProgress captures enum value "in progress"
	CreateMilestoneStateInProgress string = "in progress"

	// CreateMilestoneStateToDo captures enum value "to do"
	CreateMilestoneStateToDo string = "to do"
)

// prop value enum
func (m *CreateMilestone) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createMilestoneTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateMilestone) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateMilestone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMilestone) UnmarshalBinary(b []byte) error {
	var res CreateMilestone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
