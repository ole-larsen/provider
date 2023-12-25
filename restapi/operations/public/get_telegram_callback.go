// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTelegramCallbackHandlerFunc turns a function with the right signature into a get telegram callback handler
type GetTelegramCallbackHandlerFunc func(GetTelegramCallbackParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTelegramCallbackHandlerFunc) Handle(params GetTelegramCallbackParams) middleware.Responder {
	return fn(params)
}

// GetTelegramCallbackHandler interface for that can handle valid get telegram callback params
type GetTelegramCallbackHandler interface {
	Handle(GetTelegramCallbackParams) middleware.Responder
}

// NewGetTelegramCallback creates a new http.Handler for the get telegram callback operation
func NewGetTelegramCallback(ctx *middleware.Context, handler GetTelegramCallbackHandler) *GetTelegramCallback {
	return &GetTelegramCallback{Context: ctx, Handler: handler}
}

/*
	GetTelegramCallback swagger:route GET /telegram/callback public getTelegramCallback

This API endpoint always responds ok
*/
type GetTelegramCallback struct {
	Context *middleware.Context
	Handler GetTelegramCallbackHandler
}

func (o *GetTelegramCallback) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTelegramCallbackParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
