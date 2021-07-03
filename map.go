package main

import "fmt"

func main() {

	person:= map[string]string{
		"name": "Nabil",
		"address" : "jaksel",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	delete(person, "address")
	fmt.Println(person)

	person2:= make(map[string]string)
	fmt.Println(person2)
}
