package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// membuat alias for multiple users struct
type UsersSlice []User

/*
	To use package sort must implementt sort contract
	Len, Less, Swap

*/

func (users UsersSlice) Len() int {
	return len(users)
}

func (users UsersSlice) Less(i, j int) bool {
	return users[i].Name < users[j].Name
}

func (users UsersSlice) Swap(i, j int) {
	users[i], users[j] = users[j], users[i]
}

func main() {
	users := []User{
		{
			Name: "Tommy",
			Age:  40,
		},
		{
			Name: "Arthur",
			Age:  45,
		},
		{
			Name: "John",
			Age:  35,
		},
	}

	sort.Sort(UsersSlice(users)) // use UsersSlice because it implements sort contract

	fmt.Println(users)
}
