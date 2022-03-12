package main

import "fmt"

func main() {
	fmt.Println("Control Structure")

	number := -5

	if number > 15 {
		fmt.Println("Number is greater than 15")
	} else {
		fmt.Println("Number is less or equal than 15")
	}

	if anotherNumber := number; anotherNumber > 0 {
		fmt.Println("Number is greater than 0")
	} else if number < -10 {
		fmt.Println("Number is less or equal than 0")
	} else {
		fmt.Println("Between 0 and -10")
	}
}
