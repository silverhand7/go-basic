package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) string {
	nameFiltered := filter(name)
	return "Hello " + nameFiltered
}

type Filter func(string) string //type declaration

func sayHelloWithFilterType(name string, filter Filter) string {
	return "Hello filter " + filter(name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	fmt.Println(sayHelloWithFilter("Robb", spamFilter))

	filter := spamFilter
	fmt.Println(sayHelloWithFilter("Anjing", filter))

	fmt.Println(sayHelloWithFilterType("Nick", spamFilter))

}
