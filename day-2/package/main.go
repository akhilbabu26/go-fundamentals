package main

import (
	"myproject/calculator"
	"fmt"
)

func main(){
	fmt.Println("ADD:",calculator.Add(5,56))
	fmt.Println("SUB:",calculator.Sub(5,56))
	fmt.Println("MUL:",calculator.Mul(5,56))

	result, err := calculator.Div(5,2)
	if err != nil{
		fmt.Println("ERROR:", err)
	}else{
		fmt.Println("Div:",result)
	}
	
}