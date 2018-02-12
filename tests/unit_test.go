package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	models "github.com/Bio-core/jtree/models"
	"github.com/Bio-core/jtree/restapi"
	"github.com/Bio-core/jtree/restapi/operations"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

const server = "http://127.0.0.1:8000"

func TestMain(m *testing.M) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewJtreeMetadataAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Jtree Metadata API"
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
	//server.Host = "127.0.0.1"
	go server.Serve()

	testResults := m.Run()
	tearDown()
	os.Exit(testResults)
}

func TestUrls(t *testing.T) {
	result := true
	result = result && CheckPageResponse(server+"/Jtree/metadata/0.1.0/columns")
	result = result && CheckNoPageResponse(server+"/x")

	if result != true {
		t.Fail()
	}
}

func TestAddSamplesPOST(t *testing.T) {
	sampleidSample1 := "Sample1"
	facilitySample1 := "TGH"
	var volumeOfBloodMarrowSample1 float32
	volumeOfBloodMarrowSample1 = 14.2
	dateCollectedSample1 := "20140506"
	sampleidSample2 := "Sample2"
	facilitySample2 := "PMH"
	var volumeOfBloodMarrowSample2 float32
	volumeOfBloodMarrowSample2 = 105.67
	dateCollectedSample2 := "2020-09-08"

	sample1 := models.Sample{
		SampleID:            &sampleidSample1,
		Facility:            &facilitySample1,
		VolumeOfBloodMarrow: &volumeOfBloodMarrowSample1,
		DateCollected:       &dateCollectedSample1,
	}

	sample2 := models.Sample{
		SampleID:            &sampleidSample2,
		Facility:            &facilitySample2,
		VolumeOfBloodMarrow: &volumeOfBloodMarrowSample2,
		DateCollected:       &dateCollectedSample2,
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

	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/sample", body)
	req2, err2 := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/sample", body2)

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
func TestAddPatientsPOST(t *testing.T) {
	namePerson1 := "Mitchell"
	patientidPerson1 := "patient1"
	sampleidPerson1 := "Sample1"
	namePerson2 := "Strong"
	patientidPerson2 := "patient2"
	sampleidPerson2 := "Sample2"

	person1 := models.Patient{
		FirstName: &namePerson1,
		PatientID: &patientidPerson1,
		SampleID:  &sampleidPerson1,
	}

	person2 := models.Patient{
		FirstName: &namePerson2,
		PatientID: &patientidPerson2,
		SampleID:  &sampleidPerson2,
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

	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/patient", body)
	req2, err2 := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/patient", body2)

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

// func TestRemoveAllBiosamples(t *testing.T) {
// 	result := repos.RemoveAllBiosamples()
// 	if !result {
// 		t.Fail()
// 		return
// 	}
// }

// func TestRemoveAllIndividuals(t *testing.T) {
// 	result := repos.RemoveAllIndividuals()
// 	if !result {
// 		t.Fail()
// 		return
// 	}
// }
