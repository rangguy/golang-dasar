package helper

import "fmt"

var version = "1.0.0"      // tidak bisa diakses dari luar package
var Application = "golang" // bisa diakses dari luar package

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Helo " + name
}

func Contoh() {
	sayGoodBye("rangga")
	fmt.Println(version)
}