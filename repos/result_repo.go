package repos

import (
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//InsertResult allows users to add generic objects to a collection in the database
func InsertResult(result *models.Result) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `JTree`.`results`(`failed_regions`,`mean_depth_of_coveage`,`mlpa_pcr`,`mutation`,`overall_hotspots_threshold`,`overall_quality_threshold`,`results_id`,`sample_id`,`uid`,`verification_pcr`)VALUES(?,?,?,?,?,?,?,?,?,?);")
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
		result.SampleID,
		result.UID,
		result.VerificationPcr)
	stmt.Close()
	if err != nil {
		log.Fatal(err, outcome)
	}
	return true
}
