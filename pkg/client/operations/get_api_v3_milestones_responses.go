// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bateau-eng/chapi/v0/pkg/models"
)

// GetAPIV3MilestonesReader is a Reader for the GetAPIV3Milestones structure.
type GetAPIV3MilestonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3MilestonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3MilestonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3MilestonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3MilestonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3MilestonesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3MilestonesOK creates a GetAPIV3MilestonesOK with default headers values
func NewGetAPIV3MilestonesOK() *GetAPIV3MilestonesOK {
	return &GetAPIV3MilestonesOK{}
}

/*GetAPIV3MilestonesOK handles this case with default header values.

Resource
*/
type GetAPIV3MilestonesOK struct {
	Payload []*models.Milestone
}

func (o *GetAPIV3MilestonesOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones][%d] getApiV3MilestonesOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3MilestonesOK) GetPayload() []*models.Milestone {
	return o.Payload
}

func (o *GetAPIV3MilestonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3MilestonesBadRequest creates a GetAPIV3MilestonesBadRequest with default headers values
func NewGetAPIV3MilestonesBadRequest() *GetAPIV3MilestonesBadRequest {
	return &GetAPIV3MilestonesBadRequest{}
}

/*GetAPIV3MilestonesBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3MilestonesBadRequest struct {
}

func (o *GetAPIV3MilestonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones][%d] getApiV3MilestonesBadRequest ", 400)
}

func (o *GetAPIV3MilestonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3MilestonesNotFound creates a GetAPIV3MilestonesNotFound with default headers values
func NewGetAPIV3MilestonesNotFound() *GetAPIV3MilestonesNotFound {
	return &GetAPIV3MilestonesNotFound{}
}

/*GetAPIV3MilestonesNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3MilestonesNotFound struct {
}

func (o *GetAPIV3MilestonesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones][%d] getApiV3MilestonesNotFound ", 404)
}

func (o *GetAPIV3MilestonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3MilestonesUnprocessableEntity creates a GetAPIV3MilestonesUnprocessableEntity with default headers values
func NewGetAPIV3MilestonesUnprocessableEntity() *GetAPIV3MilestonesUnprocessableEntity {
	return &GetAPIV3MilestonesUnprocessableEntity{}
}

/*GetAPIV3MilestonesUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3MilestonesUnprocessableEntity struct {
}

func (o *GetAPIV3MilestonesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones][%d] getApiV3MilestonesUnprocessableEntity ", 422)
}

func (o *GetAPIV3MilestonesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}