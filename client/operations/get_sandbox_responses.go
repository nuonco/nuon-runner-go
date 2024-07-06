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

	"github.com/nuonco/nuon-go/models"
)

// GetSandboxReader is a Reader for the GetSandbox structure.
type GetSandboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSandboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSandboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSandboxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSandboxUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSandboxForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSandboxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSandboxInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/sandboxes/{sandbox_id}] GetSandbox", response, response.Code())
	}
}

// NewGetSandboxOK creates a GetSandboxOK with default headers values
func NewGetSandboxOK() *GetSandboxOK {
	return &GetSandboxOK{}
}

/*
GetSandboxOK describes a response with status code 200, with default header values.

OK
*/
type GetSandboxOK struct {
	Payload *models.AppSandbox
}

// IsSuccess returns true when this get sandbox o k response has a 2xx status code
func (o *GetSandboxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sandbox o k response has a 3xx status code
func (o *GetSandboxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sandbox o k response has a 4xx status code
func (o *GetSandboxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sandbox o k response has a 5xx status code
func (o *GetSandboxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sandbox o k response a status code equal to that given
func (o *GetSandboxOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sandbox o k response
func (o *GetSandboxOK) Code() int {
	return 200
}

func (o *GetSandboxOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxOK %s", 200, payload)
}

func (o *GetSandboxOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxOK %s", 200, payload)
}

func (o *GetSandboxOK) GetPayload() *models.AppSandbox {
	return o.Payload
}

func (o *GetSandboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppSandbox)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSandboxBadRequest creates a GetSandboxBadRequest with default headers values
func NewGetSandboxBadRequest() *GetSandboxBadRequest {
	return &GetSandboxBadRequest{}
}

/*
GetSandboxBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetSandboxBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get sandbox bad request response has a 2xx status code
func (o *GetSandboxBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sandbox bad request response has a 3xx status code
func (o *GetSandboxBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sandbox bad request response has a 4xx status code
func (o *GetSandboxBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sandbox bad request response has a 5xx status code
func (o *GetSandboxBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get sandbox bad request response a status code equal to that given
func (o *GetSandboxBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get sandbox bad request response
func (o *GetSandboxBadRequest) Code() int {
	return 400
}

func (o *GetSandboxBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxBadRequest %s", 400, payload)
}

func (o *GetSandboxBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxBadRequest %s", 400, payload)
}

func (o *GetSandboxBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetSandboxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSandboxUnauthorized creates a GetSandboxUnauthorized with default headers values
func NewGetSandboxUnauthorized() *GetSandboxUnauthorized {
	return &GetSandboxUnauthorized{}
}

/*
GetSandboxUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSandboxUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get sandbox unauthorized response has a 2xx status code
func (o *GetSandboxUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sandbox unauthorized response has a 3xx status code
func (o *GetSandboxUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sandbox unauthorized response has a 4xx status code
func (o *GetSandboxUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sandbox unauthorized response has a 5xx status code
func (o *GetSandboxUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get sandbox unauthorized response a status code equal to that given
func (o *GetSandboxUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get sandbox unauthorized response
func (o *GetSandboxUnauthorized) Code() int {
	return 401
}

func (o *GetSandboxUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxUnauthorized %s", 401, payload)
}

func (o *GetSandboxUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxUnauthorized %s", 401, payload)
}

func (o *GetSandboxUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetSandboxUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSandboxForbidden creates a GetSandboxForbidden with default headers values
func NewGetSandboxForbidden() *GetSandboxForbidden {
	return &GetSandboxForbidden{}
}

/*
GetSandboxForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSandboxForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get sandbox forbidden response has a 2xx status code
func (o *GetSandboxForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sandbox forbidden response has a 3xx status code
func (o *GetSandboxForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sandbox forbidden response has a 4xx status code
func (o *GetSandboxForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sandbox forbidden response has a 5xx status code
func (o *GetSandboxForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get sandbox forbidden response a status code equal to that given
func (o *GetSandboxForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get sandbox forbidden response
func (o *GetSandboxForbidden) Code() int {
	return 403
}

func (o *GetSandboxForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxForbidden %s", 403, payload)
}

func (o *GetSandboxForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxForbidden %s", 403, payload)
}

func (o *GetSandboxForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetSandboxForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSandboxNotFound creates a GetSandboxNotFound with default headers values
func NewGetSandboxNotFound() *GetSandboxNotFound {
	return &GetSandboxNotFound{}
}

/*
GetSandboxNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSandboxNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get sandbox not found response has a 2xx status code
func (o *GetSandboxNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sandbox not found response has a 3xx status code
func (o *GetSandboxNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sandbox not found response has a 4xx status code
func (o *GetSandboxNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get sandbox not found response has a 5xx status code
func (o *GetSandboxNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get sandbox not found response a status code equal to that given
func (o *GetSandboxNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get sandbox not found response
func (o *GetSandboxNotFound) Code() int {
	return 404
}

func (o *GetSandboxNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxNotFound %s", 404, payload)
}

func (o *GetSandboxNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxNotFound %s", 404, payload)
}

func (o *GetSandboxNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetSandboxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSandboxInternalServerError creates a GetSandboxInternalServerError with default headers values
func NewGetSandboxInternalServerError() *GetSandboxInternalServerError {
	return &GetSandboxInternalServerError{}
}

/*
GetSandboxInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSandboxInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get sandbox internal server error response has a 2xx status code
func (o *GetSandboxInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get sandbox internal server error response has a 3xx status code
func (o *GetSandboxInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sandbox internal server error response has a 4xx status code
func (o *GetSandboxInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sandbox internal server error response has a 5xx status code
func (o *GetSandboxInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get sandbox internal server error response a status code equal to that given
func (o *GetSandboxInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get sandbox internal server error response
func (o *GetSandboxInternalServerError) Code() int {
	return 500
}

func (o *GetSandboxInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxInternalServerError %s", 500, payload)
}

func (o *GetSandboxInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/sandboxes/{sandbox_id}][%d] getSandboxInternalServerError %s", 500, payload)
}

func (o *GetSandboxInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetSandboxInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}