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

// Story Stories are the standard unit of work in Clubhouse and represent individual features, bugs, and chores.
//
// swagger:model Story
type Story struct {

	// The Clubhouse application url for the Story.
	// Required: true
	AppURL *string `json:"app_url"`

	// True if the story has been archived or not.
	// Required: true
	Archived *bool `json:"archived"`

	// A true/false boolean indicating if the Story is currently blocked.
	// Required: true
	Blocked *bool `json:"blocked"`

	// A true/false boolean indicating if the Story is currently a blocker of another story.
	// Required: true
	Blocker *bool `json:"blocker"`

	// An array of Git branches attached to the story.
	// Required: true
	Branches []*Branch `json:"branches"`

	// An array of comments attached to the story.
	// Required: true
	Comments []*Comment `json:"comments"`

	// An array of commits attached to the story.
	// Required: true
	Commits []*Commit `json:"commits"`

	// A true/false boolean indicating if the Story has been completed.
	// Required: true
	Completed *bool `json:"completed"`

	// The time/date the Story was completed.
	// Required: true
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at"`

	// A manual override for the time/date the Story was completed.
	// Required: true
	// Format: date-time
	CompletedAtOverride *strfmt.DateTime `json:"completed_at_override"`

	// The time/date the Story was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// The cycle time (in seconds) of this story when complete.
	CycleTime int64 `json:"cycle_time,omitempty"`

	// The due date of the story.
	// Required: true
	// Format: date-time
	Deadline *strfmt.DateTime `json:"deadline"`

	// The description of the story.
	// Required: true
	Description *string `json:"description"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The ID of the epic the story belongs to.
	// Required: true
	EpicID *int64 `json:"epic_id"`

	// The numeric point estimate of the story. Can also be null, which means unestimated.
	// Required: true
	Estimate *int64 `json:"estimate"`

	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	// Required: true
	ExternalID *string `json:"external_id"`

	// An array of external link (strings) associated with a Story
	// Required: true
	ExternalLinks []string `json:"external_links"`

	// An array of External Tickets associated with a Story
	// Required: true
	ExternalTickets []*ExternalTicket `json:"external_tickets"`

	// An array of files attached to the story.
	// Required: true
	Files []*File `json:"files"`

	// An array of UUIDs for any Members listed as Followers.
	// Required: true
	FollowerIds []strfmt.UUID `json:"follower_ids"`

	// The ID of the group associated with the story.
	// Required: true
	// Format: uuid
	GroupID *strfmt.UUID `json:"group_id"`

	// An array of Group IDs that have been mentioned in the Story description.
	// Required: true
	GroupMentionIds []strfmt.UUID `json:"group_mention_ids"`

	// The unique ID of the Story.
	// Required: true
	ID *int64 `json:"id"`

	// The ID of the iteration the story belongs to.
	// Required: true
	IterationID *int64 `json:"iteration_id"`

	// An array of labels attached to the story.
	// Required: true
	Labels []*Label `json:"labels"`

	// The lead time (in seconds) of this story when complete.
	LeadTime int64 `json:"lead_time,omitempty"`

	// An array of linked files attached to the story.
	// Required: true
	LinkedFiles []*LinkedFile `json:"linked_files"`

	// An array of Member IDs that have been mentioned in the Story description.
	// Required: true
	MemberMentionIds []strfmt.UUID `json:"member_mention_ids"`

	// Deprecated: use member_mention_ids.
	// Required: true
	MentionIds []strfmt.UUID `json:"mention_ids"`

	// The time/date the Story was last changed workflow-state.
	// Required: true
	// Format: date-time
	MovedAt *strfmt.DateTime `json:"moved_at"`

	// The name of the story.
	// Required: true
	Name *string `json:"name"`

	// An array of UUIDs of the owners of this story.
	// Required: true
	OwnerIds []strfmt.UUID `json:"owner_ids"`

	// A number representing the position of the story in relation to every other story in the current project.
	// Required: true
	Position *int64 `json:"position"`

	// The IDs of the iteration the story belongs to.
	// Required: true
	PreviousIterationIds []int64 `json:"previous_iteration_ids"`

	// The ID of the project the story belongs to.
	// Required: true
	ProjectID *int64 `json:"project_id"`

	// An array of Pull/Merge Requests attached to the story.
	// Required: true
	PullRequests []*PullRequest `json:"pull_requests"`

	// The ID of the Member that requested the story.
	// Required: true
	// Format: uuid
	RequestedByID *strfmt.UUID `json:"requested_by_id"`

	// A true/false boolean indicating if the Story has been started.
	// Required: true
	Started *bool `json:"started"`

	// The time/date the Story was started.
	// Required: true
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at"`

	// A manual override for the time/date the Story was started.
	// Required: true
	// Format: date-time
	StartedAtOverride *strfmt.DateTime `json:"started_at_override"`

	// stats
	// Required: true
	Stats *StoryStats `json:"stats"`

	// An array of story links attached to the Story.
	// Required: true
	StoryLinks []*TypedStoryLink `json:"story_links"`

	// The type of story (feature, bug, chore).
	// Required: true
	StoryType *string `json:"story_type"`

	// support tickets
	// Required: true
	SupportTickets []*SupportTicket `json:"support_tickets"`

	// An array of tasks connected to the story.
	// Required: true
	Tasks []*Task `json:"tasks"`

	// The time/date the Story was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// The ID of the workflow state the story is currently in.
	// Required: true
	WorkflowStateID *int64 `json:"workflow_state_id"`
}

