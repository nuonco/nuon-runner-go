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
)

// NewGetAppConfigsParams creates a new GetAppConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppConfigsParams() *GetAppConfigsParams {
	return &GetAppConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppConfigsParamsWithTimeout creates a new GetAppConfigsParams object
// with the ability to set a timeout on a request.
func NewGetAppConfigsParamsWithTimeout(timeout time.Duration) *GetAppConfigsParams {
	return &GetAppConfigsParams{
		timeout: timeout,
	}
}

// NewGetAppConfigsParamsWithContext creates a new GetAppConfigsParams object
// with the ability to set a context for a request.
func NewGetAppConfigsParamsWithContext(ctx context.Context) *GetAppConfigsParams {
	return &GetAppConfigsParams{
		Context: ctx,
	}
}

// NewGetAppConfigsParamsWithHTTPClient creates a new GetAppConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppConfigsParamsWithHTTPClient(client *http.Client) *GetAppConfigsParams {
	return &GetAppConfigsParams{
		HTTPClient: client,
	}
}

/*
GetAppConfigsParams contains all the parameters to send to the API endpoint

	for the get app configs operation.

	Typically these are written to a http.Request.
*/
type GetAppConfigsParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppConfigsParams) WithDefaults() *GetAppConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app configs params
func (o *GetAppConfigsParams) WithTimeout(timeout time.Duration) *GetAppConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app configs params
func (o *GetAppConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app configs params
func (o *GetAppConfigsParams) WithContext(ctx context.Context) *GetAppConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app configs params
func (o *GetAppConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app configs params
func (o *GetAppConfigsParams) WithHTTPClient(client *http.Client) *GetAppConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app configs params
func (o *GetAppConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the get app configs params
func (o *GetAppConfigsParams) WithAppID(appID string) *GetAppConfigsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the get app configs params
func (o *GetAppConfigsParams) SetAppID(appID string) {
	o.AppID = appID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
