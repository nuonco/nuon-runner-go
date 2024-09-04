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

	"github.com/nuonco/nuon-runner-go/models"
)

// NewCreateRunnerJobExecutionParams creates a new CreateRunnerJobExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRunnerJobExecutionParams() *CreateRunnerJobExecutionParams {
	return &CreateRunnerJobExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunnerJobExecutionParamsWithTimeout creates a new CreateRunnerJobExecutionParams object
// with the ability to set a timeout on a request.
func NewCreateRunnerJobExecutionParamsWithTimeout(timeout time.Duration) *CreateRunnerJobExecutionParams {
	return &CreateRunnerJobExecutionParams{
		timeout: timeout,
	}
}

// NewCreateRunnerJobExecutionParamsWithContext creates a new CreateRunnerJobExecutionParams object
// with the ability to set a context for a request.
func NewCreateRunnerJobExecutionParamsWithContext(ctx context.Context) *CreateRunnerJobExecutionParams {
	return &CreateRunnerJobExecutionParams{
		Context: ctx,
	}
}

// NewCreateRunnerJobExecutionParamsWithHTTPClient creates a new CreateRunnerJobExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRunnerJobExecutionParamsWithHTTPClient(client *http.Client) *CreateRunnerJobExecutionParams {
	return &CreateRunnerJobExecutionParams{
		HTTPClient: client,
	}
}

/*
CreateRunnerJobExecutionParams contains all the parameters to send to the API endpoint

	for the create runner job execution operation.

	Typically these are written to a http.Request.
*/
type CreateRunnerJobExecutionParams struct {

	/* Req.

	   Input
	*/
	Req models.ServiceCreateRunnerJobExecutionRequest

	/* RunnerJobID.

	   runner job ID
	*/
	RunnerJobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create runner job execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunnerJobExecutionParams) WithDefaults() *CreateRunnerJobExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create runner job execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunnerJobExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) WithTimeout(timeout time.Duration) *CreateRunnerJobExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) WithContext(ctx context.Context) *CreateRunnerJobExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) WithHTTPClient(client *http.Client) *CreateRunnerJobExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) WithReq(req models.ServiceCreateRunnerJobExecutionRequest) *CreateRunnerJobExecutionParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) SetReq(req models.ServiceCreateRunnerJobExecutionRequest) {
	o.Req = req
}

// WithRunnerJobID adds the runnerJobID to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) WithRunnerJobID(runnerJobID string) *CreateRunnerJobExecutionParams {
	o.SetRunnerJobID(runnerJobID)
	return o
}

// SetRunnerJobID adds the runnerJobId to the create runner job execution params
func (o *CreateRunnerJobExecutionParams) SetRunnerJobID(runnerJobID string) {
	o.RunnerJobID = runnerJobID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunnerJobExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	// path param runner_job_id
	if err := r.SetPathParam("runner_job_id", o.RunnerJobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
