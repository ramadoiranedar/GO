// FUNCTION AS PARAMETER

package main

import "fmt"

type Filter func(string) string //Function Type Declarations

func main() {
	sayHelloWithFilter("Damar", spamFilter)
	sayHelloWithFilter("BABI", spamFilter)
}

// func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter Filter) { // With Function Type Declarations
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "BABI" {
		return ". . ."
	} else {
		return name
	}
}