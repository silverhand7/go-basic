package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second(), now.Unix())

	utc := time.Date(2020, 01, 01, 12, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1998-01-01")
	fmt.Println(parse)

}
