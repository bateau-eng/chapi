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

// CreateStoryParams Used to create multiple stories in a single request.
//
// swagger:model CreateStoryParams
type CreateStoryParams struct {

	// Controls the story's archived state.
	Archived bool `json:"archived,omitempty"`

	// An array of comments to add to the story.
	Comments []*CreateStoryCommentParams `json:"comments"`

	// A manual override for the time/date the Story was completed.
	// Format: date-time
	CompletedAtOverride strfmt.DateTime `json:"completed_at_override,omitempty"`

	// The time/date the Story was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// The due date of the story.
	// Format: date-time
	Deadline *strfmt.DateTime `json:"deadline,omitempty"`

	// The description of the story.
	Description string `json:"description,omitempty"`

	// The ID of the epic the story belongs to.
	EpicID *int64 `json:"epic_id,omitempty"`

	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate,omitempty"`

	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalID string `json:"external_id,omitempty"`

	// An array of External Tickets associated with this story. These External Tickets must have unquie external id. Duplicated External Tickets will be removed.
	ExternalTickets []*CreateExternalTicketParams `json:"external_tickets"`

	// An array of IDs of files attached to the story.
	// Unique: true
	FileIds []int64 `json:"file_ids"`

	// An array of UUIDs of the followers of this story.
	// Unique: true
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// The id of the group to associate with this story.
	// Format: uuid
	GroupID *strfmt.UUID `json:"group_id,omitempty"`

	// The ID of the iteration the story belongs to.
	IterationID *int64 `json:"iteration_id,omitempty"`

	// An array of labels attached to the story.
	Labels []*CreateLabelParams `json:"labels"`

	// An array of IDs of linked files attached to the story.
	// Unique: true
	LinkedFileIds []int64 `json:"linked_file_ids"`

	// The name of the story.
	// Required: true
	Name *string `json:"name"`

	// An array of UUIDs of the owners of this story.
	// Unique: true
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// The ID of the project the story belongs to.
	// Required: true
	ProjectID *int64 `json:"project_id"`

	// The ID of the member that requested the story.
	// Format: uuid
	RequestedByID strfmt.UUID `json:"requested_by_id,omitempty"`

	// A manual override for the time/date the Story was started.
	// Format: date-time
	StartedAtOverride strfmt.DateTime `json:"started_at_override,omitempty"`

	// An array of story links attached to the story.
	StoryLinks []*CreateStoryLinkParams `json:"story_links"`

	// The type of story (feature, bug, chore).
	// Enum: [bug chore feature]
	StoryType string `json:"story_type,omitempty"`

	// support tickets
	SupportTickets []*CreateExternalTicketParams `json:"support_tickets"`

	// An array of tasks connected to the story.
	Tasks []*CreateTaskParams `json:"tasks"`

	// The time/date the Story was updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// The ID of the workflow state the story will be in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`
}

// Validate validates this create story params
func (m *CreateStoryParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalTickets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFileIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedByID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportTickets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
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

func (m *CreateStoryParams) validateComments(formats strfmt.Registry) error {

	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	for i := 0; i < len(m.Comments); i++ {
		if swag.IsZero(m.Comments[i]) { // not required
			continue
		}

		if m.Comments[i] != nil {
			if err := m.Comments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("comments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryParams) validateCompletedAtOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateDeadline(formats strfmt.Registry) error {

	if swag.IsZero(m.Deadline) { // not required
		return nil
	}

	if err := validate.FormatOf("deadline", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateExternalTickets(formats strfmt.Registry) error {

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

func (m *CreateStoryParams) validateFileIds(formats strfmt.Registry) error {

	if swag.IsZero(m.FileIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("file_ids", "body", m.FileIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateFollowerIds(formats strfmt.Registry) error {

	if swag.IsZero(m.FollowerIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("follower_ids", "body", m.FollowerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.FollowerIds); i++ {

		if err := validate.FormatOf("follower_ids"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryParams) validateGroupID(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupID) { // not required
		return nil
	}

	if err := validate.FormatOf("group_id", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateLabels(formats strfmt.Registry) error {

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

func (m *CreateStoryParams) validateLinkedFileIds(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedFileIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("linked_file_ids", "body", m.LinkedFileIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateOwnerIds(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("owner_ids", "body", m.OwnerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.OwnerIds); i++ {

		if err := validate.FormatOf("owner_ids"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CreateStoryParams) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("project_id", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateRequestedByID(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedByID) { // not required
		return nil
	}

	if err := validate.FormatOf("requested_by_id", "body", "uuid", m.RequestedByID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateStartedAtOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAtOverride) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateStoryLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.StoryLinks) { // not required
		return nil
	}

	for i := 0; i < len(m.StoryLinks); i++ {
		if swag.IsZero(m.StoryLinks[i]) { // not required
			continue
		}

		if m.StoryLinks[i] != nil {
			if err := m.StoryLinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("story_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var createStoryParamsTypeStoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bug","chore","feature"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createStoryParamsTypeStoryTypePropEnum = append(createStoryParamsTypeStoryTypePropEnum, v)
	}
}

const (

	// CreateStoryParamsStoryTypeBug captures enum value "bug"
	CreateStoryParamsStoryTypeBug string = "bug"

	// CreateStoryParamsStoryTypeChore captures enum value "chore"
	CreateStoryParamsStoryTypeChore string = "chore"

	// CreateStoryParamsStoryTypeFeature captures enum value "feature"
	CreateStoryParamsStoryTypeFeature string = "feature"
)

// prop value enum
func (m *CreateStoryParams) validateStoryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createStoryParamsTypeStoryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateStoryParams) validateStoryType(formats strfmt.Registry) error {

	if swag.IsZero(m.StoryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateStoryTypeEnum("story_type", "body", m.StoryType); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoryParams) validateSupportTickets(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportTickets) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportTickets); i++ {
		if swag.IsZero(m.SupportTickets[i]) { // not required
			continue
		}

		if m.SupportTickets[i] != nil {
			if err := m.SupportTickets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("support_tickets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateStoryParams) validateTasks(formats strfmt.Registry) error {

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

func (m *CreateStoryParams) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateStoryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateStoryParams) UnmarshalBinary(b []byte) error {
	var res CreateStoryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
