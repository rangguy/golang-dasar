package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello, My name is ", customer.Name)
}

func main() {
	var rangga Customer
	rangga.Name = "Rangga Dwi Mahendra"
	rangga.Address = "Indonesia"
	rangga.Age = 21

	fmt.Println(rangga.Name)
	fmt.Println(rangga.Address)
	fmt.Println(rangga.Age)

	// struct literals
	dwi := Customer{
		Name:    "Rangga",
		Address: "Belgia",
		Age:     22,
	}
	fmt.Println(dwi)

	mahendra := Customer{"Dwi", "Indonesia", 23}
	fmt.Println(mahendra)

	mahendra.sayHello("rangga")
}
