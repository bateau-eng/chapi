// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutAPIV3GroupsDisableReader is a Reader for the PutAPIV3GroupsDisable structure.
type PutAPIV3GroupsDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3GroupsDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutAPIV3GroupsDisableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3GroupsDisableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3GroupsDisableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3GroupsDisableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3GroupsDisableNoContent creates a PutAPIV3GroupsDisableNoContent with default headers values
func NewPutAPIV3GroupsDisableNoContent() *PutAPIV3GroupsDisableNoContent {
	return &PutAPIV3GroupsDisableNoContent{}
}

/*PutAPIV3GroupsDisableNoContent handles this case with default header values.

No Content
*/
type PutAPIV3GroupsDisableNoContent struct {
}

func (o *PutAPIV3GroupsDisableNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/disable][%d] putApiV3GroupsDisableNoContent ", 204)
}

func (o *PutAPIV3GroupsDisableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3GroupsDisableBadRequest creates a PutAPIV3GroupsDisableBadRequest with default headers values
func NewPutAPIV3GroupsDisableBadRequest() *PutAPIV3GroupsDisableBadRequest {
	return &PutAPIV3GroupsDisableBadRequest{}
}

/*PutAPIV3GroupsDisableBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3GroupsDisableBadRequest struct {
}

func (o *PutAPIV3GroupsDisableBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/disable][%d] putApiV3GroupsDisableBadRequest ", 400)
}

func (o *PutAPIV3GroupsDisableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3GroupsDisableNotFound creates a PutAPIV3GroupsDisableNotFound with default headers values
func NewPutAPIV3GroupsDisableNotFound() *PutAPIV3GroupsDisableNotFound {
	return &PutAPIV3GroupsDisableNotFound{}
}

/*PutAPIV3GroupsDisableNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3GroupsDisableNotFound struct {
}

func (o *PutAPIV3GroupsDisableNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/disable][%d] putApiV3GroupsDisableNotFound ", 404)
}

func (o *PutAPIV3GroupsDisableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3GroupsDisableUnprocessableEntity creates a PutAPIV3GroupsDisableUnprocessableEntity with default headers values
func NewPutAPIV3GroupsDisableUnprocessableEntity() *PutAPIV3GroupsDisableUnprocessableEntity {
	return &PutAPIV3GroupsDisableUnprocessableEntity{}
}

/*PutAPIV3GroupsDisableUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3GroupsDisableUnprocessableEntity struct {
}

func (o *PutAPIV3GroupsDisableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/groups/disable][%d] putApiV3GroupsDisableUnprocessableEntity ", 422)
}

func (o *PutAPIV3GroupsDisableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
