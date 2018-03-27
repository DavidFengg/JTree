package repos

import (
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//InsertPatient allows users to add generic objects to a collection in the database
func InsertPatient(person *models.Patient) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `patients`(`first_name`,`last_name`,`initials`,`gender`,`mrn`,`dob`,`on_hcn`,`clinical_history`,`patient_type`,`se_num`,`patient_id`,`date_received`,`referring_physican`,`date_reported`,`surgical_date`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	result, err := stmt.Exec(
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
		person.DateReceived,
		person.ReferringPhysican,
		person.DateReported,
		person.SurgicalDate)
	stmt.Close()
	if err != nil {
		log.Fatal(err, result)
	}
	return true
}