// Validate validates this story
func (m *Story) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArchived(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletedAt(formats); err != nil {
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

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEpicID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalTickets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFollowerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIterationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMentionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMovedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousIterationIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedByID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAtOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
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

	if err := m.validateWorkflowStateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Story) validateAppURL(formats strfmt.Registry) error {

	if err := validate.Required("app_url", "body", m.AppURL); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateArchived(formats strfmt.Registry) error {

	if err := validate.Required("archived", "body", m.Archived); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateBlocker(formats strfmt.Registry) error {

	if err := validate.Required("blocker", "body", m.Blocker); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateBranches(formats strfmt.Registry) error {

	if err := validate.Required("branches", "body", m.Branches); err != nil {
		return err
	}

	for i := 0; i < len(m.Branches); i++ {
		if swag.IsZero(m.Branches[i]) { // not required
			continue
		}

		if m.Branches[i] != nil {
			if err := m.Branches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("branches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Story) validateComments(formats strfmt.Registry) error {

	if err := validate.Required("comments", "body", m.Comments); err != nil {
		return err
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

func (m *Story) validateCommits(formats strfmt.Registry) error {

	if err := validate.Required("commits", "body", m.Commits); err != nil {
		return err
	}

	for i := 0; i < len(m.Commits); i++ {
		if swag.IsZero(m.Commits[i]) { // not required
			continue
		}

		if m.Commits[i] != nil {
			if err := m.Commits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Story) validateCompleted(formats strfmt.Registry) error {

	if err := validate.Required("completed", "body", m.Completed); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateCompletedAt(formats strfmt.Registry) error {

	if err := validate.Required("completed_at", "body", m.CompletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateCompletedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("completed_at_override", "body", m.CompletedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("completed_at_override", "body", "date-time", m.CompletedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateDeadline(formats strfmt.Registry) error {

	if err := validate.Required("deadline", "body", m.Deadline); err != nil {
		return err
	}

	if err := validate.FormatOf("deadline", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateEpicID(formats strfmt.Registry) error {

	if err := validate.Required("epic_id", "body", m.EpicID); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateEstimate(formats strfmt.Registry) error {

	if err := validate.Required("estimate", "body", m.Estimate); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateExternalID(formats strfmt.Registry) error {

	if err := validate.Required("external_id", "body", m.ExternalID); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateExternalLinks(formats strfmt.Registry) error {

	if err := validate.Required("external_links", "body", m.ExternalLinks); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateExternalTickets(formats strfmt.Registry) error {

	if err := validate.Required("external_tickets", "body", m.ExternalTickets); err != nil {
		return err
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

func (m *Story) validateFiles(formats strfmt.Registry) error {

	if err := validate.Required("files", "body", m.Files); err != nil {
		return err
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

func (m *Story) validateFollowerIds(formats strfmt.Registry) error {

	if err := validate.Required("follower_ids", "body", m.FollowerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.FollowerIds); i++ {

		if err := validate.FormatOf("follower_ids"+"."+strconv.Itoa(i), "body", "uuid", m.FollowerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Story) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	if err := validate.FormatOf("group_id", "body", "uuid", m.GroupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateGroupMentionIds(formats strfmt.Registry) error {

	if err := validate.Required("group_mention_ids", "body", m.GroupMentionIds); err != nil {
		return err
	}

	for i := 0; i < len(m.GroupMentionIds); i++ {

		if err := validate.FormatOf("group_mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.GroupMentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Story) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateIterationID(formats strfmt.Registry) error {

	if err := validate.Required("iteration_id", "body", m.IterationID); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("labels", "body", m.Labels); err != nil {
		return err
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

func (m *Story) validateLinkedFiles(formats strfmt.Registry) error {

	if err := validate.Required("linked_files", "body", m.LinkedFiles); err != nil {
		return err
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

func (m *Story) validateMemberMentionIds(formats strfmt.Registry) error {

	if err := validate.Required("member_mention_ids", "body", m.MemberMentionIds); err != nil {
		return err
	}

	for i := 0; i < len(m.MemberMentionIds); i++ {

		if err := validate.FormatOf("member_mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MemberMentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Story) validateMentionIds(formats strfmt.Registry) error {

	if err := validate.Required("mention_ids", "body", m.MentionIds); err != nil {
		return err
	}

	for i := 0; i < len(m.MentionIds); i++ {

		if err := validate.FormatOf("mention_ids"+"."+strconv.Itoa(i), "body", "uuid", m.MentionIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Story) validateMovedAt(formats strfmt.Registry) error {

	if err := validate.Required("moved_at", "body", m.MovedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("moved_at", "body", "date-time", m.MovedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateOwnerIds(formats strfmt.Registry) error {

	if err := validate.Required("owner_ids", "body", m.OwnerIds); err != nil {
		return err
	}

	for i := 0; i < len(m.OwnerIds); i++ {

		if err := validate.FormatOf("owner_ids"+"."+strconv.Itoa(i), "body", "uuid", m.OwnerIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *Story) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *Story) validatePreviousIterationIds(formats strfmt.Registry) error {

	if err := validate.Required("previous_iteration_ids", "body", m.PreviousIterationIds); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("project_id", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *Story) validatePullRequests(formats strfmt.Registry) error {

	if err := validate.Required("pull_requests", "body", m.PullRequests); err != nil {
		return err
	}

	for i := 0; i < len(m.PullRequests); i++ {
		if swag.IsZero(m.PullRequests[i]) { // not required
			continue
		}

		if m.PullRequests[i] != nil {
			if err := m.PullRequests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pull_requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Story) validateRequestedByID(formats strfmt.Registry) error {

	if err := validate.Required("requested_by_id", "body", m.RequestedByID); err != nil {
		return err
	}

	if err := validate.FormatOf("requested_by_id", "body", "uuid", m.RequestedByID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateStarted(formats strfmt.Registry) error {

	if err := validate.Required("started", "body", m.Started); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("started_at", "body", m.StartedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateStartedAtOverride(formats strfmt.Registry) error {

	if err := validate.Required("started_at_override", "body", m.StartedAtOverride); err != nil {
		return err
	}

	if err := validate.FormatOf("started_at_override", "body", "date-time", m.StartedAtOverride.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateStats(formats strfmt.Registry) error {

	if err := validate.Required("stats", "body", m.Stats); err != nil {
		return err
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *Story) validateStoryLinks(formats strfmt.Registry) error {

	if err := validate.Required("story_links", "body", m.StoryLinks); err != nil {
		return err
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

func (m *Story) validateStoryType(formats strfmt.Registry) error {

	if err := validate.Required("story_type", "body", m.StoryType); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateSupportTickets(formats strfmt.Registry) error {

	if err := validate.Required("support_tickets", "body", m.SupportTickets); err != nil {
		return err
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

func (m *Story) validateTasks(formats strfmt.Registry) error {

	if err := validate.Required("tasks", "body", m.Tasks); err != nil {
		return err
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

func (m *Story) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Story) validateWorkflowStateID(formats strfmt.Registry) error {

	if err := validate.Required("workflow_state_id", "body", m.WorkflowStateID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Story) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Story) UnmarshalBinary(b []byte) error {
	var res Story
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}