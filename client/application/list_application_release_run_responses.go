// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// ListApplicationReleaseRunReader is a Reader for the ListApplicationReleaseRun structure.
type ListApplicationReleaseRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListApplicationReleaseRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListApplicationReleaseRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListApplicationReleaseRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListApplicationReleaseRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListApplicationReleaseRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListApplicationReleaseRunOK creates a ListApplicationReleaseRunOK with default headers values
func NewListApplicationReleaseRunOK() *ListApplicationReleaseRunOK {
	return &ListApplicationReleaseRunOK{}
}

/*
ListApplicationReleaseRunOK describes a response with status code 200, with default header values.

OK
*/
type ListApplicationReleaseRunOK struct {
	Payload *models.ReleaseRunList
}

func (o *ListApplicationReleaseRunOK) Error() string {
	return fmt.Sprintf("[GET /admin/app/{id}/releaserun][%d] listApplicationReleaseRunOK  %+v", 200, o.Payload)
}
func (o *ListApplicationReleaseRunOK) GetPayload() *models.ReleaseRunList {
	return o.Payload
}

func (o *ListApplicationReleaseRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseRunList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListApplicationReleaseRunBadRequest creates a ListApplicationReleaseRunBadRequest with default headers values
func NewListApplicationReleaseRunBadRequest() *ListApplicationReleaseRunBadRequest {
	return &ListApplicationReleaseRunBadRequest{}
}

/*
ListApplicationReleaseRunBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListApplicationReleaseRunBadRequest struct {
	Payload interface{}
}

func (o *ListApplicationReleaseRunBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/app/{id}/releaserun][%d] listApplicationReleaseRunBadRequest  %+v", 400, o.Payload)
}
func (o *ListApplicationReleaseRunBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ListApplicationReleaseRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListApplicationReleaseRunNotFound creates a ListApplicationReleaseRunNotFound with default headers values
func NewListApplicationReleaseRunNotFound() *ListApplicationReleaseRunNotFound {
	return &ListApplicationReleaseRunNotFound{}
}

/*
ListApplicationReleaseRunNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListApplicationReleaseRunNotFound struct {
	Payload interface{}
}

func (o *ListApplicationReleaseRunNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/app/{id}/releaserun][%d] listApplicationReleaseRunNotFound  %+v", 404, o.Payload)
}
func (o *ListApplicationReleaseRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListApplicationReleaseRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListApplicationReleaseRunInternalServerError creates a ListApplicationReleaseRunInternalServerError with default headers values
func NewListApplicationReleaseRunInternalServerError() *ListApplicationReleaseRunInternalServerError {
	return &ListApplicationReleaseRunInternalServerError{}
}

/*
ListApplicationReleaseRunInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ListApplicationReleaseRunInternalServerError struct {
	Payload interface{}
}

func (o *ListApplicationReleaseRunInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/app/{id}/releaserun][%d] listApplicationReleaseRunInternalServerError  %+v", 500, o.Payload)
}
func (o *ListApplicationReleaseRunInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ListApplicationReleaseRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}