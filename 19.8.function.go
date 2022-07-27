// FUNCTION VALUE

package main

import "fmt"

func main() {
	// THIS 
	// goodbye := getGoodbye
	// fmt.Println(goodbye("Jeje"))

	// OR THIS
	goodbye := getGoodbye("Jeje")
	fmt.Println(goodbye)
}

func getGoodbye(name string) string {
	return "Good Bye " + name
}