// Shape with Interfaces
// Statement
// Create a Rectangle struct and a Circle struct, both struct should have
// two Method:Area() Perimeter()
// And then create a Shape interface for those struct. After that create a
// function that will receive a Shape interface as parameter and will execute the Area()
// and the Perimeter() from each struct.

// Topics to Practice:
// Interfaces, struct, method, function, math pkg

package main

import (
	"fmt"

	"github.com/jgfiuza/golang-bootcamp/level2/03_shapeWithInterface/shapes"
)

func areaAndPerimeter(s shapes.Shape) {
	switch s.(type) {
	case shapes.Circle:
		fmt.Printf("The area of the circle is %.2f and its perimeter is %.2f\n", s.Area(), s.Perimeter())
	case shapes.Rectangle:
		fmt.Printf("The area of the rectangle is %.2f and its perimeter is %.2f\n", s.Area(), s.Perimeter())

	}
}

func main() {
	circle := shapes.Circle{5.0}
	square := shapes.Rectangle{4.0, 4.0}
	rectangle := shapes.Rectangle{3.0, 5.0}

	shapes := []shapes.Shape{circle, square, rectangle}

	for _, shape := range shapes {
		areaAndPerimeter(shape)
	}

}
