package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//InsertExperiment allows users to add generic objects to a collection in the database
func InsertExperiment(experiment *models.Experiment) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `experiments` (`experiment_id`,`study_id`,`panel_assay_screened`,`test_date`,`chip_cartridge_barcode`,`complete_date`,`pcr`,`sample_id`,`project_name`,`priority`,`opened_date`,`project_id`,`has_project_files`,`procedure_order_datetime`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?);")

	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		experiment.ExperimentID,
		experiment.StudyID,
		experiment.PanelAssayScreened,
		experiment.TestDate.Format(shortForm),
		experiment.ChipCartridgeBarcode,
		experiment.CompleteDate.Format(shortForm),
		experiment.Pcr,
		experiment.SampleID,
		experiment.ProjectName,
		experiment.Priority,
		experiment.OpenedDate.Format(shortForm),
		experiment.ProjectID,
		experiment.HasProjectFiles,
		experiment.ProcedureOrderDatetime.Format(shortForm))
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//GetExperimentByID gets all and results a list of samples
func GetExperimentByID(ID string) *models.Experiment {
	experiments := []*models.Experiment{}
	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "experiments")
	query.SelectedCondition = make([][]string, 0)
	//query.SelectedCondition = append(query.SelectedCondition, make([]string, 0))
	conditions := []string{"AND", "experiments.experiment_id", "Equal to", ID}
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryString := database.BuildQuery(query)
	err := database.DBSelect.Select(&experiments, queryString)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(experiments) == 0 {
		return nil
	}
	return experiments[0]
}

//UpdateExperiment allows users to add generic objects to a collection in the database
func UpdateExperiment(experimentID string, experiment *models.Experiment) bool {
	stmt, err := database.DBUpdate.Prepare("UPDATE `experiments` SET `experiment_id` = ?,`study_id` = ?,`panel_assay_screened` = ?,`test_date` = ?,`chip_cartridge_barcode` = ?,`complete_date` = ?,`pcr` = ?,`sample_id` = ?,`project_name` = ?,`priority` = ?,`opened_date` = ?,`project_id` = ?,`has_project_files` = ?,`procedure_order_datetime` = ? WHERE `experiment_id` = ?;")

	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(
		experimentID,
		experiment.StudyID,
		experiment.PanelAssayScreened,
		experiment.TestDate.Format(shortForm),
		experiment.ChipCartridgeBarcode,
		experiment.CompleteDate.Format(shortForm),
		experiment.Pcr,
		experiment.SampleID,
		experiment.ProjectName,
		experiment.Priority,
		experiment.OpenedDate.Format(shortForm),
		experiment.ProjectID,
		experiment.HasProjectFiles,
		experiment.ProcedureOrderDatetime.Format(shortForm),
		experimentID)
	stmt.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//DeleteExperiment removes an experiment by id
func DeleteExperiment(experimentID string) bool {
	stmt, err := database.DBUpdate.Prepare("DELETE FROM experiments WHERE experiment_id = ?")
	if err != nil {
		log.Fatal(err)
		return false
	}

	result, err := stmt.Exec(experimentID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, result)
		return false
	}
	return true
}

//HasResult returns a boolean depending on whether the experiment has results
func HasResults(ID string) bool {
	results := []*models.Result{}
	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "results")
	query.SelectedCondition = make([][]string, 0)

	conditions := []string{"AND", "results.experiment_id", "Equal to", ID}
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryString := database.BuildQuery(query)
	err := database.DBSelect.Select(&results, queryString)
	if err != nil {
		fmt.Println(err)
	}

	// returns true if length of results is greater than 0
	return len(results) > 0;
}