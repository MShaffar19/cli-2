// Code generated by go-swagger; DO NOT EDIT.

package limits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// GetOrganizationLimitsReader is a Reader for the GetOrganizationLimits structure.
type GetOrganizationLimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationLimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationLimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetOrganizationLimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationLimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetOrganizationLimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationLimitsOK creates a GetOrganizationLimitsOK with default headers values
func NewGetOrganizationLimitsOK() *GetOrganizationLimitsOK {
	return &GetOrganizationLimitsOK{}
}

/*GetOrganizationLimitsOK handles this case with default header values.

Success
*/
type GetOrganizationLimitsOK struct {
	Payload *models.Limits
}

func (o *GetOrganizationLimitsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/limits][%d] getOrganizationLimitsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Limits)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationLimitsForbidden creates a GetOrganizationLimitsForbidden with default headers values
func NewGetOrganizationLimitsForbidden() *GetOrganizationLimitsForbidden {
	return &GetOrganizationLimitsForbidden{}
}

/*GetOrganizationLimitsForbidden handles this case with default header values.

Forbidden
*/
type GetOrganizationLimitsForbidden struct {
	Payload *models.Message
}

func (o *GetOrganizationLimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/limits][%d] getOrganizationLimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationLimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationLimitsNotFound creates a GetOrganizationLimitsNotFound with default headers values
func NewGetOrganizationLimitsNotFound() *GetOrganizationLimitsNotFound {
	return &GetOrganizationLimitsNotFound{}
}

/*GetOrganizationLimitsNotFound handles this case with default header values.

Not Found
*/
type GetOrganizationLimitsNotFound struct {
	Payload *models.Message
}

func (o *GetOrganizationLimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/limits][%d] getOrganizationLimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationLimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationLimitsInternalServerError creates a GetOrganizationLimitsInternalServerError with default headers values
func NewGetOrganizationLimitsInternalServerError() *GetOrganizationLimitsInternalServerError {
	return &GetOrganizationLimitsInternalServerError{}
}

/*GetOrganizationLimitsInternalServerError handles this case with default header values.

Server Error
*/
type GetOrganizationLimitsInternalServerError struct {
	Payload *models.Message
}

func (o *GetOrganizationLimitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/limits][%d] getOrganizationLimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrganizationLimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
