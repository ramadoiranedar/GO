package main

import "fmt"

func main() {
	name := "Damar"
	var letterByte = name[0]
	var letterString = string(letterByte)

	fmt.Println(letterString)
}