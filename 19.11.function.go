// ANONYMOUS FUNCTION

package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}
	registerUser("master", blacklist)
	registerUser("admin", blacklist)
	registerUser("root", func(name string) bool {
		return name == "root"
	})
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// blacklistAdmin(name string) bool {
// 	return name == "admin"
// }

// blacklistRoot(name string) bool {
// 	return name == "admin"
// }