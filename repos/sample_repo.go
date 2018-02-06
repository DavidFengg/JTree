package repos

import (
	"fmt"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

//GetAllSamples gets all and results a list of samples
func GetAllSamples(query string) []*models.Sample {
	samples := []*models.Sample{}
	err := database.DB.Select(&samples, query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return samples
}

//GetSampleColumns gets the columns in a table
func GetSampleColumns() []string {
	rows, err := database.DB.Query("Select * from Samples where sample_id = \"err\"")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return columns
}

//InsertSample allows users to add generic objects to a collection in the database
func InsertSample(sample *models.Sample) bool {
	stmt, err := database.DB.Prepare("INSERT INTO `JTree`.`Samples` (`sample_id`, `facility`, `test_requested`, `se_num`, `date_collected`, `date_received`, `sample_type`, `material_received`, `material_received_num`, `material_received_other`, `volume_of_blood_marrow`, `surgical_num`, `tumor_site`, `historical_diagnosis`, `tumor_percnt_of_total`, `tumor_percnt_of_circled`, `reviewed_by`, `h_e_slide_location`, `non_uhn_id`, `name_of_requestor`, `dna_concentration`, `dna_volume`, `dna_location`, `rna_concentration`, `rna_volume`, `rna_location`, `wbc_location`, `plasma_location`, `cf_plasma_location`, `pb_bm_location`, `rna_lysate_location`, `sample_size`, `study_id`, `sample_name`, `date_submitted`, `container_type`, `container_name`, `container_id`, `container_well`, `copath_num`, `other_identifier`, `has_sample_files`, `dna_sample_barcode`, `dna_extraction_date`, `dna_quality`, `ffpe_qc_date`, `delta_ct_Value`, `comments`, `rnase_p_date`, `dna_quality_by_rnase_p`, `rna_quality`, `rna_extraction_date`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?); ")

	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(
		sample.SampleID,
		sample.Facility,
		sample.TestRequested,
		sample.SeNum,
		sample.DateCollected,
		sample.DateReceived,
		sample.SampleType,
		sample.MaterialReceived,
		sample.MaterialReceivedNum,
		sample.MaterialReceivedOther,
		sample.VolumeOfBloodMarrow,
		sample.SurgicalNum,
		sample.TumorSite,
		sample.HistoricalDiagnosis,
		sample.TumorPercntOfTotal,
		sample.TumorPercntOfCircled,
		sample.ReviewedBy,
		sample.HESlideLocation,
		sample.NonUhnID,
		sample.NameOfRequestor,
		sample.DnaConcentration,
		sample.DnaVolume,
		sample.DnaLocation,
		sample.RnaConcentration,
		sample.RnaVolume,
		sample.RnaLocation,
		sample.WbcLocation,
		sample.PlasmaLocation,
		sample.CfPlasmaLocation,
		sample.PbBmLocation,
		sample.RnaLysateLocation,
		sample.SampleSize,
		sample.StudyID,
		sample.SampleName,
		sample.DateSubmitted,
		sample.ContainerType,
		sample.ContainerName,
		sample.ContainerID,
		sample.ContainerWell,
		sample.CopathNum,
		sample.OtherIdentifier,
		sample.HasSampleFiles,
		sample.DnaSampleBarcode,
		sample.DnaExtractionDate,
		sample.DnaQuality,
		sample.FfpeQcDate,
		sample.DeltaCtValue,
		sample.Comments,
		sample.RnasePDate,
		sample.DnaQualityByRnaseP,
		sample.RnaQuality,
		sample.RnaExtractionDate)
	return true
}

//RemoveAllSamples will empty a collection
func RemoveAllSamples() bool {
	//Implement here
	return true
}
