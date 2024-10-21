package main

import "fmt"

// interface
type HasName interface {
	GetName() string
}

// fungsi dari interface
func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// implementasi interface
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// contoh lainnya
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Rangga"}
	SayHello(person)

	animal := Animal{Name: "Cat"}
	SayHello(animal)
}
