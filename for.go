package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Loop", counter)
		counter++
	}

	fmt.Println(counter)

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Harry", "Styles"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	names := []string{"Jon", "Snow"}

	for index, name := range names {
		fmt.Println("Index ", index, "=", name)
	}

	person := make(map[string]string)
	person["name"] = "Jaime Lannister"
	person["title"] = "Kingslayer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
