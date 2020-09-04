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

// UpdateProject update project
//
// swagger:model UpdateProject
type UpdateProject struct {

	// The Project abbreviation used in Story summaries. Should be kept to 3 characters at most.
	Abbreviation string `json:"abbreviation,omitempty"`

	// A true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived,omitempty"`

	// The color that represents the Project in the UI.
	Color string `json:"color,omitempty"`

	// The number of days before the thermometer appears in the Story summary.
	DaysToThermometer int64 `json:"days_to_thermometer,omitempty"`

	// The Project's description.
	Description string `json:"description,omitempty"`

	// An array of UUIDs for any Members you want to add as Followers.
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// The Project's name.
	Name string `json:"name,omitempty"`

	// Configuration to enable or disable thermometers in the Story summary.
	ShowThermometer bool `json:"show_thermometer,omitempty"`

	// The ID of the team the project belongs to.
	TeamID int64 `json:"team_id,omitempty"`
}

// Validate validates this update project
func (m *UpdateProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProject) validateFollowerIds(formats strfmt.Registry) error {

	if swag.IsZero(m.FollowerIds) { // not required
		return nil
	}

	for i := 0; i < len(m.FollowerIds); i++ {

		if err := validate.FormatOf("follower_ids"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProject) UnmarshalBinary(b []byte) error {
	var res UpdateProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}