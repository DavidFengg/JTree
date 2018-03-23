package dummydata

import (
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"

	models "github.com/Bio-core/jtree/models"
	repos "github.com/Bio-core/jtree/repos"
	yaml "gopkg.in/yaml.v2"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r *rand.Rand
var id1 int
var id2 int
var id3 int
var id4 int
var id5 int
var genes *GeneArray

//GeneArray is an object with an array of gene types
type GeneArray struct {
	Genes []string
}

//MakeData makes dummy data and puts it into the db
func MakeData(number int) {
	r = rand.New(rand.NewSource(99))
	genes = &GeneArray{}
	genes = genes.GetGenes()
	num1 := createPatients(number)
	num2 := createSamples(num1)
	num3 := createExperiments(num2)
	num4 := createResults(num3)
	createResultDetails(num4)
}

func makeRandomString() string {
	num := rand.Intn(50)
	value := randSeq(num)
	return value
}

func makeRandomDate() string {
	date := strconv.Itoa(rand.Intn(118)+1900) + "-" + strconv.Itoa(rand.Intn(11)+1) + "-" + strconv.Itoa(rand.Intn(27)+1)
	return date
}

func makeRandomFloat() float32 {
	num := rand.Float32()
	num += float32(rand.Intn(600))
	return num
}

func makeRandomBool() bool {
	num := rand.Intn(1)
	if num == 1 {
		return true
	}
	return false
}

func makeRandomGene() string {
	num := genrand(0, 568, 0, 568, 5)
	return genes.Genes[num]
}

func genrand(bmin, bmax, rmin, rmax, n int) int {
	const randMax = 32767
	// Generalized random number generator;
	// sum of n random variables (usually 3).
	// Bell curve spans bmin<=x<bmax; then,
	// values outside rmin<=x<rmax are rejected.
	var sum, i, u int
	sum = 0
	for {
		for i = 0; i < n; i++ {
			sum += bmin + (rand.Intn(randMax) % (bmax - bmin))
		}
		if sum < 0 {
			sum -= n - 1
		}
		u = sum / n

		if rmin <= u && u < rmax {
			break
		}
	}
	return u
}

func createPatients(number int) int {
	for i := 0; i < number; i++ {
		id1++
		tempPatient := makePatient(id1)
		repos.InsertPatient(&tempPatient)
	}
	return id1
}
func createResults(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(2) + 1
		for j := 0; j < c; j++ {
			id4++
			tempResult := makeResult(id1, id4)
			repos.InsertResult(&tempResult)
		}
	}
	return id4
}
func createResultDetails(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(2) + 1
		for j := 0; j < c; j++ {
			id5++
			tempResultDetail := makeResultDetail(id1, id5)
			repos.InsertResultDetail(&tempResultDetail)
		}
	}
	return id5
}

