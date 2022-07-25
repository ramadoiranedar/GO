package main
import "fmt"
func main() {
	var names = [3]string {
		"Raden",
		"Ario",
		"Damar",
	}
	fmt.Println(names)
	
	var lengthNames = len(names)
	fmt.Println(lengthNames)
}