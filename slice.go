package main

import (
	"fmt"
)

func main() {
	names := [...]string{"rangga", "dwi", "mahendra", "ganteng", "banget", "parah"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[:]
	fmt.Println(slice3)

	// var slice []string = values[:] -> default membuat slice

	// built-in function slice
	fmt.Println(len(slice1)) // length
	fmt.Println(cap(slice1)) // capacity

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// membuat slice dari awal
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rangga"
	newSlice[1] = "Dwi"
	// newSlice[2] = "Mahendra" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Mahendra")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Ganteng"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	// perbedaan membuat array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5, 6}
	iniSlice := []int{1,2,3,4,5,6,7}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}