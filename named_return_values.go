package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "rangga"
	middleName = "dwi"
	lastName = "mahendra"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}