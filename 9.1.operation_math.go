package main
import "fmt"

func main() {
	// math operations
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var result = 10 * 10
	fmt.Println(result)

	// augmented assigment
	result += 900
	fmt.Println(result)

	result -= 900
	fmt.Println(result)

	// unary
	var i = 0
	fmt.Println(i)
	i++
	fmt.Println(i)
	var state = !true
	fmt.Println(state)
}