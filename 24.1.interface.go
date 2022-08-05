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

// THIS CONTRACT FOR THE GETNAME! MUST BE RETURN STRING AND SAME NAME
func (person Person) GetName() string {
	return person.Name
}

func main() {
	var damar Person
	damar.Name = "Damar"
	SayHello(damar)
}
