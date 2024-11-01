package main

import "fmt"

func function1() {
	fmt.Println("Executing function1")
}

func function2() {
	fmt.Println("Executing function2")
}

func studentApproved(n1, n2 float32) bool {
	defer fmt.Println("Average calculated. Result will be returned")
	fmt.Println("Function for verifying student approval")

	average := (n1 + n2) / 2

	if average >= 6 {
		return true
	}

	return false
}

func main() {
	defer function1()
	function2()

	fmt.Println(studentApproved(7.0, 6.0))
}
