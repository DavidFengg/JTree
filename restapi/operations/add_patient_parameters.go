// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Bio-core/jtree/models"
)

// NewAddPatientParams creates a new AddPatientParams object
// with the default values initialized.
func NewAddPatientParams() AddPatientParams {
	var ()
	return AddPatientParams{}
}

// AddPatientParams contains all the bound params for the add patient operation
// typically these are obtained from a http.Request
//
// swagger:parameters addPatient
type AddPatientParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Patient
	  In: body
	*/
	Patient *models.Patient
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddPatientParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Patient
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("patient", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Patient = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
