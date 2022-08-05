package main

import "fmt"

func random() interface{} {
	return [3]string{
		"January",
		"February",
		"March",
	}
	// return 10000
	// return "INI STRING"
	// return 0.1
	// return true
}

func main() {
	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	case float64:
		fmt.Println("Value", value, "is float64")
	case [3]string:
		fmt.Println("Value", value, "is ARRAY")
		lenArr := len([3]string{
			"January",
			"February",
			"March",
		})
		fmt.Println(lenArr)
	default:
		fmt.Println("UNKOWN Type!")
	}

}
