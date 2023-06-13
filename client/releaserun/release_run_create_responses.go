// Code generated by go-swagger; DO NOT EDIT.

package releaserun

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

// ReleaseRunCreateReader is a Reader for the ReleaseRunCreate structure.
type ReleaseRunCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReleaseRunCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReleaseRunCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReleaseRunCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReleaseRunCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReleaseRunCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReleaseRunCreateOK creates a ReleaseRunCreateOK with default headers values
func NewReleaseRunCreateOK() *ReleaseRunCreateOK {
	return &ReleaseRunCreateOK{}
}

/*
ReleaseRunCreateOK describes a response with status code 200, with default header values.

OK
*/
type ReleaseRunCreateOK struct {

	/* total matching count
	 */
	XTotalCount int64

	Payload *models.ReleaseRun
}

func (o *ReleaseRunCreateOK) Error() string {
	return fmt.Sprintf("[POST /admin/releaserun][%d] releaseRunCreateOK  %+v", 200, o.Payload)
}
func (o *ReleaseRunCreateOK) GetPayload() *models.ReleaseRun {
	return o.Payload
}

func (o *ReleaseRunCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	o.Payload = new(models.ReleaseRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseRunCreateBadRequest creates a ReleaseRunCreateBadRequest with default headers values
func NewReleaseRunCreateBadRequest() *ReleaseRunCreateBadRequest {
	return &ReleaseRunCreateBadRequest{}
}

/*
ReleaseRunCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ReleaseRunCreateBadRequest struct {
	Payload interface{}
}

func (o *ReleaseRunCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/releaserun][%d] releaseRunCreateBadRequest  %+v", 400, o.Payload)
}
func (o *ReleaseRunCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ReleaseRunCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseRunCreateNotFound creates a ReleaseRunCreateNotFound with default headers values
func NewReleaseRunCreateNotFound() *ReleaseRunCreateNotFound {
	return &ReleaseRunCreateNotFound{}
}

/*
ReleaseRunCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ReleaseRunCreateNotFound struct {
	Payload interface{}
}

func (o *ReleaseRunCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /admin/releaserun][%d] releaseRunCreateNotFound  %+v", 404, o.Payload)
}
func (o *ReleaseRunCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ReleaseRunCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseRunCreateInternalServerError creates a ReleaseRunCreateInternalServerError with default headers values
func NewReleaseRunCreateInternalServerError() *ReleaseRunCreateInternalServerError {
	return &ReleaseRunCreateInternalServerError{}
}

/*
ReleaseRunCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReleaseRunCreateInternalServerError struct {
	Payload interface{}
}

func (o *ReleaseRunCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/releaserun][%d] releaseRunCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *ReleaseRunCreateInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ReleaseRunCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
