package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Loop: ", counter)
	}

	names := []string{"Nabil", "Fawwaz", "Elqayyim"}

	for _, name := range names {
		//fmt.Println("index", i, " = ", name)
		fmt.Println(name)
	}
}
