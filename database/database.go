package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

var databaseName string
var connectionString string

var session *mgo.Session
var db *mgo.Database
var err error

//Init creates a connection to the database
func Init(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring

	session, err = mgo.Dial(connectionString)
	db = session.DB(dbName)

	if err != nil {
		panic(err)
	}
}

//SetCollection changes the collection of the datbase context
func SetCollection(collection string) *mgo.Collection {
	return db.C(collection)
}

//Insert allows users to add generic objects to a collection in the database
func Insert(collection string, object interface{}) bool {
	c := SetCollection(collection)
	err := c.Insert(object)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

//RemoveAll will empty a collection
func RemoveAll(collection string) bool {
	c := SetCollection(collection)
	_, err := c.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
