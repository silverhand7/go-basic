package main

import "fmt"

func main() {
	var nilaiUjian = 80
	var nilaiAbsensi = 80

	lulusUjian := nilaiUjian >= 80
	lulusAbsensi := nilaiAbsensi >= 80

	lulus := lulusUjian && lulusAbsensi

	fmt.Println(lulus)
}
