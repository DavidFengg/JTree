package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

const server = "http://localhost:8000"

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
	if !tearDown() {
		testResults = -1
	}
	os.Exit(testResults)
}

func TestUrls(t *testing.T) {
	result := true
	result = result && CheckPageResponse(server+"/Jtree/metadata/0.1.0/columns")
	result = result && CheckNoPageResponse(server+"/x")

	if result != true {
		t.Error("Web Pages Not Successful")
	}
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

func TestAddExperimentsPOST(t *testing.T) {
	sampleidSample1 := "Sample1"
	pcrSample1 := "PCR1"
	dateCollectedSample1 := "20140506"
	sampleidSample2 := "Sample2"
	pcrSample2 := "PCR2"
	dateCollectedSample2 := "2020-09-08"

	sample1 := models.Experiment{
		SampleID:     &sampleidSample1,
		Pcr:          &pcrSample1,
		CompleteDate: &dateCollectedSample1,
	}

	sample2 := models.Experiment{
		SampleID:     &sampleidSample2,
		Pcr:          &pcrSample2,
		CompleteDate: &dateCollectedSample2,
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

	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/experiment", body)
	req2, err2 := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/experiment", body2)

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
func TestSamplesQuery(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"samples.sample_id", "samples.facility", "samples.volume_of_blood_marrow", "samples.date_collected"},
		SelectedTables:    []string{"samples"},
		SelectedCondition: [][]string{},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Sample.SampleID != "Sample1" || *results[1].Sample.SampleID != "Sample2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestPatientsQuery(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"patients.sample_id", "patients.patient_id", "patients.first_name"},
		SelectedTables:    []string{"patients"},
		SelectedCondition: [][]string{},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Patient.SampleID != "Sample1" || *results[1].Patient.SampleID != "Sample2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestExperimentsQuery(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"experiments.sample_id", "experiments.complete_date", "experiments.pcr"},
		SelectedTables:    []string{"experiments"},
		SelectedCondition: [][]string{},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Experiment.SampleID != "Sample1" || *results[1].Experiment.SampleID != "Sample2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestJoinQuery(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"patients.sample_id", "samples.sample_id", "experiments.sample_id"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Patient.SampleID != "Sample1" || *results[1].Patient.SampleID != "Sample2" {
		t.Fail()
		return
	}

	if *results[0].Sample.SampleID != "Sample1" || *results[1].Sample.SampleID != "Sample2" {
		t.Fail()
		return
	}

	if *results[0].Experiment.SampleID != "Sample1" || *results[1].Experiment.SampleID != "Sample2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestEqualTo(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.sample_id", "Equal to", "Sample1"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 1 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestNotEqualTo(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.sample_id", "Not equal to", "Sample1"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 1 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Strong" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "PMH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestGreaterThan(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.volume_of_blood_marrow", "Greater than", "14.2"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 1 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Strong" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "PMH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestLessThan(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.volume_of_blood_marrow", "Less than", "105.67"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 1 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}
func TestGreaterThanEqual(t *testing.T) {
	t.Skip()
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.volume_of_blood_marrow", "Greater or equal to", "14.2"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" || *results[1].Patient.FirstName != "Strong" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" || *results[1].Sample.Facility != "PMH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" || *results[1].Experiment.Pcr != "PCR2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestLessThanEqual(t *testing.T) {
	t.Skip()
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.volume_of_blood_marrow", "Less or equal to", "105.67"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" || *results[1].Patient.FirstName != "Strong" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" || *results[1].Sample.Facility != "PMH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" || *results[1].Experiment.Pcr != "PCR2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestBeginsWith(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.sample_id", "Begins with", "Sample"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" || *results[1].Patient.FirstName != "Strong" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" || *results[1].Sample.Facility != "PMH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" || *results[1].Experiment.Pcr != "PCR2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestNotBeginsWith(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "samples.sample_id", "Not begins with", "Sample"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 0 {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestEndsWith(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "patients.sample_id", "Ends with", "2"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 1 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Strong" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "PMH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestNotEndsWith(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "patients.sample_id", "Not ends with", "2"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 1 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestContains(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "experiments.pcr", "Contains", "pcr"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 2 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" || *results[1].Patient.FirstName != "Strong" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" || *results[1].Sample.Facility != "PMH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" || *results[1].Experiment.Pcr != "PCR2" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestNotContains(t *testing.T) {
	query := models.Query{
		SelectedFields:    []string{"*"},
		SelectedTables:    []string{"patients", "samples", "experiments"},
		SelectedCondition: [][]string{{"AND", "experiments.pcr", "Not contains", "2"}},
	}
	queryBytes, err := json.Marshal(query)

	if err != nil {
		t.Fail()
		return
	}
	body := bytes.NewReader(queryBytes)
	req, err := http.NewRequest("POST", server+"/Jtree/metadata/0.1.0/query", body)
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
	if len(results) != 1 {
		t.Fail()
		return
	}
	if *results[0].Patient.FirstName != "Mitchell" {
		t.Fail()
		return
	}

	if *results[0].Sample.Facility != "TGH" {
		t.Fail()
		return
	}

	if *results[0].Experiment.Pcr != "PCR1" {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}
