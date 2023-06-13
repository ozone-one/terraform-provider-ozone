// Code generated by go-swagger; DO NOT EDIT.

package pipelinerun

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// PipelineRunGetReader is a Reader for the PipelineRunGet structure.
type PipelineRunGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PipelineRunGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPipelineRunGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPipelineRunGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPipelineRunGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPipelineRunGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPipelineRunGetOK creates a PipelineRunGetOK with default headers values
func NewPipelineRunGetOK() *PipelineRunGetOK {
	return &PipelineRunGetOK{}
}

/*
PipelineRunGetOK describes a response with status code 200, with default header values.

OK
*/
type PipelineRunGetOK struct {

	/* total matching count
	 */
	XTotalCount int64

	Payload *models.PipelineRunList
}

func (o *PipelineRunGetOK) Error() string {
	return fmt.Sprintf("[GET /admin/pipelinerun][%d] pipelineRunGetOK  %+v", 200, o.Payload)
}
func (o *PipelineRunGetOK) GetPayload() *models.PipelineRunList {
	return o.Payload
}

func (o *PipelineRunGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	o.Payload = new(models.PipelineRunList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineRunGetBadRequest creates a PipelineRunGetBadRequest with default headers values
func NewPipelineRunGetBadRequest() *PipelineRunGetBadRequest {
	return &PipelineRunGetBadRequest{}
}

/*
PipelineRunGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PipelineRunGetBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *PipelineRunGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/pipelinerun][%d] pipelineRunGetBadRequest  %+v", 400, o.Payload)
}
func (o *PipelineRunGetBadRequest) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PipelineRunGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineRunGetNotFound creates a PipelineRunGetNotFound with default headers values
func NewPipelineRunGetNotFound() *PipelineRunGetNotFound {
	return &PipelineRunGetNotFound{}
}

/*
PipelineRunGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PipelineRunGetNotFound struct {
	Payload *models.ErrorMessage
}

func (o *PipelineRunGetNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/pipelinerun][%d] pipelineRunGetNotFound  %+v", 404, o.Payload)
}
func (o *PipelineRunGetNotFound) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PipelineRunGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPipelineRunGetInternalServerError creates a PipelineRunGetInternalServerError with default headers values
func NewPipelineRunGetInternalServerError() *PipelineRunGetInternalServerError {
	return &PipelineRunGetInternalServerError{}
}

/*
PipelineRunGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PipelineRunGetInternalServerError struct {
	Payload *models.ErrorMessage
}

func (o *PipelineRunGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/pipelinerun][%d] pipelineRunGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PipelineRunGetInternalServerError) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PipelineRunGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}