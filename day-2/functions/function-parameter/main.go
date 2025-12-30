package main

import "fmt"

func add(a float64, b float64) float64{
	return a+b
}

func main(){
	result := add(5,2)
	fmt.Println(result)
}