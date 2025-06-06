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

// GetV1InstallInstallIDOwnerPlansOwnerIDReader is a Reader for the GetV1InstallInstallIDOwnerPlansOwnerID structure.
type GetV1InstallInstallIDOwnerPlansOwnerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGetV1InstallInstallIDOwnerPlansOwnerIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV1InstallInstallIDOwnerPlansOwnerIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetV1InstallInstallIDOwnerPlansOwnerIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV1InstallInstallIDOwnerPlansOwnerIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/install/{install_id}/owner-plans/{owner_id}] GetV1InstallInstallIDOwnerPlansOwnerID", response, response.Code())
	}
}

// NewGetV1InstallInstallIDOwnerPlansOwnerIDCreated creates a GetV1InstallInstallIDOwnerPlansOwnerIDCreated with default headers values
func NewGetV1InstallInstallIDOwnerPlansOwnerIDCreated() *GetV1InstallInstallIDOwnerPlansOwnerIDCreated {
	return &GetV1InstallInstallIDOwnerPlansOwnerIDCreated{}
}

/*
GetV1InstallInstallIDOwnerPlansOwnerIDCreated describes a response with status code 201, with default header values.

Created
*/
type GetV1InstallInstallIDOwnerPlansOwnerIDCreated struct {
	Payload *models.AppInstallPlan
}

// IsSuccess returns true when this get v1 install install Id owner plans owner Id created response has a 2xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 install install Id owner plans owner Id created response has a 3xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 install install Id owner plans owner Id created response has a 4xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 install install Id owner plans owner Id created response has a 5xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 install install Id owner plans owner Id created response a status code equal to that given
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the get v1 install install Id owner plans owner Id created response
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) Code() int {
	return 201
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) Error() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdCreated  %+v", 201, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) String() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdCreated  %+v", 201, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) GetPayload() *models.AppInstallPlan {
	return o.Payload
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstallPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallInstallIDOwnerPlansOwnerIDBadRequest creates a GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest with default headers values
func NewGetV1InstallInstallIDOwnerPlansOwnerIDBadRequest() *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest {
	return &GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest{}
}

/*
GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 install install Id owner plans owner Id bad request response has a 2xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 install install Id owner plans owner Id bad request response has a 3xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 install install Id owner plans owner Id bad request response has a 4xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 install install Id owner plans owner Id bad request response has a 5xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 install install Id owner plans owner Id bad request response a status code equal to that given
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get v1 install install Id owner plans owner Id bad request response
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) Code() int {
	return 400
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized creates a GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized with default headers values
func NewGetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized() *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized {
	return &GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized{}
}

/*
GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 install install Id owner plans owner Id unauthorized response has a 2xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 install install Id owner plans owner Id unauthorized response has a 3xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 install install Id owner plans owner Id unauthorized response has a 4xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 install install Id owner plans owner Id unauthorized response has a 5xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 install install Id owner plans owner Id unauthorized response a status code equal to that given
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get v1 install install Id owner plans owner Id unauthorized response
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) Code() int {
	return 401
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallInstallIDOwnerPlansOwnerIDForbidden creates a GetV1InstallInstallIDOwnerPlansOwnerIDForbidden with default headers values
func NewGetV1InstallInstallIDOwnerPlansOwnerIDForbidden() *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden {
	return &GetV1InstallInstallIDOwnerPlansOwnerIDForbidden{}
}

/*
GetV1InstallInstallIDOwnerPlansOwnerIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetV1InstallInstallIDOwnerPlansOwnerIDForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 install install Id owner plans owner Id forbidden response has a 2xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 install install Id owner plans owner Id forbidden response has a 3xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 install install Id owner plans owner Id forbidden response has a 4xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 install install Id owner plans owner Id forbidden response has a 5xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 install install Id owner plans owner Id forbidden response a status code equal to that given
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get v1 install install Id owner plans owner Id forbidden response
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) Code() int {
	return 403
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdForbidden  %+v", 403, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallInstallIDOwnerPlansOwnerIDNotFound creates a GetV1InstallInstallIDOwnerPlansOwnerIDNotFound with default headers values
func NewGetV1InstallInstallIDOwnerPlansOwnerIDNotFound() *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound {
	return &GetV1InstallInstallIDOwnerPlansOwnerIDNotFound{}
}

/*
GetV1InstallInstallIDOwnerPlansOwnerIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetV1InstallInstallIDOwnerPlansOwnerIDNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 install install Id owner plans owner Id not found response has a 2xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 install install Id owner plans owner Id not found response has a 3xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 install install Id owner plans owner Id not found response has a 4xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get v1 install install Id owner plans owner Id not found response has a 5xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 install install Id owner plans owner Id not found response a status code equal to that given
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get v1 install install Id owner plans owner Id not found response
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) Code() int {
	return 404
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdNotFound  %+v", 404, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError creates a GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError with default headers values
func NewGetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError() *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError {
	return &GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError{}
}

/*
GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get v1 install install Id owner plans owner Id internal server error response has a 2xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get v1 install install Id owner plans owner Id internal server error response has a 3xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 install install Id owner plans owner Id internal server error response has a 4xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 install install Id owner plans owner Id internal server error response has a 5xx status code
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get v1 install install Id owner plans owner Id internal server error response a status code equal to that given
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get v1 install install Id owner plans owner Id internal server error response
func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) Code() int {
	return 500
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/install/{install_id}/owner-plans/{owner_id}][%d] getV1InstallInstallIdOwnerPlansOwnerIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetV1InstallInstallIDOwnerPlansOwnerIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
