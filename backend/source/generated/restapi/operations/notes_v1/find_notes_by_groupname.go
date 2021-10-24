// Code generated by go-swagger; DO NOT EDIT.

package notes_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindNotesByGroupnameHandlerFunc turns a function with the right signature into a find notes by groupname handler
type FindNotesByGroupnameHandlerFunc func(FindNotesByGroupnameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindNotesByGroupnameHandlerFunc) Handle(params FindNotesByGroupnameParams) middleware.Responder {
	return fn(params)
}

// FindNotesByGroupnameHandler interface for that can handle valid find notes by groupname params
type FindNotesByGroupnameHandler interface {
	Handle(FindNotesByGroupnameParams) middleware.Responder
}

// NewFindNotesByGroupname creates a new http.Handler for the find notes by groupname operation
func NewFindNotesByGroupname(ctx *middleware.Context, handler FindNotesByGroupnameHandler) *FindNotesByGroupname {
	return &FindNotesByGroupname{Context: ctx, Handler: handler}
}

/* FindNotesByGroupname swagger:route GET /v1/notes/groupname/{group_name} notesV1 findNotesByGroupname

find notes by Groupname

one groupname can be provided to search

*/
type FindNotesByGroupname struct {
	Context *middleware.Context
	Handler FindNotesByGroupnameHandler
}

func (o *FindNotesByGroupname) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindNotesByGroupnameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
