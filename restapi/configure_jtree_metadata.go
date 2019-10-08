package restapi

import (
	"crypto/tls"
	localerrors "errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// "os/exec"
	"strconv"

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
	uuid "github.com/nu7hatch/gouuid"
	"github.com/rs/cors"
)

var c config.Conf

func newID() string {
	// below command does not work since JTree is not using a linux container
	// out, err := exec.Command("uuidgen").Output()
	out, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	ID := fmt.Sprintf("%s", out)
	ID = ID[:len(ID)-1]
	return ID
}

func addPatient(patient *models.Patient) string {
	// return error if empty patient or specified patient id
	if patient == nil || patient.PatientID != nil {
		return "error"
	}

	// insert new patient
	NewID := newID()
	patient.PatientID = &NewID
	repos.InsertPatient(patient)
	return NewID
}

func updatePatient(patientID string, patient *models.Patient) string {
	// return error if empty patient model
	if patient == nil {
		return "error"
	}

	patientOLD := repos.GetPatientByID(patientID)

	if patientOLD == nil {
		return "error, patient_id not in db"
	}
	repos.UpdatePatients(patientID, patient)
	return patientID
}

func deletePatient(patientID string) string {
	// return error if patient id not specified
	if patientID == "" {
		return "error"
	}

	existingID := repos.GetPatientByID(patientID)

	if existingID == nil {
		return "error"
	}

	repos.DeletePatient(patientID)
	return patientID
}

func addSample(sample *models.Sample) string {
	// return error if empty sample or patient id not specified
	if sample == nil || *sample.PatientID == "" {
		return "error"
	}

	// insert new sample
	NewID := newID()
	sample.SampleID = &NewID
	repos.InsertSample(sample)
	return NewID
}

func updateSample(sampleID string, sample *models.Sample) string {
	// return error if empty sample model
	if sample == nil {
		return "error"
	}

	sampleOLD := repos.GetSampleByID(sampleID)

	if sampleOLD == nil {
		return "error"
	}
	repos.UpdateSample(sampleID, sample)
	return sampleID
}

func deleteSample(sampleID string) string {
	// return error if sample id not specified
	if sampleID == "" {
		return "error"
	}

	existingID := repos.GetSampleByID(sampleID)

	if existingID == nil {
		return "error"
	}

	repos.DeleteSample(sampleID)
	return sampleID
}

func addExperiment(experiment *models.Experiment) string {
	// return error if empty experiment or sample id not specified
	if experiment == nil || *experiment.SampleID == "" {
		return "error"
	}

	// insert new experiment
	NewID := newID()
	experiment.ExperimentID = &NewID
	repos.InsertExperiment(experiment)
	return NewID
}

func updateExperiment(experimentID string, experiment *models.Experiment) string {
	// return error if empty experiment model
	if experiment == nil {
		return "error"
	}

	experimentOLD := repos.GetExperimentByID(experimentID)

	if experimentOLD == nil {
		return "error"
	}
	repos.UpdateExperiment(experimentID, experiment)
	return experimentID

}

func deleteExperiment(experimentID string) string {
	// return error if experiment id not specified
	if experimentID == "" {
		return "error"
	}

	existingID := repos.GetExperimentByID(experimentID)

	if existingID == nil {
		return "error"
	}

	repos.DeleteExperiment(experimentID)
	return experimentID
}

func addResult(result *models.Result) string {
	// return error if empty result or experiment id not specified
	if result == nil || *result.ExperimentID == "" {
		return "error"
	}

	// insert new result
	NewID := newID()
	result.ResultsID = &NewID
	repos.InsertResult(result)
	return NewID
}

func updateResult(resultID string, result *models.Result) string {
	// return error if empty result model
	if result == nil {
		return "error"
	}

	resultOLD := repos.GetResultByID(resultID)

	if resultOLD == nil {
		return "error"
	}
	repos.UpdateResult(resultID, result)
	return resultID
}

