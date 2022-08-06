package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (v UserSlice) Len() int {
	return len(v)
}

func (v UserSlice) Less(i int, j int) bool {
	return v[i].Age < v[j].Age
}

func (v UserSlice) Swap(i int, j int) {
	v[i], v[j] = v[j], v[i]
}

func main() {
	users := []User{
		{"Damar", 30},
		{"Hendra", 35},
		{"Irsan", 31},
		{"Rudi", 22},
	}
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
