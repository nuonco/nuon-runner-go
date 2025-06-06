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

// GetHelmReleasesReader is a Reader for the GetHelmReleases structure.
type GetHelmReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHelmReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHelmReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHelmReleasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetHelmReleasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetHelmReleasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetHelmReleasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHelmReleasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}] GetHelmReleases", response, response.Code())
	}
}

// NewGetHelmReleasesOK creates a GetHelmReleasesOK with default headers values
func NewGetHelmReleasesOK() *GetHelmReleasesOK {
	return &GetHelmReleasesOK{}
}

/*
GetHelmReleasesOK describes a response with status code 200, with default header values.

OK
*/
type GetHelmReleasesOK struct {
	Payload []*models.HelmRelease
}

// IsSuccess returns true when this get helm releases o k response has a 2xx status code
func (o *GetHelmReleasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get helm releases o k response has a 3xx status code
func (o *GetHelmReleasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm releases o k response has a 4xx status code
func (o *GetHelmReleasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get helm releases o k response has a 5xx status code
func (o *GetHelmReleasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm releases o k response a status code equal to that given
func (o *GetHelmReleasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get helm releases o k response
func (o *GetHelmReleasesOK) Code() int {
	return 200
}

func (o *GetHelmReleasesOK) Error() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesOK  %+v", 200, o.Payload)
}

func (o *GetHelmReleasesOK) String() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesOK  %+v", 200, o.Payload)
}

func (o *GetHelmReleasesOK) GetPayload() []*models.HelmRelease {
	return o.Payload
}

func (o *GetHelmReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmReleasesBadRequest creates a GetHelmReleasesBadRequest with default headers values
func NewGetHelmReleasesBadRequest() *GetHelmReleasesBadRequest {
	return &GetHelmReleasesBadRequest{}
}

/*
GetHelmReleasesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetHelmReleasesBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get helm releases bad request response has a 2xx status code
func (o *GetHelmReleasesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm releases bad request response has a 3xx status code
func (o *GetHelmReleasesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm releases bad request response has a 4xx status code
func (o *GetHelmReleasesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm releases bad request response has a 5xx status code
func (o *GetHelmReleasesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm releases bad request response a status code equal to that given
func (o *GetHelmReleasesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get helm releases bad request response
func (o *GetHelmReleasesBadRequest) Code() int {
	return 400
}

func (o *GetHelmReleasesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesBadRequest  %+v", 400, o.Payload)
}

func (o *GetHelmReleasesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesBadRequest  %+v", 400, o.Payload)
}

func (o *GetHelmReleasesBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetHelmReleasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmReleasesUnauthorized creates a GetHelmReleasesUnauthorized with default headers values
func NewGetHelmReleasesUnauthorized() *GetHelmReleasesUnauthorized {
	return &GetHelmReleasesUnauthorized{}
}

/*
GetHelmReleasesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetHelmReleasesUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get helm releases unauthorized response has a 2xx status code
func (o *GetHelmReleasesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm releases unauthorized response has a 3xx status code
func (o *GetHelmReleasesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm releases unauthorized response has a 4xx status code
func (o *GetHelmReleasesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm releases unauthorized response has a 5xx status code
func (o *GetHelmReleasesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm releases unauthorized response a status code equal to that given
func (o *GetHelmReleasesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get helm releases unauthorized response
func (o *GetHelmReleasesUnauthorized) Code() int {
	return 401
}

func (o *GetHelmReleasesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHelmReleasesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetHelmReleasesUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetHelmReleasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmReleasesForbidden creates a GetHelmReleasesForbidden with default headers values
func NewGetHelmReleasesForbidden() *GetHelmReleasesForbidden {
	return &GetHelmReleasesForbidden{}
}

/*
GetHelmReleasesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetHelmReleasesForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get helm releases forbidden response has a 2xx status code
func (o *GetHelmReleasesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm releases forbidden response has a 3xx status code
func (o *GetHelmReleasesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm releases forbidden response has a 4xx status code
func (o *GetHelmReleasesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm releases forbidden response has a 5xx status code
func (o *GetHelmReleasesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm releases forbidden response a status code equal to that given
func (o *GetHelmReleasesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get helm releases forbidden response
func (o *GetHelmReleasesForbidden) Code() int {
	return 403
}

func (o *GetHelmReleasesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesForbidden  %+v", 403, o.Payload)
}

func (o *GetHelmReleasesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesForbidden  %+v", 403, o.Payload)
}

func (o *GetHelmReleasesForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetHelmReleasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmReleasesNotFound creates a GetHelmReleasesNotFound with default headers values
func NewGetHelmReleasesNotFound() *GetHelmReleasesNotFound {
	return &GetHelmReleasesNotFound{}
}

/*
GetHelmReleasesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetHelmReleasesNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get helm releases not found response has a 2xx status code
func (o *GetHelmReleasesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm releases not found response has a 3xx status code
func (o *GetHelmReleasesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm releases not found response has a 4xx status code
func (o *GetHelmReleasesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get helm releases not found response has a 5xx status code
func (o *GetHelmReleasesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get helm releases not found response a status code equal to that given
func (o *GetHelmReleasesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get helm releases not found response
func (o *GetHelmReleasesNotFound) Code() int {
	return 404
}

func (o *GetHelmReleasesNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesNotFound  %+v", 404, o.Payload)
}

func (o *GetHelmReleasesNotFound) String() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesNotFound  %+v", 404, o.Payload)
}

func (o *GetHelmReleasesNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetHelmReleasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHelmReleasesInternalServerError creates a GetHelmReleasesInternalServerError with default headers values
func NewGetHelmReleasesInternalServerError() *GetHelmReleasesInternalServerError {
	return &GetHelmReleasesInternalServerError{}
}

/*
GetHelmReleasesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHelmReleasesInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get helm releases internal server error response has a 2xx status code
func (o *GetHelmReleasesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get helm releases internal server error response has a 3xx status code
func (o *GetHelmReleasesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get helm releases internal server error response has a 4xx status code
func (o *GetHelmReleasesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get helm releases internal server error response has a 5xx status code
func (o *GetHelmReleasesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get helm releases internal server error response a status code equal to that given
func (o *GetHelmReleasesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get helm releases internal server error response
func (o *GetHelmReleasesInternalServerError) Code() int {
	return 500
}

func (o *GetHelmReleasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHelmReleasesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/helm-releases/{helm_chart_id}/releases/{namespace}][%d] getHelmReleasesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetHelmReleasesInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetHelmReleasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
