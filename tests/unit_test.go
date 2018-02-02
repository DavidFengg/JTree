package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	models "github.com/CanDIG/candig_mds/models"
	"github.com/CanDIG/candig_mds/repos"
	"github.com/CanDIG/candig_mds/restapi"
	"github.com/CanDIG/candig_mds/restapi/operations"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

const server = "http://localhost:8000"

func TestMain(m *testing.M) {
	Databasename = "testCandig"
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCandigMetadataAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = 8000
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Candig Metadata API"
	parser.LongDescription = "Metadata API"

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

	go server.Serve()
	testResults := m.Run()
	tearDown()
	os.Exit(testResults)
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

func TestAddBiosamplesPOST(t *testing.T) {
	nameSample1 := "Sample1"
	descriptionSample1 := "This is a test sample"
	idSample1 := "UNIT_TESTS"
	collectionageSample1 := "10"
	nameSample2 := "Sample2"
	descriptionSample2 := "This is another test sample"
	idSample2 := "UNIT_TESTS"
	collectionageSample2 := "20"

	sample1 := models.Biosample{
		Name:          &nameSample1,
		Description:   &descriptionSample1,
		IndividualID:  &idSample1,
		CollectionAge: &collectionageSample1,
	}

	sample2 := models.Biosample{
		Name:          &nameSample2,
		Description:   &descriptionSample2,
		IndividualID:  &idSample2,
		CollectionAge: &collectionageSample2,
	}

	sample1Bytes, err := json.Marshal(sample1)
	sample2Bytes, err2 := json.Marshal(sample2)

	if err != nil {
		t.Fail()
		return
	}

	if err2 != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(sample1Bytes)
	body2 := bytes.NewReader(sample2Bytes)

	req, err := http.NewRequest("POST", server+"/CanDIG/metadata/0.1.0/biosample", body)
	req2, err2 := http.NewRequest("POST", server+"/CanDIG/metadata/0.1.0/biosample", body2)

	if err != nil {
		t.Fail()
		return
	}
	if err2 != nil {
		t.Fail()
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	resp2, err2 := http.DefaultClient.Do(req2)

	if resp.Status != "201 Created" || resp2.Status != "201 Created" {
		t.Fail()
		return
	}

	if err != nil {
		t.Fail()
		return
	}
	if err2 != nil {
		t.Fail()
		return
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()

}
func TestAddIndividualsPOST(t *testing.T) {
	namePerson1 := "Sample1"
	descriptionPerson1 := "This is a test sample"
	namePerson2 := "Sample2"
	descriptionPerson2 := "This is another test sample"

	person1 := models.Individual{
		Name:        &namePerson1,
		Description: &descriptionPerson1,
	}

	person2 := models.Individual{
		Name:        &namePerson2,
		Description: &descriptionPerson2,
	}

	person1Bytes, err := json.Marshal(person1)
	person2Bytes, err2 := json.Marshal(person2)

	if err != nil {
		t.Fail()
		return
	}

	if err2 != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(person1Bytes)
	body2 := bytes.NewReader(person2Bytes)

	req, err := http.NewRequest("POST", server+"/CanDIG/metadata/0.1.0/individual", body)
	req2, err2 := http.NewRequest("POST", server+"/CanDIG/metadata/0.1.0/individual", body2)

	if err != nil {
		t.Fail()
		return
	}
	if err2 != nil {
		t.Fail()
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	resp2, err2 := http.DefaultClient.Do(req2)

	if resp.Status != "201 Created" || resp2.Status != "201 Created" {
		t.Fail()
		return
	}

	if err != nil {
		t.Fail()
		return
	}
	if err2 != nil {
		t.Fail()
		return
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()

}

func TestRemoveAllBiosamples(t *testing.T) {
	result := repos.RemoveAllBiosamples()
	if !result {
		t.Fail()
		return
	}
}

func TestRemoveAllIndividuals(t *testing.T) {
	result := repos.RemoveAllIndividuals()
	if !result {
		t.Fail()
		return
	}
}
