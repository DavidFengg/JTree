package database

import (
	"database/sql"

	// SQL driver for mysql
	_ "github.com/go-sql-driver/mysql"
)

var databaseName string
var connectionString string

var err error

//DB is the database reference
var DB *sql.DB

//Init creates a connection to the database
func Init(dbName, connectionstring string) {
	databaseName = dbName
	connectionString = connectionstring

	DB, err = sql.Open("mysql", "root:waterloo@/JTree")
	if err != nil {
		panic(err)
	}
}
