package main

import (
	"fmt"
	"test-introduction/addresses"
)

func main() {
	typeAddress := addresses.TypeOfAddress("Rua")
	fmt.Println(typeAddress)
}
