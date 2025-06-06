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

// NewGetTerraformWorkspaceStateByIDParams creates a new GetTerraformWorkspaceStateByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformWorkspaceStateByIDParams() *GetTerraformWorkspaceStateByIDParams {
	return &GetTerraformWorkspaceStateByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformWorkspaceStateByIDParamsWithTimeout creates a new GetTerraformWorkspaceStateByIDParams object
// with the ability to set a timeout on a request.
func NewGetTerraformWorkspaceStateByIDParamsWithTimeout(timeout time.Duration) *GetTerraformWorkspaceStateByIDParams {
	return &GetTerraformWorkspaceStateByIDParams{
		timeout: timeout,
	}
}

// NewGetTerraformWorkspaceStateByIDParamsWithContext creates a new GetTerraformWorkspaceStateByIDParams object
// with the ability to set a context for a request.
func NewGetTerraformWorkspaceStateByIDParamsWithContext(ctx context.Context) *GetTerraformWorkspaceStateByIDParams {
	return &GetTerraformWorkspaceStateByIDParams{
		Context: ctx,
	}
}

// NewGetTerraformWorkspaceStateByIDParamsWithHTTPClient creates a new GetTerraformWorkspaceStateByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformWorkspaceStateByIDParamsWithHTTPClient(client *http.Client) *GetTerraformWorkspaceStateByIDParams {
	return &GetTerraformWorkspaceStateByIDParams{
		HTTPClient: client,
	}
}

/*
GetTerraformWorkspaceStateByIDParams contains all the parameters to send to the API endpoint

	for the get terraform workspace state by ID operation.

	Typically these are written to a http.Request.
*/
type GetTerraformWorkspaceStateByIDParams struct {

	/* StateID.

	   state ID
	*/
	StateID string

	/* WorkspaceID.

	   workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform workspace state by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformWorkspaceStateByIDParams) WithDefaults() *GetTerraformWorkspaceStateByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform workspace state by ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformWorkspaceStateByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) WithTimeout(timeout time.Duration) *GetTerraformWorkspaceStateByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) WithContext(ctx context.Context) *GetTerraformWorkspaceStateByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) WithHTTPClient(client *http.Client) *GetTerraformWorkspaceStateByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStateID adds the stateID to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) WithStateID(stateID string) *GetTerraformWorkspaceStateByIDParams {
	o.SetStateID(stateID)
	return o
}

// SetStateID adds the stateId to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) SetStateID(stateID string) {
	o.StateID = stateID
}

// WithWorkspaceID adds the workspaceID to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) WithWorkspaceID(workspaceID string) *GetTerraformWorkspaceStateByIDParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get terraform workspace state by ID params
func (o *GetTerraformWorkspaceStateByIDParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformWorkspaceStateByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param state_id
	if err := r.SetPathParam("state_id", o.StateID); err != nil {
		return err
	}

	// path param workspace_id
	if err := r.SetPathParam("workspace_id", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
