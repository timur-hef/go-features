package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkPerimeter(t, rectangle, 36.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12, 6}, 72.0},
		{"circle", Circle{10}, 314.1592653589793},
		{"triangle", Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("shape (%b) got %g want %g", tt.shape, got, tt.want)
		}
	}
}
