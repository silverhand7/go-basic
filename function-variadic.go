package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 20, 20))

	// send slice to variadic argument
	slice := []int{10, 20, 30, 40}
	total := sumAll(slice...)
	fmt.Println(total)
}
