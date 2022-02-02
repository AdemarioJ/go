package main

import (
	"fmt"
)

func main() {
	var number int64 = -100000000000000000
	fmt.Println(number)

	var number2 uint32 = 1000
	fmt.Println(number2)

	// alias
	var number3 rune = 123456 // rune is alias for int32
	fmt.Println(number3)

	var number4 byte = 123 // byte is alias for uint8
	fmt.Println(number4)

	// float
	var numberReal1 float32 = 1.23
	fmt.Println(numberReal1)

	var numberReal2 float64 = 1.232323
	fmt.Println(numberReal2)

	numberReal3 := 0.00001
	fmt.Println(numberReal3)

	// String

	var str string = "Text"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'a'
	fmt.Println(char)
}
