package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//InsertResult allows users to add generic objects to a collection in the database
func InsertResult(result *models.Result) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `results`(`failed_regions`,`mean_depth_of_coveage`,`mlpa_pcr`,`mutation`,`overall_hotspots_threshold`,`overall_quality_threshold`,`results_id`,`uid`,`verification_pcr`, `experiment_id`)VALUES(?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	outcome, err := stmt.Exec(
		result.FailedRegions,
		result.MeanDepthOfCoveage,
		result.MlpaPcr,
		result.Mutation,
		result.OverallHotspotsThreshold,
		result.OverallQualityThreshold,
		result.ResultsID,
		result.UID,
		result.VerificationPcr,
		result.ExperimentID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, outcome)
	}
	return true
}

//GetResultByID gets all and results a list of samples
func GetResultByID(ID string) *models.Result {
	results := []*models.Result{}
	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "results")
	query.SelectedCondition = make([][]string, 0)
	//query.SelectedCondition = append(query.SelectedCondition, make([]string, 0))
	conditions := []string{"AND", "results.results_id", "Equal to", ID}
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryString := database.BuildQuery(query)
	err := database.DBSelect.Select(&results, queryString)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

//UpdateResult allows users to add generic objects to a collection in the database
func UpdateResult(resultID string, result *models.Result) bool {
	stmt, err := database.DBUpdate.Prepare("UPDATE `results` SET `failed_regions` = ?,`mean_depth_of_coveage` = ?,`mlpa_pcr` = ?,`mutation` = ?,`overall_hotspots_threshold` = ?,`overall_quality_threshold` = ?,`results_id` = ?,`uid` = ?,`verification_pcr` = ?,`experiment_id` = ? WHERE `results_id` = ?;")
	if err != nil {
		log.Fatal(err)
	}
	outcome, err := stmt.Exec(
		result.FailedRegions,
		result.MeanDepthOfCoveage,
		result.MlpaPcr,
		result.Mutation,
		result.OverallHotspotsThreshold,
		result.OverallQualityThreshold,
		resultID,
		result.UID,
		result.VerificationPcr,
		result.ExperimentID,
		resultID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, outcome)
	}
	return true
}

//DeleteResult removes a result by id
func DeleteResult(resultID string) bool {
	stmt, err := database.DBUpdate.Prepare("DELETE FROM results WHERE results_id = ?")
	if err != nil {
		log.Fatal(err)
		return false
	}

	result, err := stmt.Exec(resultID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, result)
		return false
	}
	return true
}

//HasResultDetails returns a boolean depending on whether the result has resultdetails
func HasResultDetails(ID string) bool {
	resultdetails := []*models.Resultdetails{}
	query := models.Query{}
	query.SelectedFields = make([]string, 0)
	query.SelectedFields = append(query.SelectedFields, "*")
	query.SelectedTables = make([]string, 0)
	query.SelectedTables = append(query.SelectedTables, "resultdetails")
	query.SelectedCondition = make([][]string, 0)

	conditions := []string{"AND", "resultdetails.results_id", "Equal to", ID}
	query.SelectedCondition = append(query.SelectedCondition, conditions)

	queryString := database.BuildQuery(query)
	err := database.DBSelect.Select(&resultdetails, queryString)
	if err != nil {
		fmt.Println(err)
	}

	// returns true if length of resultdetails is greater than 0
	return len(resultdetails) > 0;
}