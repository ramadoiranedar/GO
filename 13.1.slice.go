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

	// if change value between array main and slice will change value array or on the contrary !

	slice1 := months[4:7]
	lenSlice1 := len(slice1)
	capacitySlice1 := cap(slice1)

	fmt.Println("Slice1 Length = ", lenSlice1)
	fmt.Println("Slice1 Capacity = ", capacitySlice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Gember") // create new array if capacity is full
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	// Make new Slice is safety
	// it will valid for use LOOPING
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ario" 
	newSlice[1] = "Damar"
	fmt.Println(newSlice) 

	// duplicate or copy slice,
	// make sure the capacity of slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// different array and slice
	thisARR := [...]int{1, 2, 3, 4, 5}
	thisSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("this array: ", thisARR)
	fmt.Println("this slice:", thisSlice)
}