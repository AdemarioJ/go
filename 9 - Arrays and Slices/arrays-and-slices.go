package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrays()
	slices()
	arrays_internal()
}

func arrays() {
	fmt.Println("Arrays")

	// Arrays
	var array1 [5]string
	fmt.Println(array1)

	array1[0] = "Position 1"
	fmt.Println(array1)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	fmt.Println(reflect.TypeOf(array3))
}

func slices() {
	fmt.Println("Slices")

	// Slices
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18)
	fmt.Println(slice)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Position 2 - Updated"
	fmt.Println(slice2)
}

func arrays_internal() {
	// Arrays Internal
	slice := make([]float32, 10, 11)
	fmt.Println(slice)

	slice = append(slice, 5)
	slice = append(slice, 6)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice2 := make([]float32, 5)
	fmt.Println(slice2)
	slice2 = append(slice2, 10)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
