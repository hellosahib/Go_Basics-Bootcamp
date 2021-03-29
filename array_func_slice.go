package main

import "fmt"

func main() {
	// Slice is an Portion Of Array
	// It Holds Direct Reference To Array Elements

	array_1 := [...]string{"one", "two", "three"}
	fmt.Println(array_1)

	// Slice Decalration
	// [Start Index Included : End Index Excluded]
	var b []string = array_1[1:3]
	fmt.Println("Slice After Creation: ", b)

	c := array_1[0:1]
	fmt.Println(c)

	// Changing Value in Slice
	// Affects Main Array Also
	c[0] = "Changed"
	fmt.Println("Slice After Modifying: ", c)
	fmt.Println("Orignal Array After Slice Modification: ", array_1)
}
