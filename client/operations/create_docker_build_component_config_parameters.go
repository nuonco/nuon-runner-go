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

// NewCreateDockerBuildComponentConfigParams creates a new CreateDockerBuildComponentConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDockerBuildComponentConfigParams() *CreateDockerBuildComponentConfigParams {
	return &CreateDockerBuildComponentConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDockerBuildComponentConfigParamsWithTimeout creates a new CreateDockerBuildComponentConfigParams object
// with the ability to set a timeout on a request.
func NewCreateDockerBuildComponentConfigParamsWithTimeout(timeout time.Duration) *CreateDockerBuildComponentConfigParams {
	return &CreateDockerBuildComponentConfigParams{
		timeout: timeout,
	}
}

// NewCreateDockerBuildComponentConfigParamsWithContext creates a new CreateDockerBuildComponentConfigParams object
// with the ability to set a context for a request.
func NewCreateDockerBuildComponentConfigParamsWithContext(ctx context.Context) *CreateDockerBuildComponentConfigParams {
	return &CreateDockerBuildComponentConfigParams{
		Context: ctx,
	}
}

// NewCreateDockerBuildComponentConfigParamsWithHTTPClient creates a new CreateDockerBuildComponentConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDockerBuildComponentConfigParamsWithHTTPClient(client *http.Client) *CreateDockerBuildComponentConfigParams {
	return &CreateDockerBuildComponentConfigParams{
		HTTPClient: client,
	}
}

/*
CreateDockerBuildComponentConfigParams contains all the parameters to send to the API endpoint

	for the create docker build component config operation.

	Typically these are written to a http.Request.
*/
type CreateDockerBuildComponentConfigParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateDockerBuildComponentConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create docker build component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDockerBuildComponentConfigParams) WithDefaults() *CreateDockerBuildComponentConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create docker build component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDockerBuildComponentConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) WithTimeout(timeout time.Duration) *CreateDockerBuildComponentConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) WithContext(ctx context.Context) *CreateDockerBuildComponentConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) WithHTTPClient(client *http.Client) *CreateDockerBuildComponentConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) WithComponentID(componentID string) *CreateDockerBuildComponentConfigParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithReq adds the req to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) WithReq(req *models.ServiceCreateDockerBuildComponentConfigRequest) *CreateDockerBuildComponentConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create docker build component config params
func (o *CreateDockerBuildComponentConfigParams) SetReq(req *models.ServiceCreateDockerBuildComponentConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDockerBuildComponentConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
