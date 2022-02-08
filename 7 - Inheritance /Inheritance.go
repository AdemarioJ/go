package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

type student struct {
	person
	course  string
	college string
}

func main() {
	fmt.Println("Inheritance")

	p1 := person{"Davi", "Silva", 23, 175}
	fmt.Println(p1)

	student := student{p1, "Go", "UFPE"}
	fmt.Println(student)
	fmt.Println(student.name)
}
