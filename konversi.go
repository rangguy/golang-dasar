package main

import "fmt"

func main() {
	// konversi number
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// konversi char dan string
	var name = "Rangga"
	var r = name[0]
	var eString = string(r)

	fmt.Println(name) // Rangga
	fmt.Println(r) // 82
	fmt.Println(eString) // R
}
