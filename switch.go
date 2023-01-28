package main

import "fmt"

func main() {
	name := "Sansa Stark"

	switch name {
	case "Arya Stark":
		fmt.Println("Arya Stark")

	case "Sansa Stark":
		fmt.Println("Sansa Stark")
		fmt.Println("of Winterfell")

	default:
		fmt.Println(name)
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Too long")
	case false:
		fmt.Println("Alright")
	}

	switch {
	case name == "Sansa Stark":
		fmt.Println("It's Sansa Stark")
	case name == "Arya Stark":
		fmt.Println("It's Arya Stark")
	default:
		fmt.Println("No one knows!")
	}
}
