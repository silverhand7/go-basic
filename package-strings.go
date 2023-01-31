package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Thomas Shelby", "Shelby")) // true
	fmt.Println(strings.Split("Thomas Shelby", " "))         // []string{"Thomas", "Shelby"} (slice)
	fmt.Println(strings.ToUpper("thomas shelby"))
	fmt.Println(strings.ToLower("THOMAS SHELBY"))
	fmt.Println(strings.Trim("   Thomas Shelby   ", " "))
	fmt.Println(strings.ReplaceAll("Thomas Shelby is from Birmingham. Thomas is the leader of the gang.", "Thomas", "Tommy"))
}
