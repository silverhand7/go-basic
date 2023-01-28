package main

import "fmt"

func main() {
	// variable that can't be changed (just like other programming languages)
	const registerAt int = 160000

	fmt.Println(registerAt)

	const (
		address      string = "St. Mountain"
		adressNumber int    = 10
	)

}
