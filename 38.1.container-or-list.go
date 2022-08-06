package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("RADEN")
	data.PushBack("ARIO")
	data.PushBack("DAMAR")
	data.PushFront("DENDI")

	// fmt.Println(data.Front().Next().Value)
	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	// front to behind
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// from behind to front
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
