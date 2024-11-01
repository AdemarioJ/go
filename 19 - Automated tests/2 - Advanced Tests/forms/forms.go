package forms

import (
	"math"
)

type Form interface {
	Area() float64
}

type Retangle struct {
	Width  float64
	Height float64
}

func (r Retangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
