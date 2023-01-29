package main

import "fmt"

type HasName interface {
	GetName() string
}

// function that implement interface contract as parameter
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// in this language, any methods that call function name of the interface will be automatically implement the interface
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

// because it has GetName method, it will be automatically implement HasName interface
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var person Person
	person.Name = "Danaerys"
	SayHello(person)

	var animal Animal
	animal.Name = "Cat"
	SayHello(animal)
}
