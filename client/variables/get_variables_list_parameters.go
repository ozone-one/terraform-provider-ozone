// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetVariablesListParams creates a new GetVariablesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVariablesListParams() *GetVariablesListParams {
	return &GetVariablesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVariablesListParamsWithTimeout creates a new GetVariablesListParams object
// with the ability to set a timeout on a request.
func NewGetVariablesListParamsWithTimeout(timeout time.Duration) *GetVariablesListParams {
	return &GetVariablesListParams{
		timeout: timeout,
	}
}

// NewGetVariablesListParamsWithContext creates a new GetVariablesListParams object
// with the ability to set a context for a request.
func NewGetVariablesListParamsWithContext(ctx context.Context) *GetVariablesListParams {
	return &GetVariablesListParams{
		Context: ctx,
	}
}

// NewGetVariablesListParamsWithHTTPClient creates a new GetVariablesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVariablesListParamsWithHTTPClient(client *http.Client) *GetVariablesListParams {
	return &GetVariablesListParams{
		HTTPClient: client,
	}
}

/*
GetVariablesListParams contains all the parameters to send to the API endpoint

	for the get variables list operation.

	Typically these are written to a http.Request.
*/
type GetVariablesListParams struct {

	/* XWorkspaceID.

	   workspace id
	*/
	XWorkspaceID string

	/* End.

	   end
	*/
	End *int64

	/* Order.

	   sort order
	*/
	Order *string

	/* Sort.

	   sort column
	*/
	Sort *string

	/* Start.

	   start
	*/
	Start *int64

	/* CreatedBy.

	   created by id. Usage ?created_by=507f1f77bcf86cd799439011
	*/
	CreatedBy []string

	/* From.

	   from : filter for updated_at. expected format: RFC3339
	*/
	From *string

	/* Global.

	   variable type. Usage true/false
	*/
	Global *bool

	/* ID.

	   variable ids. Usage ?id=507f1f77bcf86cd799439011&id=507f1f77bcf86cd799439012
	*/
	ID []string

	/* Name.

	   strings matching name of variables
	*/
	Name *string

	/* NameSlug.

	   find by variable name
	*/
	NameSlug *string

	/* To.

	   to : filter for updated_at. expected format: RFC3339
	*/
	To *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get variables list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariablesListParams) WithDefaults() *GetVariablesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get variables list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVariablesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get variables list params
