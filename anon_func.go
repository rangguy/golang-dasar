package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func (name string) bool {
		return name == "Dwi"
	}

	registerUser("Rangga", blackList)
	registerUser("Dwi", func(name string) bool {
		return name == "Dwi"
	})
}