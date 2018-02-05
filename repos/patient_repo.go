package repos

import (
	"database/sql"
	"encoding/json"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

const individualCollection = "individual"

var starPatient = "first_name, last_name, initials, gender, mrn, dob, on_hcn, clinical_history, patient_type, se_num, patient_id, sample_id, data_received, referring_physican, date_reported, surgical_date"

//GetAllPatients gets all and results a list of individuals
func GetAllPatients() []*models.Patient {
	var list []*models.Patient
	rows, err := database.DB.Query("SELECT " + starPatient + " FROM Patients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		person := getPatientRow(rows)
		list = append(list, &person)
	}
	return list
}

//GetManyPatientsByString gets all and results a list of individuals
func GetManyPatientsByString(field, value string) []*models.Patient {
	var query = "SELECT " + starPatient + " FROM Patients Where " + field + "= '" + value + "'"
	var list []*models.Patient
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		patient := getPatientRow(rows)
		list = append(list, &patient)
	}
	return list
}

//GetOnePatientByString gets a bio sample and returns it based on the strings provided
func GetOnePatientByString(field, value string) *models.Patient {
	var patient models.Patient
	var query = "SELECT " + starPatient + " FROM Patients Where " + field + "= '" + value + "' LIMIT 1"
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		patient = getPatientRow(rows)
	}

	return &patient
}

//InsertPatient allows users to add generic objects to a collection in the database
func InsertPatient(patient *models.Patient) bool {
	j, _ := json.Marshal(&patient)
	var interpatient interface{}
	json.Unmarshal(j, &interpatient)
	//implement here
	return true
}

//RemoveAllPatients will empty a collection
func RemoveAllPatients() bool {
	//implement here
	return true
}

func getPatientRow(rows *sql.Rows) models.Patient {
	var person models.Patient
	err := rows.Scan(
		&person.FirstName,
		&person.LastName,
		&person.Initials,
		&person.Gender,
		&person.Mrn,
		&person.Dob,
		&person.OnHcn,
		&person.ClinicalHistory,
		&person.PatientType,
		&person.SeNum,
		&person.PatientID,
		&person.SampleID,
		&person.DateReceived,
		&person.ReferringPhysican,
		&person.DateReported,
		&person.SurgicalDate)
	if err != nil {
		log.Fatal(err)
	}
	return person
}
