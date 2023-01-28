package main

import "fmt"

func main() {
	var name = "Jon"

	if name == "Jon" {
		fmt.Println("Hello Jon")
	} else {
		fmt.Println("Hello", name)
	}

	if length := len(name); length > 5 { // if short statement
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Good name")
	}
}
