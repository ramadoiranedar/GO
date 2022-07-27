package main

import "fmt"

func main() {
	counter := 0 // global variable set only for 0
	increment := func() {
		fmt.Println("Increment")
		///counter := 0 // solution is declaration new variable for privating variable or closures
		counter++ // reference global variable --> counter

	}

	fmt.Println("counter global", counter) // variable global counter before
	increment()
	increment()
	fmt.Println("counter in func", counter) // variable global counter after

}