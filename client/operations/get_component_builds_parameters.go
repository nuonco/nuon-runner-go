// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetComponentBuildsParams creates a new GetComponentBuildsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComponentBuildsParams() *GetComponentBuildsParams {
	return &GetComponentBuildsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComponentBuildsParamsWithTimeout creates a new GetComponentBuildsParams object
// with the ability to set a timeout on a request.
func NewGetComponentBuildsParamsWithTimeout(timeout time.Duration) *GetComponentBuildsParams {
	return &GetComponentBuildsParams{
		timeout: timeout,
	}
}

// NewGetComponentBuildsParamsWithContext creates a new GetComponentBuildsParams object
// with the ability to set a context for a request.
func NewGetComponentBuildsParamsWithContext(ctx context.Context) *GetComponentBuildsParams {
	return &GetComponentBuildsParams{
		Context: ctx,
	}
}

// NewGetComponentBuildsParamsWithHTTPClient creates a new GetComponentBuildsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComponentBuildsParamsWithHTTPClient(client *http.Client) *GetComponentBuildsParams {
	return &GetComponentBuildsParams{
		HTTPClient: client,
	}
}

/*
GetComponentBuildsParams contains all the parameters to send to the API endpoint

	for the get component builds operation.

	Typically these are written to a http.Request.
*/
type GetComponentBuildsParams struct {

	/* AppID.

	   app id to filter by
	*/
	AppID *string

	/* ComponentID.

	   component id to filter by
	*/
	ComponentID *string

	/* Limit.

	   limit of builds to return

	   Default: 60
	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get component builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComponentBuildsParams) WithDefaults() *GetComponentBuildsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get component builds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComponentBuildsParams) SetDefaults() {
	var (
		limitDefault = int64(60)
	)

	val := GetComponentBuildsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get component builds params
func (o *GetComponentBuildsParams) WithTimeout(timeout time.Duration) *GetComponentBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get component builds params
func (o *GetComponentBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get component builds params
func (o *GetComponentBuildsParams) WithContext(ctx context.Context) *GetComponentBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get component builds params
func (o *GetComponentBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get component builds params
func (o *GetComponentBuildsParams) WithHTTPClient(client *http.Client) *GetComponentBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get component builds params
func (o *GetComponentBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get component builds params
func (o *GetComponentBuildsParams) WithAppID(appID *string) *GetComponentBuildsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get component builds params
func (o *GetComponentBuildsParams) SetAppID(appID *string) {
	o.AppID = appID
}

// WithComponentID adds the componentID to the get component builds params
func (o *GetComponentBuildsParams) WithComponentID(componentID *string) *GetComponentBuildsParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the get component builds params
func (o *GetComponentBuildsParams) SetComponentID(componentID *string) {
	o.ComponentID = componentID
}

// WithLimit adds the limit to the get component builds params
func (o *GetComponentBuildsParams) WithLimit(limit *int64) *GetComponentBuildsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get component builds params
func (o *GetComponentBuildsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *GetComponentBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppID != nil {

		// query param app_id
		var qrAppID string

		if o.AppID != nil {
			qrAppID = *o.AppID
		}
		qAppID := qrAppID
		if qAppID != "" {

			if err := r.SetQueryParam("app_id", qAppID); err != nil {
				return err
			}
		}
	}

	if o.ComponentID != nil {

		// query param component_id
		var qrComponentID string

		if o.ComponentID != nil {
			qrComponentID = *o.ComponentID
		}
		qComponentID := qrComponentID
		if qComponentID != "" {

			if err := r.SetQueryParam("component_id", qComponentID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}