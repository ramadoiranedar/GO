// INTERFACE
// BIASA DIGUNAKAN FUNCTION GENERAL YANG MENGGUNAKAN CONTRACT

package main
import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// THIS CONTRACT FOR THE NAME!
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

// THIS CONTRACT FOR THE NAME!
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// way 1
	var damar Person
	damar.Name = "Damar"
	SayHello(damar)

	// way 2
	cat := Animal{
		Name: "GENJU",
	}
	SayHello(cat)
}