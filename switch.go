package main

import "fmt"

func main() {
	name := "asasa"

	switch name {
	case "Eko":
		fmt.Println("Masuk pak Eko")
	case "Nabil":
		fmt.Println("Masuk Pak Nabil")
	default:
		fmt.Println("???")
	}

	switch len(name) > 3 {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	switch {
	case len(name) > 5:
		fmt.Println("Lebih dari 5")
	case len(name) < 5:
		fmt.Println("Kurang dari 5")
	default:
		fmt.Println("tepat 5")
	}
}
