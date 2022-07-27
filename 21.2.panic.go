package main

import "fmt"

func main() {
	runApplication(true)
}

func endApplication() {
	fmt.Println("End Application")
}

func runApplication(error bool) {
	defer endApplication()
	if error {
		panic("The Application is ERROR BRO!!!")
	}
	fmt.Println("Run Application")
}