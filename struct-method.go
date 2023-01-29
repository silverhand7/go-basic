package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHelloToCustomer(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Tommy"
	customer.Address = "Birmingham"
	customer.Age = 45

	customer.sayHelloToCustomer("John")
}
