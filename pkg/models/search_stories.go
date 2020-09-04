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

// SearchStories search stories
//
// swagger:model SearchStories
type SearchStories struct {

	// A true/false boolean indicating whether the Story is in archived state.
	Archived bool `json:"archived,omitempty"`

	// Stories should have been completed before this date.
	// Format: date-time
	CompletedAtEnd strfmt.DateTime `json:"completed_at_end,omitempty"`

	// Stories should have been competed after this date.
	// Format: date-time
	CompletedAtStart strfmt.DateTime `json:"completed_at_start,omitempty"`

	// Stories should have been created before this date.
	// Format: date-time
	CreatedAtEnd strfmt.DateTime `json:"created_at_end,omitempty"`

	// Stories should have been created after this date.
	// Format: date-time
	CreatedAtStart strfmt.DateTime `json:"created_at_start,omitempty"`

	// Stories should have a deadline before this date.
	// Format: date-time
	DeadlineEnd strfmt.DateTime `json:"deadline_end,omitempty"`

	// Stories should have a deadline after this date.
	// Format: date-time
	DeadlineStart strfmt.DateTime `json:"deadline_start,omitempty"`

	// The Epic IDs that may be associated with the Stories.
	EpicID *int64 `json:"epic_id,omitempty"`

	// The Epic IDs that may be associated with the Stories.
	// Unique: true
	EpicIds []int64 `json:"epic_ids"`

	// The number of estimate points associate with the Stories.
	Estimate int64 `json:"estimate,omitempty"`

	// An ID or URL that references an external resource. Useful during imports.
	ExternalID string `json:"external_id,omitempty"`

	// The Group ID that is associated with the Stories
	// Format: uuid
	GroupID *strfmt.UUID `json:"group_id,omitempty"`

	// The Group IDs that are associated with the Stories
	// Unique: true
	GroupIds []strfmt.UUID `json:"group_ids"`

	// Whether to include the story description in the response.
	IncludesDescription bool `json:"includes_description,omitempty"`

	// The Iteration ID that may be associated with the Stories.
	IterationID *int64 `json:"iteration_id,omitempty"`

	// The Iteration IDs that may be associated with the Stories.
	// Unique: true
	IterationIds []int64 `json:"iteration_ids"`

	// The Label IDs that may be associated with the Stories.
	// Unique: true
	LabelIds []int64 `json:"label_ids"`

	// The name of any associated Labels.
	LabelName string `json:"label_name,omitempty"`

	// An array of UUIDs for any Users who may be Owners of the Stories.
	// Format: uuid
	OwnerID *strfmt.UUID `json:"owner_id,omitempty"`

	// An array of UUIDs for any Users who may be Owners of the Stories.
	// Unique: true
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// The IDs for the Projects the Stories may be assigned to.
	ProjectID int64 `json:"project_id,omitempty"`

	// The IDs for the Projects the Stories may be assigned to.
	// Unique: true
	ProjectIds []int64 `json:"project_ids"`

	// The UUID of any Users who may have requested the Stories.
	// Format: uuid
	RequestedByID strfmt.UUID `json:"requested_by_id,omitempty"`

	// The type of Stories that you want returned.
	// Enum: [bug chore feature]
	StoryType string `json:"story_type,omitempty"`

	// Stories should have been updated before this date.
	// Format: date-time
	UpdatedAtEnd strfmt.DateTime `json:"updated_at_end,omitempty"`

	// Stories should have been updated after this date.
	// Format: date-time
	UpdatedAtStart strfmt.DateTime `json:"updated_at_start,omitempty"`

	// The unique IDs of the specific Workflow States that the Stories should be in.
	WorkflowStateID int64 `json:"workflow_state_id,omitempty"`

	// The type of Workflow State the Stories may be in.
	WorkflowStateTypes []string `json:"workflow_state_types"`
}

