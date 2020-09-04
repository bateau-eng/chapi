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

// CreateStoryContents A map of story attributes this template populates.
//
// swagger:model CreateStoryContents
type CreateStoryContents struct {

	// The due date of the story.
	// Format: date-time
	Deadline *strfmt.DateTime `json:"deadline,omitempty"`

	// The description of the story.
	Description string `json:"description,omitempty"`

	// A string description of this resource.
	EntityType string `json:"entity_type,omitempty"`

	// The ID of the epic the to be populated.
	EpicID *int64 `json:"epic_id,omitempty"`

	// The numeric point estimate to be populated.
	Estimate *int64 `json:"estimate,omitempty"`

	// An array of the external ticket IDs to be populated.
	ExternalTickets []*CreateEntityTemplateExternalTicket `json:"external_tickets"`

	// An array of the attached file IDs to be populated.
	// Unique: true
	FileIds []int64 `json:"file_ids"`

	// An array of files attached to the story.
	Files []*File `json:"files"`

	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// The ID of the iteration the to be populated.
	IterationID *int64 `json:"iteration_id,omitempty"`

	// An array of labels to be populated by the template.
	Labels []*CreateLabelParams `json:"labels"`

	// An array of the linked file IDs to be populated.
	// Unique: true
	LinkedFileIds []int64 `json:"linked_file_ids"`

	// An array of linked files attached to the story.
	LinkedFiles []*LinkedFile `json:"linked_files"`

	// The name of the story.
	Name string `json:"name,omitempty"`

	// An array of UUIDs of the owners of this story.
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// The ID of the project the story belongs to.
	ProjectID int64 `json:"project_id,omitempty"`

	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type,omitempty"`

	// An array of tasks to be populated by the template.
	Tasks []*EntityTemplateTask `json:"tasks"`

	// The ID of the workflow state the story is currently in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

// Validate validates this create story contents
func (m *CreateStoryContents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalTickets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFileIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateStoryContents) validateDeadline(formats strfmt.Registry) error {

	if swag.IsZero(m.Deadline) { // not required
		return nil
	}

	if err := validate.FormatOf("deadline", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryContents) validateExternalTickets(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalTickets) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalTickets); i++ {
		if swag.IsZero(m.ExternalTickets[i]) { // not required
			continue
		}

		if m.ExternalTickets[i] != nil {
			if err := m.ExternalTickets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_tickets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryContents) validateFileIds(formats strfmt.Registry) error {

	if swag.IsZero(m.FileIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("file_ids", "body", m.FileIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryContents) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryContents) validateFollowerIds(formats strfmt.Registry) error {

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

func (m *CreateStoryContents) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryContents) validateLinkedFileIds(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedFileIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("linked_file_ids", "body", m.LinkedFileIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryContents) validateLinkedFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedFiles) { // not required
		return nil
	}

	for i := 0; i < len(m.LinkedFiles); i++ {
		if swag.IsZero(m.LinkedFiles[i]) { // not required
			continue
		}

		if m.LinkedFiles[i] != nil {
			if err := m.LinkedFiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("linked_files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryContents) validateOwnerIds(formats strfmt.Registry) error {

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

func (m *CreateStoryContents) validateTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateStoryContents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateStoryContents) UnmarshalBinary(b []byte) error {
	var res CreateStoryContents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}