package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Method with a pointer receiver
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with a value receiver
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
}
