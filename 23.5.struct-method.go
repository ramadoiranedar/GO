package main

import (
	"fmt"
	"strconv"
)

type Customer struct {
	Name    string
	Message string
}

// STRUCT METHOD WITH FUNCTION AS PARAM AND RETURN VALUE
func (c Customer) showMessages(name string, something string) string {
	return c.Name + ": " + something
}

func main() {
	var c Customer
	c.Name = "Damar"                         // example data string from db in struct/class
	c.Message = "What is your Dream?"        // example data string from db in struct/class
	say := c.showMessages(c.Name, c.Message) // use the struct method
	fmt.Println(say)

	// testing case Messages
	lenChat := 5
	for i := 1; i <= lenChat; i++ {
		strMessage := c.Message + " " + strconv.Itoa(i)
		says := c.showMessages(c.Name, strMessage)
		fmt.Println(says)
	}
}
