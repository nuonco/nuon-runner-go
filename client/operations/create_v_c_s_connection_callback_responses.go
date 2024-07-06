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

// CreateVCSConnectionCallbackReader is a Reader for the CreateVCSConnectionCallback structure.
type CreateVCSConnectionCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVCSConnectionCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVCSConnectionCallbackCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVCSConnectionCallbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateVCSConnectionCallbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVCSConnectionCallbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVCSConnectionCallbackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVCSConnectionCallbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/vcs/connection-callback] CreateVCSConnectionCallback", response, response.Code())
	}
}

// NewCreateVCSConnectionCallbackCreated creates a CreateVCSConnectionCallbackCreated with default headers values
func NewCreateVCSConnectionCallbackCreated() *CreateVCSConnectionCallbackCreated {
	return &CreateVCSConnectionCallbackCreated{}
}

/*
CreateVCSConnectionCallbackCreated describes a response with status code 201, with default header values.

Created
*/
type CreateVCSConnectionCallbackCreated struct {
	Payload *models.AppVCSConnection
}

// IsSuccess returns true when this create v c s connection callback created response has a 2xx status code
func (o *CreateVCSConnectionCallbackCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create v c s connection callback created response has a 3xx status code
func (o *CreateVCSConnectionCallbackCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection callback created response has a 4xx status code
func (o *CreateVCSConnectionCallbackCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create v c s connection callback created response has a 5xx status code
func (o *CreateVCSConnectionCallbackCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection callback created response a status code equal to that given
func (o *CreateVCSConnectionCallbackCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create v c s connection callback created response
func (o *CreateVCSConnectionCallbackCreated) Code() int {
	return 201
}

func (o *CreateVCSConnectionCallbackCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackCreated %s", 201, payload)
}

func (o *CreateVCSConnectionCallbackCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackCreated %s", 201, payload)
}

func (o *CreateVCSConnectionCallbackCreated) GetPayload() *models.AppVCSConnection {
	return o.Payload
}

func (o *CreateVCSConnectionCallbackCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppVCSConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionCallbackBadRequest creates a CreateVCSConnectionCallbackBadRequest with default headers values
func NewCreateVCSConnectionCallbackBadRequest() *CreateVCSConnectionCallbackBadRequest {
	return &CreateVCSConnectionCallbackBadRequest{}
}

/*
CreateVCSConnectionCallbackBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateVCSConnectionCallbackBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection callback bad request response has a 2xx status code
func (o *CreateVCSConnectionCallbackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection callback bad request response has a 3xx status code
func (o *CreateVCSConnectionCallbackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection callback bad request response has a 4xx status code
func (o *CreateVCSConnectionCallbackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection callback bad request response has a 5xx status code
func (o *CreateVCSConnectionCallbackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection callback bad request response a status code equal to that given
func (o *CreateVCSConnectionCallbackBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create v c s connection callback bad request response
func (o *CreateVCSConnectionCallbackBadRequest) Code() int {
	return 400
}

func (o *CreateVCSConnectionCallbackBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackBadRequest %s", 400, payload)
}

func (o *CreateVCSConnectionCallbackBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackBadRequest %s", 400, payload)
}

func (o *CreateVCSConnectionCallbackBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionCallbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionCallbackUnauthorized creates a CreateVCSConnectionCallbackUnauthorized with default headers values
func NewCreateVCSConnectionCallbackUnauthorized() *CreateVCSConnectionCallbackUnauthorized {
	return &CreateVCSConnectionCallbackUnauthorized{}
}

/*
CreateVCSConnectionCallbackUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateVCSConnectionCallbackUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection callback unauthorized response has a 2xx status code
func (o *CreateVCSConnectionCallbackUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection callback unauthorized response has a 3xx status code
func (o *CreateVCSConnectionCallbackUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection callback unauthorized response has a 4xx status code
func (o *CreateVCSConnectionCallbackUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection callback unauthorized response has a 5xx status code
func (o *CreateVCSConnectionCallbackUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection callback unauthorized response a status code equal to that given
func (o *CreateVCSConnectionCallbackUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create v c s connection callback unauthorized response
func (o *CreateVCSConnectionCallbackUnauthorized) Code() int {
	return 401
}

func (o *CreateVCSConnectionCallbackUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackUnauthorized %s", 401, payload)
}

func (o *CreateVCSConnectionCallbackUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackUnauthorized %s", 401, payload)
}

func (o *CreateVCSConnectionCallbackUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionCallbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionCallbackForbidden creates a CreateVCSConnectionCallbackForbidden with default headers values
func NewCreateVCSConnectionCallbackForbidden() *CreateVCSConnectionCallbackForbidden {
	return &CreateVCSConnectionCallbackForbidden{}
}

/*
CreateVCSConnectionCallbackForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateVCSConnectionCallbackForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection callback forbidden response has a 2xx status code
func (o *CreateVCSConnectionCallbackForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection callback forbidden response has a 3xx status code
func (o *CreateVCSConnectionCallbackForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection callback forbidden response has a 4xx status code
func (o *CreateVCSConnectionCallbackForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection callback forbidden response has a 5xx status code
func (o *CreateVCSConnectionCallbackForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection callback forbidden response a status code equal to that given
func (o *CreateVCSConnectionCallbackForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create v c s connection callback forbidden response
func (o *CreateVCSConnectionCallbackForbidden) Code() int {
	return 403
}

func (o *CreateVCSConnectionCallbackForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackForbidden %s", 403, payload)
}

func (o *CreateVCSConnectionCallbackForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackForbidden %s", 403, payload)
}

func (o *CreateVCSConnectionCallbackForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionCallbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionCallbackNotFound creates a CreateVCSConnectionCallbackNotFound with default headers values
func NewCreateVCSConnectionCallbackNotFound() *CreateVCSConnectionCallbackNotFound {
	return &CreateVCSConnectionCallbackNotFound{}
}

/*
CreateVCSConnectionCallbackNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateVCSConnectionCallbackNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection callback not found response has a 2xx status code
func (o *CreateVCSConnectionCallbackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection callback not found response has a 3xx status code
func (o *CreateVCSConnectionCallbackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection callback not found response has a 4xx status code
func (o *CreateVCSConnectionCallbackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create v c s connection callback not found response has a 5xx status code
func (o *CreateVCSConnectionCallbackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create v c s connection callback not found response a status code equal to that given
func (o *CreateVCSConnectionCallbackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create v c s connection callback not found response
func (o *CreateVCSConnectionCallbackNotFound) Code() int {
	return 404
}

func (o *CreateVCSConnectionCallbackNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackNotFound %s", 404, payload)
}

func (o *CreateVCSConnectionCallbackNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackNotFound %s", 404, payload)
}

func (o *CreateVCSConnectionCallbackNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionCallbackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVCSConnectionCallbackInternalServerError creates a CreateVCSConnectionCallbackInternalServerError with default headers values
func NewCreateVCSConnectionCallbackInternalServerError() *CreateVCSConnectionCallbackInternalServerError {
	return &CreateVCSConnectionCallbackInternalServerError{}
}

/*
CreateVCSConnectionCallbackInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateVCSConnectionCallbackInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this create v c s connection callback internal server error response has a 2xx status code
func (o *CreateVCSConnectionCallbackInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create v c s connection callback internal server error response has a 3xx status code
func (o *CreateVCSConnectionCallbackInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create v c s connection callback internal server error response has a 4xx status code
func (o *CreateVCSConnectionCallbackInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create v c s connection callback internal server error response has a 5xx status code
func (o *CreateVCSConnectionCallbackInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create v c s connection callback internal server error response a status code equal to that given
func (o *CreateVCSConnectionCallbackInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create v c s connection callback internal server error response
func (o *CreateVCSConnectionCallbackInternalServerError) Code() int {
	return 500
}

func (o *CreateVCSConnectionCallbackInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackInternalServerError %s", 500, payload)
}

func (o *CreateVCSConnectionCallbackInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/vcs/connection-callback][%d] createVCSConnectionCallbackInternalServerError %s", 500, payload)
}

func (o *CreateVCSConnectionCallbackInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *CreateVCSConnectionCallbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}