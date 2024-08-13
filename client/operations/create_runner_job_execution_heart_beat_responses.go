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

// CreateRunnerJobExecutionHeartBeatReader is a Reader for the CreateRunnerJobExecutionHeartBeat structure.
type CreateRunnerJobExecutionHeartBeatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRunnerJobExecutionHeartBeatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRunnerJobExecutionHeartBeatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRunnerJobExecutionHeartBeatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRunnerJobExecutionHeartBeatUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRunnerJobExecutionHeartBeatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRunnerJobExecutionHeartBeatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRunnerJobExecutionHeartBeatInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats] CreateRunnerJobExecutionHeartBeat", response, response.Code())
	}
}

// NewCreateRunnerJobExecutionHeartBeatOK creates a CreateRunnerJobExecutionHeartBeatOK with default headers values
func NewCreateRunnerJobExecutionHeartBeatOK() *CreateRunnerJobExecutionHeartBeatOK {
	return &CreateRunnerJobExecutionHeartBeatOK{}
}

/*
CreateRunnerJobExecutionHeartBeatOK describes a response with status code 200, with default header values.

OK
*/
type CreateRunnerJobExecutionHeartBeatOK struct {
	Payload *models.AppRunnerJobExecutionHeartBeat
}

// IsSuccess returns true when this create runner job execution heart beat o k response has a 2xx status code
func (o *CreateRunnerJobExecutionHeartBeatOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create runner job execution heart beat o k response has a 3xx status code
func (o *CreateRunnerJobExecutionHeartBeatOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner job execution heart beat o k response has a 4xx status code
func (o *CreateRunnerJobExecutionHeartBeatOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create runner job execution heart beat o k response has a 5xx status code
func (o *CreateRunnerJobExecutionHeartBeatOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner job execution heart beat o k response a status code equal to that given
func (o *CreateRunnerJobExecutionHeartBeatOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create runner job execution heart beat o k response
func (o *CreateRunnerJobExecutionHeartBeatOK) Code() int {
	return 200
}

func (o *CreateRunnerJobExecutionHeartBeatOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatOK %s", 200, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatOK %s", 200, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatOK) GetPayload() *models.AppRunnerJobExecutionHeartBeat {
	return o.Payload
}

func (o *CreateRunnerJobExecutionHeartBeatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRunnerJobExecutionHeartBeat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerJobExecutionHeartBeatBadRequest creates a CreateRunnerJobExecutionHeartBeatBadRequest with default headers values
func NewCreateRunnerJobExecutionHeartBeatBadRequest() *CreateRunnerJobExecutionHeartBeatBadRequest {
	return &CreateRunnerJobExecutionHeartBeatBadRequest{}
}

/*
CreateRunnerJobExecutionHeartBeatBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateRunnerJobExecutionHeartBeatBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner job execution heart beat bad request response has a 2xx status code
func (o *CreateRunnerJobExecutionHeartBeatBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner job execution heart beat bad request response has a 3xx status code
func (o *CreateRunnerJobExecutionHeartBeatBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner job execution heart beat bad request response has a 4xx status code
func (o *CreateRunnerJobExecutionHeartBeatBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner job execution heart beat bad request response has a 5xx status code
func (o *CreateRunnerJobExecutionHeartBeatBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner job execution heart beat bad request response a status code equal to that given
func (o *CreateRunnerJobExecutionHeartBeatBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create runner job execution heart beat bad request response
func (o *CreateRunnerJobExecutionHeartBeatBadRequest) Code() int {
	return 400
}

func (o *CreateRunnerJobExecutionHeartBeatBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatBadRequest %s", 400, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatBadRequest %s", 400, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerJobExecutionHeartBeatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerJobExecutionHeartBeatUnauthorized creates a CreateRunnerJobExecutionHeartBeatUnauthorized with default headers values
func NewCreateRunnerJobExecutionHeartBeatUnauthorized() *CreateRunnerJobExecutionHeartBeatUnauthorized {
	return &CreateRunnerJobExecutionHeartBeatUnauthorized{}
}

/*
CreateRunnerJobExecutionHeartBeatUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateRunnerJobExecutionHeartBeatUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner job execution heart beat unauthorized response has a 2xx status code
func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner job execution heart beat unauthorized response has a 3xx status code
func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner job execution heart beat unauthorized response has a 4xx status code
func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner job execution heart beat unauthorized response has a 5xx status code
func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner job execution heart beat unauthorized response a status code equal to that given
func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create runner job execution heart beat unauthorized response
func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) Code() int {
	return 401
}

func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatUnauthorized %s", 401, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatUnauthorized %s", 401, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerJobExecutionHeartBeatUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerJobExecutionHeartBeatForbidden creates a CreateRunnerJobExecutionHeartBeatForbidden with default headers values
func NewCreateRunnerJobExecutionHeartBeatForbidden() *CreateRunnerJobExecutionHeartBeatForbidden {
	return &CreateRunnerJobExecutionHeartBeatForbidden{}
}

/*
CreateRunnerJobExecutionHeartBeatForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateRunnerJobExecutionHeartBeatForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner job execution heart beat forbidden response has a 2xx status code
func (o *CreateRunnerJobExecutionHeartBeatForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner job execution heart beat forbidden response has a 3xx status code
func (o *CreateRunnerJobExecutionHeartBeatForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner job execution heart beat forbidden response has a 4xx status code
func (o *CreateRunnerJobExecutionHeartBeatForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner job execution heart beat forbidden response has a 5xx status code
func (o *CreateRunnerJobExecutionHeartBeatForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner job execution heart beat forbidden response a status code equal to that given
func (o *CreateRunnerJobExecutionHeartBeatForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create runner job execution heart beat forbidden response
func (o *CreateRunnerJobExecutionHeartBeatForbidden) Code() int {
	return 403
}

func (o *CreateRunnerJobExecutionHeartBeatForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatForbidden %s", 403, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatForbidden %s", 403, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerJobExecutionHeartBeatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerJobExecutionHeartBeatNotFound creates a CreateRunnerJobExecutionHeartBeatNotFound with default headers values
func NewCreateRunnerJobExecutionHeartBeatNotFound() *CreateRunnerJobExecutionHeartBeatNotFound {
	return &CreateRunnerJobExecutionHeartBeatNotFound{}
}

/*
CreateRunnerJobExecutionHeartBeatNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateRunnerJobExecutionHeartBeatNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner job execution heart beat not found response has a 2xx status code
func (o *CreateRunnerJobExecutionHeartBeatNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner job execution heart beat not found response has a 3xx status code
func (o *CreateRunnerJobExecutionHeartBeatNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner job execution heart beat not found response has a 4xx status code
func (o *CreateRunnerJobExecutionHeartBeatNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner job execution heart beat not found response has a 5xx status code
func (o *CreateRunnerJobExecutionHeartBeatNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner job execution heart beat not found response a status code equal to that given
func (o *CreateRunnerJobExecutionHeartBeatNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create runner job execution heart beat not found response
func (o *CreateRunnerJobExecutionHeartBeatNotFound) Code() int {
	return 404
}

func (o *CreateRunnerJobExecutionHeartBeatNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatNotFound %s", 404, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatNotFound %s", 404, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerJobExecutionHeartBeatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerJobExecutionHeartBeatInternalServerError creates a CreateRunnerJobExecutionHeartBeatInternalServerError with default headers values
func NewCreateRunnerJobExecutionHeartBeatInternalServerError() *CreateRunnerJobExecutionHeartBeatInternalServerError {
	return &CreateRunnerJobExecutionHeartBeatInternalServerError{}
}

/*
CreateRunnerJobExecutionHeartBeatInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateRunnerJobExecutionHeartBeatInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner job execution heart beat internal server error response has a 2xx status code
func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner job execution heart beat internal server error response has a 3xx status code
func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner job execution heart beat internal server error response has a 4xx status code
func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create runner job execution heart beat internal server error response has a 5xx status code
func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create runner job execution heart beat internal server error response a status code equal to that given
func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create runner job execution heart beat internal server error response
func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) Code() int {
	return 500
}

func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatInternalServerError %s", 500, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runner-job/{runner_job_id}/executions/{runner_job_execution_id}/heart-beats][%d] createRunnerJobExecutionHeartBeatInternalServerError %s", 500, payload)
}

func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerJobExecutionHeartBeatInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
