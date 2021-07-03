package main

import (
	"fmt"
)

func main() {
	fistName, middleName, _ := getFullName()
	fmt.Println(fistName, middleName)

	fmt.Println(getCompleteName())
	fmt.Println(sumAll(1, 5, 4))

	sayHelloWithFilter("Nabil", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	//anonymous function
	blackList := func(name string) bool {
		return name == "anjing"
	}

	registerUser("nabil", blackList)
	registerUser("anjing", blackList)
}

func sayHello(name int) {
	fmt.Println(name)
}

func sayHello2(name string) string {
	return name
}

// return multiple value function
func getFullName() (string, string, string) {
	return "Nabil", "Fawwaz", "Elqayyim"
}

func getCompleteName() (firstName, middleName, lastName string) {
	return "Nabil", "Fawwaz", "Elqayyim"
}

// varargs / variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// function as parameter

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// for anonymous function

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked, ", name)
	} else {
		fmt.Println("Welcome, ", name)
	}

}

// recursive function
// without recursive
func factorialLoop(value int) int {
	result := 1
	for i := value; i < value; i-- {
		result *= i
	}
	return result
}

// with recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
