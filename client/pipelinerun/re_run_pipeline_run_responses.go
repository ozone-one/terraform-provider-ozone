// Code generated by go-swagger; DO NOT EDIT.

package pipelinerun

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ozone-one/terraform-provider-ozone/models"
)

// ReRunPipelineRunReader is a Reader for the ReRunPipelineRun structure.
type ReRunPipelineRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReRunPipelineRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReRunPipelineRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReRunPipelineRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReRunPipelineRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReRunPipelineRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReRunPipelineRunOK creates a ReRunPipelineRunOK with default headers values
func NewReRunPipelineRunOK() *ReRunPipelineRunOK {
	return &ReRunPipelineRunOK{}
}

/*
ReRunPipelineRunOK describes a response with status code 200, with default header values.

OK
*/
type ReRunPipelineRunOK struct {
	Payload *models.PipelineRun
}

func (o *ReRunPipelineRunOK) Error() string {
	return fmt.Sprintf("[POST /admin/pipelinerun/{pipeline_run_id}/rerun][%d] reRunPipelineRunOK  %+v", 200, o.Payload)
}
func (o *ReRunPipelineRunOK) GetPayload() *models.PipelineRun {
	return o.Payload
}

func (o *ReRunPipelineRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReRunPipelineRunBadRequest creates a ReRunPipelineRunBadRequest with default headers values
func NewReRunPipelineRunBadRequest() *ReRunPipelineRunBadRequest {
	return &ReRunPipelineRunBadRequest{}
}

/*
ReRunPipelineRunBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ReRunPipelineRunBadRequest struct {
	Payload interface{}
}

func (o *ReRunPipelineRunBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/pipelinerun/{pipeline_run_id}/rerun][%d] reRunPipelineRunBadRequest  %+v", 400, o.Payload)
}
func (o *ReRunPipelineRunBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ReRunPipelineRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReRunPipelineRunNotFound creates a ReRunPipelineRunNotFound with default headers values
func NewReRunPipelineRunNotFound() *ReRunPipelineRunNotFound {
	return &ReRunPipelineRunNotFound{}
}

/*
ReRunPipelineRunNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ReRunPipelineRunNotFound struct {
	Payload interface{}
}

func (o *ReRunPipelineRunNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/pipelinerun/{pipeline_run_id}/rerun][%d] reRunPipelineRunNotFound  %+v", 404, o.Payload)
}
func (o *ReRunPipelineRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ReRunPipelineRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReRunPipelineRunInternalServerError creates a ReRunPipelineRunInternalServerError with default headers values
func NewReRunPipelineRunInternalServerError() *ReRunPipelineRunInternalServerError {
	return &ReRunPipelineRunInternalServerError{}
}

/*
ReRunPipelineRunInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReRunPipelineRunInternalServerError struct {
	Payload interface{}
}

func (o *ReRunPipelineRunInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/pipelinerun/{pipeline_run_id}/rerun][%d] reRunPipelineRunInternalServerError  %+v", 500, o.Payload)
}
func (o *ReRunPipelineRunInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ReRunPipelineRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
