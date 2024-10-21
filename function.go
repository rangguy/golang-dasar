package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Rangga", "Mahendra")
}

