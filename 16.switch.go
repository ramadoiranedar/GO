package main
import "fmt"
func main() {

	name := "Dendi"

	// switch
	switch name {
	case "Damar":
		fmt.Println("Your name is Damar")
	case "Ario":
		fmt.Println("Your name is Ario")
	case "Raden":
		fmt.Println("Your name is Raden")
	default:
		fmt.Println("We need to acquaintance")
	}

	// SWITCH SHORT STATEMEN
	switch length := 3; length > 5 {
	case true:
		fmt.Println("Too long.")
	case false:
		fmt.Println("You are good.")
	}

	// SWITCH WITHOUT EXPRESSION
	switch {
		case name == "Damar":
			fmt.Println("Your name is Damar")
		case name == "Ario":
			fmt.Println("Your name is Ario")
		case name == "Raden":
			fmt.Println("Your name is Raden")
		default:
			fmt.Println("We need to acquaintance")
	}

}