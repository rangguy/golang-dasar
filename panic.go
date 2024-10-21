package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp() // tetap akan dijalankan walaupun panic

	if error {
		panic("Ups Error") // akan menghentikan program jika terjadi error
	}
}

func main() {
	runApp(true)
}