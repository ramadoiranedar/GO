// CREATE DATA STRUCT

package main
import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var damar Customer
	damar.Name = "Damar"
	damar.Address = "Indonesia"
	damar.Age = 30

	fmt.Println(damar)
	fmt.Println(damar.Name)
	fmt.Println(damar.Address)
	fmt.Println(damar.Age)
}