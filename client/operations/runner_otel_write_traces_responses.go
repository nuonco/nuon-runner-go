// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/nuon-runner-go/models"
)

// RunnerOtelWriteTracesReader is a Reader for the RunnerOtelWriteTraces structure.
type RunnerOtelWriteTracesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunnerOtelWriteTracesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunnerOtelWriteTracesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunnerOtelWriteTracesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRunnerOtelWriteTracesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRunnerOtelWriteTracesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunnerOtelWriteTracesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRunnerOtelWriteTracesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/runners/{runner_id}/traces] RunnerOtelWriteTraces", response, response.Code())
	}
}

// NewRunnerOtelWriteTracesOK creates a RunnerOtelWriteTracesOK with default headers values
func NewRunnerOtelWriteTracesOK() *RunnerOtelWriteTracesOK {
	return &RunnerOtelWriteTracesOK{}
}

/*
RunnerOtelWriteTracesOK describes a response with status code 200, with default header values.

OK
*/
type RunnerOtelWriteTracesOK struct {
	Payload string
}

// IsSuccess returns true when this runner otel write traces o k response has a 2xx status code
func (o *RunnerOtelWriteTracesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this runner otel write traces o k response has a 3xx status code
func (o *RunnerOtelWriteTracesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write traces o k response has a 4xx status code
func (o *RunnerOtelWriteTracesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this runner otel write traces o k response has a 5xx status code
func (o *RunnerOtelWriteTracesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write traces o k response a status code equal to that given
func (o *RunnerOtelWriteTracesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the runner otel write traces o k response
func (o *RunnerOtelWriteTracesOK) Code() int {
	return 200
}

func (o *RunnerOtelWriteTracesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesOK %s", 200, payload)
}

func (o *RunnerOtelWriteTracesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesOK %s", 200, payload)
}

func (o *RunnerOtelWriteTracesOK) GetPayload() string {
	return o.Payload
}

func (o *RunnerOtelWriteTracesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteTracesBadRequest creates a RunnerOtelWriteTracesBadRequest with default headers values
func NewRunnerOtelWriteTracesBadRequest() *RunnerOtelWriteTracesBadRequest {
	return &RunnerOtelWriteTracesBadRequest{}
}

/*
RunnerOtelWriteTracesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RunnerOtelWriteTracesBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write traces bad request response has a 2xx status code
func (o *RunnerOtelWriteTracesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write traces bad request response has a 3xx status code
func (o *RunnerOtelWriteTracesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write traces bad request response has a 4xx status code
func (o *RunnerOtelWriteTracesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write traces bad request response has a 5xx status code
func (o *RunnerOtelWriteTracesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write traces bad request response a status code equal to that given
func (o *RunnerOtelWriteTracesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the runner otel write traces bad request response
func (o *RunnerOtelWriteTracesBadRequest) Code() int {
	return 400
}

func (o *RunnerOtelWriteTracesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesBadRequest %s", 400, payload)
}

func (o *RunnerOtelWriteTracesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesBadRequest %s", 400, payload)
}

func (o *RunnerOtelWriteTracesBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteTracesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteTracesUnauthorized creates a RunnerOtelWriteTracesUnauthorized with default headers values
func NewRunnerOtelWriteTracesUnauthorized() *RunnerOtelWriteTracesUnauthorized {
	return &RunnerOtelWriteTracesUnauthorized{}
}

/*
RunnerOtelWriteTracesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RunnerOtelWriteTracesUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write traces unauthorized response has a 2xx status code
func (o *RunnerOtelWriteTracesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write traces unauthorized response has a 3xx status code
func (o *RunnerOtelWriteTracesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write traces unauthorized response has a 4xx status code
func (o *RunnerOtelWriteTracesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write traces unauthorized response has a 5xx status code
func (o *RunnerOtelWriteTracesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write traces unauthorized response a status code equal to that given
func (o *RunnerOtelWriteTracesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the runner otel write traces unauthorized response
func (o *RunnerOtelWriteTracesUnauthorized) Code() int {
	return 401
}

func (o *RunnerOtelWriteTracesUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesUnauthorized %s", 401, payload)
}

func (o *RunnerOtelWriteTracesUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesUnauthorized %s", 401, payload)
}

func (o *RunnerOtelWriteTracesUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteTracesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteTracesForbidden creates a RunnerOtelWriteTracesForbidden with default headers values
func NewRunnerOtelWriteTracesForbidden() *RunnerOtelWriteTracesForbidden {
	return &RunnerOtelWriteTracesForbidden{}
}

/*
RunnerOtelWriteTracesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RunnerOtelWriteTracesForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write traces forbidden response has a 2xx status code
func (o *RunnerOtelWriteTracesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write traces forbidden response has a 3xx status code
func (o *RunnerOtelWriteTracesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write traces forbidden response has a 4xx status code
func (o *RunnerOtelWriteTracesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write traces forbidden response has a 5xx status code
func (o *RunnerOtelWriteTracesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write traces forbidden response a status code equal to that given
func (o *RunnerOtelWriteTracesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the runner otel write traces forbidden response
func (o *RunnerOtelWriteTracesForbidden) Code() int {
	return 403
}

func (o *RunnerOtelWriteTracesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesForbidden %s", 403, payload)
}

func (o *RunnerOtelWriteTracesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesForbidden %s", 403, payload)
}

func (o *RunnerOtelWriteTracesForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteTracesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteTracesNotFound creates a RunnerOtelWriteTracesNotFound with default headers values
func NewRunnerOtelWriteTracesNotFound() *RunnerOtelWriteTracesNotFound {
	return &RunnerOtelWriteTracesNotFound{}
}

/*
RunnerOtelWriteTracesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RunnerOtelWriteTracesNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write traces not found response has a 2xx status code
func (o *RunnerOtelWriteTracesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write traces not found response has a 3xx status code
func (o *RunnerOtelWriteTracesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write traces not found response has a 4xx status code
func (o *RunnerOtelWriteTracesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write traces not found response has a 5xx status code
func (o *RunnerOtelWriteTracesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write traces not found response a status code equal to that given
func (o *RunnerOtelWriteTracesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the runner otel write traces not found response
func (o *RunnerOtelWriteTracesNotFound) Code() int {
	return 404
}

func (o *RunnerOtelWriteTracesNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesNotFound %s", 404, payload)
}

func (o *RunnerOtelWriteTracesNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesNotFound %s", 404, payload)
}

func (o *RunnerOtelWriteTracesNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteTracesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteTracesInternalServerError creates a RunnerOtelWriteTracesInternalServerError with default headers values
func NewRunnerOtelWriteTracesInternalServerError() *RunnerOtelWriteTracesInternalServerError {
	return &RunnerOtelWriteTracesInternalServerError{}
}

/*
RunnerOtelWriteTracesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RunnerOtelWriteTracesInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write traces internal server error response has a 2xx status code
func (o *RunnerOtelWriteTracesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write traces internal server error response has a 3xx status code
func (o *RunnerOtelWriteTracesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write traces internal server error response has a 4xx status code
func (o *RunnerOtelWriteTracesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this runner otel write traces internal server error response has a 5xx status code
func (o *RunnerOtelWriteTracesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this runner otel write traces internal server error response a status code equal to that given
func (o *RunnerOtelWriteTracesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the runner otel write traces internal server error response
func (o *RunnerOtelWriteTracesInternalServerError) Code() int {
	return 500
}

func (o *RunnerOtelWriteTracesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesInternalServerError %s", 500, payload)
}

func (o *RunnerOtelWriteTracesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/traces][%d] runnerOtelWriteTracesInternalServerError %s", 500, payload)
}

func (o *RunnerOtelWriteTracesInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteTracesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}