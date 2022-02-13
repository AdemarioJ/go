package main

import "fmt"

func calculateMath(n1, n2 int) (sum int, subtraction int) {
	sum = n1 + n2
	subtraction = n1 - n2
	return
}

func main() {
	sum, subtraction := calculateMath(10, 5)
	fmt.Println("Sum:", sum, "Subtraction:", subtraction)
}
