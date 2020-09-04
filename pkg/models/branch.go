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

// Branch Branch refers to a VCS branch. Branches are feature branches associated with Clubhouse Stories.
//
// swagger:model Branch
type Branch struct {

	// The time/date the Branch was created.
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// A true/false boolean indicating if the Branch has been deleted.
	// Required: true
	Deleted *bool `json:"deleted"`

	// A string description of this resource.
	// Required: true
	EntityType *string `json:"entity_type"`

	// The unique ID of the Branch.
	// Required: true
	ID *int64 `json:"id"`

	// The IDs of the Branches the Branch has been merged into.
	// Required: true
	MergedBranchIds []int64 `json:"merged_branch_ids"`

	// The name of the Branch.
	// Required: true
	Name *string `json:"name"`

	// A true/false boolean indicating if the Branch is persistent; e.g. master.
	// Required: true
	Persistent *bool `json:"persistent"`

	// An array of PullRequests attached to the Branch (there is usually only one).
	// Required: true
	PullRequests []*PullRequest `json:"pull_requests"`

	// The ID of the Repository that contains the Branch.
	// Required: true
	RepositoryID *int64 `json:"repository_id"`

	// The time/date the Branch was updated.
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// The URL of the Branch.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this branch
func (m *Branch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMergedBranchIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersistent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Branch) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entity_type", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateMergedBranchIds(formats strfmt.Registry) error {

	if err := validate.Required("merged_branch_ids", "body", m.MergedBranchIds); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validatePersistent(formats strfmt.Registry) error {

	if err := validate.Required("persistent", "body", m.Persistent); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validatePullRequests(formats strfmt.Registry) error {

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

func (m *Branch) validateRepositoryID(formats strfmt.Registry) error {

	if err := validate.Required("repository_id", "body", m.RepositoryID); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Branch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Branch) UnmarshalBinary(b []byte) error {
	var res Branch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
