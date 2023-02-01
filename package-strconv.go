package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	/*
		Untuk parameter kedua ParseInt
		desimal = 10
		binary = 2
		octal = 8
		hexadecimal = 16
	*/

	number, _ := strconv.ParseInt("10000", 10, 64)
	fmt.Println(number)

	number2, _ := strconv.Atoi("2000")
	fmt.Println(number2)

}
