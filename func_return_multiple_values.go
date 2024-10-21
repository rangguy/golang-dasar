package main

import "fmt"

func getFullName() (string, string, string) {
	return "rangga", "dwi", "mahendra"
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)

	// menghiraukan salah satu return value
	firstName1, _, _ := getFullName() // menggunakan "_"
	fmt.Println(firstName1)
}
