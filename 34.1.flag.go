package main

import (
	"flag"
	"fmt"
)

func main() {
	var host = flag.String("host", "local", "Put your host, default value is \"local\"")
	var user = flag.String("user", "root", "Put your user, default value is \"root\"")
	var password = flag.String("password", "root", "Put your password, default value is \"root\"")
	var pin = flag.Int("pin", 12345678, "Put your pin, default value is \"12345678\"")
	flag.Parse() // Dont forget to PARSE!!!

	fmt.Println("HOST:", *host)
	fmt.Println("USER:", *user)
	fmt.Println("PASSWORD:", *password)
	fmt.Println("PIN:", *pin)
}
