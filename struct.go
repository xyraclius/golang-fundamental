package main

import "fmt"

type Customer struct {
	Name, Address string
	Age  int
}
func main() {
	var customer Customer
	customer.Age = 25
	customer.Address = "Jakarta"
	customer.Name = "Nabil"

	fmt.Println(customer)

	// Struct Literal 1
	joko := Customer{
		Name: "Joko",
		Address: "Bandung",
		Age: 22,
	}

	fmt.Println(joko)

	// Struct Literal 2
	budi := Customer{"Budi","Bekasi",23}

	fmt.Println(budi)
}
