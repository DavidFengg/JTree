package dummydata

import (
	"math/rand"
	"strconv"

	models "github.com/Bio-core/jtree/models"
	repos "github.com/Bio-core/jtree/repos"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var r *rand.Rand

//MakeData makes dummy data and puts it into the db
func MakeData(numberPatients, numberSamples int) {
	r = rand.New(rand.NewSource(99))
	createPatients(numberPatients)
	createSamples(numberSamples)
}

func makeRandomString() string {
	num := rand.Intn(50)
	value := randSeq(num)
	return value
}

func makeRandomDate() string {
	date := strconv.Itoa(rand.Intn(12)) + "-" + strconv.Itoa(rand.Intn(12)) + "-" + strconv.Itoa(rand.Intn(20))
	return date
}

func createPatients(number int) {
	for i := 0; i < number; i++ {
		tempPatient := makePatient()
		repos.InsertPatient(&tempPatient)
	}
}

func createSamples(number int) {
	for i := 0; i < number; i++ {
		tempSample := models.Sample{}
		repos.InsertSample(&tempSample)
	}
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func makePatient() models.Patient {
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
	PatientID := makeRandomString()
	patient.PatientID = &PatientID
	PatientType := makeRandomString()
	patient.PatientType = &PatientType
	ReferringPhysican := makeRandomString()
	patient.ReferringPhysican = &ReferringPhysican
	SampleID := makeRandomString()
	patient.SampleID = &SampleID
	SeNum := makeRandomString()
	patient.SeNum = &SeNum
	SurgicalDate := makeRandomDate()
	patient.SurgicalDate = &SurgicalDate

	return patient
}
