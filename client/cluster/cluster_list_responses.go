// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ClusterListReader is a Reader for the ClusterList structure.
type ClusterListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewClusterListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClusterListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewClusterListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewClusterListOK creates a ClusterListOK with default headers values
func NewClusterListOK() *ClusterListOK {
	return &ClusterListOK{}
}

/*
ClusterListOK describes a response with status code 200, with default header values.

OK
*/
type ClusterListOK struct {

	/* total matching count
	 */
	XTotalCount int64

	Payload *models.ClusterList
}

func (o *ClusterListOK) Error() string {
	return fmt.Sprintf("[GET /admin/cluster][%d] clusterListOK  %+v", 200, o.Payload)
}
func (o *ClusterListOK) GetPayload() *models.ClusterList {
	return o.Payload
}

func (o *ClusterListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	o.Payload = new(models.ClusterList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterListBadRequest creates a ClusterListBadRequest with default headers values
func NewClusterListBadRequest() *ClusterListBadRequest {
	return &ClusterListBadRequest{}
}

/*
ClusterListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ClusterListBadRequest struct {
	Payload interface{}
}

func (o *ClusterListBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/cluster][%d] clusterListBadRequest  %+v", 400, o.Payload)
}
func (o *ClusterListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ClusterListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterListNotFound creates a ClusterListNotFound with default headers values
func NewClusterListNotFound() *ClusterListNotFound {
	return &ClusterListNotFound{}
}

/*
ClusterListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ClusterListNotFound struct {
	Payload interface{}
}

func (o *ClusterListNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/cluster][%d] clusterListNotFound  %+v", 404, o.Payload)
}
func (o *ClusterListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ClusterListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterListInternalServerError creates a ClusterListInternalServerError with default headers values
func NewClusterListInternalServerError() *ClusterListInternalServerError {
	return &ClusterListInternalServerError{}
}

/*
ClusterListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ClusterListInternalServerError struct {
	Payload interface{}
}

func (o *ClusterListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/cluster][%d] clusterListInternalServerError  %+v", 500, o.Payload)
}
func (o *ClusterListInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *ClusterListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
