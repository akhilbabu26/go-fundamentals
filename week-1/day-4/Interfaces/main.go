package main

import (
	"fmt"
	"math"
)

type Result interface {
	Area() float64
	Perimeter() float64
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Common function using interface
func result(s Result) {
	fmt.Println("Area      :", s.Area())
	fmt.Println("Perimeter :", s.Perimeter())
	fmt.Println("------")
}

func main() {
	c := Circle{Radius: 6}
	r := Rectangle{Width: 5, Height: 8}

	fmt.Println("Circle:")
	result(c)

	fmt.Println("Rectangle:")
	result(r)
}