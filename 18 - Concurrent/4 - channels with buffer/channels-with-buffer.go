package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Hello World"
	channel <- "Programming in Go"

	for message := range channel {
		fmt.Println(message)
	}
}
