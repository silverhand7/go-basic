package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)
	// data.Value = "Ring First"

	// var data2 = data.Next()
	// data2.Value = "Ring Second"

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.Itoa(i)
		data = data.Next()
	}

	fmt.Println(data)

	// don't use for loop because it will never end, instead use built-in function Do
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
