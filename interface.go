package main

import "fmt"

type Person struct {
	Name string
}
type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello, ", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

//empty interface

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var nabil Person
	nabil.Name = "Nabil"

	var data interface{} = Ups(3)
	fmt.Println(data)
}
