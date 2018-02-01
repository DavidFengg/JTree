package tests

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/CanDIG/candig_mds/restapi"
	"github.com/CanDIG/candig_mds/restapi/operations"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

const server = "http://localhost:8000"

func TestMain(m *testing.M) {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCandigMetadataAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Candig Metadata API"
	parser.LongDescription = "Metadata API"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

func TestUrls(t *testing.T) {
	result := true
	result = result && CheckPageResponse(server+"/CanDIG/metadata/0.1.0/biosample/search")
	result = result && CheckPageResponse(server+"/CanDIG/metadata/0.1.0/individual/search")
	result = result && CheckNoPageResponse(server+"/x")

	if result != true {
		t.Fail()
	}
}

//CheckPageResponse checks if a page that should respond is found correctly
func CheckPageResponse(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	if response == nil {
		return false
	}
	if response.Status == "404 Not Found" {
		return false
	}
	return true
}

//CheckNoPageResponse checks if a page that does not exist responds with a 404 Error
func CheckNoPageResponse(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return true
	}
	if response == nil {
		return true
	}
	if response.Status == "404 Not Found" {
		return true
	}
	return false
}
