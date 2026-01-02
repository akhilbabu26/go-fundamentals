package main

import (
	"mathfunc/math"
	"fmt"
)

func main(){
	fmt.Println("ADD",math.Add(5,5))
	fmt.Println("SUB",math.Sub(5,8))
	fmt.Println("MUL",math.Mul(9,5))

	result, err := math.Div(5,0)

	if err != nil{
		fmt.Println("ERROR:",err)
	}else{
		fmt.Println("Div:",result)
	}
}