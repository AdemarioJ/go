package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(1)
	generic(1.5)
	generic(true)

	fmt.Println(1, 2, "String", false, true, float64(123))
}
