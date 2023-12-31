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

// GetReleaseRunByIDReader is a Reader for the GetReleaseRunByID structure.
type GetReleaseRunByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseRunByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseRunByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReleaseRunByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReleaseRunByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReleaseRunByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReleaseRunByIDOK creates a GetReleaseRunByIDOK with default headers values
func NewGetReleaseRunByIDOK() *GetReleaseRunByIDOK {
	return &GetReleaseRunByIDOK{}
}

/*
GetReleaseRunByIDOK describes a response with status code 200, with default header values.

OK
*/
type GetReleaseRunByIDOK struct {

	/* total matching count
	 */
	XTotalCount int64

	Payload *models.ReleaseRun
}

func (o *GetReleaseRunByIDOK) Error() string {
	return fmt.Sprintf("[GET /admin/releaserun/{id}][%d] getReleaseRunByIdOK  %+v", 200, o.Payload)
}
func (o *GetReleaseRunByIDOK) GetPayload() *models.ReleaseRun {
	return o.Payload
}

func (o *GetReleaseRunByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReleaseRunByIDBadRequest creates a GetReleaseRunByIDBadRequest with default headers values
func NewGetReleaseRunByIDBadRequest() *GetReleaseRunByIDBadRequest {
	return &GetReleaseRunByIDBadRequest{}
}

/*
GetReleaseRunByIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetReleaseRunByIDBadRequest struct {
	Payload interface{}
}

func (o *GetReleaseRunByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/releaserun/{id}][%d] getReleaseRunByIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetReleaseRunByIDBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *GetReleaseRunByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseRunByIDNotFound creates a GetReleaseRunByIDNotFound with default headers values
func NewGetReleaseRunByIDNotFound() *GetReleaseRunByIDNotFound {
	return &GetReleaseRunByIDNotFound{}
}

/*
GetReleaseRunByIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetReleaseRunByIDNotFound struct {
	Payload interface{}
}

func (o *GetReleaseRunByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/releaserun/{id}][%d] getReleaseRunByIdNotFound  %+v", 404, o.Payload)
}
func (o *GetReleaseRunByIDNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetReleaseRunByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseRunByIDInternalServerError creates a GetReleaseRunByIDInternalServerError with default headers values
func NewGetReleaseRunByIDInternalServerError() *GetReleaseRunByIDInternalServerError {
	return &GetReleaseRunByIDInternalServerError{}
}

/*
GetReleaseRunByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetReleaseRunByIDInternalServerError struct {
	Payload interface{}
}

func (o *GetReleaseRunByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/releaserun/{id}][%d] getReleaseRunByIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReleaseRunByIDInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *GetReleaseRunByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
