// CREATE DATA STRUCT LITERALS

package main
import "fmt"

// RECOMMENDED DATA USING STRUCT
// LIKE BELOW
type Customer struct {
	Name, Address string
	Age int
	Married bool
}

func main() {
	damar := Customer { // RECOMMENDED USE THIS WAY
		Name: "Damar",
		Address: "Indonesia",
		Age: 30,
	}
	fmt.Println(damar)

	dendi := Customer {"Dendi", "Afghanistan", 38} // MAKE SURE STRUCTURAL OF THE STRUCT!
	fmt.Println(dendi)


}