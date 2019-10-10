package tests

import (
	"net/http"
	"log"
	"bytes"
	"io/ioutil"
	"encoding/json"

	"github.com/Bio-core/jtree/models"
)

var host = "http://localhost:8000"

//CheckPageResponse checks if a page that should respond is found correctly
func CheckPageResponse(url string) bool {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}
	response, err := client.Do(req)
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

func returnQuery(fields []string, tables []string, conditions [][]string) models.Query {
	query := models.Query{
		SelectedFields:    fields,
		SelectedTables:    tables,
		SelectedCondition: conditions,
	}
	return query
}

// GetQueryReponse returns the json response of a specific query
func GetQueryResponse(fields string, tables string, conditions []string ) []byte {
	var query models.Query

	query.SelectedFields = make([]string, 0)	
	query.SelectedFields = append(query.SelectedFields, fields)
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, tables)
	query.SelectedCondition = make([][]string, 0)
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryBytes, err := json.Marshal(query)
	if err != nil {
		log.Println(err)
	}

	body := bytes.NewReader(queryBytes)

	req, err := http.NewRequest("POST", host + "/Jtree/metadata/0.1.0/query", body)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return content
}
