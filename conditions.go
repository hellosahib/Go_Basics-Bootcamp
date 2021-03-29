package main

import (
	"fmt"
)

func main() {

	a := 20
	if a <= 20 {
		fmt.Println("A is greater than 20")
	}

	b := 20
	if b > 20 {
		fmt.Println("B is greater than 20")
	} else {
		fmt.Println("B is lesser than 20")
	}

	var x = 90
	if x < 10 {
		fmt.Println("X is less than 10")
	} else if x >= 10 && x <= 90 {
		fmt.Println("X is between 10 and 90")
	} else {
		fmt.Println("X is greater than 90")
	}
}
