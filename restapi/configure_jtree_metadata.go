package restapi

import (
	"crypto/tls"
	localerrors "errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	config "github.com/Bio-core/jtree/conf"
	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/dummydata"
	"github.com/Bio-core/jtree/models"
	"github.com/Bio-core/jtree/repos"
	keycloak "github.com/Bio-core/keycloakgo"
	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	"github.com/Bio-core/jtree/restapi/operations"
	"github.com/rs/cors"
)

var lastPatientID int64
var patientLock = &sync.Mutex{}
var lastSampleID int64
var sampleLock = &sync.Mutex{}
var c config.Conf

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
	if patient.PatientID == nil {
		patient.PatientID = &newIDString
	}
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
	if sample.SampleID == nil {
		sample.SampleID = &newIDString
	}
	repos.InsertSample(sample)

	return nil
}

func addExperiment(experiment *models.Experiment) error {
	if experiment == nil {
		return errors.New(500, "item must be present")
	}

	sampleLock.Lock()
	defer sampleLock.Unlock()

	var newID = newSampleID()
	var newIDString = strconv.FormatInt(newID, 10)
	if experiment.ExperimentID == nil {
		experiment.ExperimentID = &newIDString
	}
	repos.InsertExperiment(experiment)

	return nil
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

func getColumns() [][]string {
	columns := database.GetColumns(database.GetTables())
	columnArray := make([][]string, len(columns))
	for i, column := range columns {
		columnArray[i] = make([]string, 2)
		columnArray[i][0] = column
		columnArray[i][1] = database.Map[column]
	}
	return columnArray
}

func logout() bool {
	return true
}

func upload(file operations.PostUploadParams) error {
	if _, err := os.Stat("./uploads/" + file.Filename); !os.IsNotExist(err) {
		return localerrors.New("File already exists")
	}
	f, err := os.OpenFile("./uploads/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = io.Copy(f, file.Upfile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

var databaseFlags = struct {
	Host       string `long:"databaseHost" description:"Database Host" required:"false"`
	Name       string `long:"databaseName" description:"Database Name" required:"false"`
	SelectUser string `long:"dbUsernameSelect" description:"Database Username for Select" required:"false"`
	SelectPass string `long:"dbPasswordSelect" description:"Database Password for Select" required:"false"`
	UpdateUser string `long:"dbUsernameUpdate" description:"Database Username for Update" required:"false"`
	UpdatePass string `long:"dbPasswordUpdate" description:"Database Password for Update" required:"false"`
}{}
var keycloakFlags = struct {
	Active bool   `short:"s" description:"Use Security Bool" required:"false"`
	Host   string `long:"keycloakHost" description:"Keycloak Host" required:"false"`
}{}
var dataGenFlags = struct {
	Generate int `short:"g" description:"generate data" required:"false"`
}{}

func configureFlags(api *operations.JtreeMetadataAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Database Flags",
			LongDescription:  "",
			Options:          &databaseFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "Keycloak Flags",
			LongDescription:  "",
			Options:          &keycloakFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "Data Generation Flags",
			LongDescription:  "",
			Options:          &dataGenFlags,
		},
	}
}

func configureAPI(api *operations.JtreeMetadataAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	c.GetConf()
	setupOptions()
	models.Enums = models.GetEnums(models.Enums)
	database.Map = database.MapSuper()

	database.DBSelect = database.Init(c.Database.Host, c.Database.Selectuser+":"+c.Database.Selectpass+"@/"+c.Database.Name, database.DBSelect)
	database.DBUpdate = database.Init(c.Database.Host, c.Database.Updateuser+":"+c.Database.Updatepass+"@/"+c.Database.Name, database.DBUpdate)
	ServerName := c.App.Host + ":" + strconv.Itoa(c.App.Port)
	KeycloakserverName := c.Keycloak.Host

	if keycloakFlags.Active {
		keycloak.Init(KeycloakserverName, "http://"+ServerName, "/Jtree/metadata/0.1.0/columns", "/Jtree/metadata/0.1.0/logout")
	}
	if dataGenFlags.Generate != 0 {
		dummydata.MakeData(dataGenFlags.Generate)
	}
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})

	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.PostUploadHandler = operations.PostUploadHandlerFunc(func(params operations.PostUploadParams) middleware.Responder {
		if err := upload(params); err != nil {
			return operations.NewPostUploadConflict()
		}
		return operations.NewPostUploadOK().WithPayload(true)
	})
	api.AddExperimentHandler = operations.AddExperimentHandlerFunc(func(params operations.AddExperimentParams) middleware.Responder {
		if err := addExperiment(params.Experiment); err != nil {
			return operations.NewAddExperimentBadRequest()
		}
		return operations.NewAddExperimentCreated()
	})
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
	api.GetSamplesByQueryHandler = operations.GetSamplesByQueryHandlerFunc(func(params operations.GetSamplesByQueryParams) middleware.Responder {
		return operations.NewGetSamplesByQueryOK().WithPayload(getSamplesByQuery(params.Query))
	})
	api.LogoutHandler = operations.LogoutHandlerFunc(func(params operations.LogoutParams) middleware.Responder {
		return operations.NewLogoutOK().WithPayload(logout())
	})
	api.GetSampleColumnsHandler = operations.GetSampleColumnsHandlerFunc(func(params operations.GetSampleColumnsParams) middleware.Responder {
		return operations.NewGetSampleColumnsOK().WithPayload(getColumns())
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
	x := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
	})
	handler = x.Handler(handler)
	if keycloakFlags.Active {
		return keycloak.AuthMiddlewareHandler(handler)
	}
	return handler
}

func setupOptions() {
	if databaseFlags.Host != "" {
		c.Database.Host = databaseFlags.Host
	}
	if databaseFlags.Name != "" {
		c.Database.Name = databaseFlags.Name
	}
	if databaseFlags.SelectUser != "" {
		c.Database.Selectuser = databaseFlags.SelectUser
	}
	if databaseFlags.SelectPass != "" {
		c.Database.Selectpass = databaseFlags.SelectPass
	}
	if databaseFlags.UpdateUser != "" {
		c.Database.Updateuser = databaseFlags.UpdateUser
	}
	if databaseFlags.UpdatePass != "" {
		c.Database.Updatepass = databaseFlags.UpdatePass
	}
	if keycloakFlags.Host != "" {
		c.Keycloak.Host = keycloakFlags.Host
	}
}
