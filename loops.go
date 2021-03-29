package main

import (
	"fmt"
)

func main() {
	// Go Supports Only For Loop
	var x int
	for x = 5; x > 0; x-- {
		fmt.Println("x:", x)
	}

	// I is created in Scope of for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("i:", i)
	}

	// Can be Created again
	i := 50
	fmt.Println("i:", i)
}
