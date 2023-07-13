// Code generated by go-swagger; DO NOT EDIT.

package pipeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// PipelineGetByIDReader is a Reader for the PipelineGetByID structure.
type PipelineGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PipelineGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPipelineGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPipelineGetByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPipelineGetByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPipelineGetByIDOK creates a PipelineGetByIDOK with default headers values
func NewPipelineGetByIDOK() *PipelineGetByIDOK {
	return &PipelineGetByIDOK{}
}

/*
PipelineGetByIDOK describes a response with status code 200, with default header values.

OK
*/
type PipelineGetByIDOK struct {
	Payload *models.Pipeline
}

func (o *PipelineGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /admin/pipeline/{id}][%d] pipelineGetByIdOK  %+v", 200, o.Payload)
}
func (o *PipelineGetByIDOK) GetPayload() *models.Pipeline {
	return o.Payload
}

func (o *PipelineGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineGetByIDBadRequest creates a PipelineGetByIDBadRequest with default headers values
func NewPipelineGetByIDBadRequest() *PipelineGetByIDBadRequest {
	return &PipelineGetByIDBadRequest{}
}

/*
PipelineGetByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PipelineGetByIDBadRequest struct {
	Payload interface{}
}

func (o *PipelineGetByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/pipeline/{id}][%d] pipelineGetByIdBadRequest  %+v", 400, o.Payload)
}
func (o *PipelineGetByIDBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PipelineGetByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineGetByIDInternalServerError creates a PipelineGetByIDInternalServerError with default headers values
func NewPipelineGetByIDInternalServerError() *PipelineGetByIDInternalServerError {
	return &PipelineGetByIDInternalServerError{}
}

/*
PipelineGetByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PipelineGetByIDInternalServerError struct {
	Payload interface{}
}

func (o *PipelineGetByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/pipeline/{id}][%d] pipelineGetByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *PipelineGetByIDInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PipelineGetByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
