package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("Aplikasi Selesai")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi Berjalan")
}
func main() {
	runApp(true)
}
