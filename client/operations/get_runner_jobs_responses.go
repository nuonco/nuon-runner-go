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

// GetRunnerJobsReader is a Reader for the GetRunnerJobs structure.
type GetRunnerJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunnerJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunnerJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRunnerJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRunnerJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRunnerJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunnerJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRunnerJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/runners/{runner_id}/jobs] GetRunnerJobs", response, response.Code())
	}
}

// NewGetRunnerJobsOK creates a GetRunnerJobsOK with default headers values
func NewGetRunnerJobsOK() *GetRunnerJobsOK {
	return &GetRunnerJobsOK{}
}

/*
GetRunnerJobsOK describes a response with status code 200, with default header values.

OK
*/
type GetRunnerJobsOK struct {
	Payload []*models.AppRunnerJob
}

// IsSuccess returns true when this get runner jobs o k response has a 2xx status code
func (o *GetRunnerJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runner jobs o k response has a 3xx status code
func (o *GetRunnerJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner jobs o k response has a 4xx status code
func (o *GetRunnerJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner jobs o k response has a 5xx status code
func (o *GetRunnerJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner jobs o k response a status code equal to that given
func (o *GetRunnerJobsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runner jobs o k response
func (o *GetRunnerJobsOK) Code() int {
	return 200
}

func (o *GetRunnerJobsOK) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsOK  %+v", 200, o.Payload)
}

func (o *GetRunnerJobsOK) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsOK  %+v", 200, o.Payload)
}

func (o *GetRunnerJobsOK) GetPayload() []*models.AppRunnerJob {
	return o.Payload
}

func (o *GetRunnerJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobsBadRequest creates a GetRunnerJobsBadRequest with default headers values
func NewGetRunnerJobsBadRequest() *GetRunnerJobsBadRequest {
	return &GetRunnerJobsBadRequest{}
}

/*
GetRunnerJobsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetRunnerJobsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner jobs bad request response has a 2xx status code
func (o *GetRunnerJobsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner jobs bad request response has a 3xx status code
func (o *GetRunnerJobsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner jobs bad request response has a 4xx status code
func (o *GetRunnerJobsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner jobs bad request response has a 5xx status code
func (o *GetRunnerJobsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner jobs bad request response a status code equal to that given
func (o *GetRunnerJobsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get runner jobs bad request response
func (o *GetRunnerJobsBadRequest) Code() int {
	return 400
}

func (o *GetRunnerJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunnerJobsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRunnerJobsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobsUnauthorized creates a GetRunnerJobsUnauthorized with default headers values
func NewGetRunnerJobsUnauthorized() *GetRunnerJobsUnauthorized {
	return &GetRunnerJobsUnauthorized{}
}

/*
GetRunnerJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRunnerJobsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner jobs unauthorized response has a 2xx status code
func (o *GetRunnerJobsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner jobs unauthorized response has a 3xx status code
func (o *GetRunnerJobsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner jobs unauthorized response has a 4xx status code
func (o *GetRunnerJobsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner jobs unauthorized response has a 5xx status code
func (o *GetRunnerJobsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner jobs unauthorized response a status code equal to that given
func (o *GetRunnerJobsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get runner jobs unauthorized response
func (o *GetRunnerJobsUnauthorized) Code() int {
	return 401
}

func (o *GetRunnerJobsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRunnerJobsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRunnerJobsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobsForbidden creates a GetRunnerJobsForbidden with default headers values
func NewGetRunnerJobsForbidden() *GetRunnerJobsForbidden {
	return &GetRunnerJobsForbidden{}
}

/*
GetRunnerJobsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRunnerJobsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner jobs forbidden response has a 2xx status code
func (o *GetRunnerJobsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner jobs forbidden response has a 3xx status code
func (o *GetRunnerJobsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner jobs forbidden response has a 4xx status code
func (o *GetRunnerJobsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner jobs forbidden response has a 5xx status code
func (o *GetRunnerJobsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner jobs forbidden response a status code equal to that given
func (o *GetRunnerJobsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get runner jobs forbidden response
func (o *GetRunnerJobsForbidden) Code() int {
	return 403
}

func (o *GetRunnerJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsForbidden  %+v", 403, o.Payload)
}

func (o *GetRunnerJobsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsForbidden  %+v", 403, o.Payload)
}

func (o *GetRunnerJobsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobsNotFound creates a GetRunnerJobsNotFound with default headers values
func NewGetRunnerJobsNotFound() *GetRunnerJobsNotFound {
	return &GetRunnerJobsNotFound{}
}

/*
GetRunnerJobsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRunnerJobsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner jobs not found response has a 2xx status code
func (o *GetRunnerJobsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner jobs not found response has a 3xx status code
func (o *GetRunnerJobsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner jobs not found response has a 4xx status code
func (o *GetRunnerJobsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner jobs not found response has a 5xx status code
func (o *GetRunnerJobsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner jobs not found response a status code equal to that given
func (o *GetRunnerJobsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get runner jobs not found response
func (o *GetRunnerJobsNotFound) Code() int {
	return 404
}

func (o *GetRunnerJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerJobsNotFound) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerJobsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunnerJobsInternalServerError creates a GetRunnerJobsInternalServerError with default headers values
func NewGetRunnerJobsInternalServerError() *GetRunnerJobsInternalServerError {
	return &GetRunnerJobsInternalServerError{}
}

/*
GetRunnerJobsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRunnerJobsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get runner jobs internal server error response has a 2xx status code
func (o *GetRunnerJobsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner jobs internal server error response has a 3xx status code
func (o *GetRunnerJobsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner jobs internal server error response has a 4xx status code
func (o *GetRunnerJobsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner jobs internal server error response has a 5xx status code
func (o *GetRunnerJobsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get runner jobs internal server error response a status code equal to that given
func (o *GetRunnerJobsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get runner jobs internal server error response
func (o *GetRunnerJobsInternalServerError) Code() int {
	return 500
}

func (o *GetRunnerJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRunnerJobsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/runners/{runner_id}/jobs][%d] getRunnerJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRunnerJobsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetRunnerJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
