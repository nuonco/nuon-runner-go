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

	"github.com/nuonco/nuon-go/models"
)

// NewCreateAppConfigParams creates a new CreateAppConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAppConfigParams() *CreateAppConfigParams {
	return &CreateAppConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAppConfigParamsWithTimeout creates a new CreateAppConfigParams object
// with the ability to set a timeout on a request.
func NewCreateAppConfigParamsWithTimeout(timeout time.Duration) *CreateAppConfigParams {
	return &CreateAppConfigParams{
		timeout: timeout,
	}
}

// NewCreateAppConfigParamsWithContext creates a new CreateAppConfigParams object
// with the ability to set a context for a request.
func NewCreateAppConfigParamsWithContext(ctx context.Context) *CreateAppConfigParams {
	return &CreateAppConfigParams{
		Context: ctx,
	}
}

// NewCreateAppConfigParamsWithHTTPClient creates a new CreateAppConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAppConfigParamsWithHTTPClient(client *http.Client) *CreateAppConfigParams {
	return &CreateAppConfigParams{
		HTTPClient: client,
	}
}

/*
CreateAppConfigParams contains all the parameters to send to the API endpoint

	for the create app config operation.

	Typically these are written to a http.Request.
*/
type CreateAppConfigParams struct {

	/* AppID.

	   app ID
	*/
	AppID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateAppConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create app config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppConfigParams) WithDefaults() *CreateAppConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create app config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAppConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create app config params
func (o *CreateAppConfigParams) WithTimeout(timeout time.Duration) *CreateAppConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create app config params
func (o *CreateAppConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create app config params
func (o *CreateAppConfigParams) WithContext(ctx context.Context) *CreateAppConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create app config params
func (o *CreateAppConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create app config params
func (o *CreateAppConfigParams) WithHTTPClient(client *http.Client) *CreateAppConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create app config params
func (o *CreateAppConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the create app config params
func (o *CreateAppConfigParams) WithAppID(appID string) *CreateAppConfigParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the create app config params
func (o *CreateAppConfigParams) SetAppID(appID string) {
	o.AppID = appID
}

// WithReq adds the req to the create app config params
func (o *CreateAppConfigParams) WithReq(req *models.ServiceCreateAppConfigRequest) *CreateAppConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create app config params
func (o *CreateAppConfigParams) SetReq(req *models.ServiceCreateAppConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAppConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", o.AppID); err != nil {
		return err
	}
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}