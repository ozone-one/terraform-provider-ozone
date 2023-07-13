// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// EnvironmentListReader is a Reader for the EnvironmentList structure.
type EnvironmentListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnvironmentListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnvironmentListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnvironmentListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnvironmentListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnvironmentListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnvironmentListOK creates a EnvironmentListOK with default headers values
func NewEnvironmentListOK() *EnvironmentListOK {
	return &EnvironmentListOK{}
}

/*
EnvironmentListOK describes a response with status code 200, with default header values.

OK
*/
type EnvironmentListOK struct {

	/* total matching count
	 */
	XTotalCount int64

	Payload *models.EnvironmentList
}

func (o *EnvironmentListOK) Error() string {
	return fmt.Sprintf("[GET /admin/environment][%d] environmentListOK  %+v", 200, o.Payload)
}
func (o *EnvironmentListOK) GetPayload() *models.EnvironmentList {
	return o.Payload
}

func (o *EnvironmentListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	o.Payload = new(models.EnvironmentList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnvironmentListBadRequest creates a EnvironmentListBadRequest with default headers values
func NewEnvironmentListBadRequest() *EnvironmentListBadRequest {
	return &EnvironmentListBadRequest{}
}

/*
EnvironmentListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type EnvironmentListBadRequest struct {
	Payload interface{}
}

func (o *EnvironmentListBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/environment][%d] environmentListBadRequest  %+v", 400, o.Payload)
}
func (o *EnvironmentListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *EnvironmentListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnvironmentListNotFound creates a EnvironmentListNotFound with default headers values
func NewEnvironmentListNotFound() *EnvironmentListNotFound {
	return &EnvironmentListNotFound{}
}

/*
EnvironmentListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type EnvironmentListNotFound struct {
	Payload interface{}
}

func (o *EnvironmentListNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/environment][%d] environmentListNotFound  %+v", 404, o.Payload)
}
func (o *EnvironmentListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *EnvironmentListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnvironmentListInternalServerError creates a EnvironmentListInternalServerError with default headers values
func NewEnvironmentListInternalServerError() *EnvironmentListInternalServerError {
	return &EnvironmentListInternalServerError{}
}

/*
EnvironmentListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type EnvironmentListInternalServerError struct {
	Payload interface{}
}

func (o *EnvironmentListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/environment][%d] environmentListInternalServerError  %+v", 500, o.Payload)
}
func (o *EnvironmentListInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *EnvironmentListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}