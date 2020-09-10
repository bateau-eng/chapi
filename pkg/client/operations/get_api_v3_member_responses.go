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

// GetAPIV3MemberReader is a Reader for the GetAPIV3Member structure.
type GetAPIV3MemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3MemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV3MemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAPIV3MemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIV3MemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAPIV3MemberUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV3MemberOK creates a GetAPIV3MemberOK with default headers values
func NewGetAPIV3MemberOK() *GetAPIV3MemberOK {
	return &GetAPIV3MemberOK{}
}

/*GetAPIV3MemberOK handles this case with default header values.

Resource
*/
type GetAPIV3MemberOK struct {
	Payload *models.MemberInfo
}

func (o *GetAPIV3MemberOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getApiV3MemberOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3MemberOK) GetPayload() *models.MemberInfo {
	return o.Payload
}

func (o *GetAPIV3MemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIV3MemberBadRequest creates a GetAPIV3MemberBadRequest with default headers values
func NewGetAPIV3MemberBadRequest() *GetAPIV3MemberBadRequest {
	return &GetAPIV3MemberBadRequest{}
}

/*GetAPIV3MemberBadRequest handles this case with default header values.

Schema mismatch
*/
type GetAPIV3MemberBadRequest struct {
}

func (o *GetAPIV3MemberBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getApiV3MemberBadRequest ", 400)
}

func (o *GetAPIV3MemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3MemberNotFound creates a GetAPIV3MemberNotFound with default headers values
func NewGetAPIV3MemberNotFound() *GetAPIV3MemberNotFound {
	return &GetAPIV3MemberNotFound{}
}

/*GetAPIV3MemberNotFound handles this case with default header values.

Resource does not exist
*/
type GetAPIV3MemberNotFound struct {
}

func (o *GetAPIV3MemberNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getApiV3MemberNotFound ", 404)
}

func (o *GetAPIV3MemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIV3MemberUnprocessableEntity creates a GetAPIV3MemberUnprocessableEntity with default headers values
func NewGetAPIV3MemberUnprocessableEntity() *GetAPIV3MemberUnprocessableEntity {
	return &GetAPIV3MemberUnprocessableEntity{}
}

/*GetAPIV3MemberUnprocessableEntity handles this case with default header values.

Unprocessable
*/
type GetAPIV3MemberUnprocessableEntity struct {
}

func (o *GetAPIV3MemberUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v3/member][%d] getApiV3MemberUnprocessableEntity ", 422)
}

func (o *GetAPIV3MemberUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
