package main
import "fmt"

func main() {
	// > , < , >= , <= , == , !=
	var operand1 = 1
	var operand2 = 2

	var result = operand1 == operand2
	fmt.Println(result)

	result = operand1 != operand2
	fmt.Println(result)

	result = operand1 > operand2
	fmt.Println(result)

	result = operand1 < operand2
	fmt.Println(result)

}