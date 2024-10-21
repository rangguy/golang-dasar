package main

import "fmt"

func random() any {
	return true
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // akan terjadi panic
	// fmt.Println(resultInt)

	// menggunakan switch, lebih aman untuk type assertions
	switch value := result.(type) {
	case string:
		fmt.Println("String is",value)
	case int:
		fmt.Println("Int is", value)
	default:
		fmt.Println("Unknown", value)
	}
}