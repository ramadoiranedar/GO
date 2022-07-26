// FUNCTION RETURN VALUE

package main
import "fmt"

func main() {
	name := "Damar"
	result1 := getHello(name)
	result2 := getHello("")
	fmt.Println(result1)
	fmt.Println(result2)
}

func getHello(name string) string {
	if name == "" {
		return "Name is EMPTY"
	} else {
		return "Hello " + name
	}
}