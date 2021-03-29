package main

import "fmt"

func main() {

	array_1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array1 Creation: ", array_1)

	slice_1 := array_1[1:3]
	fmt.Println("Slice1 Creation: ", slice_1)

	array_2 := [...]int{5, 6, 7}
	fmt.Println("Array2 Creation: ", array_2)

	slice_2 := array_2[1:3]
	fmt.Println("Slice2 Creation: ", slice_2)

	// Append Function Only Works on Slices
	// To Append Additional Values to Slices
	/**
	* If you array size is greater than appending items size
	* Than Values at particular Index Gets Replaced With New Items
	* If Your Orignal Array is of lesser Size than new Items Being Appended
	* Than No Changes are done to Orignal Array
	**/
	slice_1 = append(slice_1, 10)
	fmt.Println("Appended Values Slice1: ", slice_1)
	fmt.Println("Orignal Array1 After Append: ", array_1)

	slice_2 = append(slice_2, 20, 30)
	fmt.Println("Appended Values Slice2: ", slice_2)
	fmt.Println("Orignal Array2 After Append: ", array_2)

	// To Append Two Slices use ... At end
	slice_1 = append(slice_1, slice_2...)
	fmt.Println("Updated Slice 1 After Append: ", slice_1)
}
