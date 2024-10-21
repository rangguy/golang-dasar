package main

import "fmt"

func endApp1() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApp1(error bool) {
	defer endApp1()

	if error {
		panic("Error")
	}
}

func main() {
	runApp1(true)
	fmt.Println("Rangga Dwi Mahendra")
}