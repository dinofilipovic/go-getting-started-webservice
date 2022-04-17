package main

import "fmt"

func main() {
	// Initailizing variable as empty pointer
	// var firstName *string
	// initializing pointer as string
	var firstName *string = new(string)
	// Tihs is associating into ponter address string "Arthur"
	*firstName = "Arthur"
	// This is printing reserved address in CPU memory
	fmt.Println(firstName)
	// This is dereferencing operation
	fmt.Println(*firstName)
}
