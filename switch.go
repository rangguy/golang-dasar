package main

import "fmt"

func main() {
	name := "Mahendra"

	switch name {
	case "Rangga":
		fmt.Println("Halo Rangga")
	case "Dwi":
		fmt.Println("Halo Dwi")
	default:
		fmt.Println("Siapa?")
	}

	// switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi, kondisi tidak wajib dalam switch
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}