package main

import (
	"errors"
	"fmt"
)

func main() {
	var exampleError error = errors.New("Oops..error!")
	fmt.Println(exampleError)
}
