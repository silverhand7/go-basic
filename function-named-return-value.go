package main

import "fmt"

func getFullNamed() (firstName string, lastName string) {
	firstName = "Cercei"
	lastName = "Lannister"
	return
}

func main() {
	fmt.Println(getFullNamed())

	fName, lName := getFullNamed()
	fmt.Println(fName, lName)
}
