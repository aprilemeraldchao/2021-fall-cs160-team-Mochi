// Code generated by go-swagger; DO NOT EDIT.

package user_mgmt_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateUserInfoV1HandlerFunc turns a function with the right signature into a update user info v1 handler
type UpdateUserInfoV1HandlerFunc func(UpdateUserInfoV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserInfoV1HandlerFunc) Handle(params UpdateUserInfoV1Params) middleware.Responder {
	return fn(params)
}

// UpdateUserInfoV1Handler interface for that can handle valid update user info v1 params
type UpdateUserInfoV1Handler interface {
	Handle(UpdateUserInfoV1Params) middleware.Responder
}

// NewUpdateUserInfoV1 creates a new http.Handler for the update user info v1 operation
func NewUpdateUserInfoV1(ctx *middleware.Context, handler UpdateUserInfoV1Handler) *UpdateUserInfoV1 {
	return &UpdateUserInfoV1{Context: ctx, Handler: handler}
}

/* UpdateUserInfoV1 swagger:route PATCH /v1/user UserMgmtV1 updateUserInfoV1

Update user info

update user info

*/
type UpdateUserInfoV1 struct {
	Context *middleware.Context
	Handler UpdateUserInfoV1Handler
}

func (o *UpdateUserInfoV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateUserInfoV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}