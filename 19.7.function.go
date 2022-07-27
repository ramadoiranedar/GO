// VARIADIC FUNCTION

package main

import "fmt"

func main() {
	// sample 1
	result := sumAll(0, 10, 20, 30, 40, 2000)
	fmt.Println("Result : ", result)

	// sample 2
	slice := []int{0, 10, 20, 30, 40, 1000}
	result = sumAll(slice...)
	fmt.Println("Result : ", result)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
