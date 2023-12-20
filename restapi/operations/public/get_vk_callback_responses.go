// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/olelarssen/provider/models"
)

// GetVkCallbackOKCode is the HTTP code returned for type GetVkCallbackOK
const GetVkCallbackOKCode int = 200

/*
GetVkCallbackOK ok response

swagger:response getVkCallbackOK
*/
type GetVkCallbackOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserInfo `json:"body,omitempty"`
}

// NewGetVkCallbackOK creates GetVkCallbackOK with default headers values
func NewGetVkCallbackOK() *GetVkCallbackOK {

	return &GetVkCallbackOK{}
}

// WithPayload adds the payload to the get vk callback o k response
func (o *GetVkCallbackOK) WithPayload(payload *models.UserInfo) *GetVkCallbackOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vk callback o k response
func (o *GetVkCallbackOK) SetPayload(payload *models.UserInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVkCallbackOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVkCallbackInternalServerErrorCode is the HTTP code returned for type GetVkCallbackInternalServerError
const GetVkCallbackInternalServerErrorCode int = 500

/*
GetVkCallbackInternalServerError When some error occurs

swagger:response getVkCallbackInternalServerError
*/
type GetVkCallbackInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVkCallbackInternalServerError creates GetVkCallbackInternalServerError with default headers values
func NewGetVkCallbackInternalServerError() *GetVkCallbackInternalServerError {

	return &GetVkCallbackInternalServerError{}
}

// WithPayload adds the payload to the get vk callback internal server error response
func (o *GetVkCallbackInternalServerError) WithPayload(payload *models.Error) *GetVkCallbackInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vk callback internal server error response
func (o *GetVkCallbackInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVkCallbackInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}