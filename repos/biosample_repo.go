package repos

import (
	"encoding/json"
	"log"

	database "github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"github.com/globalsign/mgo/bson"
)

const biosampleCollection = "biosample"

//GetAllBiosamples gets all and results a list of biosamples
func GetAllBiosamples() []*models.Biosample {
	c := database.SetCollection(biosampleCollection)
	var list []*models.Biosample
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//GetManyBiosamplesByString gets all and results a list of biosamples
func GetManyBiosamplesByString(field, value string) []*models.Biosample {
	c := database.SetCollection(biosampleCollection)
	var list []*models.Biosample
	err := c.Find(bson.M{field: value}).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//GetOneBiosampleByString gets a bio sample and returns it based on the strings provided
func GetOneBiosampleByString(field, value string) *models.Biosample {
	c := database.SetCollection(biosampleCollection)
	var sample *models.Biosample
	err := c.Find(bson.M{field: value}).One(&sample)
	if err != nil {
		log.Printf("%v", err)
	}
	return sample
}

//InsertBiosample allows users to add generic objects to a collection in the database
func InsertBiosample(sample *models.Biosample) bool {
	c := database.SetCollection(biosampleCollection)
	j, _ := json.Marshal(&sample)
	var intersample interface{}
	json.Unmarshal(j, &intersample)
	err := c.Insert(intersample)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//RemoveAllBiosamples will empty a collection
func RemoveAllBiosamples() bool {
	c := database.SetCollection(biosampleCollection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
