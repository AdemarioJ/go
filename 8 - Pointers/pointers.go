package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variable1 int = 10
	var variable2 int = variable1

	variable1++
	fmt.Println(variable1, variable2)

	variable2++
	fmt.Println(variable1, variable2)

	// Pointers is a memory reference
	var variable3 int
	var pointer *int

	variable3 = 100
	pointer = &variable3

	fmt.Println(variable3, *pointer) // dereferencing

	variable3 = 150
	fmt.Println(variable3, *pointer) // dereferencing

}
