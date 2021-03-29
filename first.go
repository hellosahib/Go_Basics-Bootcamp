package main

import (
	"fmt"
)

func main() {
	var x int
	x = 4
	fmt.Println("x: ", x)

	var y = 4
	fmt.Println("y: ", y)

	var z float64 = 5
	// Cannot Reassign Variable type
	// Uncomment Below Line to Get Error
	// z = "I am Not Allowed"
	fmt.Println("z: ", z)

	var g string = "Hello G"

	fmt.Println(g)
}
