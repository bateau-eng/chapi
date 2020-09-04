// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutAPIV3IterationsDisableReader is a Reader for the PutAPIV3IterationsDisable structure.
type PutAPIV3IterationsDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3IterationsDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutAPIV3IterationsDisableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3IterationsDisableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3IterationsDisableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3IterationsDisableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3IterationsDisableNoContent creates a PutAPIV3IterationsDisableNoContent with default headers values
func NewPutAPIV3IterationsDisableNoContent() *PutAPIV3IterationsDisableNoContent {
	return &PutAPIV3IterationsDisableNoContent{}
}

/*PutAPIV3IterationsDisableNoContent handles this case with default header values.

No Content
*/
type PutAPIV3IterationsDisableNoContent struct {
}

func (o *PutAPIV3IterationsDisableNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/disable][%d] putApiV3IterationsDisableNoContent ", 204)
}

func (o *PutAPIV3IterationsDisableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3IterationsDisableBadRequest creates a PutAPIV3IterationsDisableBadRequest with default headers values
func NewPutAPIV3IterationsDisableBadRequest() *PutAPIV3IterationsDisableBadRequest {
	return &PutAPIV3IterationsDisableBadRequest{}
}

/*PutAPIV3IterationsDisableBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3IterationsDisableBadRequest struct {
}

func (o *PutAPIV3IterationsDisableBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/disable][%d] putApiV3IterationsDisableBadRequest ", 400)
}

func (o *PutAPIV3IterationsDisableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3IterationsDisableNotFound creates a PutAPIV3IterationsDisableNotFound with default headers values
func NewPutAPIV3IterationsDisableNotFound() *PutAPIV3IterationsDisableNotFound {
	return &PutAPIV3IterationsDisableNotFound{}
}

/*PutAPIV3IterationsDisableNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3IterationsDisableNotFound struct {
}

func (o *PutAPIV3IterationsDisableNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/disable][%d] putApiV3IterationsDisableNotFound ", 404)
}

func (o *PutAPIV3IterationsDisableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3IterationsDisableUnprocessableEntity creates a PutAPIV3IterationsDisableUnprocessableEntity with default headers values
func NewPutAPIV3IterationsDisableUnprocessableEntity() *PutAPIV3IterationsDisableUnprocessableEntity {
	return &PutAPIV3IterationsDisableUnprocessableEntity{}
}

/*PutAPIV3IterationsDisableUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3IterationsDisableUnprocessableEntity struct {
}

func (o *PutAPIV3IterationsDisableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/iterations/disable][%d] putApiV3IterationsDisableUnprocessableEntity ", 422)
}

func (o *PutAPIV3IterationsDisableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
