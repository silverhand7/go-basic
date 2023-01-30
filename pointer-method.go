package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { // jika tidak menggunakan * (pointer), maka data struct tidak berubah
	man.Name = "Mr. " + man.Name
}

func main() {
	person := Man{"Joe"} // jika di struct method tidak memerlukan simbol &
	person.Married()
	fmt.Println(person.Name)
}
