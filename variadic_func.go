package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	// mengirim slice ke variadic function
	numbers := []int{1, 2, 3, 4, 5, 6}
	result = sumAll(numbers...)
	fmt.Println(result)
}