// Validate validates this search stories
func (m *SearchStories) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAtEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAtStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAtEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAtStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadlineEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadlineStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEpicIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIterationIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedByID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAtEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAtStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkflowStateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchStories) validateCompletedAtEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAtEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at_end", "body", "date-time", m.CompletedAtEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateCompletedAtStart(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAtStart) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at_start", "body", "date-time", m.CompletedAtStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateCreatedAtEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAtEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at_end", "body", "date-time", m.CreatedAtEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateCreatedAtStart(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAtStart) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at_start", "body", "date-time", m.CreatedAtStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateDeadlineEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.DeadlineEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("deadline_end", "body", "date-time", m.DeadlineEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateDeadlineStart(formats strfmt.Registry) error {

	if swag.IsZero(m.DeadlineStart) { // not required
		return nil
	}

	if err := validate.FormatOf("deadline_start", "body", "date-time", m.DeadlineStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateEpicIds(formats strfmt.Registry) error {

	if swag.IsZero(m.EpicIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("epic_ids", "body", m.EpicIds); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateGroupID(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupID) { // not required
		return nil
	}

	if err := validate.FormatOf("group_id", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateGroupIds(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("group_ids", "body", m.GroupIds); err != nil {
		return err
	}

	for i := 0; i < len(m.GroupIds); i++ {

		if err := validate.FormatOf("group_ids"+"."+strconv.Itoa(i), "body", "uuid", m.GroupIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *SearchStories) validateIterationIds(formats strfmt.Registry) error {

	if swag.IsZero(m.IterationIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("iteration_ids", "body", m.IterationIds); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateLabelIds(formats strfmt.Registry) error {

	if swag.IsZero(m.LabelIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("label_ids", "body", m.LabelIds); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateOwnerID(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerID) { // not required
		return nil
	}

	if err := validate.FormatOf("owner_id", "body", "uuid", m.OwnerID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateOwnerIds(formats strfmt.Registry) error {

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

func (m *SearchStories) validateProjectIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("project_ids", "body", m.ProjectIds); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateRequestedByID(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestedByID) { // not required
		return nil
	}

	if err := validate.FormatOf("requested_by_id", "body", "uuid", m.RequestedByID.String(), formats); err != nil {
		return err
	}

	return nil
}

var searchStoriesTypeStoryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bug","chore","feature"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchStoriesTypeStoryTypePropEnum = append(searchStoriesTypeStoryTypePropEnum, v)
	}
}

const (

	// SearchStoriesStoryTypeBug captures enum value "bug"
	SearchStoriesStoryTypeBug string = "bug"

	// SearchStoriesStoryTypeChore captures enum value "chore"
	SearchStoriesStoryTypeChore string = "chore"

	// SearchStoriesStoryTypeFeature captures enum value "feature"
	SearchStoriesStoryTypeFeature string = "feature"
)

// prop value enum
func (m *SearchStories) validateStoryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchStoriesTypeStoryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchStories) validateStoryType(formats strfmt.Registry) error {

	if swag.IsZero(m.StoryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateStoryTypeEnum("story_type", "body", m.StoryType); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateUpdatedAtEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAtEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at_end", "body", "date-time", m.UpdatedAtEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SearchStories) validateUpdatedAtStart(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAtStart) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at_start", "body", "date-time", m.UpdatedAtStart.String(), formats); err != nil {
		return err
	}

	return nil
}

var searchStoriesWorkflowStateTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["done","started","unstarted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchStoriesWorkflowStateTypesItemsEnum = append(searchStoriesWorkflowStateTypesItemsEnum, v)
	}
}

func (m *SearchStories) validateWorkflowStateTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchStoriesWorkflowStateTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchStories) validateWorkflowStateTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkflowStateTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkflowStateTypes); i++ {

		// value enum
		if err := m.validateWorkflowStateTypesItemsEnum("workflow_state_types"+"."+strconv.Itoa(i), "body", m.WorkflowStateTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchStories) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchStories) UnmarshalBinary(b []byte) error {
	var res SearchStories
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
