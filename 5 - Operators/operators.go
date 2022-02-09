package main

import "fmt"

func main() {
	// Arithmetic operators
	fmt.Println("-------Arithmetic operators---------")
	sum := 3 + 2
	subtraction := 3 - 2
	multiplication := 3 * 2
	division := 3 / 2
	modulo := 3 % 2

	fmt.Println(sum, subtraction, multiplication, division, modulo)

	// Attribution
	fmt.Println("-------Attribution---------")
	var variable1 string = "String"
	variable2 := "String1"
	fmt.Println(variable1, variable2)

	// Relational Operators
	fmt.Println("-------Relational Operators---------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)

	// Logical Operators
	fmt.Println("-------Logical Operators---------")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Unary Operators
	fmt.Println("-------Unary Operators---------")
	number := 10
	number++
	number += 5 // number = number + 5
	fmt.Println(number)

	number--
	number -= 5 // number = number - 5
	fmt.Println(number)

	number *= 2 // number = number * 2
	fmt.Println(number)

	number /= 2 // number = number / 2
	number %= 3 // number = number % 2
	fmt.Println(number)

	// Ternary Operator
	fmt.Println("-------Ternary Operator---------")
	var text string
	if number > 5 {
		text = "Number is greater than 5"
	} else {
		text = "Number is less than 5"
	}

	fmt.Println(text)
}
