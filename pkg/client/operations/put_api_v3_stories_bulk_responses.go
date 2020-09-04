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

// PutAPIV3StoriesBulkReader is a Reader for the PutAPIV3StoriesBulk structure.
type PutAPIV3StoriesBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3StoriesBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV3StoriesBulkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3StoriesBulkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3StoriesBulkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3StoriesBulkUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3StoriesBulkOK creates a PutAPIV3StoriesBulkOK with default headers values
func NewPutAPIV3StoriesBulkOK() *PutAPIV3StoriesBulkOK {
	return &PutAPIV3StoriesBulkOK{}
}

/*PutAPIV3StoriesBulkOK handles this case with default header values.

Resource
*/
type PutAPIV3StoriesBulkOK struct {
	Payload []*models.StorySlim
}

func (o *PutAPIV3StoriesBulkOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] putApiV3StoriesBulkOK  %+v", 200, o.Payload)
}

func (o *PutAPIV3StoriesBulkOK) GetPayload() []*models.StorySlim {
	return o.Payload
}

func (o *PutAPIV3StoriesBulkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAPIV3StoriesBulkBadRequest creates a PutAPIV3StoriesBulkBadRequest with default headers values
func NewPutAPIV3StoriesBulkBadRequest() *PutAPIV3StoriesBulkBadRequest {
	return &PutAPIV3StoriesBulkBadRequest{}
}

/*PutAPIV3StoriesBulkBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3StoriesBulkBadRequest struct {
}

func (o *PutAPIV3StoriesBulkBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] putApiV3StoriesBulkBadRequest ", 400)
}

func (o *PutAPIV3StoriesBulkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoriesBulkNotFound creates a PutAPIV3StoriesBulkNotFound with default headers values
func NewPutAPIV3StoriesBulkNotFound() *PutAPIV3StoriesBulkNotFound {
	return &PutAPIV3StoriesBulkNotFound{}
}

/*PutAPIV3StoriesBulkNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3StoriesBulkNotFound struct {
}

func (o *PutAPIV3StoriesBulkNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] putApiV3StoriesBulkNotFound ", 404)
}

func (o *PutAPIV3StoriesBulkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3StoriesBulkUnprocessableEntity creates a PutAPIV3StoriesBulkUnprocessableEntity with default headers values
func NewPutAPIV3StoriesBulkUnprocessableEntity() *PutAPIV3StoriesBulkUnprocessableEntity {
	return &PutAPIV3StoriesBulkUnprocessableEntity{}
}

/*PutAPIV3StoriesBulkUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3StoriesBulkUnprocessableEntity struct {
}

func (o *PutAPIV3StoriesBulkUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/stories/bulk][%d] putApiV3StoriesBulkUnprocessableEntity ", 422)
}

func (o *PutAPIV3StoriesBulkUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}