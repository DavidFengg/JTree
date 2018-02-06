package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//GetAllPatients gets all and results a list of individuals
func GetAllPatients(query string) []*models.Patient {
	patients := []*models.Patient{}
	err := database.DB.Select(&patients, query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return patients
}

//InsertPatient allows users to add generic objects to a collection in the database
func InsertPatient(person *models.Patient) bool {
	stmt, err := database.DB.Prepare("INSERT INTO `Patients`(`first_name`,`last_name`,`initials`,`gender`,`mrn`,`dob`,`on_hcn`,`clinical_history`,`patient_type`,`se_num`,`patient_id`,`sample_id`,`data_received`,`referring_physican`,`date_reported`,`surgical_date`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(
		person.FirstName,
		person.LastName,
		person.Initials,
		person.Gender,
		person.Mrn,
		person.Dob,
		person.OnHcn,
		person.ClinicalHistory,
		person.PatientType,
		person.SeNum,
		person.PatientID,
		person.SampleID,
		person.DateReceived,
		person.ReferringPhysican,
		person.DateReported,
		person.SurgicalDate)
	return true
}

//RemoveAllPatients will empty a collection
func RemoveAllPatients() bool {
	//implement here
	return true
}
