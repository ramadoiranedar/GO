package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	fmt.Println("Args array: ", args)
	fmt.Println("Args: ", args[1])
	fmt.Println("Args: ", args[2])
	fmt.Println("Args: ", args[3])

	myHostname, errOS := os.Hostname()
	if errOS == nil {
		fmt.Println("Hostname:", myHostname)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	myUsername := os.Getenv("APP_USERNAME")
	myPassword := os.Getenv("APP_PASSWORD")
	fmt.Println("Username", myUsername)
	fmt.Println("Password", myPassword)
}