func deleteResult(resultID string) string {
	// return error if result id not specified
	if resultID == "" {
		return "error"
	}

	existingID := repos.GetResultByID(resultID)

	if existingID == nil {
		return "error"
	}

	repos.DeleteResult(resultID)
	return resultID
}

func addResultdetail(resultdetail *models.Resultdetails) string {
	// return error if empty resultdetail or result id not specified
	if resultdetail == nil || *resultdetail.ResultsID == "" {
		return "error"
	}

	// insert new resultdetail
	NewID := newID()
	resultdetail.ResultsDetailsID = &NewID
	repos.InsertResultDetail(resultdetail)
	return NewID
}

func updateResultdetail(results_details_id string, resultdetail *models.Resultdetails) string {
	// return error if empty resultdetail model
	if resultdetail == nil {
		return "error"
	}

	resultOLD := repos.GetResultDetailByID(results_details_id)

	if resultOLD == nil {
		return "error"
	}
	repos.UpdateResultDetail(results_details_id, resultdetail)
	return results_details_id
}

func deleteResultdetail(result_details_id string) string {
	// return error if resultdetails id not specified
	if result_details_id == "" {
		return "error"
	}

	existingID := repos.GetResultDetailByID(result_details_id)

	if existingID == nil {
		return "error"
	}

	repos.DeleteResultDetail(result_details_id)
	return result_details_id
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

func getSearchable() []string {
	return models.Sefields.Searchable
}

func getUneditable() []string {
	return models.Sefields.Uneditable
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
	models.Sefields = &models.SEFields{}
	models.Sefields = models.Sefields.GetSEFields()

	database.Map = database.MapSuper()

	//database.DBSelect = database.Init(c.Database.Host, c.Database.Selectuser+":"+c.Database.Selectpass+"@/"+c.Database.Name+"?parseTime=true", database.DBSelect)
	//database.DBUpdate = database.Init(c.Database.Host, c.Database.Updateuser+":"+c.Database.Updatepass+"@/"+c.Database.Name+"?parseTime=true", database.DBUpdate)
	// database.DBSelect = database.Init(c.Database.Host, c.Database.Selectuser+":"+c.Database.Selectpass+"@tcp(172.23.0.2:3306)/"+c.Database.Name+"?parseTime=true", database.DBSelect)
	// database.DBUpdate = database.Init(c.Database.Host, c.Database.Updateuser+":"+c.Database.Updatepass+"@tcp(172.23.0.2:3306)/"+c.Database.Name+"?parseTime=true", database.DBUpdate)
	
	database.DBSelect = database.Init(c.Database.Host, c.Database.Selectuser+":"+c.Database.Selectpass+"@tcp(mysql:3306)/"+c.Database.Name+"?parseTime=true", database.DBSelect)
	database.DBUpdate = database.Init(c.Database.Host, c.Database.Updateuser+":"+c.Database.Updatepass+"@tcp(mysql:3306)/"+c.Database.Name+"?parseTime=true", database.DBUpdate)
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

	// endpoint: /patient
	api.AddPatientHandler = operations.AddPatientHandlerFunc(func(params operations.AddPatientParams) middleware.Responder {
		return operations.NewAddPatientOK().WithPayload(addPatient(params.Patient))
	})
	api.UpdatePatientHandler = operations.UpdatePatientHandlerFunc(func(params operations.UpdatePatientParams) middleware.Responder {
		return operations.NewUpdatePatientCreated().WithPayload(updatePatient(params.ID, params.Patient))
	})
	api.DeletePatientHandler = operations.DeletePatientHandlerFunc(func(params operations.DeletePatientParams) middleware.Responder {
		return operations.NewDeletePatientOK().WithPayload(deletePatient(params.ID))
	})

	// endpoint: /sample
	api.AddSampleHandler = operations.AddSampleHandlerFunc(func(params operations.AddSampleParams) middleware.Responder {
		return operations.NewAddSampleOK().WithPayload(addSample(params.Sample))
	})
	api.UpdateSampleHandler = operations.UpdateSampleHandlerFunc(func(params operations.UpdateSampleParams) middleware.Responder {
		return operations.NewUpdateSampleOK().WithPayload(updateSample(params.ID, params.Sample))
	})
	api.DeleteSampleHandler = operations.DeleteSampleHandlerFunc(func(params operations.DeleteSampleParams) middleware.Responder {
		return operations.NewDeleteSampleOK().WithPayload(deleteSample(params.ID))
	})

	// endpoint: /experiment
	api.AddExperimentHandler = operations.AddExperimentHandlerFunc(func(params operations.AddExperimentParams) middleware.Responder {
		return operations.NewAddExperimentOK().WithPayload(addExperiment(params.Experiment))
	})
	api.UpdateExperimentHandler = operations.UpdateExperimentHandlerFunc(func(params operations.UpdateExperimentParams) middleware.Responder {
		return operations.NewUpdateExperimentOK().WithPayload(updateExperiment(params.ID, params.Experiment))
	})
	api.DeleteExperimentHandler = operations.DeleteExperimentHandlerFunc(func(params operations.DeleteExperimentParams) middleware.Responder {
		return operations.NewDeleteExperimentOK().WithPayload(deleteExperiment(params.ID))
	})

	// endpoint: /result
	api.AddResultHandler = operations.AddResultHandlerFunc(func(params operations.AddResultParams) middleware.Responder {
		return operations.NewAddResultOK().WithPayload(addResult(params.Result))
	})
	api.UpdateResultHandler = operations.UpdateResultHandlerFunc(func(params operations.UpdateResultParams) middleware.Responder {
		return operations.NewUpdateResultOK().WithPayload(updateResult(params.ID, params.Result))
	})
	api.DeleteResultHandler = operations.DeleteResultHandlerFunc(func(params operations.DeleteResultParams) middleware.Responder {
		return operations.NewDeleteResultOK().WithPayload(deleteResult(params.ID))
	})

	// endpoint: /resultdetails
	api.AddResultdetailsHandler = operations.AddResultdetailsHandlerFunc(func(params operations.AddResultdetailsParams) middleware.Responder {
		return operations.NewAddResultdetailsOK().WithPayload(addResultdetail(params.Resultdetails))
	})
	api.UpdateResultdetailsHandler = operations.UpdateResultdetailsHandlerFunc(func(params operations.UpdateResultdetailsParams) middleware.Responder {
		return operations.NewUpdateResultdetailsOK().WithPayload(updateResultdetail(params.ID, params.Resultdetails))
	})
	api.DeleteResultdetailsHandler = operations.DeleteResultdetailsHandlerFunc(func(params operations.DeleteResultdetailsParams) middleware.Responder {
		return operations.NewDeleteResultdetailsOK().WithPayload(deleteResultdetail(params.ID))
	})
	
	api.GetSamplesByQueryHandler = operations.GetSamplesByQueryHandlerFunc(func(params operations.GetSamplesByQueryParams) middleware.Responder {
		return operations.NewGetSamplesByQueryOK().WithPayload(getSamplesByQuery(params.Query))
	})
	api.LogoutHandler = operations.LogoutHandlerFunc(func(params operations.LogoutParams) middleware.Responder {
		return operations.NewLogoutOK().WithPayload(logout())
	})

	// Return database fields
	api.GetSampleColumnsHandler = operations.GetSampleColumnsHandlerFunc(func(params operations.GetSampleColumnsParams) middleware.Responder {
		return operations.NewGetSampleColumnsOK().WithPayload(getColumns())
	})
	api.GetSearchableHandler = operations.GetSearchableHandlerFunc(func(params operations.GetSearchableParams) middleware.Responder {
		return operations.NewGetSearchableOK().WithPayload(getSearchable())
	})
	api.GetUneditableHandler = operations.GetUneditableHandlerFunc(func(params operations.GetUneditableParams) middleware.Responder {
		return operations.NewGetUneditableOK().WithPayload(getUneditable())
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
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
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
