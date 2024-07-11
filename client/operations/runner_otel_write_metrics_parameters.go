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

// NewRunnerOtelWriteMetricsParams creates a new RunnerOtelWriteMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRunnerOtelWriteMetricsParams() *RunnerOtelWriteMetricsParams {
	return &RunnerOtelWriteMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRunnerOtelWriteMetricsParamsWithTimeout creates a new RunnerOtelWriteMetricsParams object
// with the ability to set a timeout on a request.
func NewRunnerOtelWriteMetricsParamsWithTimeout(timeout time.Duration) *RunnerOtelWriteMetricsParams {
	return &RunnerOtelWriteMetricsParams{
		timeout: timeout,
	}
}

// NewRunnerOtelWriteMetricsParamsWithContext creates a new RunnerOtelWriteMetricsParams object
// with the ability to set a context for a request.
func NewRunnerOtelWriteMetricsParamsWithContext(ctx context.Context) *RunnerOtelWriteMetricsParams {
	return &RunnerOtelWriteMetricsParams{
		Context: ctx,
	}
}

// NewRunnerOtelWriteMetricsParamsWithHTTPClient creates a new RunnerOtelWriteMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRunnerOtelWriteMetricsParamsWithHTTPClient(client *http.Client) *RunnerOtelWriteMetricsParams {
	return &RunnerOtelWriteMetricsParams{
		HTTPClient: client,
	}
}

/*
RunnerOtelWriteMetricsParams contains all the parameters to send to the API endpoint

	for the runner otel write metrics operation.

	Typically these are written to a http.Request.
*/
type RunnerOtelWriteMetricsParams struct {

	/* Req.

	   Input
	*/
	Req interface{}

	/* RunnerID.

	   runner ID
	*/
	RunnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the runner otel write metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunnerOtelWriteMetricsParams) WithDefaults() *RunnerOtelWriteMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the runner otel write metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RunnerOtelWriteMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) WithTimeout(timeout time.Duration) *RunnerOtelWriteMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) WithContext(ctx context.Context) *RunnerOtelWriteMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) WithHTTPClient(client *http.Client) *RunnerOtelWriteMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReq adds the req to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) WithReq(req interface{}) *RunnerOtelWriteMetricsParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) SetReq(req interface{}) {
	o.Req = req
}

// WithRunnerID adds the runnerID to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) WithRunnerID(runnerID string) *RunnerOtelWriteMetricsParams {
	o.SetRunnerID(runnerID)
	return o
}

// SetRunnerID adds the runnerId to the runner otel write metrics params
func (o *RunnerOtelWriteMetricsParams) SetRunnerID(runnerID string) {
	o.RunnerID = runnerID
}

// WriteToRequest writes these params to a swagger request
func (o *RunnerOtelWriteMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	// path param runner_id
	if err := r.SetPathParam("runner_id", o.RunnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
