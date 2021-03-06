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

// GetAPIV3EpicsReader is a Reader for the GetAPIV3Epics structure.
type GetAPIV3EpicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3EpicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3EpicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3EpicsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3EpicsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3EpicsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3EpicsOK creates a GetAPIV3EpicsOK with default headers values
func NewGetAPIV3EpicsOK() *GetAPIV3EpicsOK {
	return &GetAPIV3EpicsOK{}
}

/*GetAPIV3EpicsOK handles this case with default header values.

Resource
*/
type GetAPIV3EpicsOK struct {
	Payload []*models.EpicSlim
}

func (o *GetAPIV3EpicsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] getApiV3EpicsOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3EpicsOK) GetPayload() []*models.EpicSlim {
	return o.Payload
}

func (o *GetAPIV3EpicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3EpicsBadRequest creates a GetAPIV3EpicsBadRequest with default headers values
func NewGetAPIV3EpicsBadRequest() *GetAPIV3EpicsBadRequest {
	return &GetAPIV3EpicsBadRequest{}
}

/*GetAPIV3EpicsBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3EpicsBadRequest struct {
}

func (o *GetAPIV3EpicsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] getApiV3EpicsBadRequest ", 400)
}

func (o *GetAPIV3EpicsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EpicsNotFound creates a GetAPIV3EpicsNotFound with default headers values
func NewGetAPIV3EpicsNotFound() *GetAPIV3EpicsNotFound {
	return &GetAPIV3EpicsNotFound{}
}

/*GetAPIV3EpicsNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3EpicsNotFound struct {
}

func (o *GetAPIV3EpicsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] getApiV3EpicsNotFound ", 404)
}

func (o *GetAPIV3EpicsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3EpicsUnprocessableEntity creates a GetAPIV3EpicsUnprocessableEntity with default headers values
func NewGetAPIV3EpicsUnprocessableEntity() *GetAPIV3EpicsUnprocessableEntity {
	return &GetAPIV3EpicsUnprocessableEntity{}
}

/*GetAPIV3EpicsUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3EpicsUnprocessableEntity struct {
}

func (o *GetAPIV3EpicsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/epics][%d] getApiV3EpicsUnprocessableEntity ", 422)
}

func (o *GetAPIV3EpicsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
