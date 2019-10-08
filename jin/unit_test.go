package jin

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"

	"os"
	"testing"

	"github.com/Bio-core/jtree/dummydata"

)

var host = "http://127.0.0.1:8000"

func TestMain(m *testing.M) {
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

// Passes locally but not in Travis
func TestAddPatientPOST(t *testing.T) {

	fmt.Println("Adding Patient Post")

	patient := dummydata.MakePatient(2)
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

	if err != nil {
		t.Fail()
		return
	}

	defer resp.Body.Close()

}
