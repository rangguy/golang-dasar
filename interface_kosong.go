package main

import "fmt"

func Ups() interface{} {
	// return 1
	// return true
	return "UPS"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}