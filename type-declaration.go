package main

import "fmt"

func main() {
	type noKtp = string
	type married = bool

	var noKtpEko noKtp = "317"
	fmt.Println(noKtpEko)

	var marriedStatus married = true
	fmt.Println(marriedStatus)
}
