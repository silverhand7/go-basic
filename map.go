package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"firstName": "Danaerys",
		"lastName":  "Targaryen",
	}

	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])

	person["gender"] = "Female"

	fmt.Println(person["gender"])

	book := make(map[string]string)
	book["title"] = "Game of Thrones"
	book["author"] = "George R. R. Martin"
	book["ups"] = "Delete"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
