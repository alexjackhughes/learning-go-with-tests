package shapes

import (
	"math"
)

// Shape is a dimensional object interface.
type Shape interface {
	Area() float64
}

// Triangle is a 3-sided shape.
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter takes the height and width of a triangle
// and returns the perimeter of it.
func (t Triangle) Perimeter() float64 {
	return t.Height + t.Base + t.Height
}

// Area takes the width and height and returns the
// area of the triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Rectangle is a 4-sided shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter takes the height and width of a rectangle
// and returns the perimeter of it.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area takes the width and height and returns the
// area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a 1-sided shape.
type Circle struct {
	Radius float64
}

// Perimeter takes the height and width of a rectangle
// and returns the circle of it.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area takes the width and height and returns the
// area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
