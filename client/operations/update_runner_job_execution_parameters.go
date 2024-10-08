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

// NewUpdateRunnerJobExecutionParams creates a new UpdateRunnerJobExecutionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRunnerJobExecutionParams() *UpdateRunnerJobExecutionParams {
	return &UpdateRunnerJobExecutionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRunnerJobExecutionParamsWithTimeout creates a new UpdateRunnerJobExecutionParams object
// with the ability to set a timeout on a request.
func NewUpdateRunnerJobExecutionParamsWithTimeout(timeout time.Duration) *UpdateRunnerJobExecutionParams {
	return &UpdateRunnerJobExecutionParams{
		timeout: timeout,
	}
}

// NewUpdateRunnerJobExecutionParamsWithContext creates a new UpdateRunnerJobExecutionParams object
// with the ability to set a context for a request.
func NewUpdateRunnerJobExecutionParamsWithContext(ctx context.Context) *UpdateRunnerJobExecutionParams {
	return &UpdateRunnerJobExecutionParams{
		Context: ctx,
	}
}

// NewUpdateRunnerJobExecutionParamsWithHTTPClient creates a new UpdateRunnerJobExecutionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRunnerJobExecutionParamsWithHTTPClient(client *http.Client) *UpdateRunnerJobExecutionParams {
	return &UpdateRunnerJobExecutionParams{
		HTTPClient: client,
	}
}

/*
UpdateRunnerJobExecutionParams contains all the parameters to send to the API endpoint

	for the update runner job execution operation.

	Typically these are written to a http.Request.
*/
type UpdateRunnerJobExecutionParams struct {

	/* Req.

	   Input
	*/
	Req *models.ServiceUpdateRunnerJobExecutionRequest

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

// WithDefaults hydrates default values in the update runner job execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRunnerJobExecutionParams) WithDefaults() *UpdateRunnerJobExecutionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update runner job execution params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRunnerJobExecutionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) WithTimeout(timeout time.Duration) *UpdateRunnerJobExecutionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) WithContext(ctx context.Context) *UpdateRunnerJobExecutionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) WithHTTPClient(client *http.Client) *UpdateRunnerJobExecutionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) WithReq(req *models.ServiceUpdateRunnerJobExecutionRequest) *UpdateRunnerJobExecutionParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) SetReq(req *models.ServiceUpdateRunnerJobExecutionRequest) {
	o.Req = req
}

// WithRunnerJobExecutionID adds the runnerJobExecutionID to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) WithRunnerJobExecutionID(runnerJobExecutionID string) *UpdateRunnerJobExecutionParams {
	o.SetRunnerJobExecutionID(runnerJobExecutionID)
	return o
}

// SetRunnerJobExecutionID adds the runnerJobExecutionId to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) SetRunnerJobExecutionID(runnerJobExecutionID string) {
	o.RunnerJobExecutionID = runnerJobExecutionID
}

// WithRunnerJobID adds the runnerJobID to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) WithRunnerJobID(runnerJobID string) *UpdateRunnerJobExecutionParams {
	o.SetRunnerJobID(runnerJobID)
	return o
}

// SetRunnerJobID adds the runnerJobId to the update runner job execution params
func (o *UpdateRunnerJobExecutionParams) SetRunnerJobID(runnerJobID string) {
	o.RunnerJobID = runnerJobID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRunnerJobExecutionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
