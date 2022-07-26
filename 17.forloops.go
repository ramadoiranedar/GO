package main
import "fmt"
func main() {

	counter := 1

	// for
	for counter <= 10 {
		fmt.Println("Loop - ", counter)
		counter++
	}

	fmt.Println("------------------------------------------")

	// FOR STATEMENT (post statement)
	for increment := 1; increment <= 10; increment++ {
		fmt.Println("Loop - ", increment)
	}

	fmt.Println("------------------------------------------")

	// FOR BASIC AND RANGE (array, slice, map)

	slice := []string{"Raden", "Ario", "Damar"}
	// basic
	for i := 0; i< len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("#########################################")

	// RANGE
	// for _, value := range slice { // i if not used
	for i, value := range slice { // i if used
		fmt.Println("Index", i, " = ", value)
		// fmt.Println(value)
	}

	fmt.Println("#########################################")

	person := make(map[string]string)
	person["name"] = "Ario"
	person["title"] = "WebDev"

	for key, value := range person {
		fmt.Println("Key", key, " = ", value)
	}
	

}