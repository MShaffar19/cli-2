// Code generated by go-swagger; DO NOT EDIT.

package version_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// GetCommitHistoryReader is a Reader for the GetCommitHistory structure.
type GetCommitHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCommitHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCommitHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCommitHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCommitHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCommitHistoryOK creates a GetCommitHistoryOK with default headers values
func NewGetCommitHistoryOK() *GetCommitHistoryOK {
	return &GetCommitHistoryOK{}
}

/*GetCommitHistoryOK handles this case with default header values.

Get commit history starting from the given commit
*/
type GetCommitHistoryOK struct {
	Payload models.CommitHistory
}

func (o *GetCommitHistoryOK) Error() string {
	return fmt.Sprintf("[GET /vcs/history/{commitID}][%d] getCommitHistoryOK  %+v", 200, o.Payload)
}

func (o *GetCommitHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommitHistoryNotFound creates a GetCommitHistoryNotFound with default headers values
func NewGetCommitHistoryNotFound() *GetCommitHistoryNotFound {
	return &GetCommitHistoryNotFound{}
}

/*GetCommitHistoryNotFound handles this case with default header values.

commit was not found
*/
type GetCommitHistoryNotFound struct {
	Payload *models.Message
}

func (o *GetCommitHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /vcs/history/{commitID}][%d] getCommitHistoryNotFound  %+v", 404, o.Payload)
}

func (o *GetCommitHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommitHistoryInternalServerError creates a GetCommitHistoryInternalServerError with default headers values
func NewGetCommitHistoryInternalServerError() *GetCommitHistoryInternalServerError {
	return &GetCommitHistoryInternalServerError{}
}

/*GetCommitHistoryInternalServerError handles this case with default header values.

Server Error
*/
type GetCommitHistoryInternalServerError struct {
	Payload *models.Message
}

func (o *GetCommitHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /vcs/history/{commitID}][%d] getCommitHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCommitHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
