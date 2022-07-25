package main
import "fmt"
func main() {
	const (
		firstName = "Raden Ario" // will not error if didn't use
		lastName 	= "Damar"
		value 		= 100
	)
	
	// fmt.Println(firstName) 
	fmt.Println(lastName)
	fmt.Println(value)

	// value = 3000 // will error
	fmt.Println(value)

}