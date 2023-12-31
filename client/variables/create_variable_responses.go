// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// CreateVariableReader is a Reader for the CreateVariable structure.
type CreateVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVariableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVariableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVariableOK creates a CreateVariableOK with default headers values
func NewCreateVariableOK() *CreateVariableOK {
	return &CreateVariableOK{}
}

/*
CreateVariableOK describes a response with status code 200, with default header values.

OK
*/
type CreateVariableOK struct {
	Payload *models.VariableView
}

func (o *CreateVariableOK) Error() string {
	return fmt.Sprintf("[POST /admin/variables][%d] createVariableOK  %+v", 200, o.Payload)
}
func (o *CreateVariableOK) GetPayload() *models.VariableView {
	return o.Payload
}

func (o *CreateVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableBadRequest creates a CreateVariableBadRequest with default headers values
func NewCreateVariableBadRequest() *CreateVariableBadRequest {
	return &CreateVariableBadRequest{}
}

/*
CreateVariableBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateVariableBadRequest struct {
	Payload interface{}
}

func (o *CreateVariableBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/variables][%d] createVariableBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVariableBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateVariableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableNotFound creates a CreateVariableNotFound with default headers values
func NewCreateVariableNotFound() *CreateVariableNotFound {
	return &CreateVariableNotFound{}
}

/*
CreateVariableNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateVariableNotFound struct {
	Payload interface{}
}

func (o *CreateVariableNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/variables][%d] createVariableNotFound  %+v", 404, o.Payload)
}
func (o *CreateVariableNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableInternalServerError creates a CreateVariableInternalServerError with default headers values
func NewCreateVariableInternalServerError() *CreateVariableInternalServerError {
	return &CreateVariableInternalServerError{}
}

/*
CreateVariableInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateVariableInternalServerError struct {
	Payload interface{}
}

func (o *CreateVariableInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/variables][%d] createVariableInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVariableInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateVariableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
