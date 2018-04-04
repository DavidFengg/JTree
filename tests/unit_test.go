package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Bio-core/Jtree/dummydata"
	"github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
	"github.com/Bio-core/jtree/repos"
	"github.com/Bio-core/jtree/restapi"
	"github.com/Bio-core/jtree/restapi/operations"
	"github.com/go-openapi/loads"
)

var host = "http://127.0.0.1:8000"

func TestMain(m *testing.M) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewJtreeMetadataAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()

	go server.Serve()

	testResults := m.Run()
	os.Exit(testResults)
}

func TestUrls(t *testing.T) {
	result := true
	result = result && CheckPageResponse(host+"/Jtree/metadata/0.1.0/columns")
	result = result && CheckPageResponse(host+"/Jtree/metadata/0.1.0/uneditable")
	result = result && CheckPageResponse(host+"/Jtree/metadata/0.1.0/searchable")
	result = result && !CheckPageResponse(host+"/x")
	result = result && !CheckNoPageResponse(host+"/Jtree/metadata/0.1.0/searchable")
	result = result && CheckNoPageResponse(host+"/x")

	if result != true {
		t.Error("Web Pages Not Successful")
	}
}

func TestGenerateDummyData(t *testing.T) {
	dummydata.MakeData(100)

	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "patients")
	query.SelectedCondition = make([][]string, 0)
	querystring := database.BuildQuery(query)
	if len(repos.GetAllSamples(querystring)) != 100 {
		t.Fail()
	}
	query.SelectedTables[0] = "samples"
	querystring = database.BuildQuery(query)
	if len(repos.GetAllSamples(querystring)) != 287 {
		t.Fail()
	}
	query.SelectedTables[0] = "experiments"
	querystring = database.BuildQuery(query)
	if len(repos.GetAllSamples(querystring)) != 866 {
		t.Fail()
	}
	query.SelectedTables[0] = "results"
	querystring = database.BuildQuery(query)
	if len(repos.GetAllSamples(querystring)) != 1282 {
		t.Fail()
	}
	query.SelectedTables[0] = "resultdetails"
	querystring = database.BuildQuery(query)
	if len(repos.GetAllSamples(querystring)) != 1899 {
		t.Fail()
	}
	return
}

func TestAddPatientsPOST(t *testing.T) {

	patient := dummydata.MakePatient(-1)
	person1Bytes, err := json.Marshal(patient)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(person1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/patient", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if resp.Status != "201 Created" {
		t.Fail()
		return
	}

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

// func TestAddSamplesPOST(t *testing.T) {
// 	sampleidSample1 := "Sample1"
// 	facilitySample1 := "TGH"
// 	var volumeOfBloodMarrowSample1 float32
// 	volumeOfBloodMarrowSample1 = 14.2
// 	dateCollectedSample1 := "20140506"
// 	sampleidSample2 := "Sample2"
// 	facilitySample2 := "PMH"
// 	var volumeOfBloodMarrowSample2 float32
// 	volumeOfBloodMarrowSample2 = 105.67
// 	dateCollectedSample2 := "2020-09-08"

// 	sample1 := models.Sample{
// 		SampleID:            &sampleidSample1,
// 		Facility:            &facilitySample1,
// 		VolumeOfBloodMarrow: &volumeOfBloodMarrowSample1,
// 		DateCollected:       &dateCollectedSample1,
// 	}

// 	sample2 := models.Sample{
// 		SampleID:            &sampleidSample2,
// 		Facility:            &facilitySample2,
// 		VolumeOfBloodMarrow: &volumeOfBloodMarrowSample2,
// 		DateCollected:       &dateCollectedSample2,
// 	}

// 	sample1Bytes, err := json.Marshal(sample1)
// 	sample2Bytes, err2 := json.Marshal(sample2)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}

// 	if err2 != nil {
// 		t.Fail()
// 		return
// 	}

// 	body := bytes.NewReader(sample1Bytes)
// 	body2 := bytes.NewReader(sample2Bytes)

// 	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/sample", body)
// 	req2, err2 := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/sample", body2)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if err2 != nil {
// 		t.Fail()
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req2.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)
// 	resp2, err2 := http.DefaultClient.Do(req2)

// 	if resp.Status != "201 Created" || resp2.Status != "201 Created" {
// 		t.Fail()
// 		return
// 	}

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if err2 != nil {
// 		t.Fail()
// 		return
// 	}
// 	defer resp.Body.Close()
// 	defer resp2.Body.Close()

// }

// func TestAddExperimentsPOST(t *testing.T) {
// 	sampleidSample1 := "Sample1"
// 	pcrSample1 := "PCR1"
// 	dateCollectedSample1 := "20140506"
// 	sampleidSample2 := "Sample2"
// 	pcrSample2 := "PCR2"
// 	dateCollectedSample2 := "2020-09-08"

// 	sample1 := models.Experiment{
// 		SampleID:     &sampleidSample1,
// 		Pcr:          &pcrSample1,
// 		CompleteDate: &dateCollectedSample1,
// 	}

// 	sample2 := models.Experiment{
// 		SampleID:     &sampleidSample2,
// 		Pcr:          &pcrSample2,
// 		CompleteDate: &dateCollectedSample2,
// 	}

// 	sample1Bytes, err := json.Marshal(sample1)
// 	sample2Bytes, err2 := json.Marshal(sample2)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}

// 	if err2 != nil {
// 		t.Fail()
// 		return
// 	}

// 	body := bytes.NewReader(sample1Bytes)
// 	body2 := bytes.NewReader(sample2Bytes)

// 	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/experiment", body)
// 	req2, err2 := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/experiment", body2)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if err2 != nil {
// 		t.Fail()
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req2.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)
// 	resp2, err2 := http.DefaultClient.Do(req2)

