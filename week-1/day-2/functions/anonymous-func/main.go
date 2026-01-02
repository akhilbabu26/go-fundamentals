package main

import "fmt"

// function expression 
	var add = func(a float64, b float64)float64{
		return a+b
	}

func main(){
	//IIFE
	func(){
		fmt.Println("Hello")
	}()

	// function expression 
	fmt.Println(add(5,3))
}