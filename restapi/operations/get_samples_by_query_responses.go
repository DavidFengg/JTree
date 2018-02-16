// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Bio-core/jtree/models"
)

// GetSamplesByQueryOKCode is the HTTP code returned for type GetSamplesByQueryOK
const GetSamplesByQueryOKCode int = 200

/*GetSamplesByQueryOK OK

swagger:response getSamplesByQueryOK
*/
type GetSamplesByQueryOK struct {

	/*
	  In: Body
	*/
	Payload models.GetSamplesByQueryOKBody `json:"body,omitempty"`
}

// NewGetSamplesByQueryOK creates GetSamplesByQueryOK with default headers values
func NewGetSamplesByQueryOK() *GetSamplesByQueryOK {
	return &GetSamplesByQueryOK{}
}

// WithPayload adds the payload to the get samples by query o k response
func (o *GetSamplesByQueryOK) WithPayload(payload models.GetSamplesByQueryOKBody) *GetSamplesByQueryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get samples by query o k response
func (o *GetSamplesByQueryOK) SetPayload(payload models.GetSamplesByQueryOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSamplesByQueryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers","Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GetSamplesByQueryOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetSamplesByQueryNotFoundCode is the HTTP code returned for type GetSamplesByQueryNotFound
const GetSamplesByQueryNotFoundCode int = 404

/*GetSamplesByQueryNotFound Sample not found

swagger:response getSamplesByQueryNotFound
*/
type GetSamplesByQueryNotFound struct {
}

// NewGetSamplesByQueryNotFound creates GetSamplesByQueryNotFound with default headers values
func NewGetSamplesByQueryNotFound() *GetSamplesByQueryNotFound {
	return &GetSamplesByQueryNotFound{}
}

// WriteResponse to the client
func (o *GetSamplesByQueryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
