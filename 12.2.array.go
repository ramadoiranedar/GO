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

	var dummy = [1000]string {

	}
	var lengthDummy = len(dummy)
	fmt.Println(lengthDummy)
}