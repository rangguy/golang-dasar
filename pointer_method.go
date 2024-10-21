package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rangga := Man{"Rangga"}
	rangga.Married()

	fmt.Println(rangga.Name)
}