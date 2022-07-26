// FUNCTION RETURNING MULTIPLE VALUES

package main
import "fmt"

func main() {
	firstName, middleName, lastName := getFullname()
	fmt.Println("Fullname: ", firstName, middleName, lastName)
}

func getFullname() (string, string, string) {
	return "Raden", "Ario", "Damar"
}