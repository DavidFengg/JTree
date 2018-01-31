package repos

import (
	"encoding/json"
	"log"

	database "github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
)

//GetAllIndividuals gets all and results a list of biosamples
func GetAllIndividuals(collection string) []*models.Individual {
	c := database.SetCollection(collection)
	var list []*models.Individual
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//InsertIndividual allows users to add generic objects to a collection in the database
func InsertIndividual(collection string, individual *models.Individual) bool {
	c := database.SetCollection(collection)
	j, _ := json.Marshal(&individual)
	var interindividual interface{}
	json.Unmarshal(j, &interindividual)
	err := c.Insert(interindividual)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//RemoveAllIndividuals will empty a collection
func RemoveAllIndividuals(collection string) bool {
	c := database.SetCollection(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
