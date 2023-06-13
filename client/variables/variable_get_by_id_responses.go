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

// VariableGetByIDReader is a Reader for the VariableGetByID structure.
type VariableGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VariableGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVariableGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVariableGetByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVariableGetByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVariableGetByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVariableGetByIDOK creates a VariableGetByIDOK with default headers values
func NewVariableGetByIDOK() *VariableGetByIDOK {
	return &VariableGetByIDOK{}
}

/*
VariableGetByIDOK describes a response with status code 200, with default header values.

OK
*/
type VariableGetByIDOK struct {
	Payload []*models.VariableView
}

func (o *VariableGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /admin/variables/{id}][%d] variableGetByIdOK  %+v", 200, o.Payload)
}
func (o *VariableGetByIDOK) GetPayload() []*models.VariableView {
	return o.Payload
}

func (o *VariableGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVariableGetByIDBadRequest creates a VariableGetByIDBadRequest with default headers values
func NewVariableGetByIDBadRequest() *VariableGetByIDBadRequest {
	return &VariableGetByIDBadRequest{}
}

/*
VariableGetByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VariableGetByIDBadRequest struct {
	Payload interface{}
}

func (o *VariableGetByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/variables/{id}][%d] variableGetByIdBadRequest  %+v", 400, o.Payload)
}
func (o *VariableGetByIDBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *VariableGetByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVariableGetByIDNotFound creates a VariableGetByIDNotFound with default headers values
func NewVariableGetByIDNotFound() *VariableGetByIDNotFound {
	return &VariableGetByIDNotFound{}
}

/*
VariableGetByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type VariableGetByIDNotFound struct {
	Payload interface{}
}

func (o *VariableGetByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/variables/{id}][%d] variableGetByIdNotFound  %+v", 404, o.Payload)
}
func (o *VariableGetByIDNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *VariableGetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVariableGetByIDInternalServerError creates a VariableGetByIDInternalServerError with default headers values
func NewVariableGetByIDInternalServerError() *VariableGetByIDInternalServerError {
	return &VariableGetByIDInternalServerError{}
}

/*
VariableGetByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type VariableGetByIDInternalServerError struct {
	Payload interface{}
}

func (o *VariableGetByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/variables/{id}][%d] variableGetByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *VariableGetByIDInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *VariableGetByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
