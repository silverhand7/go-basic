package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	name, err := os.Hostname()
	fmt.Println(name, err)

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username, password)
}
