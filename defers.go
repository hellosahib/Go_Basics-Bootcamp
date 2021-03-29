package main

import "fmt"

//Defer statements are used to defer the execution of a function call
//until the function that contains the defer statement completes execution.

func display(num int) {
	fmt.Println(num)
}

// Defer works in LIFO order
// Here Defer Functions will be called after main Function is completed

func main() {
	defer display(1)
	defer display(2)
	defer display(3)
	fmt.Println(4)
}
