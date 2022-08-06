package main

import (
	"fmt"
	"go-basic/helper"
)

func main() {
	helper.SayHello()
	// helper.sayGoodbye() // ACCESS MODIFIER: Unexported/tidak bisa diexport, dan di ubah dari luar package
	fmt.Println(helper.VersionCode)
}
