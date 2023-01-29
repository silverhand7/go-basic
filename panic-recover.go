package main

import "fmt"

func endApp() {
	message := recover() // recover to send message error, and program will continue
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi Selesai")

}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
	runApp(true)
	fmt.Println("Continue the program") // program continued (program not stopped) because of the recover
}
