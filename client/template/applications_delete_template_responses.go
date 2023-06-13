// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ApplicationsDeleteTemplateReader is a Reader for the ApplicationsDeleteTemplate structure.
type ApplicationsDeleteTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationsDeleteTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationsDeleteTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplicationsDeleteTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewApplicationsDeleteTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplicationsDeleteTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewApplicationsDeleteTemplateOK creates a ApplicationsDeleteTemplateOK with default headers values
func NewApplicationsDeleteTemplateOK() *ApplicationsDeleteTemplateOK {
	return &ApplicationsDeleteTemplateOK{}
}

/*
ApplicationsDeleteTemplateOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationsDeleteTemplateOK struct {
	Payload interface{}
}

func (o *ApplicationsDeleteTemplateOK) Error() string {
	return fmt.Sprintf("[DELETE /admin/app/template][%d] applicationsDeleteTemplateOK  %+v", 200, o.Payload)
}
func (o *ApplicationsDeleteTemplateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationsDeleteTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationsDeleteTemplateBadRequest creates a ApplicationsDeleteTemplateBadRequest with default headers values
func NewApplicationsDeleteTemplateBadRequest() *ApplicationsDeleteTemplateBadRequest {
	return &ApplicationsDeleteTemplateBadRequest{}
}

/*
ApplicationsDeleteTemplateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ApplicationsDeleteTemplateBadRequest struct {
	Payload interface{}
}

func (o *ApplicationsDeleteTemplateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /admin/app/template][%d] applicationsDeleteTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *ApplicationsDeleteTemplateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationsDeleteTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationsDeleteTemplateNotFound creates a ApplicationsDeleteTemplateNotFound with default headers values
func NewApplicationsDeleteTemplateNotFound() *ApplicationsDeleteTemplateNotFound {
	return &ApplicationsDeleteTemplateNotFound{}
}

/*
ApplicationsDeleteTemplateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ApplicationsDeleteTemplateNotFound struct {
	Payload interface{}
}

func (o *ApplicationsDeleteTemplateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /admin/app/template][%d] applicationsDeleteTemplateNotFound  %+v", 404, o.Payload)
}
func (o *ApplicationsDeleteTemplateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationsDeleteTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationsDeleteTemplateInternalServerError creates a ApplicationsDeleteTemplateInternalServerError with default headers values
func NewApplicationsDeleteTemplateInternalServerError() *ApplicationsDeleteTemplateInternalServerError {
	return &ApplicationsDeleteTemplateInternalServerError{}
}

/*
ApplicationsDeleteTemplateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ApplicationsDeleteTemplateInternalServerError struct {
	Payload interface{}
}

func (o *ApplicationsDeleteTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /admin/app/template][%d] applicationsDeleteTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *ApplicationsDeleteTemplateInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationsDeleteTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
