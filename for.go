package main

import "fmt"

func main() {
	
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke ", counter)
	// }

	// fmt.Println("Selesai")

	// mengakses manual dengan for 
	name := []string{"rangga", "dwi", "mahendra"}
	// for i := 0; i < len(name); i++{
	// 	fmt.Println(name[i])
	// }
	
	// mengakses dengan for range
	for index, names := range name{
		fmt.Println("index ", index, "=", names)
	}

	// jika tidak memakai index
	for _, names := range name{
		fmt.Println( names)
	}
	
}