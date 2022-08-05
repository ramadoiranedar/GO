// CREATE STRUCT  WITH FIELD TAGS
package main

import (
	"encoding/json"
	"fmt"
)

type Customer struct {
	//field - datatype - tags
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
}

func main() {
	myStringJSON := `{
		"name": "Rocky",
		"age": 22
	}`

	customer1 := new(Customer)
	json.Unmarshal([]byte(myStringJSON), customer1)
	fmt.Println(customer1)

	customer2 := new(Customer)
	customer2.Name = "Rehan"
	customer2.Age = 24
	stringJSON, _ := json.Marshal(customer2)
	fmt.Printf("%s\n", stringJSON)
}
