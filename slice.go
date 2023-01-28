package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bulan Baru") // append slice2 which already the maximum capacity of the months array will create a new array
	slice3[1] = "Desember Baru"
	fmt.Println(slice3)

	fmt.Println(slice2) // that's why this is not affected the new data because it uses different array source
	fmt.Println(months) // this still like the same

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Thomas"
	newSlice[1] = "Shelby"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice)) // if the capacity  is less than the source slice, the copy slice will trimmed
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3} //but not defined the capacity
	iniSlice := []int{1, 2, 3}
	fmt.Println(iniArray, len(iniArray), cap(iniArray))
	fmt.Println(iniSlice, len(iniSlice), cap(iniSlice))
}
