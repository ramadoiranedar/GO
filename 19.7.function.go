// VARIADIC FUNCTION

package main

import "fmt"

func main() {
	result := sumAll(0, 10, 20, 30, 40 ,60)
	fmt.Println("Result : ", result)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}