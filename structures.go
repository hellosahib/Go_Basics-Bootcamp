package main

import "fmt"

func main() {
	// This Will Initalize empData1 With Default Values
	var empData1 emp
	displayAge(empData1)

	empData1.name = "John Doe"
	empData1.address = "Buckingam Palace"
	empData1.age = 10
	displayAge(empData1)

	// Other Way of Creating Struct variable and Initalizing it
	// You need to pass values for all fields in this way
	empData2 := emp{"Raj", "Paris", 25}
	displayName(empData2)
	displayAge(empData2)
}

type emp struct {
	name    string
	address string
	age     int
}

func displayName(e emp) {
	fmt.Println(e.name)
}

func displayAge(e emp) {
	fmt.Println(e.age)
}
