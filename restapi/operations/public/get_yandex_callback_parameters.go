// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetYandexCallbackParams creates a new GetYandexCallbackParams object
//
// There are no default values defined in the spec.
func NewGetYandexCallbackParams() GetYandexCallbackParams {

	return GetYandexCallbackParams{}
}

// GetYandexCallbackParams contains all the bound params for the get yandex callback operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetYandexCallback
type GetYandexCallbackParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*client identity
	  In: query
	*/
	Authuser *string
	/*consent
	  In: query
	*/
	Prompt *string
	/*client scope
	  In: query
	*/
	Scope *string
	/*client state
	  In: query
	*/
	State *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetYandexCallbackParams() beforehand.
func (o *GetYandexCallbackParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAuthuser, qhkAuthuser, _ := qs.GetOK("authuser")
	if err := o.bindAuthuser(qAuthuser, qhkAuthuser, route.Formats); err != nil {
		res = append(res, err)
	}

	qPrompt, qhkPrompt, _ := qs.GetOK("prompt")
	if err := o.bindPrompt(qPrompt, qhkPrompt, route.Formats); err != nil {
		res = append(res, err)
	}

	qScope, qhkScope, _ := qs.GetOK("scope")
	if err := o.bindScope(qScope, qhkScope, route.Formats); err != nil {
		res = append(res, err)
	}

	qState, qhkState, _ := qs.GetOK("state")
	if err := o.bindState(qState, qhkState, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthuser binds and validates parameter Authuser from query.
func (o *GetYandexCallbackParams) bindAuthuser(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Authuser = &raw

	return nil
}

// bindPrompt binds and validates parameter Prompt from query.
func (o *GetYandexCallbackParams) bindPrompt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Prompt = &raw

	return nil
}

// bindScope binds and validates parameter Scope from query.
func (o *GetYandexCallbackParams) bindScope(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Scope = &raw

	return nil
}

// bindState binds and validates parameter State from query.
func (o *GetYandexCallbackParams) bindState(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.State = &raw

	return nil
}