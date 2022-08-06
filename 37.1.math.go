package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.77))
	fmt.Println(math.Round(1.33))
	fmt.Println(math.Floor(1.77))
	fmt.Println(math.Floor(1.33))
	fmt.Println(math.Ceil(1.77))
	fmt.Println(math.Ceil(1.33))
	fmt.Println(math.Max(100, 200))
	fmt.Println(math.Min(100, 200))
}
