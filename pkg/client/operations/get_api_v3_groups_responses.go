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

// GetAPIV3GroupsReader is a Reader for the GetAPIV3Groups structure.
type GetAPIV3GroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3GroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3GroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3GroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3GroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3GroupsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3GroupsOK creates a GetAPIV3GroupsOK with default headers values
func NewGetAPIV3GroupsOK() *GetAPIV3GroupsOK {
	return &GetAPIV3GroupsOK{}
}

/*GetAPIV3GroupsOK handles this case with default header values.

Resource
*/
type GetAPIV3GroupsOK struct {
	Payload []*models.Group
}

func (o *GetAPIV3GroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups][%d] getApiV3GroupsOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3GroupsOK) GetPayload() []*models.Group {
	return o.Payload
}

func (o *GetAPIV3GroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3GroupsBadRequest creates a GetAPIV3GroupsBadRequest with default headers values
func NewGetAPIV3GroupsBadRequest() *GetAPIV3GroupsBadRequest {
	return &GetAPIV3GroupsBadRequest{}
}

/*GetAPIV3GroupsBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3GroupsBadRequest struct {
}

func (o *GetAPIV3GroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups][%d] getApiV3GroupsBadRequest ", 400)
}

func (o *GetAPIV3GroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3GroupsNotFound creates a GetAPIV3GroupsNotFound with default headers values
func NewGetAPIV3GroupsNotFound() *GetAPIV3GroupsNotFound {
	return &GetAPIV3GroupsNotFound{}
}

/*GetAPIV3GroupsNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3GroupsNotFound struct {
}

func (o *GetAPIV3GroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups][%d] getApiV3GroupsNotFound ", 404)
}

func (o *GetAPIV3GroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3GroupsUnprocessableEntity creates a GetAPIV3GroupsUnprocessableEntity with default headers values
func NewGetAPIV3GroupsUnprocessableEntity() *GetAPIV3GroupsUnprocessableEntity {
	return &GetAPIV3GroupsUnprocessableEntity{}
}

/*GetAPIV3GroupsUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3GroupsUnprocessableEntity struct {
}

func (o *GetAPIV3GroupsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/groups][%d] getApiV3GroupsUnprocessableEntity ", 422)
}

func (o *GetAPIV3GroupsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
