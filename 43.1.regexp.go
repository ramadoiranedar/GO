package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`d([a-z][a-z][a-z])r`)

	fmt.Println(regex.MatchString("damar"))
	fmt.Println(regex.MatchString("danar"))
	fmt.Println(regex.MatchString("dimar"))

	search := regex.FindAllString("damar dendi dimar denny damor dinar", -1) // -1 all finds
	fmt.Println(search)
}
