// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/olelarssen/provider/models"
)

// GetCredentialsReader is a Reader for the GetCredentials structure.
type GetCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCredentialsOK creates a GetCredentialsOK with default headers values
func NewGetCredentialsOK() *GetCredentialsOK {
	return &GetCredentialsOK{}
}

/*
GetCredentialsOK describes a response with status code 200, with default header values.

get client id and client secret
*/
type GetCredentialsOK struct {
	Payload *models.Credentials
}

// IsSuccess returns true when this get credentials o k response has a 2xx status code
func (o *GetCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get credentials o k response has a 3xx status code
func (o *GetCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials o k response has a 4xx status code
func (o *GetCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials o k response has a 5xx status code
func (o *GetCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get credentials o k response a status code equal to that given
func (o *GetCredentialsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /credentials][%d] getCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetCredentialsOK) String() string {
	return fmt.Sprintf("[GET /credentials][%d] getCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetCredentialsOK) GetPayload() *models.Credentials {
	return o.Payload
}

func (o *GetCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Credentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsInternalServerError creates a GetCredentialsInternalServerError with default headers values
func NewGetCredentialsInternalServerError() *GetCredentialsInternalServerError {
	return &GetCredentialsInternalServerError{}
}

/*
GetCredentialsInternalServerError describes a response with status code 500, with default header values.

When some error occurs
*/
type GetCredentialsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get credentials internal server error response has a 2xx status code
func (o *GetCredentialsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get credentials internal server error response has a 3xx status code
func (o *GetCredentialsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credentials internal server error response has a 4xx status code
func (o *GetCredentialsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credentials internal server error response has a 5xx status code
func (o *GetCredentialsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get credentials internal server error response a status code equal to that given
func (o *GetCredentialsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /credentials][%d] getCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialsInternalServerError) String() string {
	return fmt.Sprintf("[GET /credentials][%d] getCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
