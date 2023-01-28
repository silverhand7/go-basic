package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // int8 can't handle value 100000, so it will returns different number because the variable will start over again

	name := "Johnny"
	s := name[0]         // it will returned as byte
	sString := string(s) // it will convert back to string

	fmt.Println(s)
	fmt.Println(sString)
}
