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

// LogStreamWriteLogsReader is a Reader for the LogStreamWriteLogs structure.
type LogStreamWriteLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogStreamWriteLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLogStreamWriteLogsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLogStreamWriteLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLogStreamWriteLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLogStreamWriteLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLogStreamWriteLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLogStreamWriteLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/log-streams/{log_stream_id}/logs] LogStreamWriteLogs", response, response.Code())
	}
}

// NewLogStreamWriteLogsCreated creates a LogStreamWriteLogsCreated with default headers values
func NewLogStreamWriteLogsCreated() *LogStreamWriteLogsCreated {
	return &LogStreamWriteLogsCreated{}
}

/*
LogStreamWriteLogsCreated describes a response with status code 201, with default header values.

Created
*/
type LogStreamWriteLogsCreated struct {
	Payload string
}

// IsSuccess returns true when this log stream write logs created response has a 2xx status code
func (o *LogStreamWriteLogsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this log stream write logs created response has a 3xx status code
func (o *LogStreamWriteLogsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream write logs created response has a 4xx status code
func (o *LogStreamWriteLogsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this log stream write logs created response has a 5xx status code
func (o *LogStreamWriteLogsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream write logs created response a status code equal to that given
func (o *LogStreamWriteLogsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the log stream write logs created response
func (o *LogStreamWriteLogsCreated) Code() int {
	return 201
}

func (o *LogStreamWriteLogsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsCreated  %+v", 201, o.Payload)
}

func (o *LogStreamWriteLogsCreated) String() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsCreated  %+v", 201, o.Payload)
}

func (o *LogStreamWriteLogsCreated) GetPayload() string {
	return o.Payload
}

func (o *LogStreamWriteLogsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamWriteLogsBadRequest creates a LogStreamWriteLogsBadRequest with default headers values
func NewLogStreamWriteLogsBadRequest() *LogStreamWriteLogsBadRequest {
	return &LogStreamWriteLogsBadRequest{}
}

/*
LogStreamWriteLogsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type LogStreamWriteLogsBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream write logs bad request response has a 2xx status code
func (o *LogStreamWriteLogsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream write logs bad request response has a 3xx status code
func (o *LogStreamWriteLogsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream write logs bad request response has a 4xx status code
func (o *LogStreamWriteLogsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream write logs bad request response has a 5xx status code
func (o *LogStreamWriteLogsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream write logs bad request response a status code equal to that given
func (o *LogStreamWriteLogsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the log stream write logs bad request response
func (o *LogStreamWriteLogsBadRequest) Code() int {
	return 400
}

func (o *LogStreamWriteLogsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsBadRequest  %+v", 400, o.Payload)
}

func (o *LogStreamWriteLogsBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsBadRequest  %+v", 400, o.Payload)
}

func (o *LogStreamWriteLogsBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamWriteLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamWriteLogsUnauthorized creates a LogStreamWriteLogsUnauthorized with default headers values
func NewLogStreamWriteLogsUnauthorized() *LogStreamWriteLogsUnauthorized {
	return &LogStreamWriteLogsUnauthorized{}
}

/*
LogStreamWriteLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type LogStreamWriteLogsUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream write logs unauthorized response has a 2xx status code
func (o *LogStreamWriteLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream write logs unauthorized response has a 3xx status code
func (o *LogStreamWriteLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream write logs unauthorized response has a 4xx status code
func (o *LogStreamWriteLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream write logs unauthorized response has a 5xx status code
func (o *LogStreamWriteLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream write logs unauthorized response a status code equal to that given
func (o *LogStreamWriteLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the log stream write logs unauthorized response
func (o *LogStreamWriteLogsUnauthorized) Code() int {
	return 401
}

func (o *LogStreamWriteLogsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *LogStreamWriteLogsUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *LogStreamWriteLogsUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamWriteLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamWriteLogsForbidden creates a LogStreamWriteLogsForbidden with default headers values
func NewLogStreamWriteLogsForbidden() *LogStreamWriteLogsForbidden {
	return &LogStreamWriteLogsForbidden{}
}

/*
LogStreamWriteLogsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type LogStreamWriteLogsForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream write logs forbidden response has a 2xx status code
func (o *LogStreamWriteLogsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream write logs forbidden response has a 3xx status code
func (o *LogStreamWriteLogsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream write logs forbidden response has a 4xx status code
func (o *LogStreamWriteLogsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream write logs forbidden response has a 5xx status code
func (o *LogStreamWriteLogsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream write logs forbidden response a status code equal to that given
func (o *LogStreamWriteLogsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the log stream write logs forbidden response
func (o *LogStreamWriteLogsForbidden) Code() int {
	return 403
}

func (o *LogStreamWriteLogsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsForbidden  %+v", 403, o.Payload)
}

func (o *LogStreamWriteLogsForbidden) String() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsForbidden  %+v", 403, o.Payload)
}

func (o *LogStreamWriteLogsForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamWriteLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamWriteLogsNotFound creates a LogStreamWriteLogsNotFound with default headers values
func NewLogStreamWriteLogsNotFound() *LogStreamWriteLogsNotFound {
	return &LogStreamWriteLogsNotFound{}
}

/*
LogStreamWriteLogsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type LogStreamWriteLogsNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream write logs not found response has a 2xx status code
func (o *LogStreamWriteLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream write logs not found response has a 3xx status code
func (o *LogStreamWriteLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream write logs not found response has a 4xx status code
func (o *LogStreamWriteLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this log stream write logs not found response has a 5xx status code
func (o *LogStreamWriteLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this log stream write logs not found response a status code equal to that given
func (o *LogStreamWriteLogsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the log stream write logs not found response
func (o *LogStreamWriteLogsNotFound) Code() int {
	return 404
}

func (o *LogStreamWriteLogsNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsNotFound  %+v", 404, o.Payload)
}

func (o *LogStreamWriteLogsNotFound) String() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsNotFound  %+v", 404, o.Payload)
}

func (o *LogStreamWriteLogsNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamWriteLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogStreamWriteLogsInternalServerError creates a LogStreamWriteLogsInternalServerError with default headers values
func NewLogStreamWriteLogsInternalServerError() *LogStreamWriteLogsInternalServerError {
	return &LogStreamWriteLogsInternalServerError{}
}

/*
LogStreamWriteLogsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type LogStreamWriteLogsInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this log stream write logs internal server error response has a 2xx status code
func (o *LogStreamWriteLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this log stream write logs internal server error response has a 3xx status code
func (o *LogStreamWriteLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this log stream write logs internal server error response has a 4xx status code
func (o *LogStreamWriteLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this log stream write logs internal server error response has a 5xx status code
func (o *LogStreamWriteLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this log stream write logs internal server error response a status code equal to that given
func (o *LogStreamWriteLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the log stream write logs internal server error response
func (o *LogStreamWriteLogsInternalServerError) Code() int {
	return 500
}

func (o *LogStreamWriteLogsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *LogStreamWriteLogsInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/log-streams/{log_stream_id}/logs][%d] logStreamWriteLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *LogStreamWriteLogsInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *LogStreamWriteLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
