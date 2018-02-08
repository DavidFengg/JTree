package restapi

import (
	"crypto/tls"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
	"github.com/Bio-core/jtree/repos"
	keycloak "github.com/Bio-core/keycloakgo"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/Bio-core/jtree/restapi/operations"
)

var lastPatientID int64
var patientLock = &sync.Mutex{}

var lastSampleID int64
var sampleLock = &sync.Mutex{}

func newPatientID() int64 {
	return atomic.AddInt64(&lastPatientID, 1)
}

func newSampleID() int64 {
	return atomic.AddInt64(&lastSampleID, 1)
}

func addPatient(patient *models.Patient) error {
	if patient == nil {
		return errors.New(500, "item must be present")
	}

	patientLock.Lock()
	defer patientLock.Unlock()

	var newID = newPatientID()
	var newIDString = strconv.FormatInt(newID, 10)
	patient.PatientID = &newIDString
	repos.InsertPatient(patient)

	return nil
}

func addSample(sample *models.Sample) error {
	if sample == nil {
		return errors.New(500, "item must be present")
	}

	sampleLock.Lock()
	defer sampleLock.Unlock()

	var newID = newSampleID()
	var newIDString = strconv.FormatInt(newID, 10)
	sample.SampleID = &newIDString
	repos.InsertSample(sample)

	return nil
}

func allPatients(query string) (result []*models.Patient) {
	if query == "search" || query == "" {
		query = "SELECT * FROM Patients"
	}
	patients := repos.GetAllPatients(query)
	result = make([]*models.Patient, 0)
	for _, item := range patients {
		result = append(result, item)
	}
	return
}

func allSamples(query string) (result []*models.Record) {
	if query == "search" || query == "" {
		query = "SELECT * FROM Samples"
	}
	list := repos.GetAllSamples(query)
	result = make([]*models.Record, 0)
	for _, item := range list {
		result = append(result, item)
	}
	return
}

func getSamplesByQuery(query *models.Query) []*models.Record {
	queryString := database.BuildQuery(*query)
	return allSamples(queryString)
}

func getColumns() []string {
	return database.GetColumns(database.GetTables())
}

func logout() models.LogoutOKBody {
	keycloak.LogoutUser()
	return true
}

func configureFlags(api *operations.JtreeMetadataAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.JtreeMetadataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	database.Init("", "")

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AddPatientHandler = operations.AddPatientHandlerFunc(func(params operations.AddPatientParams) middleware.Responder {
		if err := addPatient(params.Patient); err != nil {
			return operations.NewAddPatientBadRequest()
		}
		return operations.NewAddPatientCreated()
	})
	api.AddSampleHandler = operations.AddSampleHandlerFunc(func(params operations.AddSampleParams) middleware.Responder {
		if err := addSample(params.Sample); err != nil {
			return operations.NewAddSampleBadRequest()
		}
		return operations.NewAddSampleCreated()
	})
	api.GetPatientHandler = operations.GetPatientHandlerFunc(func(params operations.GetPatientParams) middleware.Responder {
		return operations.NewGetPatientOK().WithPayload(allPatients(params.PatientID))
	})
	api.GetSampleHandler = operations.GetSampleHandlerFunc(func(params operations.GetSampleParams) middleware.Responder {
		return operations.NewGetSampleOK().WithPayload(allSamples(params.SampleID))
	})
	api.SearchPatientHandler = operations.SearchPatientHandlerFunc(func(params operations.SearchPatientParams) middleware.Responder {
		return operations.NewGetPatientOK().WithPayload(allPatients(""))
	})
	api.SearchSampleHandler = operations.SearchSampleHandlerFunc(func(params operations.SearchSampleParams) middleware.Responder {
		return operations.NewGetSampleOK().WithPayload(allSamples(""))
	})
	api.GetSamplesByQueryHandler = operations.GetSamplesByQueryHandlerFunc(func(params operations.GetSamplesByQueryParams) middleware.Responder {
		return operations.NewGetSampleOK().WithPayload(getSamplesByQuery(params.Query))
	})
	api.GetSampleColumnsHandler = operations.GetSampleColumnsHandlerFunc(func(params operations.GetSampleColumnsParams) middleware.Responder {
		return operations.NewGetSampleColumnsOK().WithPayload(getColumns())
	})

	api.LogoutHandler = operations.LogoutHandlerFunc(func(params operations.LogoutParams) middleware.Responder {
		return operations.NewLogoutOK().WithPayload(logout())
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return keycloak.AuthMiddlewareHandler(handler)
}
