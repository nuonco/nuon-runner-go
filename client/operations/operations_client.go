// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new operations API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new operations API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRunnerHealthCheck(params *CreateRunnerHealthCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerHealthCheckCreated, error)

	CreateRunnerHeartBeat(params *CreateRunnerHeartBeatParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerHeartBeatCreated, error)

	CreateRunnerJobExecution(params *CreateRunnerJobExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerJobExecutionCreated, error)

	CreateRunnerJobExecutionResult(params *CreateRunnerJobExecutionResultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerJobExecutionResultCreated, error)

	GetRunner(params *GetRunnerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerOK, error)

	GetRunnerJob(params *GetRunnerJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobOK, error)

	GetRunnerJobExecutions(params *GetRunnerJobExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobExecutionsOK, error)

	GetRunnerJobPlan(params *GetRunnerJobPlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobPlanOK, error)

	GetRunnerJobs(params *GetRunnerJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobsOK, error)

	GetRunnerSettings(params *GetRunnerSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerSettingsOK, error)

	LogStreamWriteLogs(params *LogStreamWriteLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LogStreamWriteLogsCreated, error)

	RunnerOtelWriteLogs(params *RunnerOtelWriteLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunnerOtelWriteLogsCreated, error)

	RunnerOtelWriteMetrics(params *RunnerOtelWriteMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunnerOtelWriteMetricsCreated, error)

	RunnerOtelWriteTraces(params *RunnerOtelWriteTracesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunnerOtelWriteTracesCreated, error)

	UpdateRunnerJobExecution(params *UpdateRunnerJobExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRunnerJobExecutionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRunnerHealthCheck creates a runner health check

Create a runner health check.
*/
func (a *Client) CreateRunnerHealthCheck(params *CreateRunnerHealthCheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerHealthCheckCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRunnerHealthCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRunnerHealthCheck",
		Method:             "POST",
		PathPattern:        "/v1/runners/{runner_id}/health-checks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRunnerHealthCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRunnerHealthCheckCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateRunnerHealthCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateRunnerHeartBeat creates a runner heart beat

Create heart beat.
*/
func (a *Client) CreateRunnerHeartBeat(params *CreateRunnerHeartBeatParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerHeartBeatCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRunnerHeartBeatParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRunnerHeartBeat",
		Method:             "POST",
		PathPattern:        "/v1/runners/{runner_id}/heart-beats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRunnerHeartBeatReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRunnerHeartBeatCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateRunnerHeartBeat: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateRunnerJobExecution creates runner job execution

Create an execution for a runner job.
*/
func (a *Client) CreateRunnerJobExecution(params *CreateRunnerJobExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerJobExecutionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRunnerJobExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRunnerJobExecution",
		Method:             "POST",
		PathPattern:        "/v1/runner-jobs/{runner_job_id}/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRunnerJobExecutionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRunnerJobExecutionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateRunnerJobExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateRunnerJobExecutionResult creates a runner job execution result

Write a runner job execution result, for both success/failure modes.
*/
func (a *Client) CreateRunnerJobExecutionResult(params *CreateRunnerJobExecutionResultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRunnerJobExecutionResultCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRunnerJobExecutionResultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRunnerJobExecutionResult",
		Method:             "POST",
		PathPattern:        "/v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}/result",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRunnerJobExecutionResultReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRunnerJobExecutionResultCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateRunnerJobExecutionResult: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRunner gets a runner

Return a runner.
*/
func (a *Client) GetRunner(params *GetRunnerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunnerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunner",
		Method:             "GET",
		PathPattern:        "/v1/runners/{runner_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunnerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRunnerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRunner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRunnerJob gets runner job

Return a runner job.
*/
func (a *Client) GetRunnerJob(params *GetRunnerJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunnerJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunnerJob",
		Method:             "GET",
		PathPattern:        "/v1/runner-jobs/{runner_job_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunnerJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRunnerJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRunnerJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRunnerJobExecutions gets runner job executions

Return executions for a runner job.
*/
func (a *Client) GetRunnerJobExecutions(params *GetRunnerJobExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobExecutionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunnerJobExecutionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunnerJobExecutions",
		Method:             "GET",
		PathPattern:        "/v1/runner-jobs/{runner_job_id}/executions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunnerJobExecutionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRunnerJobExecutionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRunnerJobExecutions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRunnerJobPlan gets runner job plan

Return a plan for a runner job.
*/
func (a *Client) GetRunnerJobPlan(params *GetRunnerJobPlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunnerJobPlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunnerJobPlan",
		Method:             "GET",
		PathPattern:        "/v1/runner-jobs/{runner_job_id}/plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunnerJobPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRunnerJobPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRunnerJobPlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRunnerJobs gets runner jobs

Fetch runner jobs by type and status.
*/
func (a *Client) GetRunnerJobs(params *GetRunnerJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunnerJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunnerJobs",
		Method:             "GET",
		PathPattern:        "/v1/runners/{runner_id}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunnerJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRunnerJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRunnerJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRunnerSettings gets runner settings

Return runner settings for the provided runner.
*/
func (a *Client) GetRunnerSettings(params *GetRunnerSettingsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetRunnerSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRunnerSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRunnerSettings",
		Method:             "GET",
		PathPattern:        "/v1/runners/{runner_id}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRunnerSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRunnerSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRunnerSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
LogStreamWriteLogs logs stream write logs

Write logs to a log stream. Can be used as an OTEL exporter endpoint.
*/
func (a *Client) LogStreamWriteLogs(params *LogStreamWriteLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LogStreamWriteLogsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogStreamWriteLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LogStreamWriteLogs",
		Method:             "POST",
		PathPattern:        "/v1/log-streams/{log_stream_id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogStreamWriteLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogStreamWriteLogsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for LogStreamWriteLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RunnerOtelWriteLogs runners write logs

	OTEL Exporter compatible endpoint for writing logs. Designed to work with the

custom, runner otel collector/exporter stack. The inbound data schema is the
schema defined by the spec. We use the official otel go SDK for this.

This endpoint accepts protobuffs, but not JSON.
*/
func (a *Client) RunnerOtelWriteLogs(params *RunnerOtelWriteLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunnerOtelWriteLogsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunnerOtelWriteLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunnerOtelWriteLogs",
		Method:             "POST",
		PathPattern:        "/v1/runners/{runner_id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RunnerOtelWriteLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunnerOtelWriteLogsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RunnerOtelWriteLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RunnerOtelWriteMetrics runners write metrics

	OTEL Exporter compatible endpoint for writing metrics. Designed to work with the custom, runner otel collector/exporter

stack.
*/
func (a *Client) RunnerOtelWriteMetrics(params *RunnerOtelWriteMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunnerOtelWriteMetricsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunnerOtelWriteMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunnerOtelWriteMetrics",
		Method:             "POST",
		PathPattern:        "/v1/runners/{runner_id}/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RunnerOtelWriteMetricsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunnerOtelWriteMetricsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RunnerOtelWriteMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	RunnerOtelWriteTraces runners write traces

	OTEL Exporter compatible endpoint for writing traces. Designed to work with the custom, runner otel collector/exporter

stack.
*/
func (a *Client) RunnerOtelWriteTraces(params *RunnerOtelWriteTracesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RunnerOtelWriteTracesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRunnerOtelWriteTracesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RunnerOtelWriteTraces",
		Method:             "POST",
		PathPattern:        "/v1/runners/{runner_id}/traces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RunnerOtelWriteTracesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RunnerOtelWriteTracesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RunnerOtelWriteTraces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateRunnerJobExecution updates a runner job execution

Update a runner job execution to mark the job status.
*/
func (a *Client) UpdateRunnerJobExecution(params *UpdateRunnerJobExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateRunnerJobExecutionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRunnerJobExecutionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateRunnerJobExecution",
		Method:             "PATCH",
		PathPattern:        "/v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRunnerJobExecutionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRunnerJobExecutionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateRunnerJobExecution: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
