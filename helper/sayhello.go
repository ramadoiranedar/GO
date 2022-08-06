package helper

import "fmt"

var versionCode int64 = 1
var VersionStr string = "1.0"

func SayHello() {
	fmt.Println("Hello World!")
}

// ACCESS MODIFIER
// case sensitive diawali huruf kecil
// tidak bisa diexport, dan di ubah dari luar package
func sayGoodbye() {
	fmt.Println("Good bye world!")
}
