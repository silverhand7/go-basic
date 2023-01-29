package main

import "fmt"

func getFullName() (string, string) {
	return "Thomas", "Shelby"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
