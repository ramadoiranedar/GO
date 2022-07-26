// NAMED RETURNING MULTIPLE VALUES WITH SEVERAL UNUSED RETURN VALUES
package main
import "fmt"

func main() {
	firstName, _, _ := getFullname()
	fmt.Println("Firstname: ", firstName)
}

func getFullname() (string, string, string) {
	return "Raden", "Ario", "Damar"
}