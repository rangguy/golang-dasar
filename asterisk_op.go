package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 
	address2.City = "Bandung"
	fmt.Println(address1) // berubah menjadi bandung
	fmt.Println(address2)

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // tetap tidak berubah
	fmt.Println(address2)
	
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // berubah menjadi jakarta
	fmt.Println(address2)
}