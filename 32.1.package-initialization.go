package main

import (
	"fmt"
	"go-basic/database"
)

func main() {
	dbName := database.GetDatabaseName()
	dbConnection := database.ConnectDatabase()

	fmt.Println(dbName)
	fmt.Println(dbConnection)
}
