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

// GetAPIV3EpicWorkflowReader is a Reader for the GetAPIV3EpicWorkflow structure.
type GetAPIV3EpicWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3EpicWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3EpicWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3EpicWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3EpicWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3EpicWorkflowUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3EpicWorkflowOK creates a GetAPIV3EpicWorkflowOK with default headers values
func NewGetAPIV3EpicWorkflowOK() *GetAPIV3EpicWorkflowOK {
	return &GetAPIV3EpicWorkflowOK{}
}

/*GetAPIV3EpicWorkflowOK handles this case with default header values.

Resource
*/
type GetAPIV3EpicWorkflowOK struct {
	Payload *models.EpicWorkflow
}

func (o *GetAPIV3EpicWorkflowOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/epic-workflow][%d] getApiV3EpicWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3EpicWorkflowOK) GetPayload() *models.EpicWorkflow {
	return o.Payload
}

func (o *GetAPIV3EpicWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EpicWorkflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3EpicWorkflowBadRequest creates a GetAPIV3EpicWorkflowBadRequest with default headers values
func NewGetAPIV3EpicWorkflowBadRequest() *GetAPIV3EpicWorkflowBadRequest {
	return &GetAPIV3EpicWorkflowBadRequest{}
}

/*GetAPIV3EpicWorkflowBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3EpicWorkflowBadRequest struct {
}

func (o *GetAPIV3EpicWorkflowBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/epic-workflow][%d] getApiV3EpicWorkflowBadRequest ", 400)
}

func (o *GetAPIV3EpicWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EpicWorkflowNotFound creates a GetAPIV3EpicWorkflowNotFound with default headers values
func NewGetAPIV3EpicWorkflowNotFound() *GetAPIV3EpicWorkflowNotFound {
	return &GetAPIV3EpicWorkflowNotFound{}
}

/*GetAPIV3EpicWorkflowNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3EpicWorkflowNotFound struct {
}

func (o *GetAPIV3EpicWorkflowNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/epic-workflow][%d] getApiV3EpicWorkflowNotFound ", 404)
}

func (o *GetAPIV3EpicWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EpicWorkflowUnprocessableEntity creates a GetAPIV3EpicWorkflowUnprocessableEntity with default headers values
func NewGetAPIV3EpicWorkflowUnprocessableEntity() *GetAPIV3EpicWorkflowUnprocessableEntity {
	return &GetAPIV3EpicWorkflowUnprocessableEntity{}
}

/*GetAPIV3EpicWorkflowUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3EpicWorkflowUnprocessableEntity struct {
}

func (o *GetAPIV3EpicWorkflowUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/epic-workflow][%d] getApiV3EpicWorkflowUnprocessableEntity ", 422)
}

func (o *GetAPIV3EpicWorkflowUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
