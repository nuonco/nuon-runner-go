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

// NewCreateJobComponentConfigParams creates a new CreateJobComponentConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateJobComponentConfigParams() *CreateJobComponentConfigParams {
	return &CreateJobComponentConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateJobComponentConfigParamsWithTimeout creates a new CreateJobComponentConfigParams object
// with the ability to set a timeout on a request.
func NewCreateJobComponentConfigParamsWithTimeout(timeout time.Duration) *CreateJobComponentConfigParams {
	return &CreateJobComponentConfigParams{
		timeout: timeout,
	}
}

// NewCreateJobComponentConfigParamsWithContext creates a new CreateJobComponentConfigParams object
// with the ability to set a context for a request.
func NewCreateJobComponentConfigParamsWithContext(ctx context.Context) *CreateJobComponentConfigParams {
	return &CreateJobComponentConfigParams{
		Context: ctx,
	}
}

// NewCreateJobComponentConfigParamsWithHTTPClient creates a new CreateJobComponentConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateJobComponentConfigParamsWithHTTPClient(client *http.Client) *CreateJobComponentConfigParams {
	return &CreateJobComponentConfigParams{
		HTTPClient: client,
	}
}

/*
CreateJobComponentConfigParams contains all the parameters to send to the API endpoint

	for the create job component config operation.

	Typically these are written to a http.Request.
*/
type CreateJobComponentConfigParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateJobComponentConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create job component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateJobComponentConfigParams) WithDefaults() *CreateJobComponentConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create job component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateJobComponentConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create job component config params
func (o *CreateJobComponentConfigParams) WithTimeout(timeout time.Duration) *CreateJobComponentConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create job component config params
func (o *CreateJobComponentConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create job component config params
func (o *CreateJobComponentConfigParams) WithContext(ctx context.Context) *CreateJobComponentConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create job component config params
func (o *CreateJobComponentConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create job component config params
func (o *CreateJobComponentConfigParams) WithHTTPClient(client *http.Client) *CreateJobComponentConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create job component config params
func (o *CreateJobComponentConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the create job component config params
func (o *CreateJobComponentConfigParams) WithComponentID(componentID string) *CreateJobComponentConfigParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the create job component config params
func (o *CreateJobComponentConfigParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithReq adds the req to the create job component config params
func (o *CreateJobComponentConfigParams) WithReq(req *models.ServiceCreateJobComponentConfigRequest) *CreateJobComponentConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create job component config params
func (o *CreateJobComponentConfigParams) SetReq(req *models.ServiceCreateJobComponentConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateJobComponentConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
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