func createSamples(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(5) + 1
		for j := 0; j < c; j++ {
			id2++
			tempSample := makeSample(id1, id2)
			repos.InsertSample(&tempSample)
		}
	}
	return id2
}
func createExperiments(number int) int {
	id1 = 0
	for i := 0; i < number; i++ {
		id1++
		c := rand.Intn(5) + 1
		for j := 0; j < c; j++ {
			id3++
			tempExperiment := makeExperiment(id1, id3)
			repos.InsertExperiment(&tempExperiment)
		}
	}
	return id3
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func makePatient(patientID int) models.Patient {
	patient := models.Patient{}
	ClinicalHistory := makeRandomString()
	patient.ClinicalHistory = &ClinicalHistory
	DateReceived := makeRandomDate()
	patient.DateReceived = &DateReceived
	DateReported := makeRandomDate()
	patient.DateReported = &DateReported
	Dob := makeRandomDate()
	patient.Dob = &Dob
	FirstName := makeRandomString()
	patient.FirstName = &FirstName
	Gender := makeRandomString()
	patient.Gender = &Gender
	Initials := makeRandomString()
	patient.Initials = &Initials
	LastName := makeRandomString()
	patient.LastName = &LastName
	Mrn := makeRandomString()
	patient.Mrn = &Mrn
	OnHcn := makeRandomString()
	patient.OnHcn = &OnHcn
	PatientID := strconv.Itoa(patientID)
	patient.PatientID = &PatientID
	PatientType := makeRandomString()
	patient.PatientType = &PatientType
	ReferringPhysican := makeRandomString()
	patient.ReferringPhysican = &ReferringPhysican
	SeNum := makeRandomString()
	patient.SeNum = &SeNum
	SurgicalDate := makeRandomDate()
	patient.SurgicalDate = &SurgicalDate

	return patient
}

func makeSample(patientID int, sampleID int) models.Sample {
	sample := models.Sample{}
	SampleID := strconv.Itoa(sampleID)
	sample.SampleID = &SampleID
	Facility := makeRandomString()
	sample.Facility = &Facility
	TestRequested := makeRandomString()
	sample.TestRequested = &TestRequested
	SeNum := makeRandomString()
	sample.SeNum = &SeNum
	DateCollected := makeRandomDate()
	sample.DateCollected = &DateCollected
	DateReceived := makeRandomDate()
	sample.DateReceived = &DateReceived
	SampleType := makeRandomString()
	sample.SampleType = &SampleType
	MaterialReceived := makeRandomString()
	sample.MaterialReceived = &MaterialReceived
	MaterialReceivedNum := makeRandomString()
	sample.MaterialReceivedNum = &MaterialReceivedNum
	MaterialReceivedOther := makeRandomString()
	sample.MaterialReceivedOther = &MaterialReceivedOther
	VolumeOfBloodMarrow := makeRandomFloat()
	sample.VolumeOfBloodMarrow = &VolumeOfBloodMarrow
	SurgicalNum := makeRandomString()
	sample.SurgicalNum = &SurgicalNum
	TumorSite := makeRandomString()
	sample.TumorSite = &TumorSite
	HistoricalDiagnosis := makeRandomString()
	sample.HistoricalDiagnosis = &HistoricalDiagnosis
	TumorPercntOfTotal := makeRandomFloat()
	sample.TumorPercntOfTotal = &TumorPercntOfTotal
	TumorPercntOfCircled := makeRandomFloat()
	sample.TumorPercntOfCircled = &TumorPercntOfCircled
	ReviewedBy := makeRandomString()
	sample.ReviewedBy = &ReviewedBy
	HESlideLocation := makeRandomString()
	sample.HESlideLocation = &HESlideLocation
	NonUhnID := makeRandomString()
	sample.NonUhnID = &NonUhnID
	NameOfRequestor := makeRandomString()
	sample.NameOfRequestor = &NameOfRequestor
	DnaConcentration := makeRandomFloat()
	sample.DnaConcentration = &DnaConcentration
	DnaVolume := makeRandomFloat()
	sample.DnaVolume = &DnaVolume
	DnaLocation := makeRandomString()
	sample.DnaLocation = &DnaLocation
	RnaConcentration := makeRandomFloat()
	sample.RnaConcentration = &RnaConcentration
	RnaVolume := makeRandomFloat()
	sample.RnaVolume = &RnaVolume
	RnaLocation := makeRandomString()
	sample.RnaLocation = &RnaLocation
	WbcLocation := makeRandomString()
	sample.WbcLocation = &WbcLocation
	PlasmaLocation := makeRandomString()
	sample.PlasmaLocation = &PlasmaLocation
	CfPlasmaLocation := makeRandomString()
	sample.CfPlasmaLocation = &CfPlasmaLocation
	PbBmLocation := makeRandomString()
	sample.PbBmLocation = &PbBmLocation
	RnaLysateLocation := makeRandomString()
	sample.RnaLysateLocation = &RnaLysateLocation
	SampleSize := makeRandomString()
	sample.SampleSize = &SampleSize
	StudyID := makeRandomString()
	sample.StudyID = &StudyID
	SampleName := makeRandomString()
	sample.SampleName = &SampleName
	DateSubmitted := makeRandomDate()
	sample.DateSubmitted = &DateSubmitted
	ContainerType := makeRandomString()
	sample.ContainerType = &ContainerType
	ContainerID := makeRandomString()
	sample.ContainerID = &ContainerID
	ContainerWell := makeRandomString()
	sample.ContainerWell = &ContainerWell
	CopathNum := makeRandomString()
	sample.CopathNum = &CopathNum
	OtherIdentifier := makeRandomString()
	sample.OtherIdentifier = &OtherIdentifier
	HasSampleFiles := makeRandomBool()
	sample.HasSampleFiles = &HasSampleFiles
	DnaSampleBarcode := makeRandomString()
	sample.DnaSampleBarcode = &DnaSampleBarcode
	DnaExtractionDate := makeRandomDate()
	sample.DnaExtractionDate = &DnaExtractionDate
	DnaQuality := makeRandomString()
	sample.DnaQuality = &DnaQuality
	FfpeQcDate := makeRandomDate()
	sample.FfpeQcDate = &FfpeQcDate
	DeltaCtValue := makeRandomFloat()
	sample.DeltaCtValue = &DeltaCtValue
	Comments := makeRandomString()
	sample.Comments = &Comments
	RnasePDate := makeRandomDate()
	sample.RnasePDate = &RnasePDate
	DnaQualityByRnaseP := makeRandomFloat()
	sample.DnaQualityByRnaseP = &DnaQualityByRnaseP
	RnaQuality := makeRandomFloat()
	sample.RnaQuality = &RnaQuality
	RnaExtractionDate := makeRandomDate()
	sample.RnaExtractionDate = &RnaExtractionDate
	PatientID := strconv.Itoa(patientID)
	sample.PatientID = &PatientID

	return sample
}

func makeExperiment(sampleID int, experimentID int) models.Experiment {
	experiment := models.Experiment{}
	ChipCartridgeBarcode := makeRandomString()
	experiment.ChipCartridgeBarcode = &ChipCartridgeBarcode
	CompleteDate := makeRandomDate()
	experiment.CompleteDate = &CompleteDate
	ExperimentID := strconv.Itoa(experimentID)
	experiment.ExperimentID = &ExperimentID
	HasProjectFiles := makeRandomBool()
	experiment.HasProjectFiles = &HasProjectFiles
	OpenedDate := makeRandomDate()
	experiment.OpenedDate = &OpenedDate
	PanelAssayScreened := makeRandomString()
	experiment.PanelAssayScreened = &PanelAssayScreened
	Pcr := makeRandomString()
	experiment.Pcr = &Pcr
	Priority := makeRandomString()
	experiment.Priority = &Priority
	ProcedureOrderDatetime := makeRandomDate()
	experiment.ProcedureOrderDatetime = &ProcedureOrderDatetime
	ProjectID := makeRandomString()
	experiment.ProjectID = &ProjectID
	ProjectName := makeRandomString()
	experiment.ProjectName = &ProjectName
	SampleID := strconv.Itoa(sampleID)
	experiment.SampleID = &SampleID
	StudyID := makeRandomString()
	experiment.StudyID = &StudyID
	TestDate := makeRandomDate()
	experiment.TestDate = &TestDate

	return experiment
}

func makeResult(experimentID int, resultID int) models.Result {
	result := models.Result{}
	FailedRegions := makeRandomString()
	result.FailedRegions = &FailedRegions
	MeanDepthOfCoveage := makeRandomFloat()
	result.MeanDepthOfCoveage = &MeanDepthOfCoveage
	MlpaPcr := makeRandomString()
	result.MlpaPcr = &MlpaPcr
	Mutation := makeRandomString()
	result.Mutation = &Mutation
	OverallHotspotsThreshold := makeRandomFloat()
	result.OverallHotspotsThreshold = &OverallHotspotsThreshold
	OverallQualityThreshold := makeRandomFloat()
	result.OverallQualityThreshold = &OverallQualityThreshold
	ResultsID := strconv.Itoa(resultID)
	result.ResultsID = &ResultsID
	ExperimentID := strconv.Itoa(experimentID)
	result.ExperimentID = &ExperimentID
	UID := makeRandomString()
	result.UID = &UID
	VerificationPcr := makeRandomString()
	result.VerificationPcr = &VerificationPcr

	return result
}

func makeResultDetail(resultID int, resultdetailID int) models.Resultdetails {
	resultdetail := models.Resultdetails{}
	CNomenclature := makeRandomString()
	resultdetail.CNomenclature = &CNomenclature
	Coverage := int64(rand.Intn(1000))
	resultdetail.Coverage = &Coverage
	Exon := int64(rand.Intn(1000))
	resultdetail.Exon = &Exon
	Gene := makeRandomGene()
	resultdetail.Gene = &Gene
	Pcr := makeRandomString()
	resultdetail.Pcr = &Pcr
	PNomenclature := makeRandomString()
	resultdetail.PNomenclature = &PNomenclature
	QualityScore := makeRandomFloat()
	resultdetail.QualityScore = &QualityScore
	Result := makeRandomString()
	resultdetail.Result = &Result
	ResultsDetailsID := strconv.Itoa(resultdetailID)
	resultdetail.ResultsDetailsID = &ResultsDetailsID
	ResultsID := strconv.Itoa(resultID)
	resultdetail.ResultsID = &ResultsID
	RiskScore := makeRandomFloat()
	resultdetail.RiskScore = &RiskScore
	UID := makeRandomString()
	resultdetail.UID = &UID
	VAF := makeRandomFloat()
	resultdetail.VAF = &VAF

	return resultdetail
}

//GetGenes fills the gene array struct
func (g *GeneArray) GetGenes() *GeneArray {
	path, _ := filepath.Abs("./models/genes.yaml")
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, g)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return g
}
