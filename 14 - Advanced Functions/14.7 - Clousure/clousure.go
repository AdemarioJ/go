package main

import "fmt"

func clousere() func() {
	text := "Inside function clousere"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Inside function main"
	fmt.Println(text)

	newFunction := clousere()
	newFunction()
}
