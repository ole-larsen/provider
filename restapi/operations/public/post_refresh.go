// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostRefreshHandlerFunc turns a function with the right signature into a post refresh handler
type PostRefreshHandlerFunc func(PostRefreshParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRefreshHandlerFunc) Handle(params PostRefreshParams) middleware.Responder {
	return fn(params)
}

// PostRefreshHandler interface for that can handle valid post refresh params
type PostRefreshHandler interface {
	Handle(PostRefreshParams) middleware.Responder
}

// NewPostRefresh creates a new http.Handler for the post refresh operation
func NewPostRefresh(ctx *middleware.Context, handler PostRefreshHandler) *PostRefresh {
	return &PostRefresh{Context: ctx, Handler: handler}
}

/*
	PostRefresh swagger:route POST /refresh public postRefresh

This API endpoint create, store and returns credentials for new user
*/
type PostRefresh struct {
	Context *middleware.Context
	Handler PostRefreshHandler
}

func (o *PostRefresh) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostRefreshParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
