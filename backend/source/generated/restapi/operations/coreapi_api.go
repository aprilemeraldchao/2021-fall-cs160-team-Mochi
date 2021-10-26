// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/comments_v1"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/groups_v1"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/notes_v1"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_images_v1"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations/user_mgmt_v1"
)

// NewCoreapiAPI creates a new Coreapi instance
func NewCoreapiAPI(spec *loads.Document) *CoreapiAPI {
	return &CoreapiAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		GroupsV1AddGroupUsersV1Handler: groups_v1.AddGroupUsersV1HandlerFunc(func(params groups_v1.AddGroupUsersV1Params) middleware.Responder {
			return middleware.NotImplemented("operation groups_v1.AddGroupUsersV1 has not yet been implemented")
		}),
		GroupsV1CreateGroupV1Handler: groups_v1.CreateGroupV1HandlerFunc(func(params groups_v1.CreateGroupV1Params) middleware.Responder {
			return middleware.NotImplemented("operation groups_v1.CreateGroupV1 has not yet been implemented")
		}),
		GroupsV1DeleteGroupV1Handler: groups_v1.DeleteGroupV1HandlerFunc(func(params groups_v1.DeleteGroupV1Params) middleware.Responder {
			return middleware.NotImplemented("operation groups_v1.DeleteGroupV1 has not yet been implemented")
		}),
		NotesV1DeleteNoteHandler: notes_v1.DeleteNoteHandlerFunc(func(params notes_v1.DeleteNoteParams) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.DeleteNote has not yet been implemented")
		}),
		NotesV1FindNotesByGroupnameHandler: notes_v1.FindNotesByGroupnameHandlerFunc(func(params notes_v1.FindNotesByGroupnameParams) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.FindNotesByGroupname has not yet been implemented")
		}),
		NotesV1FindNotesByTagsHandler: notes_v1.FindNotesByTagsHandlerFunc(func(params notes_v1.FindNotesByTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.FindNotesByTags has not yet been implemented")
		}),
		NotesV1FindNotesByUsernameHandler: notes_v1.FindNotesByUsernameHandlerFunc(func(params notes_v1.FindNotesByUsernameParams) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.FindNotesByUsername has not yet been implemented")
		}),
		NotesV1GetFileV1Handler: notes_v1.GetFileV1HandlerFunc(func(params notes_v1.GetFileV1Params) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.GetFileV1 has not yet been implemented")
		}),
		GroupsV1GetGroupInfoV1Handler: groups_v1.GetGroupInfoV1HandlerFunc(func(params groups_v1.GetGroupInfoV1Params) middleware.Responder {
			return middleware.NotImplemented("operation groups_v1.GetGroupInfoV1 has not yet been implemented")
		}),
		GroupsV1GetGroupUsersV1Handler: groups_v1.GetGroupUsersV1HandlerFunc(func(params groups_v1.GetGroupUsersV1Params) middleware.Responder {
			return middleware.NotImplemented("operation groups_v1.GetGroupUsersV1 has not yet been implemented")
		}),
		GroupsV1GetGroupsV1Handler: groups_v1.GetGroupsV1HandlerFunc(func(params groups_v1.GetGroupsV1Params) middleware.Responder {
			return middleware.NotImplemented("operation groups_v1.GetGroupsV1 has not yet been implemented")
		}),
		NotesV1GetMultipleFilesV1Handler: notes_v1.GetMultipleFilesV1HandlerFunc(func(params notes_v1.GetMultipleFilesV1Params) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.GetMultipleFilesV1 has not yet been implemented")
		}),
		NotesV1GetNoteCommentsHandler: notes_v1.GetNoteCommentsHandlerFunc(func(params notes_v1.GetNoteCommentsParams) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.GetNoteComments has not yet been implemented")
		}),
		NotesV1GetNoteMembersHandler: notes_v1.GetNoteMembersHandlerFunc(func(params notes_v1.GetNoteMembersParams) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.GetNoteMembers has not yet been implemented")
		}),
		UserImagesV1GetUserImagesV1Handler: user_images_v1.GetUserImagesV1HandlerFunc(func(params user_images_v1.GetUserImagesV1Params) middleware.Responder {
			return middleware.NotImplemented("operation user_images_v1.GetUserImagesV1 has not yet been implemented")
		}),
		UserMgmtV1GetUserV1Handler: user_mgmt_v1.GetUserV1HandlerFunc(func(params user_mgmt_v1.GetUserV1Params) middleware.Responder {
			return middleware.NotImplemented("operation user_mgmt_v1.GetUserV1 has not yet been implemented")
		}),
		UserMgmtV1LoginV1Handler: user_mgmt_v1.LoginV1HandlerFunc(func(params user_mgmt_v1.LoginV1Params) middleware.Responder {
			return middleware.NotImplemented("operation user_mgmt_v1.LoginV1 has not yet been implemented")
		}),
		CommentsV1PostCommentsV1Handler: comments_v1.PostCommentsV1HandlerFunc(func(params comments_v1.PostCommentsV1Params) middleware.Responder {
			return middleware.NotImplemented("operation comments_v1.PostCommentsV1 has not yet been implemented")
		}),
		NotesV1PostFileV1Handler: notes_v1.PostFileV1HandlerFunc(func(params notes_v1.PostFileV1Params) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.PostFileV1 has not yet been implemented")
		}),
		UserImagesV1PostUserImagesV1Handler: user_images_v1.PostUserImagesV1HandlerFunc(func(params user_images_v1.PostUserImagesV1Params) middleware.Responder {
			return middleware.NotImplemented("operation user_images_v1.PostUserImagesV1 has not yet been implemented")
		}),
		CommentsV1RemoveComnentV1Handler: comments_v1.RemoveComnentV1HandlerFunc(func(params comments_v1.RemoveComnentV1Params) middleware.Responder {
			return middleware.NotImplemented("operation comments_v1.RemoveComnentV1 has not yet been implemented")
		}),
		GroupsV1RemoveGroupUsersV1Handler: groups_v1.RemoveGroupUsersV1HandlerFunc(func(params groups_v1.RemoveGroupUsersV1Params) middleware.Responder {
			return middleware.NotImplemented("operation groups_v1.RemoveGroupUsersV1 has not yet been implemented")
		}),
		NotesV1UpdateNoteHandler: notes_v1.UpdateNoteHandlerFunc(func(params notes_v1.UpdateNoteParams) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.UpdateNote has not yet been implemented")
		}),
		UserMgmtV1UpdatePasswordV1Handler: user_mgmt_v1.UpdatePasswordV1HandlerFunc(func(params user_mgmt_v1.UpdatePasswordV1Params) middleware.Responder {
			return middleware.NotImplemented("operation user_mgmt_v1.UpdatePasswordV1 has not yet been implemented")
		}),
		UserMgmtV1UpdateUserInfoV1Handler: user_mgmt_v1.UpdateUserInfoV1HandlerFunc(func(params user_mgmt_v1.UpdateUserInfoV1Params) middleware.Responder {
			return middleware.NotImplemented("operation user_mgmt_v1.UpdateUserInfoV1 has not yet been implemented")
		}),
		NotesV1UploadNoteV1Handler: notes_v1.UploadNoteV1HandlerFunc(func(params notes_v1.UploadNoteV1Params) middleware.Responder {
			return middleware.NotImplemented("operation notes_v1.UploadNoteV1 has not yet been implemented")
		}),
	}
}

