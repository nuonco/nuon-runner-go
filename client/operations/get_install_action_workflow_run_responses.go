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

// GetInstallActionWorkflowRunReader is a Reader for the GetInstallActionWorkflowRun structure.
type GetInstallActionWorkflowRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallActionWorkflowRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallActionWorkflowRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallActionWorkflowRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallActionWorkflowRunUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallActionWorkflowRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallActionWorkflowRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallActionWorkflowRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}] GetInstallActionWorkflowRun", response, response.Code())
	}
}

// NewGetInstallActionWorkflowRunOK creates a GetInstallActionWorkflowRunOK with default headers values
func NewGetInstallActionWorkflowRunOK() *GetInstallActionWorkflowRunOK {
	return &GetInstallActionWorkflowRunOK{}
}

/*
GetInstallActionWorkflowRunOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallActionWorkflowRunOK struct {
	Payload *models.AppInstallActionWorkflowRun
}

// IsSuccess returns true when this get install action workflow run o k response has a 2xx status code
func (o *GetInstallActionWorkflowRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install action workflow run o k response has a 3xx status code
func (o *GetInstallActionWorkflowRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install action workflow run o k response has a 4xx status code
func (o *GetInstallActionWorkflowRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install action workflow run o k response has a 5xx status code
func (o *GetInstallActionWorkflowRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install action workflow run o k response a status code equal to that given
func (o *GetInstallActionWorkflowRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install action workflow run o k response
func (o *GetInstallActionWorkflowRunOK) Code() int {
	return 200
}

func (o *GetInstallActionWorkflowRunOK) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunOK  %+v", 200, o.Payload)
}

func (o *GetInstallActionWorkflowRunOK) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunOK  %+v", 200, o.Payload)
}

func (o *GetInstallActionWorkflowRunOK) GetPayload() *models.AppInstallActionWorkflowRun {
	return o.Payload
}

func (o *GetInstallActionWorkflowRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallActionWorkflowRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallActionWorkflowRunBadRequest creates a GetInstallActionWorkflowRunBadRequest with default headers values
func NewGetInstallActionWorkflowRunBadRequest() *GetInstallActionWorkflowRunBadRequest {
	return &GetInstallActionWorkflowRunBadRequest{}
}

/*
GetInstallActionWorkflowRunBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallActionWorkflowRunBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install action workflow run bad request response has a 2xx status code
func (o *GetInstallActionWorkflowRunBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install action workflow run bad request response has a 3xx status code
func (o *GetInstallActionWorkflowRunBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install action workflow run bad request response has a 4xx status code
func (o *GetInstallActionWorkflowRunBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install action workflow run bad request response has a 5xx status code
func (o *GetInstallActionWorkflowRunBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install action workflow run bad request response a status code equal to that given
func (o *GetInstallActionWorkflowRunBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install action workflow run bad request response
func (o *GetInstallActionWorkflowRunBadRequest) Code() int {
	return 400
}

func (o *GetInstallActionWorkflowRunBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallActionWorkflowRunBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallActionWorkflowRunBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallActionWorkflowRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallActionWorkflowRunUnauthorized creates a GetInstallActionWorkflowRunUnauthorized with default headers values
func NewGetInstallActionWorkflowRunUnauthorized() *GetInstallActionWorkflowRunUnauthorized {
	return &GetInstallActionWorkflowRunUnauthorized{}
}

/*
GetInstallActionWorkflowRunUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallActionWorkflowRunUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install action workflow run unauthorized response has a 2xx status code
func (o *GetInstallActionWorkflowRunUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install action workflow run unauthorized response has a 3xx status code
func (o *GetInstallActionWorkflowRunUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install action workflow run unauthorized response has a 4xx status code
func (o *GetInstallActionWorkflowRunUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install action workflow run unauthorized response has a 5xx status code
func (o *GetInstallActionWorkflowRunUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install action workflow run unauthorized response a status code equal to that given
func (o *GetInstallActionWorkflowRunUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install action workflow run unauthorized response
func (o *GetInstallActionWorkflowRunUnauthorized) Code() int {
	return 401
}

func (o *GetInstallActionWorkflowRunUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallActionWorkflowRunUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallActionWorkflowRunUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallActionWorkflowRunUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallActionWorkflowRunForbidden creates a GetInstallActionWorkflowRunForbidden with default headers values
func NewGetInstallActionWorkflowRunForbidden() *GetInstallActionWorkflowRunForbidden {
	return &GetInstallActionWorkflowRunForbidden{}
}

/*
GetInstallActionWorkflowRunForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallActionWorkflowRunForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install action workflow run forbidden response has a 2xx status code
func (o *GetInstallActionWorkflowRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install action workflow run forbidden response has a 3xx status code
func (o *GetInstallActionWorkflowRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install action workflow run forbidden response has a 4xx status code
func (o *GetInstallActionWorkflowRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install action workflow run forbidden response has a 5xx status code
func (o *GetInstallActionWorkflowRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install action workflow run forbidden response a status code equal to that given
func (o *GetInstallActionWorkflowRunForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install action workflow run forbidden response
func (o *GetInstallActionWorkflowRunForbidden) Code() int {
	return 403
}

func (o *GetInstallActionWorkflowRunForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallActionWorkflowRunForbidden) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallActionWorkflowRunForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallActionWorkflowRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallActionWorkflowRunNotFound creates a GetInstallActionWorkflowRunNotFound with default headers values
func NewGetInstallActionWorkflowRunNotFound() *GetInstallActionWorkflowRunNotFound {
	return &GetInstallActionWorkflowRunNotFound{}
}

/*
GetInstallActionWorkflowRunNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallActionWorkflowRunNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install action workflow run not found response has a 2xx status code
func (o *GetInstallActionWorkflowRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install action workflow run not found response has a 3xx status code
func (o *GetInstallActionWorkflowRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install action workflow run not found response has a 4xx status code
func (o *GetInstallActionWorkflowRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install action workflow run not found response has a 5xx status code
func (o *GetInstallActionWorkflowRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install action workflow run not found response a status code equal to that given
func (o *GetInstallActionWorkflowRunNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install action workflow run not found response
func (o *GetInstallActionWorkflowRunNotFound) Code() int {
	return 404
}

func (o *GetInstallActionWorkflowRunNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallActionWorkflowRunNotFound) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallActionWorkflowRunNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallActionWorkflowRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallActionWorkflowRunInternalServerError creates a GetInstallActionWorkflowRunInternalServerError with default headers values
func NewGetInstallActionWorkflowRunInternalServerError() *GetInstallActionWorkflowRunInternalServerError {
	return &GetInstallActionWorkflowRunInternalServerError{}
}

/*
GetInstallActionWorkflowRunInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallActionWorkflowRunInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install action workflow run internal server error response has a 2xx status code
func (o *GetInstallActionWorkflowRunInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install action workflow run internal server error response has a 3xx status code
func (o *GetInstallActionWorkflowRunInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install action workflow run internal server error response has a 4xx status code
func (o *GetInstallActionWorkflowRunInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install action workflow run internal server error response has a 5xx status code
func (o *GetInstallActionWorkflowRunInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install action workflow run internal server error response a status code equal to that given
func (o *GetInstallActionWorkflowRunInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install action workflow run internal server error response
func (o *GetInstallActionWorkflowRunInternalServerError) Code() int {
	return 500
}

func (o *GetInstallActionWorkflowRunInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallActionWorkflowRunInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/installs/{install_id}/action-workflows/runs/{run_id}][%d] getInstallActionWorkflowRunInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallActionWorkflowRunInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallActionWorkflowRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