// 	if resp.Status != "201 Created" || resp2.Status != "201 Created" {
// 		t.Fail()
// 		return
// 	}

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if err2 != nil {
// 		t.Fail()
// 		return
// 	}
// 	defer resp.Body.Close()
// 	defer resp2.Body.Close()

// }
// func TestSamplesQuery(t *testing.T) {
// 	query := models.Query{
// 		SelectedFields:    []string{"samples.sample_id", "samples.facility", "samples.volume_of_blood_marrow", "samples.date_collected"},
// 		SelectedTables:    []string{"samples"},
// 		SelectedCondition: [][]string{},
// 	}
// 	queryBytes, err := json.Marshal(query)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	body := bytes.NewReader(queryBytes)
// 	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	content, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	var results []models.Record
// 	err = json.Unmarshal(content, &results)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if len(results) != 2 {
// 		t.Fail()
// 		return
// 	}
// 	if *results[0].Sample.SampleID != "Sample1" || *results[1].Sample.SampleID != "Sample2" {
// 		t.Fail()
// 		return
// 	}

// 	defer resp.Body.Close()

// }

// func TestPatientsQuery(t *testing.T) {
// 	query := models.Query{
// 		SelectedFields:    []string{"patients.sample_id", "patients.patient_id", "patients.first_name"},
// 		SelectedTables:    []string{"patients"},
// 		SelectedCondition: [][]string{},
// 	}
// 	queryBytes, err := json.Marshal(query)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	body := bytes.NewReader(queryBytes)
// 	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	content, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	var results []models.Record
// 	err = json.Unmarshal(content, &results)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if len(results) != 2 {
// 		t.Fail()
// 		return
// 	}
// 	if *results[0].Patient.SampleID != "Sample1" || *results[1].Patient.SampleID != "Sample2" {
// 		t.Fail()
// 		return
// 	}

// 	defer resp.Body.Close()

// }

// func TestExperimentsQuery(t *testing.T) {
// 	query := models.Query{
// 		SelectedFields:    []string{"experiments.sample_id", "experiments.complete_date", "experiments.pcr"},
// 		SelectedTables:    []string{"experiments"},
// 		SelectedCondition: [][]string{},
// 	}
// 	queryBytes, err := json.Marshal(query)

// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	body := bytes.NewReader(queryBytes)
// 	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	content, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	var results []models.Record
// 	err = json.Unmarshal(content, &results)
// 	if err != nil {
// 		t.Fail()
// 		return
// 	}
// 	if len(results) != 2 {
// 		t.Fail()
// 		return
// 	}
// 	if *results[0].Experiment.SampleID != "Sample1" || *results[1].Experiment.SampleID != "Sample2" {
// 		t.Fail()
// 		return
// 	}

// 	defer resp.Body.Close()

// }

func TestQueries(t *testing.T) {

	patient := dummydata.MakePatient(-1)
	mrn := "Test123"
	id := "TestID"
	dob, _ := time.Parse("2006-01-02", "2034-01-02")
	patient.Mrn = &mrn
	patient.PatientID = &id
	patient.Dob = &dob

	repos.InsertPatient(&patient)
	queriesList := make([]models.Query, 0)
	expected := 102
	//Return all
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{}))
	//Test Equal to
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Equal to", "Test123"}}))
	//Test not equal to
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Not equal to", "Test123"}}))
	//Test Begins with
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Begins with", "Te"}}))
	//Test Not begins with
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Not begins with", "Te"}}))
	//Test Ends with
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Ends with", "123"}}))
	//Test Not ends with
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Not ends with", "123"}}))
	//Test Contains
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Contains", "est1"}}))
	//Test Not contains
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.mrn", "Not contains", "est1"}}))
	//Test Greater than
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.dob", "Greater than", "2034-01-01"}}))
	//Test Less than
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.dob", "Less than", "2034-01-01"}}))
	//Test Greater or equal to
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.dob", "Greater or equal to", "2034-01-02"}}))
	//Test Less than
	queriesList = append(queriesList, returnQuery([]string{"*"}, []string{"patients"}, [][]string{{"AND", "patients.dob", "Less or equal to", "2034-01-02"}}))

	for i, q := range queriesList {
		queryBytes, err := json.Marshal(q)

		if err != nil {
			t.Fail()
			return
		}
		body := bytes.NewReader(queryBytes)
		req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/query", body)
		if err != nil {
			t.Fail()
			return
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fail()
			return
		}
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fail()
			return
		}
		var results []models.Record
		err = json.Unmarshal(content, &results)
		if err != nil {
			t.Fail()
			return
		}
		defer resp.Body.Close()
		if len(results) != expected {
			t.Error("Query #", i+1, " failed - Expected:", expected, " Got:", len(results))
		}
		if i%2 == 0 {
			expected = 1
		} else {
			expected = 101
		}
		if i+2 == len(queriesList) {
			expected = 102
		}
	}
	return
}
