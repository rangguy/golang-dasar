package main

import "fmt"

func main() {
	name := "Rangga Dwi"

	if name == "Rangga" {
		fmt.Println("Halo Rangga")
	} else if name == "Dwi" {
		fmt.Println("Halo Dwi")
	} else if name == "Mahendra" {
		fmt.Println("Halo Mahendra")
	} else {
		fmt.Println("bukan Rangga")
	}

	// if dengan short statement sebelum kondisi
	if length := len(name);length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
