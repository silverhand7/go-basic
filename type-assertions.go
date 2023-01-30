package main

import (
	"fmt"
	"math/rand"
)

func random() interface{} {
	return "Random"
}

func randomElement() interface{} {
	var elements []interface{} = []interface{}{1, 2, "apple", true, false}
	return elements[rand.Intn(len(elements))]
}

func main() {
	var result interface{} = random()         // data type return interface
	var resultString string = result.(string) // convert interface return type to string

	fmt.Println(resultString)

	randomEl := randomElement()

	switch value := randomEl.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is integer")
	default:
		fmt.Println("Value", value, "unknown type")
	}
}
