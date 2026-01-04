package main

import "fmt"

func main(){
	x := 10
	p := &x // p is a pointer &helps to store memory address

	fmt.Println("x value:", x)
	fmt.Println("x address:", p)

	
}