// FUNCTION PARAMETER 

package main
import "fmt"

func main() {
	lastName := "Damar"
	sayHello("Raden Ario", lastName)
}

func sayHello(firstName string, lastname string) {
	fmt.Println("Hello", firstName, lastname, "!!!")
}