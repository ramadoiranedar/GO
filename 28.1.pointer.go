// pointer
// * sumber data
// & pengambil data
// Waring: Disarankan pada Struct data besar (banyak Data) lebih baik menggunakan Pointer irit Memory
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// addr1 := Address{"Subang", "Jakarta", "Bandung"}
	// addr2 := addr1 // without Pointer
	// addr2 := &addr1 // with Pointer

	var addr1 Address = Address{"Subang", "Jakarta", "Bandung"}
	// var addr2 Address = addr1 // without Pointer
	var addr2 *Address = &addr1 // with Pointer // same address addr1

	var addr3 *Address = &addr1 // with Pointer // same address addr1

	addr2.City = "Yogyakarta"

	*addr2 = Address{"Solo", "Malang", "Purwokerto"} // new address addr2 refer (FORCE POINTER TO THIS NEW ADDRESS)

	fmt.Println(addr1)
	fmt.Println(addr2)
	fmt.Println(addr3)

	// with Pointer
	// {Yogyakarta Jakarta Bandung}
	// &{Yogyakarta Jakarta Bandung}

	// withouth Pointer
	// {Subang Jakarta Bandung}
	// {Yogyakarta Jakarta Bandung}

	// with force Pointer
	// {Solo Malang Purwokerto}
	// &{Solo Malang Purwokerto}
	// &{Solo Malang Purwokerto}

	// new pointer reference
	var addr4 *Address = new(Address)
	addr4.City = "Surabaya"
	fmt.Println(addr4)

	// {Solo Malang Purwokerto}
	// &{Solo Malang Purwokerto}
	// &{Solo Malang Purwokerto}
	// &{Surabaya  }
}
