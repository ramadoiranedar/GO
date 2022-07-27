// RECURSIVE FUNCTION

package main

import "fmt"

func main() {
	value := 10
	var factorial int

	// using loop
	factorial = factorialLoop(value)
	fmt.Println(factorial)
	
	// using recursive function
	factorial = factorialRecursive(value)
	fmt.Println(factorial)
}

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}