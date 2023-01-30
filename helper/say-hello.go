package helper

import "fmt"

var Application = "Belajar Golang"
var version = 1 //tidak bisa diakses dari luar

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) { // tidak bisa diakses dari luar
	fmt.Println("Goodbye", name)
}
