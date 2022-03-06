package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

func setArea(f form) {
	fmt.Println("The area of the rectangle is:", f.area())
}

type retangle struct {
	width  float64
	height float64
}

func (r retangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := retangle{width: 10, height: 5}
	setArea(r)

	c := circle{radius: 5}
	setArea(c)
}