func (o *GetVariablesListParams) WithTimeout(timeout time.Duration) *GetVariablesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get variables list params
func (o *GetVariablesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get variables list params
func (o *GetVariablesListParams) WithContext(ctx context.Context) *GetVariablesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get variables list params
func (o *GetVariablesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get variables list params
func (o *GetVariablesListParams) WithHTTPClient(client *http.Client) *GetVariablesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get variables list params
func (o *GetVariablesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXWorkspaceID adds the xWorkspaceID to the get variables list params
func (o *GetVariablesListParams) WithXWorkspaceID(xWorkspaceID string) *GetVariablesListParams {
	o.SetXWorkspaceID(xWorkspaceID)
	return o
}

// SetXWorkspaceID adds the xWorkspaceId to the get variables list params
func (o *GetVariablesListParams) SetXWorkspaceID(xWorkspaceID string) {
	o.XWorkspaceID = xWorkspaceID
}

// WithEnd adds the end to the get variables list params
func (o *GetVariablesListParams) WithEnd(end *int64) *GetVariablesListParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get variables list params
func (o *GetVariablesListParams) SetEnd(end *int64) {
	o.End = end
}

// WithOrder adds the order to the get variables list params
func (o *GetVariablesListParams) WithOrder(order *string) *GetVariablesListParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the get variables list params
func (o *GetVariablesListParams) SetOrder(order *string) {
	o.Order = order
}

// WithSort adds the sort to the get variables list params
func (o *GetVariablesListParams) WithSort(sort *string) *GetVariablesListParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get variables list params
func (o *GetVariablesListParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStart adds the start to the get variables list params
func (o *GetVariablesListParams) WithStart(start *int64) *GetVariablesListParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get variables list params
func (o *GetVariablesListParams) SetStart(start *int64) {
	o.Start = start
}

// WithCreatedBy adds the createdBy to the get variables list params
func (o *GetVariablesListParams) WithCreatedBy(createdBy []string) *GetVariablesListParams {
	o.SetCreatedBy(createdBy)
	return o
}

// SetCreatedBy adds the createdBy to the get variables list params
func (o *GetVariablesListParams) SetCreatedBy(createdBy []string) {
	o.CreatedBy = createdBy
}

// WithFrom adds the from to the get variables list params
func (o *GetVariablesListParams) WithFrom(from *string) *GetVariablesListParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get variables list params
func (o *GetVariablesListParams) SetFrom(from *string) {
	o.From = from
}

// WithGlobal adds the global to the get variables list params
func (o *GetVariablesListParams) WithGlobal(global *bool) *GetVariablesListParams {
	o.SetGlobal(global)
	return o
}

// SetGlobal adds the global to the get variables list params
func (o *GetVariablesListParams) SetGlobal(global *bool) {
	o.Global = global
}

// WithID adds the id to the get variables list params
func (o *GetVariablesListParams) WithID(id []string) *GetVariablesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get variables list params
func (o *GetVariablesListParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get variables list params
func (o *GetVariablesListParams) WithName(name *string) *GetVariablesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get variables list params
func (o *GetVariablesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameSlug adds the nameSlug to the get variables list params
func (o *GetVariablesListParams) WithNameSlug(nameSlug *string) *GetVariablesListParams {
	o.SetNameSlug(nameSlug)
	return o
}

// SetNameSlug adds the nameSlug to the get variables list params
func (o *GetVariablesListParams) SetNameSlug(nameSlug *string) {
	o.NameSlug = nameSlug
}

// WithTo adds the to to the get variables list params
func (o *GetVariablesListParams) WithTo(to *string) *GetVariablesListParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get variables list params
func (o *GetVariablesListParams) SetTo(to *string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetVariablesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Workspace-ID
	if err := r.SetHeaderParam("X-Workspace-ID", o.XWorkspaceID); err != nil {
		return err
	}

	if o.End != nil {

		// query param _end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("_end", qEnd); err != nil {
				return err
			}
		}
	}

	if o.Order != nil {

		// query param _order
		var qrOrder string

		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {

			if err := r.SetQueryParam("_order", qOrder); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param _sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("_sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Start != nil {

		// query param _start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("_start", qStart); err != nil {
				return err
			}
		}
	}

	if o.CreatedBy != nil {

		// binding items for created_by
		joinedCreatedBy := o.bindParamCreatedBy(reg)

		// query array param created_by
		if err := r.SetQueryParam("created_by", joinedCreatedBy...); err != nil {
			return err
		}
	}

	if o.From != nil {

		// query param from
		var qrFrom string

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := qrFrom
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	if o.Global != nil {

		// query param global
		var qrGlobal bool

		if o.Global != nil {
			qrGlobal = *o.Global
		}
		qGlobal := swag.FormatBool(qrGlobal)
		if qGlobal != "" {

			if err := r.SetQueryParam("global", qGlobal); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameSlug != nil {

		// query param name_slug
		var qrNameSlug string

		if o.NameSlug != nil {
			qrNameSlug = *o.NameSlug
		}
		qNameSlug := qrNameSlug
		if qNameSlug != "" {

			if err := r.SetQueryParam("name_slug", qNameSlug); err != nil {
				return err
			}
		}
	}

	if o.To != nil {

		// query param to
		var qrTo string

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := qrTo
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetVariablesList binds the parameter created_by
func (o *GetVariablesListParams) bindParamCreatedBy(formats strfmt.Registry) []string {
	createdByIR := o.CreatedBy

	var createdByIC []string
	for _, createdByIIR := range createdByIR { // explode []string

		createdByIIV := createdByIIR // string as string
		createdByIC = append(createdByIC, createdByIIV)
	}

	// items.CollectionFormat: "multi"
	createdByIS := swag.JoinByFormat(createdByIC, "multi")

	return createdByIS
}

// bindParamGetVariablesList binds the parameter id
func (o *GetVariablesListParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
