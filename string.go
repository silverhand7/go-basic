package main

import "fmt"

func main() {
	fmt.Println("String")
	fmt.Println("Jumlah Karakter", len("String"))
	fmt.Println("Huruf pertama dari string", "String"[0]) //will return byte number, not the string "s"
}
