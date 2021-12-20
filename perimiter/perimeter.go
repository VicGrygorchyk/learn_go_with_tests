package perimeter

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle) Area() float64 {
    return rect.Width * rect.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Perimeter() float64 {
	return 2 * circle.Radius * 3.14
}

func (circle Circle) Area() float64 {
    return circle.Radius * circle.Radius * math.Pi
}
