// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// RepoListReader is a Reader for the RepoList structure.
type RepoListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepoListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRepoListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRepoListOK creates a RepoListOK with default headers values
func NewRepoListOK() *RepoListOK {
	return &RepoListOK{}
}

/*
RepoListOK describes a response with status code 200, with default header values.

OK
*/
type RepoListOK struct {

	/* total matching count
	 */
	XTotalCount int64

	Payload *models.RepositoryList
}

func (o *RepoListOK) Error() string {
	return fmt.Sprintf("[GET /admin/repository][%d] repoListOK  %+v", 200, o.Payload)
}
func (o *RepoListOK) GetPayload() *models.RepositoryList {
	return o.Payload
}

func (o *RepoListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	o.Payload = new(models.RepositoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListBadRequest creates a RepoListBadRequest with default headers values
func NewRepoListBadRequest() *RepoListBadRequest {
	return &RepoListBadRequest{}
}

/*
RepoListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RepoListBadRequest struct {
	Payload interface{}
}

func (o *RepoListBadRequest) Error() string {
	return fmt.Sprintf("[GET /admin/repository][%d] repoListBadRequest  %+v", 400, o.Payload)
}
func (o *RepoListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *RepoListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListNotFound creates a RepoListNotFound with default headers values
func NewRepoListNotFound() *RepoListNotFound {
	return &RepoListNotFound{}
}

/*
RepoListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RepoListNotFound struct {
	Payload interface{}
}

func (o *RepoListNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/repository][%d] repoListNotFound  %+v", 404, o.Payload)
}
func (o *RepoListNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *RepoListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListInternalServerError creates a RepoListInternalServerError with default headers values
func NewRepoListInternalServerError() *RepoListInternalServerError {
	return &RepoListInternalServerError{}
}

/*
RepoListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RepoListInternalServerError struct {
	Payload interface{}
}

func (o *RepoListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/repository][%d] repoListInternalServerError  %+v", 500, o.Payload)
}
func (o *RepoListInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *RepoListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
