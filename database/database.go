package database

import (

	// SQL driver for mysql
	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

var databaseName string
var connectionString string

var err error

//DBSelect is the database reference for selects
var DBSelect *sqlx.DB

//DBUpdate is the database reference for updates
var DBUpdate *sqlx.DB

//Init creates a connection to the database
func Init(dbName, connectionstring string, DB *sqlx.DB) *sqlx.DB {
	databaseName = dbName
	connectionString = connectionstring
	DB, err = sqlx.Connect(databaseName, connectionString)
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(100)
	return DB
}
