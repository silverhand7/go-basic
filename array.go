package main

import "fmt"

func main() {
	var names [3]string //number of array, data type of the array's element
	names[0] = "Thomas"
	names[1] = "Shelby"
	names[2] = "Birmingham"
	fullName := names[0] + " " + names[1]
	fmt.Println(names)
	fmt.Println(fullName)

	values := [4]int{
		90, 80, 10, 20,
	}
	fmt.Println(values)
	fmt.Print(len(values)) // 4
}
