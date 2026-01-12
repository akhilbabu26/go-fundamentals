package main

import (
	"packages_and_methods/calculator"
)

func main(){
	area := calculator.Circle{
		Pi: 3.14159,
		Radius: 5,
	}

	area.Area()
}