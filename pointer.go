package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Denpasar", "Bali", "Indonesia"}
	address2 := address1
	address2.City = "Tabanan"

	address3 := &address1     // pointer ke address 1
	address3.City = "Gianyar" // City Address 1 akan berubah

	fmt.Println(address1) // City tetap Denpasar
	fmt.Println(address2) // City Tabanan
	fmt.Println(address3)

	address4 := Address{"Malang", "Jawa Timur", "Indonesia"}
	address5 := &address4
	address5.City = "Jember"
	address6 := &address4
	*address5 = Address{"Surabaya", "Jawa Timur", "Indonesia"} // semua yg mereference ke address4 (karena address5 reference ke address4) akan berubah
	fmt.Println(address4)
	fmt.Println(address5)
	fmt.Println(address6)

	// var address7 *Address = &Address{"Subang", "Jawa Barag", "Indonesia"}
	var address7 *Address = new(Address)
	address8 := address7
	address8.City = "Buleleng"
	fmt.Println(address7) // address7 akan berubah juga
	fmt.Println(address8)

	var alamat = Address{
		City:     "Denpasar",
		Province: "Bali",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat) // Jika tidak menggunakan & dan * maka data tidak akan berubah

	fmt.Println(alamat)

}
