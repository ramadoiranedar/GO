package main
import "fmt"
func main() {

	name := "Budi"

	// if , else if , else
	if name == "Damar" {
		fmt.Println("Hello Damar. . .")
	} else if name == "Ario" {
		fmt.Println("Hello Ario. . .")
	} else if name == "Raden" {
		fmt.Println("Hello Raden. . .")
	} else {
		fmt.Println("We need to acquaintance!")
	}

	// IF SHORT STATEMENT
	if length := 1; length > 5 {
		fmt.Println("Too long.")
	} else {
		fmt.Println("You are good.")
	}

}