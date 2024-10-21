package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Rangga"
	names[1] = "Dwi"
	names[2] = "Mahendra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array langsung
	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)

	var values_undefined = [...]int{
		90,
		80,
	}

	// builtin function array
	fmt.Println(values_undefined)
	fmt.Println(len(values_undefined)) // length of array
	values_undefined[1] = 95
	fmt.Println(values_undefined)
}
