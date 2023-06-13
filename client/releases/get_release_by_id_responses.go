// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// GetReleaseByIDReader is a Reader for the GetReleaseByID structure.
type GetReleaseByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReleaseByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReleaseByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReleaseByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReleaseByIDOK creates a GetReleaseByIDOK with default headers values
func NewGetReleaseByIDOK() *GetReleaseByIDOK {
	return &GetReleaseByIDOK{}
}

/*
GetReleaseByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetReleaseByIDOK struct {
	Payload *models.AppRelease
}

func (o *GetReleaseByIDOK) Error() string {
	return fmt.Sprintf("[GET /admin/release/{id}][%d] getReleaseByIdOK  %+v", 200, o.Payload)
}
func (o *GetReleaseByIDOK) GetPayload() *models.AppRelease {
	return o.Payload
}

func (o *GetReleaseByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseByIDBadRequest creates a GetReleaseByIDBadRequest with default headers values
func NewGetReleaseByIDBadRequest() *GetReleaseByIDBadRequest {
	return &GetReleaseByIDBadRequest{}
}

/*
GetReleaseByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetReleaseByIDBadRequest struct {
	Payload interface{}
}

func (o *GetReleaseByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/release/{id}][%d] getReleaseByIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetReleaseByIDBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *GetReleaseByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseByIDNotFound creates a GetReleaseByIDNotFound with default headers values
func NewGetReleaseByIDNotFound() *GetReleaseByIDNotFound {
	return &GetReleaseByIDNotFound{}
}

/*
GetReleaseByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetReleaseByIDNotFound struct {
	Payload interface{}
}

func (o *GetReleaseByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/release/{id}][%d] getReleaseByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetReleaseByIDNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetReleaseByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseByIDInternalServerError creates a GetReleaseByIDInternalServerError with default headers values
func NewGetReleaseByIDInternalServerError() *GetReleaseByIDInternalServerError {
	return &GetReleaseByIDInternalServerError{}
}

/*
GetReleaseByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReleaseByIDInternalServerError struct {
	Payload interface{}
}

func (o *GetReleaseByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/release/{id}][%d] getReleaseByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReleaseByIDInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *GetReleaseByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}