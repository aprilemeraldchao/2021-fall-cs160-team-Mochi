// Code generated by go-swagger; DO NOT EDIT.

package comments_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// PostCommentsV1OKCode is the HTTP code returned for type PostCommentsV1OK
const PostCommentsV1OKCode int = 200

/*PostCommentsV1OK Success

swagger:response postCommentsV1OK
*/
type PostCommentsV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.CommentResponse `json:"body,omitempty"`
}

// NewPostCommentsV1OK creates PostCommentsV1OK with default headers values
func NewPostCommentsV1OK() *PostCommentsV1OK {

	return &PostCommentsV1OK{}
}

// WithPayload adds the payload to the post comments v1 o k response
func (o *PostCommentsV1OK) WithPayload(payload *models.CommentResponse) *PostCommentsV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments v1 o k response
func (o *PostCommentsV1OK) SetPayload(payload *models.CommentResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCommentsV1BadRequestCode is the HTTP code returned for type PostCommentsV1BadRequest
const PostCommentsV1BadRequestCode int = 400

/*PostCommentsV1BadRequest Bad Request

swagger:response postCommentsV1BadRequest
*/
type PostCommentsV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewPostCommentsV1BadRequest creates PostCommentsV1BadRequest with default headers values
func NewPostCommentsV1BadRequest() *PostCommentsV1BadRequest {

	return &PostCommentsV1BadRequest{}
}

// WithPayload adds the payload to the post comments v1 bad request response
func (o *PostCommentsV1BadRequest) WithPayload(payload *models.ErrResponse) *PostCommentsV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments v1 bad request response
func (o *PostCommentsV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCommentsV1UnauthorizedCode is the HTTP code returned for type PostCommentsV1Unauthorized
const PostCommentsV1UnauthorizedCode int = 401

/*PostCommentsV1Unauthorized Unauthorized

swagger:response postCommentsV1Unauthorized
*/
type PostCommentsV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewPostCommentsV1Unauthorized creates PostCommentsV1Unauthorized with default headers values
func NewPostCommentsV1Unauthorized() *PostCommentsV1Unauthorized {

	return &PostCommentsV1Unauthorized{}
}

// WithPayload adds the payload to the post comments v1 unauthorized response
func (o *PostCommentsV1Unauthorized) WithPayload(payload *models.ErrResponse) *PostCommentsV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments v1 unauthorized response
func (o *PostCommentsV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCommentsV1ForbiddenCode is the HTTP code returned for type PostCommentsV1Forbidden
const PostCommentsV1ForbiddenCode int = 403

/*PostCommentsV1Forbidden Forbidden

swagger:response postCommentsV1Forbidden
*/
type PostCommentsV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewPostCommentsV1Forbidden creates PostCommentsV1Forbidden with default headers values
func NewPostCommentsV1Forbidden() *PostCommentsV1Forbidden {

	return &PostCommentsV1Forbidden{}
}

// WithPayload adds the payload to the post comments v1 forbidden response
func (o *PostCommentsV1Forbidden) WithPayload(payload *models.ErrResponse) *PostCommentsV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments v1 forbidden response
func (o *PostCommentsV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCommentsV1NotFoundCode is the HTTP code returned for type PostCommentsV1NotFound
const PostCommentsV1NotFoundCode int = 404

/*PostCommentsV1NotFound Not Found

swagger:response postCommentsV1NotFound
*/
type PostCommentsV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewPostCommentsV1NotFound creates PostCommentsV1NotFound with default headers values
func NewPostCommentsV1NotFound() *PostCommentsV1NotFound {

	return &PostCommentsV1NotFound{}
}

// WithPayload adds the payload to the post comments v1 not found response
func (o *PostCommentsV1NotFound) WithPayload(payload *models.ErrResponse) *PostCommentsV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments v1 not found response
func (o *PostCommentsV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCommentsV1ConflictCode is the HTTP code returned for type PostCommentsV1Conflict
const PostCommentsV1ConflictCode int = 409

/*PostCommentsV1Conflict Conflict

swagger:response postCommentsV1Conflict
*/
type PostCommentsV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewPostCommentsV1Conflict creates PostCommentsV1Conflict with default headers values
func NewPostCommentsV1Conflict() *PostCommentsV1Conflict {

	return &PostCommentsV1Conflict{}
}

// WithPayload adds the payload to the post comments v1 conflict response
func (o *PostCommentsV1Conflict) WithPayload(payload *models.ErrResponse) *PostCommentsV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments v1 conflict response
func (o *PostCommentsV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCommentsV1InternalServerErrorCode is the HTTP code returned for type PostCommentsV1InternalServerError
const PostCommentsV1InternalServerErrorCode int = 500

/*PostCommentsV1InternalServerError Internal Server Error

swagger:response postCommentsV1InternalServerError
*/
type PostCommentsV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewPostCommentsV1InternalServerError creates PostCommentsV1InternalServerError with default headers values
func NewPostCommentsV1InternalServerError() *PostCommentsV1InternalServerError {

	return &PostCommentsV1InternalServerError{}
}

// WithPayload adds the payload to the post comments v1 internal server error response
func (o *PostCommentsV1InternalServerError) WithPayload(payload *models.ErrResponse) *PostCommentsV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post comments v1 internal server error response
func (o *PostCommentsV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCommentsV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}