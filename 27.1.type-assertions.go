package main

import "fmt"

func random() interface{} {
	return "OK"
	// return true // will throw panic if not using TYPE ASSERTIONS SWITCH like 27.2
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)
}
