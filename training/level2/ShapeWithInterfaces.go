package main

import (
	"fmt"
	"math"
)

/*
Create a Rectangle struct and a Circle struct, both struct should have two Method:Area() Perimeter()
And then create a Shape interface for those struct.
After that create a function that will receive a Shape interface as parameter and will execute the Area() and the Perimeter() from each struct.

Topics to Practice:
Interfaces, struct, method, function, math pkg
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radious float64
}

type Rectangle struct {
	length float64
	width float64
}

func (c Circle) Area() (a float64) {
	return math.Pi * math.Pow(c.radious, 2)
}
func (c Circle) Perimeter() (a float64) {
	return 2 * math.Pi * c.radious
}

func (c Rectangle) Area() (a float64) {
	return c.length * c.width
}

func (c Rectangle) Perimeter() (a float64) {
	return c.length * c.width * 2
}

func ExecuteArea(s Shape) {
	fmt.Println(s.Area())
}