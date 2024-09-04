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

// CreateRunnerHealthCheckReader is a Reader for the CreateRunnerHealthCheck structure.
type CreateRunnerHealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRunnerHealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRunnerHealthCheckCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRunnerHealthCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRunnerHealthCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRunnerHealthCheckForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRunnerHealthCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRunnerHealthCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/runners/{runner_id}/health-checks] CreateRunnerHealthCheck", response, response.Code())
	}
}

// NewCreateRunnerHealthCheckCreated creates a CreateRunnerHealthCheckCreated with default headers values
func NewCreateRunnerHealthCheckCreated() *CreateRunnerHealthCheckCreated {
	return &CreateRunnerHealthCheckCreated{}
}

/*
CreateRunnerHealthCheckCreated describes a response with status code 201, with default header values.

Created
*/
type CreateRunnerHealthCheckCreated struct {
	Payload *models.AppRunnerHealthCheck
}

// IsSuccess returns true when this create runner health check created response has a 2xx status code
func (o *CreateRunnerHealthCheckCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create runner health check created response has a 3xx status code
func (o *CreateRunnerHealthCheckCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner health check created response has a 4xx status code
func (o *CreateRunnerHealthCheckCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create runner health check created response has a 5xx status code
func (o *CreateRunnerHealthCheckCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner health check created response a status code equal to that given
func (o *CreateRunnerHealthCheckCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create runner health check created response
func (o *CreateRunnerHealthCheckCreated) Code() int {
	return 201
}

func (o *CreateRunnerHealthCheckCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckCreated %s", 201, payload)
}

func (o *CreateRunnerHealthCheckCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckCreated %s", 201, payload)
}

func (o *CreateRunnerHealthCheckCreated) GetPayload() *models.AppRunnerHealthCheck {
	return o.Payload
}

func (o *CreateRunnerHealthCheckCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRunnerHealthCheck)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerHealthCheckBadRequest creates a CreateRunnerHealthCheckBadRequest with default headers values
func NewCreateRunnerHealthCheckBadRequest() *CreateRunnerHealthCheckBadRequest {
	return &CreateRunnerHealthCheckBadRequest{}
}

/*
CreateRunnerHealthCheckBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateRunnerHealthCheckBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner health check bad request response has a 2xx status code
func (o *CreateRunnerHealthCheckBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner health check bad request response has a 3xx status code
func (o *CreateRunnerHealthCheckBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner health check bad request response has a 4xx status code
func (o *CreateRunnerHealthCheckBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner health check bad request response has a 5xx status code
func (o *CreateRunnerHealthCheckBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner health check bad request response a status code equal to that given
func (o *CreateRunnerHealthCheckBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create runner health check bad request response
func (o *CreateRunnerHealthCheckBadRequest) Code() int {
	return 400
}

func (o *CreateRunnerHealthCheckBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckBadRequest %s", 400, payload)
}

func (o *CreateRunnerHealthCheckBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckBadRequest %s", 400, payload)
}

func (o *CreateRunnerHealthCheckBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerHealthCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerHealthCheckUnauthorized creates a CreateRunnerHealthCheckUnauthorized with default headers values
func NewCreateRunnerHealthCheckUnauthorized() *CreateRunnerHealthCheckUnauthorized {
	return &CreateRunnerHealthCheckUnauthorized{}
}

/*
CreateRunnerHealthCheckUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateRunnerHealthCheckUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner health check unauthorized response has a 2xx status code
func (o *CreateRunnerHealthCheckUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner health check unauthorized response has a 3xx status code
func (o *CreateRunnerHealthCheckUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner health check unauthorized response has a 4xx status code
func (o *CreateRunnerHealthCheckUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner health check unauthorized response has a 5xx status code
func (o *CreateRunnerHealthCheckUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner health check unauthorized response a status code equal to that given
func (o *CreateRunnerHealthCheckUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create runner health check unauthorized response
func (o *CreateRunnerHealthCheckUnauthorized) Code() int {
	return 401
}

func (o *CreateRunnerHealthCheckUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckUnauthorized %s", 401, payload)
}

func (o *CreateRunnerHealthCheckUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckUnauthorized %s", 401, payload)
}

func (o *CreateRunnerHealthCheckUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerHealthCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerHealthCheckForbidden creates a CreateRunnerHealthCheckForbidden with default headers values
func NewCreateRunnerHealthCheckForbidden() *CreateRunnerHealthCheckForbidden {
	return &CreateRunnerHealthCheckForbidden{}
}

/*
CreateRunnerHealthCheckForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateRunnerHealthCheckForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner health check forbidden response has a 2xx status code
func (o *CreateRunnerHealthCheckForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner health check forbidden response has a 3xx status code
func (o *CreateRunnerHealthCheckForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner health check forbidden response has a 4xx status code
func (o *CreateRunnerHealthCheckForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner health check forbidden response has a 5xx status code
func (o *CreateRunnerHealthCheckForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner health check forbidden response a status code equal to that given
func (o *CreateRunnerHealthCheckForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create runner health check forbidden response
func (o *CreateRunnerHealthCheckForbidden) Code() int {
	return 403
}

func (o *CreateRunnerHealthCheckForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckForbidden %s", 403, payload)
}

func (o *CreateRunnerHealthCheckForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckForbidden %s", 403, payload)
}

func (o *CreateRunnerHealthCheckForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerHealthCheckForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerHealthCheckNotFound creates a CreateRunnerHealthCheckNotFound with default headers values
func NewCreateRunnerHealthCheckNotFound() *CreateRunnerHealthCheckNotFound {
	return &CreateRunnerHealthCheckNotFound{}
}

/*
CreateRunnerHealthCheckNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateRunnerHealthCheckNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner health check not found response has a 2xx status code
func (o *CreateRunnerHealthCheckNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner health check not found response has a 3xx status code
func (o *CreateRunnerHealthCheckNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner health check not found response has a 4xx status code
func (o *CreateRunnerHealthCheckNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create runner health check not found response has a 5xx status code
func (o *CreateRunnerHealthCheckNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create runner health check not found response a status code equal to that given
func (o *CreateRunnerHealthCheckNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create runner health check not found response
func (o *CreateRunnerHealthCheckNotFound) Code() int {
	return 404
}

func (o *CreateRunnerHealthCheckNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckNotFound %s", 404, payload)
}

func (o *CreateRunnerHealthCheckNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckNotFound %s", 404, payload)
}

func (o *CreateRunnerHealthCheckNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerHealthCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunnerHealthCheckInternalServerError creates a CreateRunnerHealthCheckInternalServerError with default headers values
func NewCreateRunnerHealthCheckInternalServerError() *CreateRunnerHealthCheckInternalServerError {
	return &CreateRunnerHealthCheckInternalServerError{}
}

/*
CreateRunnerHealthCheckInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateRunnerHealthCheckInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create runner health check internal server error response has a 2xx status code
func (o *CreateRunnerHealthCheckInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create runner health check internal server error response has a 3xx status code
func (o *CreateRunnerHealthCheckInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create runner health check internal server error response has a 4xx status code
func (o *CreateRunnerHealthCheckInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create runner health check internal server error response has a 5xx status code
func (o *CreateRunnerHealthCheckInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create runner health check internal server error response a status code equal to that given
func (o *CreateRunnerHealthCheckInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create runner health check internal server error response
func (o *CreateRunnerHealthCheckInternalServerError) Code() int {
	return 500
}

func (o *CreateRunnerHealthCheckInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckInternalServerError %s", 500, payload)
}

func (o *CreateRunnerHealthCheckInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/runners/{runner_id}/health-checks][%d] createRunnerHealthCheckInternalServerError %s", 500, payload)
}

func (o *CreateRunnerHealthCheckInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateRunnerHealthCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
