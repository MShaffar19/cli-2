// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// SearchEmailsReader is a Reader for the SearchEmails structure.
type SearchEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewSearchEmailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSearchEmailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchEmailsOK creates a SearchEmailsOK with default headers values
func NewSearchEmailsOK() *SearchEmailsOK {
	return &SearchEmailsOK{}
}

/*SearchEmailsOK handles this case with default header values.

Search for users matching the given search string
*/
type SearchEmailsOK struct {
	Payload []*models.User
}

func (o *SearchEmailsOK) Error() string {
	return fmt.Sprintf("[POST /users/search_emails][%d] searchEmailsOK  %+v", 200, o.Payload)
}

func (o *SearchEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchEmailsForbidden creates a SearchEmailsForbidden with default headers values
func NewSearchEmailsForbidden() *SearchEmailsForbidden {
	return &SearchEmailsForbidden{}
}

/*SearchEmailsForbidden handles this case with default header values.

Forbidden
*/
type SearchEmailsForbidden struct {
	Payload *models.Message
}

func (o *SearchEmailsForbidden) Error() string {
	return fmt.Sprintf("[POST /users/search_emails][%d] searchEmailsForbidden  %+v", 403, o.Payload)
}

func (o *SearchEmailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchEmailsInternalServerError creates a SearchEmailsInternalServerError with default headers values
func NewSearchEmailsInternalServerError() *SearchEmailsInternalServerError {
	return &SearchEmailsInternalServerError{}
}

/*SearchEmailsInternalServerError handles this case with default header values.

Server Error
*/
type SearchEmailsInternalServerError struct {
	Payload *models.Message
}

func (o *SearchEmailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/search_emails][%d] searchEmailsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchEmailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SearchEmailsBody search emails body
swagger:model SearchEmailsBody
*/
type SearchEmailsBody struct {

	// The search query
	Query string `json:"query,omitempty"`
}

// Validate validates this search emails body
func (o *SearchEmailsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchEmailsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchEmailsBody) UnmarshalBinary(b []byte) error {
	var res SearchEmailsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
