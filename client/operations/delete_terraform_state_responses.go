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

// DeleteTerraformStateReader is a Reader for the DeleteTerraformState structure.
type DeleteTerraformStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTerraformStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTerraformStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTerraformStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTerraformStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTerraformStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTerraformStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTerraformStateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /v1/terraform-backend] DeleteTerraformState", response, response.Code())
	}
}

// NewDeleteTerraformStateOK creates a DeleteTerraformStateOK with default headers values
func NewDeleteTerraformStateOK() *DeleteTerraformStateOK {
	return &DeleteTerraformStateOK{}
}

/*
DeleteTerraformStateOK describes a response with status code 200, with default header values.

OK
*/
type DeleteTerraformStateOK struct {
	Payload string
}

// IsSuccess returns true when this delete terraform state o k response has a 2xx status code
func (o *DeleteTerraformStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete terraform state o k response has a 3xx status code
func (o *DeleteTerraformStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform state o k response has a 4xx status code
func (o *DeleteTerraformStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete terraform state o k response has a 5xx status code
func (o *DeleteTerraformStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform state o k response a status code equal to that given
func (o *DeleteTerraformStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete terraform state o k response
func (o *DeleteTerraformStateOK) Code() int {
	return 200
}

func (o *DeleteTerraformStateOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateOK  %+v", 200, o.Payload)
}

func (o *DeleteTerraformStateOK) String() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateOK  %+v", 200, o.Payload)
}

func (o *DeleteTerraformStateOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteTerraformStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerraformStateBadRequest creates a DeleteTerraformStateBadRequest with default headers values
func NewDeleteTerraformStateBadRequest() *DeleteTerraformStateBadRequest {
	return &DeleteTerraformStateBadRequest{}
}

/*
DeleteTerraformStateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteTerraformStateBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete terraform state bad request response has a 2xx status code
func (o *DeleteTerraformStateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform state bad request response has a 3xx status code
func (o *DeleteTerraformStateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform state bad request response has a 4xx status code
func (o *DeleteTerraformStateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete terraform state bad request response has a 5xx status code
func (o *DeleteTerraformStateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform state bad request response a status code equal to that given
func (o *DeleteTerraformStateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete terraform state bad request response
func (o *DeleteTerraformStateBadRequest) Code() int {
	return 400
}

func (o *DeleteTerraformStateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTerraformStateBadRequest) String() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTerraformStateBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteTerraformStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerraformStateUnauthorized creates a DeleteTerraformStateUnauthorized with default headers values
func NewDeleteTerraformStateUnauthorized() *DeleteTerraformStateUnauthorized {
	return &DeleteTerraformStateUnauthorized{}
}

/*
DeleteTerraformStateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTerraformStateUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete terraform state unauthorized response has a 2xx status code
func (o *DeleteTerraformStateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform state unauthorized response has a 3xx status code
func (o *DeleteTerraformStateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform state unauthorized response has a 4xx status code
func (o *DeleteTerraformStateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete terraform state unauthorized response has a 5xx status code
func (o *DeleteTerraformStateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform state unauthorized response a status code equal to that given
func (o *DeleteTerraformStateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete terraform state unauthorized response
func (o *DeleteTerraformStateUnauthorized) Code() int {
	return 401
}

func (o *DeleteTerraformStateUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTerraformStateUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTerraformStateUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteTerraformStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerraformStateForbidden creates a DeleteTerraformStateForbidden with default headers values
func NewDeleteTerraformStateForbidden() *DeleteTerraformStateForbidden {
	return &DeleteTerraformStateForbidden{}
}

/*
DeleteTerraformStateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTerraformStateForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete terraform state forbidden response has a 2xx status code
func (o *DeleteTerraformStateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform state forbidden response has a 3xx status code
func (o *DeleteTerraformStateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform state forbidden response has a 4xx status code
func (o *DeleteTerraformStateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete terraform state forbidden response has a 5xx status code
func (o *DeleteTerraformStateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform state forbidden response a status code equal to that given
func (o *DeleteTerraformStateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete terraform state forbidden response
func (o *DeleteTerraformStateForbidden) Code() int {
	return 403
}

func (o *DeleteTerraformStateForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTerraformStateForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTerraformStateForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteTerraformStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerraformStateNotFound creates a DeleteTerraformStateNotFound with default headers values
func NewDeleteTerraformStateNotFound() *DeleteTerraformStateNotFound {
	return &DeleteTerraformStateNotFound{}
}

/*
DeleteTerraformStateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteTerraformStateNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete terraform state not found response has a 2xx status code
func (o *DeleteTerraformStateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform state not found response has a 3xx status code
func (o *DeleteTerraformStateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform state not found response has a 4xx status code
func (o *DeleteTerraformStateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete terraform state not found response has a 5xx status code
func (o *DeleteTerraformStateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete terraform state not found response a status code equal to that given
func (o *DeleteTerraformStateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete terraform state not found response
func (o *DeleteTerraformStateNotFound) Code() int {
	return 404
}

func (o *DeleteTerraformStateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTerraformStateNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTerraformStateNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteTerraformStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTerraformStateInternalServerError creates a DeleteTerraformStateInternalServerError with default headers values
func NewDeleteTerraformStateInternalServerError() *DeleteTerraformStateInternalServerError {
	return &DeleteTerraformStateInternalServerError{}
}

/*
DeleteTerraformStateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteTerraformStateInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this delete terraform state internal server error response has a 2xx status code
func (o *DeleteTerraformStateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete terraform state internal server error response has a 3xx status code
func (o *DeleteTerraformStateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete terraform state internal server error response has a 4xx status code
func (o *DeleteTerraformStateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete terraform state internal server error response has a 5xx status code
func (o *DeleteTerraformStateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete terraform state internal server error response a status code equal to that given
func (o *DeleteTerraformStateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete terraform state internal server error response
func (o *DeleteTerraformStateInternalServerError) Code() int {
	return 500
}

func (o *DeleteTerraformStateInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTerraformStateInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/terraform-backend][%d] deleteTerraformStateInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTerraformStateInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *DeleteTerraformStateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
