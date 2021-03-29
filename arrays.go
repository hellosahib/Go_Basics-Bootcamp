package main

import "fmt"

func main() {
	// Arrays Can be Decalred of Specific Size And Specific Type

	var numbers_word [3]string
	numbers_word[0] = "One"
	numbers_word[1] = "Two"
	numbers_word[2] = "Three"

	// Going out Of Bounds Will Throw An Error
	// Uncomment Below Line To Get Error
	// numbers[3] = "Four"

	// Type Cannot be Changed while Assignment
	// Uncomment Below Line To Get Error
	// numbers[0] = 1

	fmt.Println(numbers_word[0])
	// To Get Length Of Array
	fmt.Println(len(numbers_word))
	// Print Whole Array
	fmt.Println(numbers_word)

	// Value Can be Changed After Initalization Also
	numbers_word[0] = "I m the One"
	fmt.Println(numbers_word[0])

	// Creating Array Without Size Specified But With Values

	numbers_numeric := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers_numeric)
	fmt.Println(len(numbers_numeric))

	// Type is Required While Declaring Array
	// Uncomment Below Line To Get Error
	// directions := [...] {"North","West","South","East"}

}
