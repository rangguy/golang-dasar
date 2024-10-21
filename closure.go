package main

import "fmt"

func main() {
	counter := 0
	
	increment := func() {
		fmt.Println("Incremenet")
		counter++
	}

	increment()
	increment()
	increment()

	
	fmt.Println(counter)
}