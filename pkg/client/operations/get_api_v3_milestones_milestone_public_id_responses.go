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

// GetAPIV3MilestonesMilestonePublicIDReader is a Reader for the GetAPIV3MilestonesMilestonePublicID structure.
type GetAPIV3MilestonesMilestonePublicIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3MilestonesMilestonePublicIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3MilestonesMilestonePublicIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3MilestonesMilestonePublicIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3MilestonesMilestonePublicIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3MilestonesMilestonePublicIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3MilestonesMilestonePublicIDOK creates a GetAPIV3MilestonesMilestonePublicIDOK with default headers values
func NewGetAPIV3MilestonesMilestonePublicIDOK() *GetAPIV3MilestonesMilestonePublicIDOK {
	return &GetAPIV3MilestonesMilestonePublicIDOK{}
}

/*GetAPIV3MilestonesMilestonePublicIDOK handles this case with default header values.

Resource
*/
type GetAPIV3MilestonesMilestonePublicIDOK struct {
	Payload *models.Milestone
}

func (o *GetAPIV3MilestonesMilestonePublicIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getApiV3MilestonesMilestonePublicIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3MilestonesMilestonePublicIDOK) GetPayload() *models.Milestone {
	return o.Payload
}

func (o *GetAPIV3MilestonesMilestonePublicIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Milestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3MilestonesMilestonePublicIDBadRequest creates a GetAPIV3MilestonesMilestonePublicIDBadRequest with default headers values
func NewGetAPIV3MilestonesMilestonePublicIDBadRequest() *GetAPIV3MilestonesMilestonePublicIDBadRequest {
	return &GetAPIV3MilestonesMilestonePublicIDBadRequest{}
}

/*GetAPIV3MilestonesMilestonePublicIDBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3MilestonesMilestonePublicIDBadRequest struct {
}

func (o *GetAPIV3MilestonesMilestonePublicIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getApiV3MilestonesMilestonePublicIdBadRequest ", 400)
}

func (o *GetAPIV3MilestonesMilestonePublicIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3MilestonesMilestonePublicIDNotFound creates a GetAPIV3MilestonesMilestonePublicIDNotFound with default headers values
func NewGetAPIV3MilestonesMilestonePublicIDNotFound() *GetAPIV3MilestonesMilestonePublicIDNotFound {
	return &GetAPIV3MilestonesMilestonePublicIDNotFound{}
}

/*GetAPIV3MilestonesMilestonePublicIDNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3MilestonesMilestonePublicIDNotFound struct {
}

func (o *GetAPIV3MilestonesMilestonePublicIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getApiV3MilestonesMilestonePublicIdNotFound ", 404)
}

func (o *GetAPIV3MilestonesMilestonePublicIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3MilestonesMilestonePublicIDUnprocessableEntity creates a GetAPIV3MilestonesMilestonePublicIDUnprocessableEntity with default headers values
func NewGetAPIV3MilestonesMilestonePublicIDUnprocessableEntity() *GetAPIV3MilestonesMilestonePublicIDUnprocessableEntity {
	return &GetAPIV3MilestonesMilestonePublicIDUnprocessableEntity{}
}

/*GetAPIV3MilestonesMilestonePublicIDUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3MilestonesMilestonePublicIDUnprocessableEntity struct {
}

func (o *GetAPIV3MilestonesMilestonePublicIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/milestones/{milestone-public-id}][%d] getApiV3MilestonesMilestonePublicIdUnprocessableEntity ", 422)
}

func (o *GetAPIV3MilestonesMilestonePublicIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
