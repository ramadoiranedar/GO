package main

import "fmt"

func main() {
	runApplication(10)
}

func logging() {
	fmt.Println("End called function")
}

func runApplication(value int) {
	// use defer first
	defer logging() // logging() still run after error in some line
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}