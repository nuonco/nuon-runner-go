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

// NewCreateRunnerJobExecutionResultParams creates a new CreateRunnerJobExecutionResultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRunnerJobExecutionResultParams() *CreateRunnerJobExecutionResultParams {
	return &CreateRunnerJobExecutionResultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunnerJobExecutionResultParamsWithTimeout creates a new CreateRunnerJobExecutionResultParams object
// with the ability to set a timeout on a request.
func NewCreateRunnerJobExecutionResultParamsWithTimeout(timeout time.Duration) *CreateRunnerJobExecutionResultParams {
	return &CreateRunnerJobExecutionResultParams{
		timeout: timeout,
	}
}

// NewCreateRunnerJobExecutionResultParamsWithContext creates a new CreateRunnerJobExecutionResultParams object
// with the ability to set a context for a request.
func NewCreateRunnerJobExecutionResultParamsWithContext(ctx context.Context) *CreateRunnerJobExecutionResultParams {
	return &CreateRunnerJobExecutionResultParams{
		Context: ctx,
	}
}

// NewCreateRunnerJobExecutionResultParamsWithHTTPClient creates a new CreateRunnerJobExecutionResultParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRunnerJobExecutionResultParamsWithHTTPClient(client *http.Client) *CreateRunnerJobExecutionResultParams {
	return &CreateRunnerJobExecutionResultParams{
		HTTPClient: client,
	}
}

/*
CreateRunnerJobExecutionResultParams contains all the parameters to send to the API endpoint

	for the create runner job execution result operation.

	Typically these are written to a http.Request.
*/
type CreateRunnerJobExecutionResultParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateRunnerJobExecutionResultRequest

	/* RunnerJobExecutionID.

	   runner job execution ID
	*/
	RunnerJobExecutionID string

	/* RunnerJobID.

	   runner job ID
	*/
	RunnerJobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create runner job execution result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunnerJobExecutionResultParams) WithDefaults() *CreateRunnerJobExecutionResultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create runner job execution result params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunnerJobExecutionResultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) WithTimeout(timeout time.Duration) *CreateRunnerJobExecutionResultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) WithContext(ctx context.Context) *CreateRunnerJobExecutionResultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) WithHTTPClient(client *http.Client) *CreateRunnerJobExecutionResultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) WithReq(req *models.ServiceCreateRunnerJobExecutionResultRequest) *CreateRunnerJobExecutionResultParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) SetReq(req *models.ServiceCreateRunnerJobExecutionResultRequest) {
	o.Req = req
}

// WithRunnerJobExecutionID adds the runnerJobExecutionID to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) WithRunnerJobExecutionID(runnerJobExecutionID string) *CreateRunnerJobExecutionResultParams {
	o.SetRunnerJobExecutionID(runnerJobExecutionID)
	return o
}

// SetRunnerJobExecutionID adds the runnerJobExecutionId to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) SetRunnerJobExecutionID(runnerJobExecutionID string) {
	o.RunnerJobExecutionID = runnerJobExecutionID
}

// WithRunnerJobID adds the runnerJobID to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) WithRunnerJobID(runnerJobID string) *CreateRunnerJobExecutionResultParams {
	o.SetRunnerJobID(runnerJobID)
	return o
}

// SetRunnerJobID adds the runnerJobId to the create runner job execution result params
func (o *CreateRunnerJobExecutionResultParams) SetRunnerJobID(runnerJobID string) {
	o.RunnerJobID = runnerJobID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunnerJobExecutionResultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	// path param runner_job_execution_id
	if err := r.SetPathParam("runner_job_execution_id", o.RunnerJobExecutionID); err != nil {
		return err
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
