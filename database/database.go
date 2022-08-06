package database

import "fmt"

var dbName string
var dbConnect bool

// like Construct in CI automate running without call this  init() func
func init() {
	fmt.Println("Init DATABASE")
	dbName = "PostgreSQL"
	dbConnect = true
}

func GetDatabaseName() string {
	return dbName
}

func ConnectDatabase() bool {
	return dbConnect
}
