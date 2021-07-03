package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Nabil")
	fmt.Println("Nabil Fawwaz")
	fmt.Println("Nabil Fawwaz Elqayyim")
	fmt.Println(len("Nabil"))
	fmt.Println("Nabil"[1])

	fmt.Println(strings.Contains("Nabil Fawwaz Elqayyim", "Nabil"))
	fmt.Println(strings.Split("Nabil Fawwaz Elqayyim", " "))
	fmt.Println(strings.ToLower("Nabil Fawwaz Elqayyim"))
	fmt.Println(strings.ToUpper("Nabil Fawwaz Elqayyim"))
	fmt.Println(strings.Trim("    Nabil Fawwaz Elqayyim    ", " "))
	fmt.Println(strings.ReplaceAll("Nabil Nabil", "Nabil", "Fawwaz"))
}
