package main
import "fmt"
func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}
	slice := months[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	var slice2 = months[10:]
	fmt.Println(slice)

	slice2 = append(slice2, "MY MONTH")
	fmt.Println(slice2)
}