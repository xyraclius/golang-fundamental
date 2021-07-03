package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("N([\\w])l")
	fmt.Println(regex.MatchString("NAl"))
}
