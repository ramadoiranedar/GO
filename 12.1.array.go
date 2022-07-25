package main
import "fmt"
func main() {
	var names [3]string
	names[0] = "Raden"
	names[1] = "Ario"
	names[2] = "Damar"
	// names[3] = "Damar" // will error out of bounds max names [3]string

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
}