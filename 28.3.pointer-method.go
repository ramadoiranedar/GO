// pointer METHOD
// & pengambil data
// * sumber data
// Waring: Disarankan pada Struct data besar (banyak Data) lebih baik menggunakan Pointer irit Memory
// cukup tambahkan Bintang (*) pada struct method referen
package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("D1 METHOD", man.Name)
}

func main() {
	damar := Man{"Damar"}
	damar.Married()
	fmt.Println(damar)
	fmt.Println(damar.Name)
}
