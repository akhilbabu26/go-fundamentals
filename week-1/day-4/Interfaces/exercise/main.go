package main

import (
	"math"
	"fmt"
)

type Shape interface{
	Area() float64
	Perimeter() float64
}

// circle shape = area and perimeter
type Circle struct{
	Radius float64
}

func (c Circle)Area()float64{
	return math.Pi * c.Radius * c.Radius
}

func (c Circle)Perimeter()float64{
	return 2 * math.Pi * c.Radius
}

// rectangle shape = area and perimeter
type Rectangle struct{
	Width, Height float64
}

func (r Rectangle)Area()float64{
	return r.Height * r.Width
}

func (r Rectangle)Perimeter()float64{
	return 2 *(r.Height+r.Width)
}

// helper
func result(s Shape){
	fmt.Println("Area:",s.Area())
	fmt.Println("Parimeter:",s.Perimeter())
	fmt.Println("----------------")
}

func main(){
	c := Circle{Radius: 5}
	r := Rectangle{Width: 5, Height: 10}

	fmt.Println("Circle shape:")
	result(c)

	fmt.Println("Reactangle shape:")
	result(r)
}