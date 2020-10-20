package shapes

import "math"

// Rectangle ...
type Rectangle struct {
	Base   float64
	Height float64
}

// Area ...
func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}

// Perimeter ...
func (r Rectangle) Perimeter() float64 {
	return 2.0 * (r.Base + r.Height)
}

// Circle ...
type Circle struct {
	Radius float64
}

// Area ...
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Perimeter ...
func (c Circle) Perimeter() float64 {
	return 2.0 * math.Pi * c.Radius
}

// Shape ...
type Shape interface {
	Area() float64
	Perimeter() float64
}
