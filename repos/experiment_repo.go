package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//InsertExperiment allows users to add generic objects to a collection in the database
func InsertExperiment(experiment *models.Experiment) bool {
	stmt, err := database.DB.Prepare("INSERT INTO `experiments` (`experiment_id`,`study_id`,`panel_assay_screened`,`test_date`,`chip_cartridge_barcode`,`complete_date`,`pcr`,`sample_id`,`project_name`,`priority`,`opened_date`,`project_id`,`has_project_files`,`procedure_order_datetime`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?);")

	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		experiment.ExperimentID,
		experiment.StudyID,
		experiment.PanelAssayScreened,
		experiment.TestDate,
		experiment.ChipCartridgeBarcode,
		experiment.CompleteDate,
		experiment.Pcr,
		experiment.SampleID,
		experiment.ProjectName,
		experiment.Priority,
		experiment.OpenedDate,
		experiment.ProjectID,
		experiment.HasProjectFiles,
		experiment.ProcedureOrderDatetime)
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// //RemoveUnitTestSamples will empty a collection
// func RemoveUnitTestSamples() bool {
// 	_, err := database.DB.Query("Delete from samples where sample_id LIKE \"%Sample%\"")
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }
