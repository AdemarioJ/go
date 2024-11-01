package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Attempt to recover execution!")
		fmt.Println("Recovered from panic:", r)
	}
}

func studentApproved(n1, n2 float64) bool {
	defer recoverExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("The average is exactly 6.0!")
}

func main() {
	fmt.Println(studentApproved(6.0, 6.0))
	fmt.Println("Post execution")
}
