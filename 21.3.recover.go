package main

import "fmt"

func main() {
	runApplication(true)
	fmt.Println("Still running not error. . .")
}

func endApplication() {
	errorMessage := recover() // MAKE SURE THE RECOVER() IS SET BEFORE THE DEFER FUNCTION OR IN THE DEFER FUNCTION REFERENCES
	if errorMessage != nil {
		fmt.Println("Errors Message:", errorMessage)
	}
	fmt.Println("End Application")
}

func runApplication(error bool) {
	defer endApplication() // DEFER FUNCTION
	if error {
		panic("The Application is ERROR BRO!!!")
	}
	fmt.Println("Run Application")
}