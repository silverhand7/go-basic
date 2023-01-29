package main

import "fmt"

func random() interface{} {
	return "Random"
}

func randomElement() interface{} {
	// var elements interface{} = []interface{}{1, 2, "apple", true, false} // gagal, I want to make array with different data types and return random element
	// return elements
	// return "String"
	return 1
}

func main() {
	var result interface{} = random()         // data type return interface
	var resultString string = result.(string) // convert interface return type to string

	fmt.Println(resultString)

	var randomEl interface{} = randomElement()

	switch value := randomEl.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Value", value, "unknown type")
	}
}
