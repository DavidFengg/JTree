// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewJtreeMetadataAPI creates a new JtreeMetadata instance
func NewJtreeMetadataAPI(spec *loads.Document) *JtreeMetadataAPI {
	return &JtreeMetadataAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		customConsumers:       make(map[string]runtime.Consumer),
		customProducers:       make(map[string]runtime.Producer),
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		BasicAuthenticator:    security.BasicAuth,
		APIKeyAuthenticator:   security.APIKeyAuth,
		BearerAuthenticator:   security.BearerAuth,
		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		JSONProducer:          runtime.JSONProducer(),
		PostUploadHandler: PostUploadHandlerFunc(func(params PostUploadParams) middleware.Responder {
			return middleware.NotImplemented("operation PostUpload has not yet been implemented")
		}),
		AddExperimentHandler: AddExperimentHandlerFunc(func(params AddExperimentParams) middleware.Responder {
			return middleware.NotImplemented("operation AddExperiment has not yet been implemented")
		}),
		AddPatientHandler: AddPatientHandlerFunc(func(params AddPatientParams) middleware.Responder {
			return middleware.NotImplemented("operation AddPatient has not yet been implemented")
		}),
		AddResultHandler: AddResultHandlerFunc(func(params AddResultParams) middleware.Responder {
			return middleware.NotImplemented("operation AddResult has not yet been implemented")
		}),
		AddResultdetailsHandler: AddResultdetailsHandlerFunc(func(params AddResultdetailsParams) middleware.Responder {
			return middleware.NotImplemented("operation AddResultdetails has not yet been implemented")
		}),
		AddSampleHandler: AddSampleHandlerFunc(func(params AddSampleParams) middleware.Responder {
			return middleware.NotImplemented("operation AddSample has not yet been implemented")
		}),
		DeletePatientHandler: DeletePatientHandlerFunc(func(params DeletePatientParams) middleware.Responder {
			return middleware.NotImplemented("operation DeletePatient has not yet been implemented")
		}),
		GetSampleColumnsHandler: GetSampleColumnsHandlerFunc(func(params GetSampleColumnsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSampleColumns has not yet been implemented")
		}),
		GetSamplesByQueryHandler: GetSamplesByQueryHandlerFunc(func(params GetSamplesByQueryParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSamplesByQuery has not yet been implemented")
		}),
		LogoutHandler: LogoutHandlerFunc(func(params LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation Logout has not yet been implemented")
		}),
		UpdatePatientHandler: UpdatePatientHandlerFunc(func(params UpdatePatientParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdatePatient has not yet been implemented")
		}),
	}
}

/*JtreeMetadataAPI Metadata API */
type JtreeMetadataAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// PostUploadHandler sets the operation handler for the post upload operation
	PostUploadHandler PostUploadHandler
	// AddExperimentHandler sets the operation handler for the add experiment operation
	AddExperimentHandler AddExperimentHandler
	// AddPatientHandler sets the operation handler for the add patient operation
	AddPatientHandler AddPatientHandler
	// AddResultHandler sets the operation handler for the add result operation
	AddResultHandler AddResultHandler
	// AddResultdetailsHandler sets the operation handler for the add resultdetails operation
	AddResultdetailsHandler AddResultdetailsHandler
	// AddSampleHandler sets the operation handler for the add sample operation
	AddSampleHandler AddSampleHandler
	// DeletePatientHandler sets the operation handler for the delete patient operation
	DeletePatientHandler DeletePatientHandler
	// GetSampleColumnsHandler sets the operation handler for the get sample columns operation
	GetSampleColumnsHandler GetSampleColumnsHandler
	// GetSamplesByQueryHandler sets the operation handler for the get samples by query operation
	GetSamplesByQueryHandler GetSamplesByQueryHandler
	// GetSearchableHandler sets the operation handler for the get searchable operation
	GetSearchableHandler GetSearchableHandler
	// GetUneditableHandler sets the operation handler for the get uneditable operation
	GetUneditableHandler GetUneditableHandler
	// LogoutHandler sets the operation handler for the logout operation
	LogoutHandler LogoutHandler
	// UpdatePatientHandler sets the operation handler for the update patient operation
	UpdatePatientHandler UpdatePatientHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *JtreeMetadataAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *JtreeMetadataAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *JtreeMetadataAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *JtreeMetadataAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *JtreeMetadataAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *JtreeMetadataAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *JtreeMetadataAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the JtreeMetadataAPI
func (o *JtreeMetadataAPI) Validate() error {
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

	if o.PostUploadHandler == nil {
		unregistered = append(unregistered, "PostUploadHandler")
	}

	if o.AddExperimentHandler == nil {
		unregistered = append(unregistered, "AddExperimentHandler")
	}

	if o.AddPatientHandler == nil {
		unregistered = append(unregistered, "AddPatientHandler")
	}

	if o.AddResultHandler == nil {
		unregistered = append(unregistered, "AddResultHandler")
	}

	if o.AddResultdetailsHandler == nil {
		unregistered = append(unregistered, "AddResultdetailsHandler")
	}

	if o.AddSampleHandler == nil {
		unregistered = append(unregistered, "AddSampleHandler")
	}

	if o.DeletePatientHandler == nil {
		unregistered = append(unregistered, "DeletePatientHandler")
	}

	if o.GetSampleColumnsHandler == nil {
		unregistered = append(unregistered, "GetSampleColumnsHandler")
	}

	if o.GetSamplesByQueryHandler == nil {
		unregistered = append(unregistered, "GetSamplesByQueryHandler")
	}

	if o.LogoutHandler == nil {
		unregistered = append(unregistered, "LogoutHandler")
	}

	if o.UpdatePatientHandler == nil {
		unregistered = append(unregistered, "UpdatePatientHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *JtreeMetadataAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *JtreeMetadataAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *JtreeMetadataAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *JtreeMetadataAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
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

// ProducersFor gets the producers for the specified media types
func (o *JtreeMetadataAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
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
func (o *JtreeMetadataAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the jtree metadata API
func (o *JtreeMetadataAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *JtreeMetadataAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/upload"] = NewPostUpload(o.context, o.PostUploadHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/experiment"] = NewAddExperiment(o.context, o.AddExperimentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/patient"] = NewAddPatient(o.context, o.AddPatientHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/result"] = NewAddResult(o.context, o.AddResultHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resultdetails"] = NewAddResultdetails(o.context, o.AddResultdetailsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sample"] = NewAddSample(o.context, o.AddSampleHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/patient/{id}"] = NewDeletePatient(o.context, o.DeletePatientHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/columns"] = NewGetSampleColumns(o.context, o.GetSampleColumnsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/query"] = NewGetSamplesByQuery(o.context, o.GetSamplesByQueryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/logout"] = NewLogout(o.context, o.LogoutHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/patient/{id}"] = NewUpdatePatient(o.context, o.UpdatePatientHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *JtreeMetadataAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *JtreeMetadataAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *JtreeMetadataAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *JtreeMetadataAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