/*CoreapiAPI This is the Public API for MochiNote */
type CoreapiAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// GroupsV1AddGroupUsersV1Handler sets the operation handler for the add group users v1 operation
	GroupsV1AddGroupUsersV1Handler groups_v1.AddGroupUsersV1Handler
	// GroupsV1CreateGroupV1Handler sets the operation handler for the create group v1 operation
	GroupsV1CreateGroupV1Handler groups_v1.CreateGroupV1Handler
	// GroupsV1DeleteGroupV1Handler sets the operation handler for the delete group v1 operation
	GroupsV1DeleteGroupV1Handler groups_v1.DeleteGroupV1Handler
	// NotesV1DeleteNoteHandler sets the operation handler for the delete note operation
	NotesV1DeleteNoteHandler notes_v1.DeleteNoteHandler
	// NotesV1FindNotesByGroupnameHandler sets the operation handler for the find notes by groupname operation
	NotesV1FindNotesByGroupnameHandler notes_v1.FindNotesByGroupnameHandler
	// NotesV1FindNotesByTagsHandler sets the operation handler for the find notes by tags operation
	NotesV1FindNotesByTagsHandler notes_v1.FindNotesByTagsHandler
	// NotesV1FindNotesByUsernameHandler sets the operation handler for the find notes by username operation
	NotesV1FindNotesByUsernameHandler notes_v1.FindNotesByUsernameHandler
	// NotesV1GetFileV1Handler sets the operation handler for the get file v1 operation
	NotesV1GetFileV1Handler notes_v1.GetFileV1Handler
	// GroupsV1GetGroupInfoV1Handler sets the operation handler for the get group info v1 operation
	GroupsV1GetGroupInfoV1Handler groups_v1.GetGroupInfoV1Handler
	// GroupsV1GetGroupUsersV1Handler sets the operation handler for the get group users v1 operation
	GroupsV1GetGroupUsersV1Handler groups_v1.GetGroupUsersV1Handler
	// GroupsV1GetGroupsV1Handler sets the operation handler for the get groups v1 operation
	GroupsV1GetGroupsV1Handler groups_v1.GetGroupsV1Handler
	// NotesV1GetMultipleFilesV1Handler sets the operation handler for the get multiple files v1 operation
	NotesV1GetMultipleFilesV1Handler notes_v1.GetMultipleFilesV1Handler
	// NotesV1GetNoteCommentsHandler sets the operation handler for the get note comments operation
	NotesV1GetNoteCommentsHandler notes_v1.GetNoteCommentsHandler
	// NotesV1GetNoteMembersHandler sets the operation handler for the get note members operation
	NotesV1GetNoteMembersHandler notes_v1.GetNoteMembersHandler
	// UserImagesV1GetUserImagesV1Handler sets the operation handler for the get user images v1 operation
	UserImagesV1GetUserImagesV1Handler user_images_v1.GetUserImagesV1Handler
	// UserMgmtV1GetUserV1Handler sets the operation handler for the get user v1 operation
	UserMgmtV1GetUserV1Handler user_mgmt_v1.GetUserV1Handler
	// UserMgmtV1LoginV1Handler sets the operation handler for the login v1 operation
	UserMgmtV1LoginV1Handler user_mgmt_v1.LoginV1Handler
	// CommentsV1PostCommentsV1Handler sets the operation handler for the post comments v1 operation
	CommentsV1PostCommentsV1Handler comments_v1.PostCommentsV1Handler
	// NotesV1PostFileV1Handler sets the operation handler for the post file v1 operation
	NotesV1PostFileV1Handler notes_v1.PostFileV1Handler
	// UserImagesV1PostUserImagesV1Handler sets the operation handler for the post user images v1 operation
	UserImagesV1PostUserImagesV1Handler user_images_v1.PostUserImagesV1Handler
	// CommentsV1RemoveComnentV1Handler sets the operation handler for the remove comnent v1 operation
	CommentsV1RemoveComnentV1Handler comments_v1.RemoveComnentV1Handler
	// GroupsV1RemoveGroupUsersV1Handler sets the operation handler for the remove group users v1 operation
	GroupsV1RemoveGroupUsersV1Handler groups_v1.RemoveGroupUsersV1Handler
	// NotesV1UpdateNoteHandler sets the operation handler for the update note operation
	NotesV1UpdateNoteHandler notes_v1.UpdateNoteHandler
	// UserMgmtV1UpdatePasswordV1Handler sets the operation handler for the update password v1 operation
	UserMgmtV1UpdatePasswordV1Handler user_mgmt_v1.UpdatePasswordV1Handler
	// UserMgmtV1UpdateUserInfoV1Handler sets the operation handler for the update user info v1 operation
	UserMgmtV1UpdateUserInfoV1Handler user_mgmt_v1.UpdateUserInfoV1Handler
	// NotesV1UploadNoteV1Handler sets the operation handler for the upload note v1 operation
	NotesV1UploadNoteV1Handler notes_v1.UploadNoteV1Handler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *CoreapiAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *CoreapiAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *CoreapiAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CoreapiAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CoreapiAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CoreapiAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CoreapiAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CoreapiAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CoreapiAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CoreapiAPI
