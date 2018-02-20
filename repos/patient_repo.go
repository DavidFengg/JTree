package repos

import (
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//GetAllPatients gets all and results a list of individuals
// func GetAllPatients(query string) []*models.Patient {
// 	patients := []*models.Patient{}
// 	err := database.DB.Select(&patients, query)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	return patients
// }

//GetPatientColumns gets the columns in a table
// func GetPatientColumns() []string {
// 	rows, err := database.DB.Query("Select * from patients where patient_id = \"err\"")
// 	defer rows.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	columns, err := rows.Columns()
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	return columns
// }

//InsertPatient allows users to add generic objects to a collection in the database
func InsertPatient(person *models.Patient) bool {
	stmt, err := database.DB.Prepare("INSERT INTO `patients`(`first_name`,`last_name`,`initials`,`gender`,`mrn`,`dob`,`on_hcn`,`clinical_history`,`patient_type`,`se_num`,`patient_id`,`sample_id`,`date_received`,`referring_physican`,`date_reported`,`surgical_date`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
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
		person.SampleID,
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

// //RemoveAllPatients will empty a collection
// func RemoveAllPatients() bool {
// 	//implement here
// 	return true
// }

//RemoveUnitTestPatients will empty a collection
func RemoveUnitTestPatients() bool {
	_, err := database.DB.Query("Delete from patients where patient_id LIKE \"%patient%\"")
	if err != nil {
		return false
	}
	return true
}
