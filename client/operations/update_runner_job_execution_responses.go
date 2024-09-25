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

// UpdateRunnerJobExecutionReader is a Reader for the UpdateRunnerJobExecution structure.
type UpdateRunnerJobExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunnerJobExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRunnerJobExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunnerJobExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRunnerJobExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRunnerJobExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRunnerJobExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRunnerJobExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}] UpdateRunnerJobExecution", response, response.Code())
	}
}

// NewUpdateRunnerJobExecutionOK creates a UpdateRunnerJobExecutionOK with default headers values
func NewUpdateRunnerJobExecutionOK() *UpdateRunnerJobExecutionOK {
	return &UpdateRunnerJobExecutionOK{}
}

/*
UpdateRunnerJobExecutionOK describes a response with status code 200, with default header values.

OK
*/
type UpdateRunnerJobExecutionOK struct {
	Payload *models.AppRunnerJobExecution
}

// IsSuccess returns true when this update runner job execution o k response has a 2xx status code
func (o *UpdateRunnerJobExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update runner job execution o k response has a 3xx status code
func (o *UpdateRunnerJobExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update runner job execution o k response has a 4xx status code
func (o *UpdateRunnerJobExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update runner job execution o k response has a 5xx status code
func (o *UpdateRunnerJobExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update runner job execution o k response a status code equal to that given
func (o *UpdateRunnerJobExecutionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update runner job execution o k response
func (o *UpdateRunnerJobExecutionOK) Code() int {
	return 200
}

func (o *UpdateRunnerJobExecutionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionOK %s", 200, payload)
}

func (o *UpdateRunnerJobExecutionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionOK %s", 200, payload)
}

func (o *UpdateRunnerJobExecutionOK) GetPayload() *models.AppRunnerJobExecution {
	return o.Payload
}

func (o *UpdateRunnerJobExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRunnerJobExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunnerJobExecutionBadRequest creates a UpdateRunnerJobExecutionBadRequest with default headers values
func NewUpdateRunnerJobExecutionBadRequest() *UpdateRunnerJobExecutionBadRequest {
	return &UpdateRunnerJobExecutionBadRequest{}
}

/*
UpdateRunnerJobExecutionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateRunnerJobExecutionBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update runner job execution bad request response has a 2xx status code
func (o *UpdateRunnerJobExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update runner job execution bad request response has a 3xx status code
func (o *UpdateRunnerJobExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update runner job execution bad request response has a 4xx status code
func (o *UpdateRunnerJobExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update runner job execution bad request response has a 5xx status code
func (o *UpdateRunnerJobExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update runner job execution bad request response a status code equal to that given
func (o *UpdateRunnerJobExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update runner job execution bad request response
func (o *UpdateRunnerJobExecutionBadRequest) Code() int {
	return 400
}

func (o *UpdateRunnerJobExecutionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionBadRequest %s", 400, payload)
}

func (o *UpdateRunnerJobExecutionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionBadRequest %s", 400, payload)
}

func (o *UpdateRunnerJobExecutionBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateRunnerJobExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunnerJobExecutionUnauthorized creates a UpdateRunnerJobExecutionUnauthorized with default headers values
func NewUpdateRunnerJobExecutionUnauthorized() *UpdateRunnerJobExecutionUnauthorized {
	return &UpdateRunnerJobExecutionUnauthorized{}
}

/*
UpdateRunnerJobExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateRunnerJobExecutionUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update runner job execution unauthorized response has a 2xx status code
func (o *UpdateRunnerJobExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update runner job execution unauthorized response has a 3xx status code
func (o *UpdateRunnerJobExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update runner job execution unauthorized response has a 4xx status code
func (o *UpdateRunnerJobExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update runner job execution unauthorized response has a 5xx status code
func (o *UpdateRunnerJobExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update runner job execution unauthorized response a status code equal to that given
func (o *UpdateRunnerJobExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update runner job execution unauthorized response
func (o *UpdateRunnerJobExecutionUnauthorized) Code() int {
	return 401
}

func (o *UpdateRunnerJobExecutionUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionUnauthorized %s", 401, payload)
}

func (o *UpdateRunnerJobExecutionUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionUnauthorized %s", 401, payload)
}

func (o *UpdateRunnerJobExecutionUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateRunnerJobExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunnerJobExecutionForbidden creates a UpdateRunnerJobExecutionForbidden with default headers values
func NewUpdateRunnerJobExecutionForbidden() *UpdateRunnerJobExecutionForbidden {
	return &UpdateRunnerJobExecutionForbidden{}
}

/*
UpdateRunnerJobExecutionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRunnerJobExecutionForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update runner job execution forbidden response has a 2xx status code
func (o *UpdateRunnerJobExecutionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update runner job execution forbidden response has a 3xx status code
func (o *UpdateRunnerJobExecutionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update runner job execution forbidden response has a 4xx status code
func (o *UpdateRunnerJobExecutionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update runner job execution forbidden response has a 5xx status code
func (o *UpdateRunnerJobExecutionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update runner job execution forbidden response a status code equal to that given
func (o *UpdateRunnerJobExecutionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update runner job execution forbidden response
func (o *UpdateRunnerJobExecutionForbidden) Code() int {
	return 403
}

func (o *UpdateRunnerJobExecutionForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionForbidden %s", 403, payload)
}

func (o *UpdateRunnerJobExecutionForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionForbidden %s", 403, payload)
}

func (o *UpdateRunnerJobExecutionForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateRunnerJobExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunnerJobExecutionNotFound creates a UpdateRunnerJobExecutionNotFound with default headers values
func NewUpdateRunnerJobExecutionNotFound() *UpdateRunnerJobExecutionNotFound {
	return &UpdateRunnerJobExecutionNotFound{}
}

/*
UpdateRunnerJobExecutionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateRunnerJobExecutionNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update runner job execution not found response has a 2xx status code
func (o *UpdateRunnerJobExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update runner job execution not found response has a 3xx status code
func (o *UpdateRunnerJobExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update runner job execution not found response has a 4xx status code
func (o *UpdateRunnerJobExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update runner job execution not found response has a 5xx status code
func (o *UpdateRunnerJobExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update runner job execution not found response a status code equal to that given
func (o *UpdateRunnerJobExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update runner job execution not found response
func (o *UpdateRunnerJobExecutionNotFound) Code() int {
	return 404
}

func (o *UpdateRunnerJobExecutionNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionNotFound %s", 404, payload)
}

func (o *UpdateRunnerJobExecutionNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionNotFound %s", 404, payload)
}

func (o *UpdateRunnerJobExecutionNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateRunnerJobExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunnerJobExecutionInternalServerError creates a UpdateRunnerJobExecutionInternalServerError with default headers values
func NewUpdateRunnerJobExecutionInternalServerError() *UpdateRunnerJobExecutionInternalServerError {
	return &UpdateRunnerJobExecutionInternalServerError{}
}

/*
UpdateRunnerJobExecutionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateRunnerJobExecutionInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this update runner job execution internal server error response has a 2xx status code
func (o *UpdateRunnerJobExecutionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update runner job execution internal server error response has a 3xx status code
func (o *UpdateRunnerJobExecutionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update runner job execution internal server error response has a 4xx status code
func (o *UpdateRunnerJobExecutionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update runner job execution internal server error response has a 5xx status code
func (o *UpdateRunnerJobExecutionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update runner job execution internal server error response a status code equal to that given
func (o *UpdateRunnerJobExecutionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update runner job execution internal server error response
func (o *UpdateRunnerJobExecutionInternalServerError) Code() int {
	return 500
}

func (o *UpdateRunnerJobExecutionInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionInternalServerError %s", 500, payload)
}

func (o *UpdateRunnerJobExecutionInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /v1/runner-jobs/{runner_job_id}/executions/{runner_job_execution_id}][%d] updateRunnerJobExecutionInternalServerError %s", 500, payload)
}

func (o *UpdateRunnerJobExecutionInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *UpdateRunnerJobExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}