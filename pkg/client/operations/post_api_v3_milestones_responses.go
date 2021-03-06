// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bateau-eng/chapi/pkg/models"
)

// PostAPIV3MilestonesReader is a Reader for the PostAPIV3Milestones structure.
type PostAPIV3MilestonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3MilestonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostAPIV3MilestonesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAPIV3MilestonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAPIV3MilestonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostAPIV3MilestonesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV3MilestonesCreated creates a PostAPIV3MilestonesCreated with default headers values
func NewPostAPIV3MilestonesCreated() *PostAPIV3MilestonesCreated {
	return &PostAPIV3MilestonesCreated{}
}

/*PostAPIV3MilestonesCreated handles this case with default header values.

Resource
*/
type PostAPIV3MilestonesCreated struct {
	Payload *models.Milestone
}

func (o *PostAPIV3MilestonesCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] postApiV3MilestonesCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3MilestonesCreated) GetPayload() *models.Milestone {
	return o.Payload
}

func (o *PostAPIV3MilestonesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Milestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIV3MilestonesBadRequest creates a PostAPIV3MilestonesBadRequest with default headers values
func NewPostAPIV3MilestonesBadRequest() *PostAPIV3MilestonesBadRequest {
	return &PostAPIV3MilestonesBadRequest{}
}

/*PostAPIV3MilestonesBadRequest handles this case with default header values.

Schema mismatch
*/
type PostAPIV3MilestonesBadRequest struct {
}

func (o *PostAPIV3MilestonesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] postApiV3MilestonesBadRequest ", 400)
}

func (o *PostAPIV3MilestonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3MilestonesNotFound creates a PostAPIV3MilestonesNotFound with default headers values
func NewPostAPIV3MilestonesNotFound() *PostAPIV3MilestonesNotFound {
	return &PostAPIV3MilestonesNotFound{}
}

/*PostAPIV3MilestonesNotFound handles this case with default header values.

Resource does not exist
*/
type PostAPIV3MilestonesNotFound struct {
}

func (o *PostAPIV3MilestonesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] postApiV3MilestonesNotFound ", 404)
}

func (o *PostAPIV3MilestonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIV3MilestonesUnprocessableEntity creates a PostAPIV3MilestonesUnprocessableEntity with default headers values
func NewPostAPIV3MilestonesUnprocessableEntity() *PostAPIV3MilestonesUnprocessableEntity {
	return &PostAPIV3MilestonesUnprocessableEntity{}
}

/*PostAPIV3MilestonesUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PostAPIV3MilestonesUnprocessableEntity struct {
}

func (o *PostAPIV3MilestonesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v3/milestones][%d] postApiV3MilestonesUnprocessableEntity ", 422)
}

func (o *PostAPIV3MilestonesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
