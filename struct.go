package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Tommy"
	customer.Address = "Birmingham"
	customer.Age = 45

	fmt.Println(customer)

	customer2 := Customer{
		Name:    "Varys",
		Address: "Dorne",
		Age:     55,
	}

	fmt.Println(customer2)
}
