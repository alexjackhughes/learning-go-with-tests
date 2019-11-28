package main

import "fmt"

// Perimeter takes the height and width of a rectangle
// and returns the perimeter of it.
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area takes the width and height and returns the
// area of the rectangle.
func Area(width, height float64) float64 {
	return width * height
}

func main() {
	fmt.Println(Perimeter(10.0, 10.0))
}
