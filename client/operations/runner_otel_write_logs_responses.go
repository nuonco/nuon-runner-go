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

// RunnerOtelWriteLogsReader is a Reader for the RunnerOtelWriteLogs structure.
type RunnerOtelWriteLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunnerOtelWriteLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunnerOtelWriteLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunnerOtelWriteLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRunnerOtelWriteLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRunnerOtelWriteLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunnerOtelWriteLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRunnerOtelWriteLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/runners/{runner_id}/logs] RunnerOtelWriteLogs", response, response.Code())
	}
}

// NewRunnerOtelWriteLogsOK creates a RunnerOtelWriteLogsOK with default headers values
func NewRunnerOtelWriteLogsOK() *RunnerOtelWriteLogsOK {
	return &RunnerOtelWriteLogsOK{}
}

/*
RunnerOtelWriteLogsOK describes a response with status code 200, with default header values.

OK
*/
type RunnerOtelWriteLogsOK struct {
	Payload string
}

// IsSuccess returns true when this runner otel write logs o k response has a 2xx status code
func (o *RunnerOtelWriteLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this runner otel write logs o k response has a 3xx status code
func (o *RunnerOtelWriteLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write logs o k response has a 4xx status code
func (o *RunnerOtelWriteLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this runner otel write logs o k response has a 5xx status code
func (o *RunnerOtelWriteLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write logs o k response a status code equal to that given
func (o *RunnerOtelWriteLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the runner otel write logs o k response
func (o *RunnerOtelWriteLogsOK) Code() int {
	return 200
}

func (o *RunnerOtelWriteLogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsOK %s", 200, payload)
}

func (o *RunnerOtelWriteLogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsOK %s", 200, payload)
}

func (o *RunnerOtelWriteLogsOK) GetPayload() string {
	return o.Payload
}

func (o *RunnerOtelWriteLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteLogsBadRequest creates a RunnerOtelWriteLogsBadRequest with default headers values
func NewRunnerOtelWriteLogsBadRequest() *RunnerOtelWriteLogsBadRequest {
	return &RunnerOtelWriteLogsBadRequest{}
}

/*
RunnerOtelWriteLogsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RunnerOtelWriteLogsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write logs bad request response has a 2xx status code
func (o *RunnerOtelWriteLogsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write logs bad request response has a 3xx status code
func (o *RunnerOtelWriteLogsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write logs bad request response has a 4xx status code
func (o *RunnerOtelWriteLogsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write logs bad request response has a 5xx status code
func (o *RunnerOtelWriteLogsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write logs bad request response a status code equal to that given
func (o *RunnerOtelWriteLogsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the runner otel write logs bad request response
func (o *RunnerOtelWriteLogsBadRequest) Code() int {
	return 400
}

func (o *RunnerOtelWriteLogsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsBadRequest %s", 400, payload)
}

func (o *RunnerOtelWriteLogsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsBadRequest %s", 400, payload)
}

func (o *RunnerOtelWriteLogsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteLogsUnauthorized creates a RunnerOtelWriteLogsUnauthorized with default headers values
func NewRunnerOtelWriteLogsUnauthorized() *RunnerOtelWriteLogsUnauthorized {
	return &RunnerOtelWriteLogsUnauthorized{}
}

/*
RunnerOtelWriteLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RunnerOtelWriteLogsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write logs unauthorized response has a 2xx status code
func (o *RunnerOtelWriteLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write logs unauthorized response has a 3xx status code
func (o *RunnerOtelWriteLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write logs unauthorized response has a 4xx status code
func (o *RunnerOtelWriteLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write logs unauthorized response has a 5xx status code
func (o *RunnerOtelWriteLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write logs unauthorized response a status code equal to that given
func (o *RunnerOtelWriteLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the runner otel write logs unauthorized response
func (o *RunnerOtelWriteLogsUnauthorized) Code() int {
	return 401
}

func (o *RunnerOtelWriteLogsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsUnauthorized %s", 401, payload)
}

func (o *RunnerOtelWriteLogsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsUnauthorized %s", 401, payload)
}

func (o *RunnerOtelWriteLogsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteLogsForbidden creates a RunnerOtelWriteLogsForbidden with default headers values
func NewRunnerOtelWriteLogsForbidden() *RunnerOtelWriteLogsForbidden {
	return &RunnerOtelWriteLogsForbidden{}
}

/*
RunnerOtelWriteLogsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RunnerOtelWriteLogsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write logs forbidden response has a 2xx status code
func (o *RunnerOtelWriteLogsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write logs forbidden response has a 3xx status code
func (o *RunnerOtelWriteLogsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write logs forbidden response has a 4xx status code
func (o *RunnerOtelWriteLogsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write logs forbidden response has a 5xx status code
func (o *RunnerOtelWriteLogsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write logs forbidden response a status code equal to that given
func (o *RunnerOtelWriteLogsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the runner otel write logs forbidden response
func (o *RunnerOtelWriteLogsForbidden) Code() int {
	return 403
}

func (o *RunnerOtelWriteLogsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsForbidden %s", 403, payload)
}

func (o *RunnerOtelWriteLogsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsForbidden %s", 403, payload)
}

func (o *RunnerOtelWriteLogsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteLogsNotFound creates a RunnerOtelWriteLogsNotFound with default headers values
func NewRunnerOtelWriteLogsNotFound() *RunnerOtelWriteLogsNotFound {
	return &RunnerOtelWriteLogsNotFound{}
}

/*
RunnerOtelWriteLogsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RunnerOtelWriteLogsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write logs not found response has a 2xx status code
func (o *RunnerOtelWriteLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write logs not found response has a 3xx status code
func (o *RunnerOtelWriteLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write logs not found response has a 4xx status code
func (o *RunnerOtelWriteLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this runner otel write logs not found response has a 5xx status code
func (o *RunnerOtelWriteLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this runner otel write logs not found response a status code equal to that given
func (o *RunnerOtelWriteLogsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the runner otel write logs not found response
func (o *RunnerOtelWriteLogsNotFound) Code() int {
	return 404
}

func (o *RunnerOtelWriteLogsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsNotFound %s", 404, payload)
}

func (o *RunnerOtelWriteLogsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsNotFound %s", 404, payload)
}

func (o *RunnerOtelWriteLogsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunnerOtelWriteLogsInternalServerError creates a RunnerOtelWriteLogsInternalServerError with default headers values
func NewRunnerOtelWriteLogsInternalServerError() *RunnerOtelWriteLogsInternalServerError {
	return &RunnerOtelWriteLogsInternalServerError{}
}

/*
RunnerOtelWriteLogsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RunnerOtelWriteLogsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this runner otel write logs internal server error response has a 2xx status code
func (o *RunnerOtelWriteLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this runner otel write logs internal server error response has a 3xx status code
func (o *RunnerOtelWriteLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this runner otel write logs internal server error response has a 4xx status code
func (o *RunnerOtelWriteLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this runner otel write logs internal server error response has a 5xx status code
func (o *RunnerOtelWriteLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this runner otel write logs internal server error response a status code equal to that given
func (o *RunnerOtelWriteLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the runner otel write logs internal server error response
func (o *RunnerOtelWriteLogsInternalServerError) Code() int {
	return 500
}

func (o *RunnerOtelWriteLogsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsInternalServerError %s", 500, payload)
}

func (o *RunnerOtelWriteLogsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/logs][%d] runnerOtelWriteLogsInternalServerError %s", 500, payload)
}

func (o *RunnerOtelWriteLogsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *RunnerOtelWriteLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
