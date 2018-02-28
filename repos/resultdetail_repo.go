package repos

import (
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//InsertResultDetail allows users to add generic objects to a collection in the database
func InsertResultDetail(result *models.Resultdetails) bool {
	stmt, err := database.DBUpdate.Prepare("INSERT INTO `JTree`.`resultdetails`(`VAF`,`c_nomenclature`,`coverage`,`exon`,`gene`,`p_nomenclature`,`pcr`,`quality_score`,`result`,`results_details_id`,`results_id`,`risk_score`,`sample_id`,`uid`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	outcome, err := stmt.Exec(
		result.VAF,
		result.CNomenclature,
		result.Coverage,
		result.Exon,
		result.Gene,
		result.PNomenclature,
		result.Pcr,
		result.QualityScore,
		result.Result,
		result.ResultsDetailsID,
		result.ResultsID,
		result.RiskScore,
		result.SampleID,
		result.UID)
	stmt.Close()
	if err != nil {
		log.Fatal(err, outcome)
	}
	return true
}
