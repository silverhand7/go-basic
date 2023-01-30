package main

import (
	"fmt"
	"go-basic/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
