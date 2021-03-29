package main

import "fmt"

func main() {

	// In Go & operator Tells Address of a variable
	a := 20
	fmt.Println("Address: ", &a)
	fmt.Println("Value: ", a)

	// Pointer Variable
	var b *int = &a
	fmt.Println("Address Pointer B: ", b)
	// To Print Value of Pointer
	fmt.Println("Value of Pointer B: ", *b)

	// Increment Value of Variable A using Pointers
	*b = *b + 1

	// Printing New Values
	fmt.Println("New Value of Pointer B :", *b)
	fmt.Println("New Value of A:", a)
}
