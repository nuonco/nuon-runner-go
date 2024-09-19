// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/nuon-runner-go/models"
)

// RunnerOtelWriteMetricsReader is a Reader for the RunnerOtelWriteMetrics structure.
type RunnerOtelWriteMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunnerOtelWriteMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRunnerOtelWriteMetricsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunnerOtelWriteMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRunnerOtelWriteMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRunnerOtelWriteMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunnerOtelWriteMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRunnerOtelWriteMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/runners/{runner_id}/metrics] RunnerOtelWriteMetrics", response, response.Code())
	}
}

// NewRunnerOtelWriteMetricsCreated creates a RunnerOtelWriteMetricsCreated with default headers values
func NewRunnerOtelWriteMetricsCreated() *RunnerOtelWriteMetricsCreated {
	return &RunnerOtelWriteMetricsCreated{}
}

/*
RunnerOtelWriteMetricsCreated describes a response with status code 201, with default header values.

Created
*/
type RunnerOtelWriteMetricsCreated struct {
	Payload string
}

// IsSuccess returns true when this runner otel write metrics created response has a 2xx status code
func (o *RunnerOtelWriteMetricsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this runner otel write metrics created response has a 3xx status code
func (o *RunnerOtelWriteMetricsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write metrics created response has a 4xx status code
func (o *RunnerOtelWriteMetricsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this runner otel write metrics created response has a 5xx status code
func (o *RunnerOtelWriteMetricsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write metrics created response a status code equal to that given
func (o *RunnerOtelWriteMetricsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the runner otel write metrics created response
func (o *RunnerOtelWriteMetricsCreated) Code() int {
	return 201
}

func (o *RunnerOtelWriteMetricsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsCreated  %+v", 201, o.Payload)
}

func (o *RunnerOtelWriteMetricsCreated) String() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsCreated  %+v", 201, o.Payload)
}

func (o *RunnerOtelWriteMetricsCreated) GetPayload() string {
	return o.Payload
}

func (o *RunnerOtelWriteMetricsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteMetricsBadRequest creates a RunnerOtelWriteMetricsBadRequest with default headers values
func NewRunnerOtelWriteMetricsBadRequest() *RunnerOtelWriteMetricsBadRequest {
	return &RunnerOtelWriteMetricsBadRequest{}
}

/*
RunnerOtelWriteMetricsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RunnerOtelWriteMetricsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write metrics bad request response has a 2xx status code
func (o *RunnerOtelWriteMetricsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write metrics bad request response has a 3xx status code
func (o *RunnerOtelWriteMetricsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write metrics bad request response has a 4xx status code
func (o *RunnerOtelWriteMetricsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write metrics bad request response has a 5xx status code
func (o *RunnerOtelWriteMetricsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write metrics bad request response a status code equal to that given
func (o *RunnerOtelWriteMetricsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the runner otel write metrics bad request response
func (o *RunnerOtelWriteMetricsBadRequest) Code() int {
	return 400
}

func (o *RunnerOtelWriteMetricsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *RunnerOtelWriteMetricsBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *RunnerOtelWriteMetricsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteMetricsUnauthorized creates a RunnerOtelWriteMetricsUnauthorized with default headers values
func NewRunnerOtelWriteMetricsUnauthorized() *RunnerOtelWriteMetricsUnauthorized {
	return &RunnerOtelWriteMetricsUnauthorized{}
}

/*
RunnerOtelWriteMetricsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RunnerOtelWriteMetricsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write metrics unauthorized response has a 2xx status code
func (o *RunnerOtelWriteMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write metrics unauthorized response has a 3xx status code
func (o *RunnerOtelWriteMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write metrics unauthorized response has a 4xx status code
func (o *RunnerOtelWriteMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write metrics unauthorized response has a 5xx status code
func (o *RunnerOtelWriteMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write metrics unauthorized response a status code equal to that given
func (o *RunnerOtelWriteMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the runner otel write metrics unauthorized response
func (o *RunnerOtelWriteMetricsUnauthorized) Code() int {
	return 401
}

func (o *RunnerOtelWriteMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *RunnerOtelWriteMetricsUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *RunnerOtelWriteMetricsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteMetricsForbidden creates a RunnerOtelWriteMetricsForbidden with default headers values
func NewRunnerOtelWriteMetricsForbidden() *RunnerOtelWriteMetricsForbidden {
	return &RunnerOtelWriteMetricsForbidden{}
}

/*
RunnerOtelWriteMetricsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RunnerOtelWriteMetricsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write metrics forbidden response has a 2xx status code
func (o *RunnerOtelWriteMetricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write metrics forbidden response has a 3xx status code
func (o *RunnerOtelWriteMetricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write metrics forbidden response has a 4xx status code
func (o *RunnerOtelWriteMetricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write metrics forbidden response has a 5xx status code
func (o *RunnerOtelWriteMetricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write metrics forbidden response a status code equal to that given
func (o *RunnerOtelWriteMetricsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the runner otel write metrics forbidden response
func (o *RunnerOtelWriteMetricsForbidden) Code() int {
	return 403
}

func (o *RunnerOtelWriteMetricsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsForbidden  %+v", 403, o.Payload)
}

func (o *RunnerOtelWriteMetricsForbidden) String() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsForbidden  %+v", 403, o.Payload)
}

func (o *RunnerOtelWriteMetricsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteMetricsNotFound creates a RunnerOtelWriteMetricsNotFound with default headers values
func NewRunnerOtelWriteMetricsNotFound() *RunnerOtelWriteMetricsNotFound {
	return &RunnerOtelWriteMetricsNotFound{}
}

/*
RunnerOtelWriteMetricsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RunnerOtelWriteMetricsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write metrics not found response has a 2xx status code
func (o *RunnerOtelWriteMetricsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write metrics not found response has a 3xx status code
func (o *RunnerOtelWriteMetricsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write metrics not found response has a 4xx status code
func (o *RunnerOtelWriteMetricsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write metrics not found response has a 5xx status code
func (o *RunnerOtelWriteMetricsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write metrics not found response a status code equal to that given
func (o *RunnerOtelWriteMetricsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the runner otel write metrics not found response
func (o *RunnerOtelWriteMetricsNotFound) Code() int {
	return 404
}

func (o *RunnerOtelWriteMetricsNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsNotFound  %+v", 404, o.Payload)
}

func (o *RunnerOtelWriteMetricsNotFound) String() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsNotFound  %+v", 404, o.Payload)
}

func (o *RunnerOtelWriteMetricsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteMetricsInternalServerError creates a RunnerOtelWriteMetricsInternalServerError with default headers values
func NewRunnerOtelWriteMetricsInternalServerError() *RunnerOtelWriteMetricsInternalServerError {
	return &RunnerOtelWriteMetricsInternalServerError{}
}

/*
RunnerOtelWriteMetricsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RunnerOtelWriteMetricsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write metrics internal server error response has a 2xx status code
func (o *RunnerOtelWriteMetricsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write metrics internal server error response has a 3xx status code
func (o *RunnerOtelWriteMetricsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write metrics internal server error response has a 4xx status code
func (o *RunnerOtelWriteMetricsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this runner otel write metrics internal server error response has a 5xx status code
func (o *RunnerOtelWriteMetricsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this runner otel write metrics internal server error response a status code equal to that given
func (o *RunnerOtelWriteMetricsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the runner otel write metrics internal server error response
func (o *RunnerOtelWriteMetricsInternalServerError) Code() int {
	return 500
}

func (o *RunnerOtelWriteMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *RunnerOtelWriteMetricsInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/metrics][%d] runnerOtelWriteMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *RunnerOtelWriteMetricsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
