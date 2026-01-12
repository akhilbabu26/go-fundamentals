package calculator

import "fmt"

type Circle struct{
	Pi float64
	Radius float64
}

func (a Circle)Area(){
	result := a.Pi*(a.Radius*a.Radius)
	fmt.Println("Area:",result)
}