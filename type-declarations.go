package main

import "fmt"

func main() {
	type NoKTP string

	var personKTP NoKTP = "128582948924842"
	fmt.Println(personKTP)

	type IsMarried bool

	var marriedStatus IsMarried = true
	fmt.Println(marriedStatus)
}
