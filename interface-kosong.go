package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func ArbitraryParameter(parameter ...interface{}) interface{} {
	return parameter
}

func main() {
	var data interface{} = Ups(24)
	fmt.Println(data)

	fmt.Println(ArbitraryParameter("Terserah", 1313, false, "Bebas", "lol"))
}
