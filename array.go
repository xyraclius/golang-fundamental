package main

import "fmt"

func main() {
	var names[3]string
	names[0] = "Nabil"
	names[1] = "Fawwaz"
	names[2] = "Elqayyim"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [4]int{
		1,
		2,
		3,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
