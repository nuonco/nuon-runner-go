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

// GetOrgReader is a Reader for the GetOrg structure.
type GetOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/orgs/current] GetOrg", response, response.Code())
	}
}

// NewGetOrgOK creates a GetOrgOK with default headers values
func NewGetOrgOK() *GetOrgOK {
	return &GetOrgOK{}
}

/*
GetOrgOK describes a response with status code 200, with default header values.

OK
*/
type GetOrgOK struct {
	Payload *models.AppOrg
}

// IsSuccess returns true when this get org o k response has a 2xx status code
func (o *GetOrgOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org o k response has a 3xx status code
func (o *GetOrgOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org o k response has a 4xx status code
func (o *GetOrgOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org o k response has a 5xx status code
func (o *GetOrgOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org o k response a status code equal to that given
func (o *GetOrgOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org o k response
func (o *GetOrgOK) Code() int {
	return 200
}

func (o *GetOrgOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgOK %s", 200, payload)
}

func (o *GetOrgOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgOK %s", 200, payload)
}

func (o *GetOrgOK) GetPayload() *models.AppOrg {
	return o.Payload
}

func (o *GetOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppOrg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgBadRequest creates a GetOrgBadRequest with default headers values
func NewGetOrgBadRequest() *GetOrgBadRequest {
	return &GetOrgBadRequest{}
}

/*
GetOrgBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetOrgBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org bad request response has a 2xx status code
func (o *GetOrgBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org bad request response has a 3xx status code
func (o *GetOrgBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org bad request response has a 4xx status code
func (o *GetOrgBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org bad request response has a 5xx status code
func (o *GetOrgBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get org bad request response a status code equal to that given
func (o *GetOrgBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get org bad request response
func (o *GetOrgBadRequest) Code() int {
	return 400
}

func (o *GetOrgBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgBadRequest %s", 400, payload)
}

func (o *GetOrgBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgBadRequest %s", 400, payload)
}

func (o *GetOrgBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgUnauthorized creates a GetOrgUnauthorized with default headers values
func NewGetOrgUnauthorized() *GetOrgUnauthorized {
	return &GetOrgUnauthorized{}
}

/*
GetOrgUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrgUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org unauthorized response has a 2xx status code
func (o *GetOrgUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org unauthorized response has a 3xx status code
func (o *GetOrgUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org unauthorized response has a 4xx status code
func (o *GetOrgUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org unauthorized response has a 5xx status code
func (o *GetOrgUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get org unauthorized response a status code equal to that given
func (o *GetOrgUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get org unauthorized response
func (o *GetOrgUnauthorized) Code() int {
	return 401
}

func (o *GetOrgUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgUnauthorized %s", 401, payload)
}

func (o *GetOrgUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgUnauthorized %s", 401, payload)
}

func (o *GetOrgUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgForbidden creates a GetOrgForbidden with default headers values
func NewGetOrgForbidden() *GetOrgForbidden {
	return &GetOrgForbidden{}
}

/*
GetOrgForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrgForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org forbidden response has a 2xx status code
func (o *GetOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org forbidden response has a 3xx status code
func (o *GetOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org forbidden response has a 4xx status code
func (o *GetOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org forbidden response has a 5xx status code
func (o *GetOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org forbidden response a status code equal to that given
func (o *GetOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org forbidden response
func (o *GetOrgForbidden) Code() int {
	return 403
}

func (o *GetOrgForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgForbidden %s", 403, payload)
}

func (o *GetOrgForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgForbidden %s", 403, payload)
}

func (o *GetOrgForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgNotFound creates a GetOrgNotFound with default headers values
func NewGetOrgNotFound() *GetOrgNotFound {
	return &GetOrgNotFound{}
}

/*
GetOrgNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrgNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org not found response has a 2xx status code
func (o *GetOrgNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org not found response has a 3xx status code
func (o *GetOrgNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org not found response has a 4xx status code
func (o *GetOrgNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org not found response has a 5xx status code
func (o *GetOrgNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get org not found response a status code equal to that given
func (o *GetOrgNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get org not found response
func (o *GetOrgNotFound) Code() int {
	return 404
}

func (o *GetOrgNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgNotFound %s", 404, payload)
}

func (o *GetOrgNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgNotFound %s", 404, payload)
}

func (o *GetOrgNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgInternalServerError creates a GetOrgInternalServerError with default headers values
func NewGetOrgInternalServerError() *GetOrgInternalServerError {
	return &GetOrgInternalServerError{}
}

/*
GetOrgInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOrgInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get org internal server error response has a 2xx status code
func (o *GetOrgInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org internal server error response has a 3xx status code
func (o *GetOrgInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org internal server error response has a 4xx status code
func (o *GetOrgInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org internal server error response has a 5xx status code
func (o *GetOrgInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org internal server error response a status code equal to that given
func (o *GetOrgInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org internal server error response
func (o *GetOrgInternalServerError) Code() int {
	return 500
}

func (o *GetOrgInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgInternalServerError %s", 500, payload)
}

func (o *GetOrgInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/orgs/current][%d] getOrgInternalServerError %s", 500, payload)
}

func (o *GetOrgInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}