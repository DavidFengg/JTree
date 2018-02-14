package database

import (

	// SQL driver for mysql
	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

var databaseName string
var connectionString string

var err error

//DB is the database reference
var DB *sqlx.DB

//Init creates a connection to the database
func Init(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring
	DB, err = sqlx.Connect(databaseName, connectionString)
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(100)
}
