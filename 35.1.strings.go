package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("RADEN ARIO DAMAR", "damar"))
	fmt.Println(strings.Split("RADEN ARIO DAMAR", " "))
	fmt.Println(strings.ToLower("RADEN ARIO DAMAR"))
	fmt.Println(strings.ToUpper("raden ario damar"))
	fmt.Println(strings.ToTitle("raden ario damar"))
	fmt.Println(strings.Trim("         RADEN          ARIO DAMAR           ", " "))
	fmt.Println(strings.ReplaceAll("RADEN ARIO DAMAR", "RADEN ARIO DAMAR", "DAMAR GANTENG"))
}
