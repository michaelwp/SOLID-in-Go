package main

import "fmt"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle struct
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// AreaCalculator calculates the area of a shape
func AreaCalculator(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 7}

	AreaCalculator(r)
	AreaCalculator(c)
}