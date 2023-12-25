// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/olelarssen/provider/models"
)

// GetTelegramCallbackOKCode is the HTTP code returned for type GetTelegramCallbackOK
const GetTelegramCallbackOKCode int = 200

/*
GetTelegramCallbackOK ok response

swagger:response getTelegramCallbackOK
*/
type GetTelegramCallbackOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserInfo `json:"body,omitempty"`
}

// NewGetTelegramCallbackOK creates GetTelegramCallbackOK with default headers values
func NewGetTelegramCallbackOK() *GetTelegramCallbackOK {

	return &GetTelegramCallbackOK{}
}

// WithPayload adds the payload to the get telegram callback o k response
func (o *GetTelegramCallbackOK) WithPayload(payload *models.UserInfo) *GetTelegramCallbackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get telegram callback o k response
func (o *GetTelegramCallbackOK) SetPayload(payload *models.UserInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTelegramCallbackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTelegramCallbackInternalServerErrorCode is the HTTP code returned for type GetTelegramCallbackInternalServerError
const GetTelegramCallbackInternalServerErrorCode int = 500

/*
GetTelegramCallbackInternalServerError When some error occurs

swagger:response getTelegramCallbackInternalServerError
*/
type GetTelegramCallbackInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTelegramCallbackInternalServerError creates GetTelegramCallbackInternalServerError with default headers values
func NewGetTelegramCallbackInternalServerError() *GetTelegramCallbackInternalServerError {

	return &GetTelegramCallbackInternalServerError{}
}

// WithPayload adds the payload to the get telegram callback internal server error response
func (o *GetTelegramCallbackInternalServerError) WithPayload(payload *models.Error) *GetTelegramCallbackInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get telegram callback internal server error response
func (o *GetTelegramCallbackInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTelegramCallbackInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
