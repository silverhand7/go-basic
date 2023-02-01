package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack(12)
	data.PushBack(99)
	data.PushBack(52)
	data.PushBack(37)
	data.PushFront(9)

	fmt.Println("First:", data.Front().Value) // first
	fmt.Println("Last:", data.Back().Value)   // last

	fmt.Println("Second:", data.Front().Next().Value)                    // 99 (data kedua)
	fmt.Println("Third:", data.Front().Next().Next().Value)              // 52 (data ketiga)
	fmt.Println("First Again (Prev):", data.Front().Next().Prev().Value) // 9 (pertama)

	// front to back
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// back to front
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
