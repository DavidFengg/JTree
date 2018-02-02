package repos

import (
	"encoding/json"
	"log"

	database "github.com/CanDIG/candig_mds/database"
	"github.com/CanDIG/candig_mds/models"
	"github.com/globalsign/mgo/bson"
)

const individualCollection = "individual"

//GetAllIndividuals gets all and results a list of individuals
func GetAllIndividuals() []*models.Individual {
	c := database.SetCollection(individualCollection)
	var list []*models.Individual
	err := c.Find(nil).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//GetManyIndividualsByString gets all and results a list of individuals
func GetManyIndividualsByString(field, value string) []*models.Individual {
	c := database.SetCollection(individualCollection)
	var list []*models.Individual
	err := c.Find(bson.M{field: value}).All(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

//GetOneIndividualByString gets a bio sample and returns it based on the strings provided
func GetOneIndividualByString(field, value string) *models.Individual {
	c := database.SetCollection(individualCollection)
	var individual *models.Individual
	err := c.Find(bson.M{field: value}).One(&individual)
	if err != nil {
		log.Printf("%v", err)
	}
	return individual
}

//InsertIndividual allows users to add generic objects to a collection in the database
func InsertIndividual(individual *models.Individual) bool {
	c := database.SetCollection(individualCollection)
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
func RemoveAllIndividuals() bool {
	c := database.SetCollection(individualCollection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
