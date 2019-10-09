package tests

import (
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
	"flag"

	"github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/dummydata"
	"github.com/Bio-core/jtree/models"
	"github.com/Bio-core/jtree/repos"
	"github.com/Bio-core/jtree/restapi"
	"github.com/Bio-core/jtree/restapi/operations"
	"github.com/go-openapi/loads"
)

var host = "http://localhost:8000"

var patientID = flag.String("patientID", "2d8812d6-5c12-40bf-462c-47693b8d869", "Unique patient id")

func TestMain(m *testing.M) {
	testResults := m.Run()
	os.Exit(testResults)
}

func TestSetupSever(t *testing.T) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		t.Errorf("%v", err)
		log.Fatalln(err)
	}

	api := operations.NewJtreeMetadataAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.ConfigureAPI()

	go server.Serve()
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

func TestAddPatient(t *testing.T) {
	
	patient := dummydata.MakePatient(1)
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

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}
	fmt.Println(string(content))

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestUpdatePatient(t *testing.T) {
	// fmt.Println(*patientID)

	patient := repos.GetPatientByID(*patientID)
	if (patient == nil) {
		t.Fail()
		return
	}
	// first := "Mitchell"
	// last := "Strong"
	// patient.FirstName = &first
	// patient.LastName = &last

	// person1Bytes, err := json.Marshal(patient)

	// if err != nil {
	// 	t.Fail()
	// 	return
	// }

	// fmt.Println(person1Bytes)

	// body := bytes.NewReader(person1Bytes)

	// req, err := http.NewRequest("PUT", host+"/Jtree/metadata/0.1.0/patient/" + patientID, body)

	// if err != nil {
	// 	t.Fail()
	// 	return
	// }

	// req.Header.Set("Content-Type", "application/json")

	// resp, err := http.DefaultClient.Do(req)
	// if (err != nil) {
	// 	t.Fail()
	// 	return
	// }

	// if resp.Status != "201 Created" {
	// 	t.Fail()
	// 	return
	// }

	// defer resp.Body.Close()

	// patientNew := repos.GetPatientByID(patientID)

	// if *patientNew.FirstName != first || *patientNew.LastName != last {
	// 	t.Fail()
	// 	return
	// }

	// return
}

func TestDeletePatient(t *testing.T) {

	req, err := http.NewRequest("DELETE", host + "/Jtree/metadata/0.1.0/patient/" + *patientID, nil)
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

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Fail()
		return
	}
}

func TestAddSample(t *testing.T) {

	sample := dummydata.MakeSample(1, 1)
	sample1Bytes, err := json.Marshal(sample)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(sample1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/sample", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}
	fmt.Println(string(content))

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestUpdateSample(t *testing.T) {

	sample := repos.GetSampleByID("1")
	comments := "updated"
	sample.Comments = &comments
	sample1Bytes, err := json.Marshal(sample)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(sample1Bytes)

	req, err := http.NewRequest("PUT", host+"/Jtree/metadata/0.1.0/sample/1", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

	sampleNew := repos.GetSampleByID("1")

	if *sampleNew.Comments != comments {
		t.Fail()
		return
	}

	return
}

func TestDeleteSample(t *testing.T) {

	req, err := http.NewRequest("DELETE", host + "/Jtree/metadata/0.1.0/sample/1", nil)
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

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Fail()
		return
	}
}

func TestAddExperiment(t *testing.T) {

	experiment := dummydata.MakeExperiment(1, 1)
	experiment1Bytes, err := json.Marshal(experiment)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(experiment1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/experiment", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}
	fmt.Println(string(content))

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestUpdateExperiment(t *testing.T) {

	experiment := repos.GetExperimentByID("1")
	projectName := "updated"
	experiment.ProjectName = &projectName
	experiment1Bytes, err := json.Marshal(experiment)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(experiment1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/experiment/1", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

	experimentNew := repos.GetExperimentByID("1")

	if *experimentNew.ProjectName != projectName {
		t.Fail()
		return
	}

	return
}

func TestDeleteExperiment(t *testing.T) {

	req, err := http.NewRequest("DELETE", host + "/Jtree/metadata/0.1.0/experiment/1", nil)
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

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Fail()
		return
	}
}

func TestAddResult(t *testing.T) {

	result := dummydata.MakeResult(1, 1)
	result1Bytes, err := json.Marshal(result)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(result1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/result", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}
	fmt.Println(string(content))

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestUpdateResult(t *testing.T) {

	result := repos.GetResultByID("1")
	uid := "updated"
	result.UID = &uid
	result1Bytes, err := json.Marshal(result)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(result1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/result/1", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}
	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

	resultNew := repos.GetResultByID("1")

	if *resultNew.UID != uid {
		t.Fail()
		return
	}

	return
}

func TestDeleteResult(t *testing.T) {

	req, err := http.NewRequest("DELETE", host + "/Jtree/metadata/0.1.0/result/1", nil)
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

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Fail()
		return
	}
}

func TestAddResDetail(t *testing.T) {

	result := dummydata.MakeResultDetail(1, 1)
	result1Bytes, err := json.Marshal(result)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(result1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/resultdetails", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}
	fmt.Println(string(content))

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}

func TestUpdateResDetail(t *testing.T) {

	result := repos.GetResultDetailByID("1")
	uid := "updated"
	result.UID = &uid
	result1Bytes, err := json.Marshal(result)

	if err != nil {
		t.Fail()
		return
	}

	body := bytes.NewReader(result1Bytes)

	req, err := http.NewRequest("POST", host+"/Jtree/metadata/0.1.0/resultdetails/1", body)

	if err != nil {
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
		return
	}
	if resp.Status != "200 OK" && string(content) != "error" {
		t.Fail()
		return
	}
	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

	resultNew := repos.GetResultByID("1")

	if *resultNew.UID != uid {
		t.Fail()
		return
	}

	return
}

func TestDeleteResDetail(t *testing.T) {

	req, err := http.NewRequest("DELETE", host + "/Jtree/metadata/0.1.0/resultdetails/1", nil)
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

	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		t.Fail()
		return
	}
}

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
