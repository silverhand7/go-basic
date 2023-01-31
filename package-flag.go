package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your db username")
	var password *string = flag.String("password", "root", "Put your db password")

	flag.Parse() // we can run this with go run flag.go -host=yourhost -user=youruser -password=yourpassword

	fmt.Println("Host:", *host) // because it's a pointer, we need to add * otherwise it will shows the pointer address (0x...)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
}
