package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()

	fmt.Println("Run application")

	// error, logging tidak akan dieksekusi
	// logging()
}

func main() {
	runApplication()
}
