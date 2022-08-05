package main

import "fmt"

//available data nil: interface, function, map, slice, pointer, channel

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("damar")
	// var person map[string]string
	// fmt.Println(person)

	if person == nil {
		fmt.Println("DATA IS EMPTY")
	} else{
		fmt.Println(person)
	}
}
