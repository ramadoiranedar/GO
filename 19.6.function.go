// NAMED RETURN VALUES

package main
import "fmt"

func main() {
	a, b, c := getFullname()
	fmt.Println("Firstname: ", a)
	fmt.Println("Middle Name: ", b)
	fmt.Println("Last Name: ", c)
}

//func getFullname() (firstName, middleName, lastName string) { // SIMPLY IF SAME DATA TYPE
func getFullname() (firstName string, middleName string, lastName string) {
	firstName = "Raden"
	middleName = "Ario"
	lastName = "Damar"
	return firstName, middleName, lastName
}