package main

import "fmt"

var n int

func init() {
	fmt.Println("Executing the function init")
	n = 10
}

func main() {
	fmt.Println("Executing the function main")
	fmt.Println(n)
}
