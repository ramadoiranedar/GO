package main
import "fmt"

func main() {
	// && , || , !
	var result = (1 < 2) && (1 == 1)
	fmt.Println(result)

	result = (1 < 2) || (1 == 1)
	fmt.Println(result)

	result = !(1 < 2)
	fmt.Println(result)
}