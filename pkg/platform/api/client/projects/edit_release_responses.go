// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// EditReleaseReader is a Reader for the EditRelease structure.
type EditReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditReleaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEditReleaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEditReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEditReleaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditReleaseOK creates a EditReleaseOK with default headers values
func NewEditReleaseOK() *EditReleaseOK {
	return &EditReleaseOK{}
}

/*EditReleaseOK handles this case with default header values.

Release updated
*/
type EditReleaseOK struct {
	Payload *models.Release
}

func (o *EditReleaseOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] editReleaseOK  %+v", 200, o.Payload)
}

func (o *EditReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Release)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditReleaseBadRequest creates a EditReleaseBadRequest with default headers values
func NewEditReleaseBadRequest() *EditReleaseBadRequest {
	return &EditReleaseBadRequest{}
}

/*EditReleaseBadRequest handles this case with default header values.

Bad Request
*/
type EditReleaseBadRequest struct {
	Payload *models.Message
}

func (o *EditReleaseBadRequest) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] editReleaseBadRequest  %+v", 400, o.Payload)
}

func (o *EditReleaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditReleaseForbidden creates a EditReleaseForbidden with default headers values
func NewEditReleaseForbidden() *EditReleaseForbidden {
	return &EditReleaseForbidden{}
}

/*EditReleaseForbidden handles this case with default header values.

Unauthorized
*/
type EditReleaseForbidden struct {
	Payload *models.Message
}

func (o *EditReleaseForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] editReleaseForbidden  %+v", 403, o.Payload)
}

func (o *EditReleaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditReleaseNotFound creates a EditReleaseNotFound with default headers values
func NewEditReleaseNotFound() *EditReleaseNotFound {
	return &EditReleaseNotFound{}
}

/*EditReleaseNotFound handles this case with default header values.

Not Found
*/
type EditReleaseNotFound struct {
	Payload *models.Message
}

func (o *EditReleaseNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] editReleaseNotFound  %+v", 404, o.Payload)
}

func (o *EditReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditReleaseInternalServerError creates a EditReleaseInternalServerError with default headers values
func NewEditReleaseInternalServerError() *EditReleaseInternalServerError {
	return &EditReleaseInternalServerError{}
}

/*EditReleaseInternalServerError handles this case with default header values.

Server Error
*/
type EditReleaseInternalServerError struct {
	Payload *models.Message
}

func (o *EditReleaseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}][%d] editReleaseInternalServerError  %+v", 500, o.Payload)
}

func (o *EditReleaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
