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

// NewGetValidateParams creates a new GetValidateParams object
//
// There are no default values defined in the spec.
func NewGetValidateParams() GetValidateParams {

	return GetValidateParams{}
}

// GetValidateParams contains all the bound params for the get validate operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetValidate
type GetValidateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*access token
	  In: query
	*/
	AccessToken *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetValidateParams() beforehand.
func (o *GetValidateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAccessToken, qhkAccessToken, _ := qs.GetOK("access_token")
	if err := o.bindAccessToken(qAccessToken, qhkAccessToken, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAccessToken binds and validates parameter AccessToken from query.
func (o *GetValidateParams) bindAccessToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.AccessToken = &raw

	return nil
}
