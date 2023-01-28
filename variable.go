package main

import "fmt"

func main() {
	var name string

	name = "Refo Junior"
	fmt.Println(name)

	name = "Silverhand"
	fmt.Println(name)

	var age = 10
	fmt.Println(age)

	hobby := "coding" // without var
	fmt.Println(hobby)

	hobby = "gaming" // re-assign value without colon
	fmt.Println(hobby)

	// declare multiple variable
	var (
		firstName string = "Geralt"
		lastName         = "of Rivia"
	)

	fmt.Println(firstName, lastName)
}
