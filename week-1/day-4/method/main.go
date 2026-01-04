package main

import "fmt"

type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle)Area(){
	fmt.Println("Area:", r.Width*r.Height)
}

func main(){
	rect := Rectangle{
		Width: 25,
		Height: 2,
	}
	rect.Area()
}