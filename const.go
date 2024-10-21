package main

import "fmt"

func main(){
	// jika tidak digunakan tidak akan error, tidak seperti variable biasa
	const firstName string = "Rangga"
	const lastName = "Mahendra"

	// error, tidak bisa mengubah nilai dari variable constant
	// firstName = "Tidak bisa diubah"
	// lastName = "tidak bisa diubah"

	// deklarasi multiple constant
	const (
		firstName1 string = "Rangga"
		lastName2 = "Mahendra"
	)

	fmt.Println(firstName1, lastName2)
}