// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGoogleLoginHandlerFunc turns a function with the right signature into a get google login handler
type GetGoogleLoginHandlerFunc func(GetGoogleLoginParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGoogleLoginHandlerFunc) Handle(params GetGoogleLoginParams) middleware.Responder {
	return fn(params)
}

// GetGoogleLoginHandler interface for that can handle valid get google login params
type GetGoogleLoginHandler interface {
	Handle(GetGoogleLoginParams) middleware.Responder
}

// NewGetGoogleLogin creates a new http.Handler for the get google login operation
func NewGetGoogleLogin(ctx *middleware.Context, handler GetGoogleLoginHandler) *GetGoogleLogin {
	return &GetGoogleLogin{Context: ctx, Handler: handler}
}

/*
	GetGoogleLogin swagger:route GET /google/login public getGoogleLogin

This API endpoint always responds ok
*/
type GetGoogleLogin struct {
	Context *middleware.Context
	Handler GetGoogleLoginHandler
}

func (o *GetGoogleLogin) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetGoogleLoginParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
