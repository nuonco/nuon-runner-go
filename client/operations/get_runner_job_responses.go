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

// GetRunnerJobReader is a Reader for the GetRunnerJob structure.
type GetRunnerJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunnerJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunnerJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunnerJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRunnerJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRunnerJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunnerJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRunnerJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/runner-jobs/{runner_job_id}] GetRunnerJob", response, response.Code())
	}
}

// NewGetRunnerJobOK creates a GetRunnerJobOK with default headers values
func NewGetRunnerJobOK() *GetRunnerJobOK {
	return &GetRunnerJobOK{}
}

/*
GetRunnerJobOK describes a response with status code 200, with default header values.

OK
*/
type GetRunnerJobOK struct {
	Payload *models.AppRunnerJob
}

// IsSuccess returns true when this get runner job o k response has a 2xx status code
func (o *GetRunnerJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runner job o k response has a 3xx status code
func (o *GetRunnerJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job o k response has a 4xx status code
func (o *GetRunnerJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner job o k response has a 5xx status code
func (o *GetRunnerJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job o k response a status code equal to that given
func (o *GetRunnerJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runner job o k response
func (o *GetRunnerJobOK) Code() int {
	return 200
}

func (o *GetRunnerJobOK) Error() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobOK  %+v", 200, o.Payload)
}

func (o *GetRunnerJobOK) String() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobOK  %+v", 200, o.Payload)
}

func (o *GetRunnerJobOK) GetPayload() *models.AppRunnerJob {
	return o.Payload
}

func (o *GetRunnerJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRunnerJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobBadRequest creates a GetRunnerJobBadRequest with default headers values
func NewGetRunnerJobBadRequest() *GetRunnerJobBadRequest {
	return &GetRunnerJobBadRequest{}
}

/*
GetRunnerJobBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRunnerJobBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job bad request response has a 2xx status code
func (o *GetRunnerJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job bad request response has a 3xx status code
func (o *GetRunnerJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job bad request response has a 4xx status code
func (o *GetRunnerJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job bad request response has a 5xx status code
func (o *GetRunnerJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job bad request response a status code equal to that given
func (o *GetRunnerJobBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get runner job bad request response
func (o *GetRunnerJobBadRequest) Code() int {
	return 400
}

func (o *GetRunnerJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunnerJobBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunnerJobBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobUnauthorized creates a GetRunnerJobUnauthorized with default headers values
func NewGetRunnerJobUnauthorized() *GetRunnerJobUnauthorized {
	return &GetRunnerJobUnauthorized{}
}

/*
GetRunnerJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRunnerJobUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job unauthorized response has a 2xx status code
func (o *GetRunnerJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job unauthorized response has a 3xx status code
func (o *GetRunnerJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job unauthorized response has a 4xx status code
func (o *GetRunnerJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job unauthorized response has a 5xx status code
func (o *GetRunnerJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job unauthorized response a status code equal to that given
func (o *GetRunnerJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get runner job unauthorized response
func (o *GetRunnerJobUnauthorized) Code() int {
	return 401
}

func (o *GetRunnerJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRunnerJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRunnerJobUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobForbidden creates a GetRunnerJobForbidden with default headers values
func NewGetRunnerJobForbidden() *GetRunnerJobForbidden {
	return &GetRunnerJobForbidden{}
}

/*
GetRunnerJobForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRunnerJobForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job forbidden response has a 2xx status code
func (o *GetRunnerJobForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job forbidden response has a 3xx status code
func (o *GetRunnerJobForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job forbidden response has a 4xx status code
func (o *GetRunnerJobForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job forbidden response has a 5xx status code
func (o *GetRunnerJobForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job forbidden response a status code equal to that given
func (o *GetRunnerJobForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get runner job forbidden response
func (o *GetRunnerJobForbidden) Code() int {
	return 403
}

func (o *GetRunnerJobForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobForbidden  %+v", 403, o.Payload)
}

func (o *GetRunnerJobForbidden) String() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobForbidden  %+v", 403, o.Payload)
}

func (o *GetRunnerJobForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobNotFound creates a GetRunnerJobNotFound with default headers values
func NewGetRunnerJobNotFound() *GetRunnerJobNotFound {
	return &GetRunnerJobNotFound{}
}

/*
GetRunnerJobNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRunnerJobNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job not found response has a 2xx status code
func (o *GetRunnerJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job not found response has a 3xx status code
func (o *GetRunnerJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job not found response has a 4xx status code
func (o *GetRunnerJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner job not found response has a 5xx status code
func (o *GetRunnerJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner job not found response a status code equal to that given
func (o *GetRunnerJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get runner job not found response
func (o *GetRunnerJobNotFound) Code() int {
	return 404
}

func (o *GetRunnerJobNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerJobNotFound) String() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerJobNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobInternalServerError creates a GetRunnerJobInternalServerError with default headers values
func NewGetRunnerJobInternalServerError() *GetRunnerJobInternalServerError {
	return &GetRunnerJobInternalServerError{}
}

/*
GetRunnerJobInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRunnerJobInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner job internal server error response has a 2xx status code
func (o *GetRunnerJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner job internal server error response has a 3xx status code
func (o *GetRunnerJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner job internal server error response has a 4xx status code
func (o *GetRunnerJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner job internal server error response has a 5xx status code
func (o *GetRunnerJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get runner job internal server error response a status code equal to that given
func (o *GetRunnerJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get runner job internal server error response
func (o *GetRunnerJobInternalServerError) Code() int {
	return 500
}

func (o *GetRunnerJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRunnerJobInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/runner-jobs/{runner_job_id}][%d] getRunnerJobInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRunnerJobInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
