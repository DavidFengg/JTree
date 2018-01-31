package repos

import (
	"encoding/json"
	"log"

	database "github.com/candig/candig_mds/database"
	"github.com/candig/candig_mds/models"
)

//GetAllBiosamples gets all and results a list of biosamples
func GetAllBiosamples(collection string) []*models.Biosample {
	c := database.SetCollection(collection)
	var list []*models.Biosample
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//InsertBiosample allows users to add generic objects to a collection in the database
func InsertBiosample(collection string, sample *models.Biosample) bool {
	c := database.SetCollection(collection)
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
func RemoveAllBiosamples(collection string) bool {
	c := database.SetCollection(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
