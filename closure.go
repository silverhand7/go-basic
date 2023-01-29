package main

import "fmt"

func main() {
	counter := 0
	name := "Jordan"

	increment := func() {
		name := "Milner"
		fmt.Println("increment")
		counter++
		fmt.Println(name) // Milner
	}

	increment()
	increment()
	fmt.Println(counter) // 2
	fmt.Println(name)    // Jordan
}
