package main

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func invertSignalPointer(number *int) {
	*number *= -1
}

func main() {
	number := 20
	numberInverted := invertSignal(number)
	fmt.Println("Number:", number, "Inverted:", numberInverted)

	newNumber := 40
	invertSignalPointer(&newNumber)
	fmt.Println("Number:", newNumber)
}
