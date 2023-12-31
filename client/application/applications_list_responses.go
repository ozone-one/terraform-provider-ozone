// Code generated by go-swagger; DO NOT EDIT.

package application

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

// ApplicationsListReader is a Reader for the ApplicationsList structure.
type ApplicationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewApplicationsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewApplicationsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewApplicationsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewApplicationsListOK creates a ApplicationsListOK with default headers values
func NewApplicationsListOK() *ApplicationsListOK {
	return &ApplicationsListOK{}
}

/*
ApplicationsListOK describes a response with status code 200, with default header values.

OK
*/
type ApplicationsListOK struct {

	/* total matching count
	 */
	XTotalCount int64

	Payload *models.ApplicationList
}

func (o *ApplicationsListOK) Error() string {
	return fmt.Sprintf("[GET /admin/app][%d] applicationsListOK  %+v", 200, o.Payload)
}
func (o *ApplicationsListOK) GetPayload() *models.ApplicationList {
	return o.Payload
}

func (o *ApplicationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	o.Payload = new(models.ApplicationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationsListBadRequest creates a ApplicationsListBadRequest with default headers values
func NewApplicationsListBadRequest() *ApplicationsListBadRequest {
	return &ApplicationsListBadRequest{}
}

/*
ApplicationsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ApplicationsListBadRequest struct {
	Payload interface{}
}

func (o *ApplicationsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/app][%d] applicationsListBadRequest  %+v", 400, o.Payload)
}
func (o *ApplicationsListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationsListNotFound creates a ApplicationsListNotFound with default headers values
func NewApplicationsListNotFound() *ApplicationsListNotFound {
	return &ApplicationsListNotFound{}
}

/*
ApplicationsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ApplicationsListNotFound struct {
	Payload interface{}
}

func (o *ApplicationsListNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/app][%d] applicationsListNotFound  %+v", 404, o.Payload)
}
func (o *ApplicationsListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationsListInternalServerError creates a ApplicationsListInternalServerError with default headers values
func NewApplicationsListInternalServerError() *ApplicationsListInternalServerError {
	return &ApplicationsListInternalServerError{}
}

/*
ApplicationsListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ApplicationsListInternalServerError struct {
	Payload interface{}
}

func (o *ApplicationsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/app][%d] applicationsListInternalServerError  %+v", 500, o.Payload)
}
func (o *ApplicationsListInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ApplicationsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
