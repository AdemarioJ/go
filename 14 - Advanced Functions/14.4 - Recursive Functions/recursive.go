package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	fmt.Println("Recursive function")

	position := uint(12)

	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonacci(i))
	}

	result := fibonacci(10)
	fmt.Println("Result:", result)
}