func (o *CoreapiAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GroupsV1AddGroupUsersV1Handler == nil {
		unregistered = append(unregistered, "groups_v1.AddGroupUsersV1Handler")
	}
	if o.GroupsV1CreateGroupV1Handler == nil {
		unregistered = append(unregistered, "groups_v1.CreateGroupV1Handler")
	}
	if o.GroupsV1DeleteGroupV1Handler == nil {
		unregistered = append(unregistered, "groups_v1.DeleteGroupV1Handler")
	}
	if o.NotesV1DeleteNoteHandler == nil {
		unregistered = append(unregistered, "notes_v1.DeleteNoteHandler")
	}
	if o.NotesV1FindNotesByGroupnameHandler == nil {
		unregistered = append(unregistered, "notes_v1.FindNotesByGroupnameHandler")
	}
	if o.NotesV1FindNotesByTagsHandler == nil {
		unregistered = append(unregistered, "notes_v1.FindNotesByTagsHandler")
	}
	if o.NotesV1FindNotesByUsernameHandler == nil {
		unregistered = append(unregistered, "notes_v1.FindNotesByUsernameHandler")
	}
	if o.NotesV1GetFileV1Handler == nil {
		unregistered = append(unregistered, "notes_v1.GetFileV1Handler")
	}
	if o.GroupsV1GetGroupInfoV1Handler == nil {
		unregistered = append(unregistered, "groups_v1.GetGroupInfoV1Handler")
	}
	if o.GroupsV1GetGroupUsersV1Handler == nil {
		unregistered = append(unregistered, "groups_v1.GetGroupUsersV1Handler")
	}
	if o.GroupsV1GetGroupsV1Handler == nil {
		unregistered = append(unregistered, "groups_v1.GetGroupsV1Handler")
	}
	if o.NotesV1GetMultipleFilesV1Handler == nil {
		unregistered = append(unregistered, "notes_v1.GetMultipleFilesV1Handler")
	}
	if o.NotesV1GetNoteCommentsHandler == nil {
		unregistered = append(unregistered, "notes_v1.GetNoteCommentsHandler")
	}
	if o.NotesV1GetNoteMembersHandler == nil {
		unregistered = append(unregistered, "notes_v1.GetNoteMembersHandler")
	}
	if o.UserImagesV1GetUserImagesV1Handler == nil {
		unregistered = append(unregistered, "user_images_v1.GetUserImagesV1Handler")
	}
	if o.UserMgmtV1GetUserV1Handler == nil {
		unregistered = append(unregistered, "user_mgmt_v1.GetUserV1Handler")
	}
	if o.UserMgmtV1LoginV1Handler == nil {
		unregistered = append(unregistered, "user_mgmt_v1.LoginV1Handler")
	}
	if o.CommentsV1PostCommentsV1Handler == nil {
		unregistered = append(unregistered, "comments_v1.PostCommentsV1Handler")
	}
	if o.NotesV1PostFileV1Handler == nil {
		unregistered = append(unregistered, "notes_v1.PostFileV1Handler")
	}
	if o.UserImagesV1PostUserImagesV1Handler == nil {
		unregistered = append(unregistered, "user_images_v1.PostUserImagesV1Handler")
	}
	if o.CommentsV1RemoveComnentV1Handler == nil {
		unregistered = append(unregistered, "comments_v1.RemoveComnentV1Handler")
	}
	if o.GroupsV1RemoveGroupUsersV1Handler == nil {
		unregistered = append(unregistered, "groups_v1.RemoveGroupUsersV1Handler")
	}
	if o.NotesV1UpdateNoteHandler == nil {
		unregistered = append(unregistered, "notes_v1.UpdateNoteHandler")
	}
	if o.UserMgmtV1UpdatePasswordV1Handler == nil {
		unregistered = append(unregistered, "user_mgmt_v1.UpdatePasswordV1Handler")
	}
	if o.UserMgmtV1UpdateUserInfoV1Handler == nil {
		unregistered = append(unregistered, "user_mgmt_v1.UpdateUserInfoV1Handler")
	}
	if o.NotesV1UploadNoteV1Handler == nil {
		unregistered = append(unregistered, "notes_v1.UploadNoteV1Handler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CoreapiAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CoreapiAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *CoreapiAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *CoreapiAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *CoreapiAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CoreapiAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the coreapi API
func (o *CoreapiAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CoreapiAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/groups/{group_id}/members"] = groups_v1.NewAddGroupUsersV1(o.context, o.GroupsV1AddGroupUsersV1Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/groups"] = groups_v1.NewCreateGroupV1(o.context, o.GroupsV1CreateGroupV1Handler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/groups/{group_id}"] = groups_v1.NewDeleteGroupV1(o.context, o.GroupsV1DeleteGroupV1Handler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/notes/{id}"] = notes_v1.NewDeleteNote(o.context, o.NotesV1DeleteNoteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/notes/groupname/{group_name}"] = notes_v1.NewFindNotesByGroupname(o.context, o.NotesV1FindNotesByGroupnameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/notes/tag/{tag}"] = notes_v1.NewFindNotesByTags(o.context, o.NotesV1FindNotesByTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/notes/username"] = notes_v1.NewFindNotesByUsername(o.context, o.NotesV1FindNotesByUsernameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/notes/file/{path}"] = notes_v1.NewGetFileV1(o.context, o.NotesV1GetFileV1Handler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/groups/{group_id}"] = groups_v1.NewGetGroupInfoV1(o.context, o.GroupsV1GetGroupInfoV1Handler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/groups/{group_id}/members"] = groups_v1.NewGetGroupUsersV1(o.context, o.GroupsV1GetGroupUsersV1Handler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/groups"] = groups_v1.NewGetGroupsV1(o.context, o.GroupsV1GetGroupsV1Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/notes/files"] = notes_v1.NewGetMultipleFilesV1(o.context, o.NotesV1GetMultipleFilesV1Handler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/notes/{note_id}/comments"] = notes_v1.NewGetNoteComments(o.context, o.NotesV1GetNoteCommentsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/notes/{id}/members"] = notes_v1.NewGetNoteMembers(o.context, o.NotesV1GetNoteMembersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/images"] = user_images_v1.NewGetUserImagesV1(o.context, o.UserImagesV1GetUserImagesV1Handler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v1/user"] = user_mgmt_v1.NewGetUserV1(o.context, o.UserMgmtV1GetUserV1Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/login"] = user_mgmt_v1.NewLoginV1(o.context, o.UserMgmtV1LoginV1Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/comments"] = comments_v1.NewPostCommentsV1(o.context, o.CommentsV1PostCommentsV1Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/notes/file"] = notes_v1.NewPostFileV1(o.context, o.NotesV1PostFileV1Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/images"] = user_images_v1.NewPostUserImagesV1(o.context, o.UserImagesV1PostUserImagesV1Handler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/comments/{comment_id}"] = comments_v1.NewRemoveComnentV1(o.context, o.CommentsV1RemoveComnentV1Handler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v1/groups/{group_id}/members"] = groups_v1.NewRemoveGroupUsersV1(o.context, o.GroupsV1RemoveGroupUsersV1Handler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v1/notes/{id}"] = notes_v1.NewUpdateNote(o.context, o.NotesV1UpdateNoteHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/v1/password/{password}"] = user_mgmt_v1.NewUpdatePasswordV1(o.context, o.UserMgmtV1UpdatePasswordV1Handler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/v1/user"] = user_mgmt_v1.NewUpdateUserInfoV1(o.context, o.UserMgmtV1UpdateUserInfoV1Handler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v1/notes"] = notes_v1.NewUploadNoteV1(o.context, o.NotesV1UploadNoteV1Handler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CoreapiAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CoreapiAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CoreapiAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CoreapiAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *CoreapiAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
