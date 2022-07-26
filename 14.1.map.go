package main
import "fmt"
func main() {
	// full detail declaration
	// var person map[string]string = map[string]string{
	// 	"name":	"Ario",
	// 	"address":	"Bekasi",
	// }

	// simple declaration
	person := map[string]string{
		"name":	"Ario",
		"address":	"Bekasi",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(person)

	// add data map
	person["title"] = "Programmer"
	fmt.Println(person["title"])

	// edit data map
	person["title"] = "Web Developer"
	fmt.Println(person["title"])

	fmt.Println(person)

	// delete data map
	delete(person, "title")

	fmt.Println(person)

	// get Length of map
	lenPerson := len(person)
	fmt.Println(lenPerson)

	// make new map
	student := make(map[string]string)
	student["nim"] = "2017310032"
	student["major"] = "Technology Information"
	fmt.Println(student)
}