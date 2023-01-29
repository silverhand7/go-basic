package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func hello(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello return value " + name
}

func main() {
	sayHello()
	hello("Tyrion")
	var getHello string = getHello("Shelby")
	fmt.Println(getHello)
}
