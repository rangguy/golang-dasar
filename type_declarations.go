package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRangga NoKTP = "111111"

	var contoh string = "222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRangga)
	fmt.Println(contohKtp)
}