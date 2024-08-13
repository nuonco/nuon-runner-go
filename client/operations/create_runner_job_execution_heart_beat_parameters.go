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

// NewCreateRunnerJobExecutionHeartBeatParams creates a new CreateRunnerJobExecutionHeartBeatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateRunnerJobExecutionHeartBeatParams() *CreateRunnerJobExecutionHeartBeatParams {
	return &CreateRunnerJobExecutionHeartBeatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRunnerJobExecutionHeartBeatParamsWithTimeout creates a new CreateRunnerJobExecutionHeartBeatParams object
// with the ability to set a timeout on a request.
func NewCreateRunnerJobExecutionHeartBeatParamsWithTimeout(timeout time.Duration) *CreateRunnerJobExecutionHeartBeatParams {
	return &CreateRunnerJobExecutionHeartBeatParams{
		timeout: timeout,
	}
}

// NewCreateRunnerJobExecutionHeartBeatParamsWithContext creates a new CreateRunnerJobExecutionHeartBeatParams object
// with the ability to set a context for a request.
func NewCreateRunnerJobExecutionHeartBeatParamsWithContext(ctx context.Context) *CreateRunnerJobExecutionHeartBeatParams {
	return &CreateRunnerJobExecutionHeartBeatParams{
		Context: ctx,
	}
}

// NewCreateRunnerJobExecutionHeartBeatParamsWithHTTPClient creates a new CreateRunnerJobExecutionHeartBeatParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateRunnerJobExecutionHeartBeatParamsWithHTTPClient(client *http.Client) *CreateRunnerJobExecutionHeartBeatParams {
	return &CreateRunnerJobExecutionHeartBeatParams{
		HTTPClient: client,
	}
}

/*
CreateRunnerJobExecutionHeartBeatParams contains all the parameters to send to the API endpoint

	for the create runner job execution heart beat operation.

	Typically these are written to a http.Request.
*/
type CreateRunnerJobExecutionHeartBeatParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateRunnerJobExecutionHeartBeatRequest

	/* RunnerJobExecutionID.

	   runner job ID
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

// WithDefaults hydrates default values in the create runner job execution heart beat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunnerJobExecutionHeartBeatParams) WithDefaults() *CreateRunnerJobExecutionHeartBeatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create runner job execution heart beat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateRunnerJobExecutionHeartBeatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) WithTimeout(timeout time.Duration) *CreateRunnerJobExecutionHeartBeatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) WithContext(ctx context.Context) *CreateRunnerJobExecutionHeartBeatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) WithHTTPClient(client *http.Client) *CreateRunnerJobExecutionHeartBeatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) WithReq(req *models.ServiceCreateRunnerJobExecutionHeartBeatRequest) *CreateRunnerJobExecutionHeartBeatParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) SetReq(req *models.ServiceCreateRunnerJobExecutionHeartBeatRequest) {
	o.Req = req
}

// WithRunnerJobExecutionID adds the runnerJobExecutionID to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) WithRunnerJobExecutionID(runnerJobExecutionID string) *CreateRunnerJobExecutionHeartBeatParams {
	o.SetRunnerJobExecutionID(runnerJobExecutionID)
	return o
}

// SetRunnerJobExecutionID adds the runnerJobExecutionId to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) SetRunnerJobExecutionID(runnerJobExecutionID string) {
	o.RunnerJobExecutionID = runnerJobExecutionID
}

// WithRunnerJobID adds the runnerJobID to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) WithRunnerJobID(runnerJobID string) *CreateRunnerJobExecutionHeartBeatParams {
	o.SetRunnerJobID(runnerJobID)
	return o
}

// SetRunnerJobID adds the runnerJobId to the create runner job execution heart beat params
func (o *CreateRunnerJobExecutionHeartBeatParams) SetRunnerJobID(runnerJobID string) {
	o.RunnerJobID = runnerJobID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRunnerJobExecutionHeartBeatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
