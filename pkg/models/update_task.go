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

// UpdateTask update task
//
// swagger:model UpdateTask
type UpdateTask struct {

	// Move task after this task ID.
	AfterID int64 `json:"after_id,omitempty"`

	// Move task before this task ID.
	BeforeID int64 `json:"before_id,omitempty"`

	// A true/false boolean indicating whether the task is complete.
	Complete bool `json:"complete,omitempty"`

	// The Task's description.
	Description string `json:"description,omitempty"`

	// An array of UUIDs of the owners of this story.
	OwnerIds []strfmt.UUID `json:"owner_ids"`
}

// Validate validates this update task
func (m *UpdateTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTask) validateOwnerIds(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerIds) { // not required
		return nil
	}

	for i := 0; i < len(m.OwnerIds); i++ {

		if err := validate.FormatOf("owner_ids"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTask) UnmarshalBinary(b []byte) error {
	var res UpdateTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
