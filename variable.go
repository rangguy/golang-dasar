package main

import "fmt"

func main() {
	// var + nama variabel + tipe variabel
	// var name string

	// name = "Rangga"
	// fmt.Println(name)

	// name = "Rangga Mahendra"
	// fmt.Println(name)

	// atau var + nama variabel + value
	var name = "Rangga"
	fmt.Println(name)

	name = "Rangga Mahendra"
	fmt.Println(name)

	// tidak menggunakan "Var", hanya bisa digunakan untuk deklarasi awal saja
	// namaVariabel := value
	name2 := "Rangga"
	fmt.Println(name2)

	// membuat banyak variabel sekaligus
	var (
		firstName = "Rangga"
		lastName = "Mahendra"
	)

	fmt.Println(firstName, lastName)
}
