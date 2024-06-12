package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {
	c := Circle{Radius: 5}
	fmt.Println("Circle properties:")
	PrintShapeProperties(c)

	r := Rectangle{Width: 5, Height: 10}
	fmt.Println("Rectangle properties:")
	PrintShapeProperties(r)
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func PrintShapeProperties(s Shape) {
	println("Area: ", s.Area())
	println("Perimeter: ", s.Perimeter())
}
