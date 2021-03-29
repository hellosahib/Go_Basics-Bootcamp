package main

import (
	"fmt"
)

func main() {
	// Declare A Constant
	const a = 10
	fmt.Println("a:", a)

	const b float64 = 99.00
	fmt.Println("b:", b)

	// Constants cannot be reassigned
	// Uncomment Below Line to Get Error
	// b = 10.98
}
