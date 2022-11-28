//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/semi-technologies/weaviate/models"
)

// BatchObjectsDeleteOKCode is the HTTP code returned for type BatchObjectsDeleteOK
const BatchObjectsDeleteOKCode int = 200

/*
BatchObjectsDeleteOK Request succeeded, see response body to get detailed information about each batched item.

swagger:response batchObjectsDeleteOK
*/
type BatchObjectsDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.BatchDeleteResponse `json:"body,omitempty"`
}

// NewBatchObjectsDeleteOK creates BatchObjectsDeleteOK with default headers values
func NewBatchObjectsDeleteOK() *BatchObjectsDeleteOK {

	return &BatchObjectsDeleteOK{}
}

// WithPayload adds the payload to the batch objects delete o k response
func (o *BatchObjectsDeleteOK) WithPayload(payload *models.BatchDeleteResponse) *BatchObjectsDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects delete o k response
func (o *BatchObjectsDeleteOK) SetPayload(payload *models.BatchDeleteResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsDeleteUnauthorizedCode is the HTTP code returned for type BatchObjectsDeleteUnauthorized
const BatchObjectsDeleteUnauthorizedCode int = 401

/*
BatchObjectsDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response batchObjectsDeleteUnauthorized
*/
type BatchObjectsDeleteUnauthorized struct {
}

// NewBatchObjectsDeleteUnauthorized creates BatchObjectsDeleteUnauthorized with default headers values
func NewBatchObjectsDeleteUnauthorized() *BatchObjectsDeleteUnauthorized {

	return &BatchObjectsDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *BatchObjectsDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BatchObjectsDeleteForbiddenCode is the HTTP code returned for type BatchObjectsDeleteForbidden
const BatchObjectsDeleteForbiddenCode int = 403

/*
BatchObjectsDeleteForbidden Forbidden

swagger:response batchObjectsDeleteForbidden
*/
type BatchObjectsDeleteForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsDeleteForbidden creates BatchObjectsDeleteForbidden with default headers values
func NewBatchObjectsDeleteForbidden() *BatchObjectsDeleteForbidden {

	return &BatchObjectsDeleteForbidden{}
}

// WithPayload adds the payload to the batch objects delete forbidden response
func (o *BatchObjectsDeleteForbidden) WithPayload(payload *models.ErrorResponse) *BatchObjectsDeleteForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects delete forbidden response
func (o *BatchObjectsDeleteForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsDeleteUnprocessableEntityCode is the HTTP code returned for type BatchObjectsDeleteUnprocessableEntity
const BatchObjectsDeleteUnprocessableEntityCode int = 422

/*
BatchObjectsDeleteUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response batchObjectsDeleteUnprocessableEntity
*/
type BatchObjectsDeleteUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsDeleteUnprocessableEntity creates BatchObjectsDeleteUnprocessableEntity with default headers values
func NewBatchObjectsDeleteUnprocessableEntity() *BatchObjectsDeleteUnprocessableEntity {

	return &BatchObjectsDeleteUnprocessableEntity{}
}

// WithPayload adds the payload to the batch objects delete unprocessable entity response
func (o *BatchObjectsDeleteUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *BatchObjectsDeleteUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects delete unprocessable entity response
func (o *BatchObjectsDeleteUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsDeleteUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BatchObjectsDeleteInternalServerErrorCode is the HTTP code returned for type BatchObjectsDeleteInternalServerError
const BatchObjectsDeleteInternalServerErrorCode int = 500

/*
BatchObjectsDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response batchObjectsDeleteInternalServerError
*/
type BatchObjectsDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewBatchObjectsDeleteInternalServerError creates BatchObjectsDeleteInternalServerError with default headers values
func NewBatchObjectsDeleteInternalServerError() *BatchObjectsDeleteInternalServerError {

	return &BatchObjectsDeleteInternalServerError{}
}

// WithPayload adds the payload to the batch objects delete internal server error response
func (o *BatchObjectsDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *BatchObjectsDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the batch objects delete internal server error response
func (o *BatchObjectsDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BatchObjectsDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
