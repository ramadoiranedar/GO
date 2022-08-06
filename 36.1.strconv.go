package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, errBool := strconv.ParseBool("true")
	// boolean, errBool := strconv.ParseBool("salah")
	if errBool == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(errBool.Error())
	}

	integer, errInt := strconv.ParseInt("1000000", 2, 64) // 2 binary base int
	// integer, errInt := strconv.ParseInt("salah", 10, 64)
	if errInt == nil {
		fmt.Println(integer)
	} else {
		fmt.Println(errInt.Error())
	}

	val := strconv.FormatInt(10000, 2) // 2 format binary
	fmt.Println(val)

	val2 := strconv.Itoa(10000)
	fmt.Println(val2)

	val3, _ := strconv.Atoi("10000")
	fmt.Println(val3)
}
