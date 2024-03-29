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

// GetPingReader is a Reader for the GetPing structure.
type GetPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPingOK creates a GetPingOK with default headers values
func NewGetPingOK() *GetPingOK {
	return &GetPingOK{}
}

/*
GetPingOK describes a response with status code 200, with default header values.

ping response
*/
type GetPingOK struct {
	Payload *models.Ping
}

// IsSuccess returns true when this get ping o k response has a 2xx status code
func (o *GetPingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ping o k response has a 3xx status code
func (o *GetPingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ping o k response has a 4xx status code
func (o *GetPingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ping o k response has a 5xx status code
func (o *GetPingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get ping o k response a status code equal to that given
func (o *GetPingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPingOK) Error() string {
	return fmt.Sprintf("[GET /ping][%d] getPingOK  %+v", 200, o.Payload)
}

func (o *GetPingOK) String() string {
	return fmt.Sprintf("[GET /ping][%d] getPingOK  %+v", 200, o.Payload)
}

func (o *GetPingOK) GetPayload() *models.Ping {
	return o.Payload
}

func (o *GetPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPingInternalServerError creates a GetPingInternalServerError with default headers values
func NewGetPingInternalServerError() *GetPingInternalServerError {
	return &GetPingInternalServerError{}
}

/*
GetPingInternalServerError describes a response with status code 500, with default header values.

When some error occurs
*/
type GetPingInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get ping internal server error response has a 2xx status code
func (o *GetPingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get ping internal server error response has a 3xx status code
func (o *GetPingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ping internal server error response has a 4xx status code
func (o *GetPingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ping internal server error response has a 5xx status code
func (o *GetPingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get ping internal server error response a status code equal to that given
func (o *GetPingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ping][%d] getPingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPingInternalServerError) String() string {
	return fmt.Sprintf("[GET /ping][%d] getPingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPingInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
