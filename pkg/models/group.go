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

// Group A Group.
//
// swagger:model Group
type Group struct {

	// The Clubhouse application url for the Group.
	// Required: true
	AppURL *string `json:"app_url"`

	// Whether or not the Group is archived.
	// Required: true
	Archived *bool `json:"archived"`

	// The description of the Group.
	// Required: true
	Description *string `json:"description"`

	// display icon
	// Required: true
	DisplayIcon *Icon `json:"display_icon"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The id of the Group.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// The Member IDs contain within the Group.
	// Required: true
	MemberIds []strfmt.UUID `json:"member_ids"`

	// The mention name of the Group.
	// Required: true
	MentionName *string `json:"mention_name"`

	// The name of the Group.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this group
func (m *Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Group) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateDisplayIcon(formats strfmt.Registry) error {

	if err := validate.Required("display_icon", "body", m.DisplayIcon); err != nil {
		return err
	}

	if m.DisplayIcon != nil {
		if err := m.DisplayIcon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display_icon")
			}
			return err
		}
	}

	return nil
}

func (m *Group) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateMemberIds(formats strfmt.Registry) error {

	if err := validate.Required("member_ids", "body", m.MemberIds); err != nil {
		return err
	}

	for i := 0; i < len(m.MemberIds); i++ {

		if err := validate.FormatOf("member_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MemberIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Group) validateMentionName(formats strfmt.Registry) error {

	if err := validate.Required("mention_name", "body", m.MentionName); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Group) UnmarshalBinary(b []byte) error {
	var res Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
