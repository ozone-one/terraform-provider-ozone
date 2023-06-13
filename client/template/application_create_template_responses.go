// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// ApplicationCreateTemplateReader is a Reader for the ApplicationCreateTemplate structure.
type ApplicationCreateTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationCreateTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationCreateTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplicationCreateTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewApplicationCreateTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplicationCreateTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewApplicationCreateTemplateOK creates a ApplicationCreateTemplateOK with default headers values
func NewApplicationCreateTemplateOK() *ApplicationCreateTemplateOK {
	return &ApplicationCreateTemplateOK{}
}

/*
ApplicationCreateTemplateOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationCreateTemplateOK struct {
	Payload *models.ApplicationTemplates
}

func (o *ApplicationCreateTemplateOK) Error() string {
	return fmt.Sprintf("[POST /admin/app/template][%d] applicationCreateTemplateOK  %+v", 200, o.Payload)
}
func (o *ApplicationCreateTemplateOK) GetPayload() *models.ApplicationTemplates {
	return o.Payload
}

func (o *ApplicationCreateTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationTemplates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationCreateTemplateBadRequest creates a ApplicationCreateTemplateBadRequest with default headers values
func NewApplicationCreateTemplateBadRequest() *ApplicationCreateTemplateBadRequest {
	return &ApplicationCreateTemplateBadRequest{}
}

/*
ApplicationCreateTemplateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ApplicationCreateTemplateBadRequest struct {
	Payload interface{}
}

func (o *ApplicationCreateTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/app/template][%d] applicationCreateTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *ApplicationCreateTemplateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationCreateTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationCreateTemplateNotFound creates a ApplicationCreateTemplateNotFound with default headers values
func NewApplicationCreateTemplateNotFound() *ApplicationCreateTemplateNotFound {
	return &ApplicationCreateTemplateNotFound{}
}

/*
ApplicationCreateTemplateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ApplicationCreateTemplateNotFound struct {
	Payload interface{}
}

func (o *ApplicationCreateTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/app/template][%d] applicationCreateTemplateNotFound  %+v", 404, o.Payload)
}
func (o *ApplicationCreateTemplateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationCreateTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationCreateTemplateInternalServerError creates a ApplicationCreateTemplateInternalServerError with default headers values
func NewApplicationCreateTemplateInternalServerError() *ApplicationCreateTemplateInternalServerError {
	return &ApplicationCreateTemplateInternalServerError{}
}

/*
ApplicationCreateTemplateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ApplicationCreateTemplateInternalServerError struct {
	Payload interface{}
}

func (o *ApplicationCreateTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/app/template][%d] applicationCreateTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *ApplicationCreateTemplateInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationCreateTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
