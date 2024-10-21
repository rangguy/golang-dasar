package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Rangga",
		"address": "Tangerang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	// tidak ada key "umur"
	fmt.Println(person["umur"])
	
	// built-in function map
	book := make(map[string]string)
	book["title"] = "Pemrograman Go"
	book["author"] = "Rangga"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	
	fmt.Println(book)
}