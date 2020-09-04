// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutAPIV3EntityTemplatesDisableReader is a Reader for the PutAPIV3EntityTemplatesDisable structure.
type PutAPIV3EntityTemplatesDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3EntityTemplatesDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutAPIV3EntityTemplatesDisableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3EntityTemplatesDisableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3EntityTemplatesDisableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3EntityTemplatesDisableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3EntityTemplatesDisableNoContent creates a PutAPIV3EntityTemplatesDisableNoContent with default headers values
func NewPutAPIV3EntityTemplatesDisableNoContent() *PutAPIV3EntityTemplatesDisableNoContent {
	return &PutAPIV3EntityTemplatesDisableNoContent{}
}

/*PutAPIV3EntityTemplatesDisableNoContent handles this case with default header values.

No Content
*/
type PutAPIV3EntityTemplatesDisableNoContent struct {
}

func (o *PutAPIV3EntityTemplatesDisableNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] putApiV3EntityTemplatesDisableNoContent ", 204)
}

func (o *PutAPIV3EntityTemplatesDisableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3EntityTemplatesDisableBadRequest creates a PutAPIV3EntityTemplatesDisableBadRequest with default headers values
func NewPutAPIV3EntityTemplatesDisableBadRequest() *PutAPIV3EntityTemplatesDisableBadRequest {
	return &PutAPIV3EntityTemplatesDisableBadRequest{}
}

/*PutAPIV3EntityTemplatesDisableBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3EntityTemplatesDisableBadRequest struct {
}

func (o *PutAPIV3EntityTemplatesDisableBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] putApiV3EntityTemplatesDisableBadRequest ", 400)
}

func (o *PutAPIV3EntityTemplatesDisableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3EntityTemplatesDisableNotFound creates a PutAPIV3EntityTemplatesDisableNotFound with default headers values
func NewPutAPIV3EntityTemplatesDisableNotFound() *PutAPIV3EntityTemplatesDisableNotFound {
	return &PutAPIV3EntityTemplatesDisableNotFound{}
}

/*PutAPIV3EntityTemplatesDisableNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3EntityTemplatesDisableNotFound struct {
}

func (o *PutAPIV3EntityTemplatesDisableNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] putApiV3EntityTemplatesDisableNotFound ", 404)
}

func (o *PutAPIV3EntityTemplatesDisableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3EntityTemplatesDisableUnprocessableEntity creates a PutAPIV3EntityTemplatesDisableUnprocessableEntity with default headers values
func NewPutAPIV3EntityTemplatesDisableUnprocessableEntity() *PutAPIV3EntityTemplatesDisableUnprocessableEntity {
	return &PutAPIV3EntityTemplatesDisableUnprocessableEntity{}
}

/*PutAPIV3EntityTemplatesDisableUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3EntityTemplatesDisableUnprocessableEntity struct {
}

func (o *PutAPIV3EntityTemplatesDisableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/disable][%d] putApiV3EntityTemplatesDisableUnprocessableEntity ", 422)
}

func (o *PutAPIV3EntityTemplatesDisableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}