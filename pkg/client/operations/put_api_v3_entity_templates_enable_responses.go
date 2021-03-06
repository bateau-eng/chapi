// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutAPIV3EntityTemplatesEnableReader is a Reader for the PutAPIV3EntityTemplatesEnable structure.
type PutAPIV3EntityTemplatesEnableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3EntityTemplatesEnableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutAPIV3EntityTemplatesEnableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAPIV3EntityTemplatesEnableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAPIV3EntityTemplatesEnableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPutAPIV3EntityTemplatesEnableUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAPIV3EntityTemplatesEnableNoContent creates a PutAPIV3EntityTemplatesEnableNoContent with default headers values
func NewPutAPIV3EntityTemplatesEnableNoContent() *PutAPIV3EntityTemplatesEnableNoContent {
	return &PutAPIV3EntityTemplatesEnableNoContent{}
}

/*PutAPIV3EntityTemplatesEnableNoContent handles this case with default header values.

No Content
*/
type PutAPIV3EntityTemplatesEnableNoContent struct {
}

func (o *PutAPIV3EntityTemplatesEnableNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/enable][%d] putApiV3EntityTemplatesEnableNoContent ", 204)
}

func (o *PutAPIV3EntityTemplatesEnableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3EntityTemplatesEnableBadRequest creates a PutAPIV3EntityTemplatesEnableBadRequest with default headers values
func NewPutAPIV3EntityTemplatesEnableBadRequest() *PutAPIV3EntityTemplatesEnableBadRequest {
	return &PutAPIV3EntityTemplatesEnableBadRequest{}
}

/*PutAPIV3EntityTemplatesEnableBadRequest handles this case with default header values.

Schema mismatch
*/
type PutAPIV3EntityTemplatesEnableBadRequest struct {
}

func (o *PutAPIV3EntityTemplatesEnableBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/enable][%d] putApiV3EntityTemplatesEnableBadRequest ", 400)
}

func (o *PutAPIV3EntityTemplatesEnableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3EntityTemplatesEnableNotFound creates a PutAPIV3EntityTemplatesEnableNotFound with default headers values
func NewPutAPIV3EntityTemplatesEnableNotFound() *PutAPIV3EntityTemplatesEnableNotFound {
	return &PutAPIV3EntityTemplatesEnableNotFound{}
}

/*PutAPIV3EntityTemplatesEnableNotFound handles this case with default header values.

Resource does not exist
*/
type PutAPIV3EntityTemplatesEnableNotFound struct {
}

func (o *PutAPIV3EntityTemplatesEnableNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/enable][%d] putApiV3EntityTemplatesEnableNotFound ", 404)
}

func (o *PutAPIV3EntityTemplatesEnableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutAPIV3EntityTemplatesEnableUnprocessableEntity creates a PutAPIV3EntityTemplatesEnableUnprocessableEntity with default headers values
func NewPutAPIV3EntityTemplatesEnableUnprocessableEntity() *PutAPIV3EntityTemplatesEnableUnprocessableEntity {
	return &PutAPIV3EntityTemplatesEnableUnprocessableEntity{}
}

/*PutAPIV3EntityTemplatesEnableUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type PutAPIV3EntityTemplatesEnableUnprocessableEntity struct {
}

func (o *PutAPIV3EntityTemplatesEnableUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/v3/entity-templates/enable][%d] putApiV3EntityTemplatesEnableUnprocessableEntity ", 422)
}

func (o *PutAPIV3EntityTemplatesEnableUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
