package main

import "fmt"

func main() {
	var name1 = "hello"
	var name2 = "world"
	var result bool = name1 == name2

	fmt.Println(result)

	number1 := 10
	number2 := 12

	fmt.Println(number1 > number2)
	fmt.Println(number1 < number2)
	fmt.Println(number1 != number2)
}
