package main

import "fmt"

// Calc is the function Name which accepts two integers
// and returns sum and difference

// Functions can return more than 1 value without Arrays

/**
* Important Note
** If function is private use camelCasing for naming
** If function is public use First Word as Capital
** eg add_two will make function private
** eg Add_two will make function public
**/
func calc(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func printOne() {
	fmt.Println("Called Print One Functions")
}

func main() {
	printOne()

	x, y := 10, 20

	// Calling the Function and assigning values with return values
	sum, diff := calc(x, y)
	fmt.Println("Sum: ", sum)
	fmt.Println("Diff: ", diff)
}
