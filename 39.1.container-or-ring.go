package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	len := 5
	data := ring.New(len)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data - " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	fmt.Println(*data)
	data.Do(func(v interface{}) {
		fmt.Println(v)
	})
}
