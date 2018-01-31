package database

import (
	"log"

	"gopkg.in/mgo.v2"
)

var databaseName string
var connectionString string

var session *mgo.Session
var err error

//DatabaseInit creates a connection to the database
func Init(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring + dbName

	session, err = mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
}

//SetCollection changes the collection of the datbase context
func SetCollection(collection string) *mgo.Collection {
	return session.DB(databaseName).C(collection)
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
