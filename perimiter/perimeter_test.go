package perimeter_test

import (
	"github.com/vicgrygorchyk/example_hello/perimiter"
	"testing"
)

func TestArea(t *testing.T) {

	assertArea := func(t testing.TB, shape perimeter.Shape, expected float64) {
		t.Helper()
		got := shape.Area()
        if got != expected {
            t.Errorf("got %g want %g", got, expected)
        }
	}

    t.Run("rectangles", func(t *testing.T) {
        rectangle := perimeter.Rectangle{5, 5}
        assertArea(t, rectangle, 25)
    })

    t.Run("circles", func(t *testing.T) {
        circle := perimeter.Circle{10}
        want := 314.1592653589793
		assertArea(t, circle, want)
    })

}

func TestPerimeter(t *testing.T) {

	assertArea := func(t testing.TB, shape perimeter.Shape, expected float64) {
		t.Helper()
		got := shape.Perimeter()
        if got != expected {
            t.Errorf("got %g want %g", got, expected)
        }
	}

    t.Run("rectangles", func(t *testing.T) {
        rectangle := perimeter.Rectangle{12, 6}
		want := 36.0
        assertArea(t, rectangle, want)
    })

    t.Run("circles", func(t *testing.T) {
        circle := perimeter.Circle{10}
        want := 62.800000000000004
		assertArea(t, circle, want)
    })

}

