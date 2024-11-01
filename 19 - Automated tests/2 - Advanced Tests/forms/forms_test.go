package forms

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangle", func(t *testing.T) {
		ret := Retangle{10, 12}
		area := float64(120)
		areaReceived := ret.Area()

		if area != areaReceived {
			t.Fatalf("Expected %f but got %f", area, areaReceived)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		area := float64(math.Pi * 100)
		areaReceived := circle.Area()

		if area != areaReceived {
			t.Fatalf("Expected %f but got %f", area, areaReceived)
		}
	})
}
