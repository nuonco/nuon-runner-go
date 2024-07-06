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

// GetInstallLatestDeployReader is a Reader for the GetInstallLatestDeploy structure.
type GetInstallLatestDeployReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallLatestDeployReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallLatestDeployOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallLatestDeployBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallLatestDeployUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallLatestDeployForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallLatestDeployNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallLatestDeployInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/installs/{install_id}/deploys/latest] GetInstallLatestDeploy", response, response.Code())
	}
}

// NewGetInstallLatestDeployOK creates a GetInstallLatestDeployOK with default headers values
func NewGetInstallLatestDeployOK() *GetInstallLatestDeployOK {
	return &GetInstallLatestDeployOK{}
}

/*
GetInstallLatestDeployOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallLatestDeployOK struct {
	Payload *models.AppInstallDeploy
}

// IsSuccess returns true when this get install latest deploy o k response has a 2xx status code
func (o *GetInstallLatestDeployOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install latest deploy o k response has a 3xx status code
func (o *GetInstallLatestDeployOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install latest deploy o k response has a 4xx status code
func (o *GetInstallLatestDeployOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install latest deploy o k response has a 5xx status code
func (o *GetInstallLatestDeployOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install latest deploy o k response a status code equal to that given
func (o *GetInstallLatestDeployOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install latest deploy o k response
func (o *GetInstallLatestDeployOK) Code() int {
	return 200
}

func (o *GetInstallLatestDeployOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployOK %s", 200, payload)
}

func (o *GetInstallLatestDeployOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployOK %s", 200, payload)
}

func (o *GetInstallLatestDeployOK) GetPayload() *models.AppInstallDeploy {
	return o.Payload
}

func (o *GetInstallLatestDeployOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallDeploy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallLatestDeployBadRequest creates a GetInstallLatestDeployBadRequest with default headers values
func NewGetInstallLatestDeployBadRequest() *GetInstallLatestDeployBadRequest {
	return &GetInstallLatestDeployBadRequest{}
}

/*
GetInstallLatestDeployBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallLatestDeployBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install latest deploy bad request response has a 2xx status code
func (o *GetInstallLatestDeployBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install latest deploy bad request response has a 3xx status code
func (o *GetInstallLatestDeployBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install latest deploy bad request response has a 4xx status code
func (o *GetInstallLatestDeployBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install latest deploy bad request response has a 5xx status code
func (o *GetInstallLatestDeployBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install latest deploy bad request response a status code equal to that given
func (o *GetInstallLatestDeployBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install latest deploy bad request response
func (o *GetInstallLatestDeployBadRequest) Code() int {
	return 400
}

func (o *GetInstallLatestDeployBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployBadRequest %s", 400, payload)
}

func (o *GetInstallLatestDeployBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployBadRequest %s", 400, payload)
}

func (o *GetInstallLatestDeployBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallLatestDeployBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallLatestDeployUnauthorized creates a GetInstallLatestDeployUnauthorized with default headers values
func NewGetInstallLatestDeployUnauthorized() *GetInstallLatestDeployUnauthorized {
	return &GetInstallLatestDeployUnauthorized{}
}

/*
GetInstallLatestDeployUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallLatestDeployUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install latest deploy unauthorized response has a 2xx status code
func (o *GetInstallLatestDeployUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install latest deploy unauthorized response has a 3xx status code
func (o *GetInstallLatestDeployUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install latest deploy unauthorized response has a 4xx status code
func (o *GetInstallLatestDeployUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install latest deploy unauthorized response has a 5xx status code
func (o *GetInstallLatestDeployUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install latest deploy unauthorized response a status code equal to that given
func (o *GetInstallLatestDeployUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install latest deploy unauthorized response
func (o *GetInstallLatestDeployUnauthorized) Code() int {
	return 401
}

func (o *GetInstallLatestDeployUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployUnauthorized %s", 401, payload)
}

func (o *GetInstallLatestDeployUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployUnauthorized %s", 401, payload)
}

func (o *GetInstallLatestDeployUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallLatestDeployUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallLatestDeployForbidden creates a GetInstallLatestDeployForbidden with default headers values
func NewGetInstallLatestDeployForbidden() *GetInstallLatestDeployForbidden {
	return &GetInstallLatestDeployForbidden{}
}

/*
GetInstallLatestDeployForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallLatestDeployForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install latest deploy forbidden response has a 2xx status code
func (o *GetInstallLatestDeployForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install latest deploy forbidden response has a 3xx status code
func (o *GetInstallLatestDeployForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install latest deploy forbidden response has a 4xx status code
func (o *GetInstallLatestDeployForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install latest deploy forbidden response has a 5xx status code
func (o *GetInstallLatestDeployForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install latest deploy forbidden response a status code equal to that given
func (o *GetInstallLatestDeployForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install latest deploy forbidden response
func (o *GetInstallLatestDeployForbidden) Code() int {
	return 403
}

func (o *GetInstallLatestDeployForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployForbidden %s", 403, payload)
}

func (o *GetInstallLatestDeployForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployForbidden %s", 403, payload)
}

func (o *GetInstallLatestDeployForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallLatestDeployForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallLatestDeployNotFound creates a GetInstallLatestDeployNotFound with default headers values
func NewGetInstallLatestDeployNotFound() *GetInstallLatestDeployNotFound {
	return &GetInstallLatestDeployNotFound{}
}

/*
GetInstallLatestDeployNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallLatestDeployNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install latest deploy not found response has a 2xx status code
func (o *GetInstallLatestDeployNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install latest deploy not found response has a 3xx status code
func (o *GetInstallLatestDeployNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install latest deploy not found response has a 4xx status code
func (o *GetInstallLatestDeployNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install latest deploy not found response has a 5xx status code
func (o *GetInstallLatestDeployNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install latest deploy not found response a status code equal to that given
func (o *GetInstallLatestDeployNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install latest deploy not found response
func (o *GetInstallLatestDeployNotFound) Code() int {
	return 404
}

func (o *GetInstallLatestDeployNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployNotFound %s", 404, payload)
}

func (o *GetInstallLatestDeployNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployNotFound %s", 404, payload)
}

func (o *GetInstallLatestDeployNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallLatestDeployNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallLatestDeployInternalServerError creates a GetInstallLatestDeployInternalServerError with default headers values
func NewGetInstallLatestDeployInternalServerError() *GetInstallLatestDeployInternalServerError {
	return &GetInstallLatestDeployInternalServerError{}
}

/*
GetInstallLatestDeployInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallLatestDeployInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install latest deploy internal server error response has a 2xx status code
func (o *GetInstallLatestDeployInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install latest deploy internal server error response has a 3xx status code
func (o *GetInstallLatestDeployInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install latest deploy internal server error response has a 4xx status code
func (o *GetInstallLatestDeployInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install latest deploy internal server error response has a 5xx status code
func (o *GetInstallLatestDeployInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install latest deploy internal server error response a status code equal to that given
func (o *GetInstallLatestDeployInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install latest deploy internal server error response
func (o *GetInstallLatestDeployInternalServerError) Code() int {
	return 500
}

func (o *GetInstallLatestDeployInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployInternalServerError %s", 500, payload)
}

func (o *GetInstallLatestDeployInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/installs/{install_id}/deploys/latest][%d] getInstallLatestDeployInternalServerError %s", 500, payload)
}

func (o *GetInstallLatestDeployInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallLatestDeployInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}