package main

import "fmt"

func main() {
	firstName := "Arthur"
	fmt.Println(firstName)

	// Adress of operator
	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}
