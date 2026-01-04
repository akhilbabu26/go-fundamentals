package main

import "fmt"

func main(){
	//creating a pointer
	x1 := 10
	p1 := &x1 // it stores memory of x1 

	fmt.Println("Value of x1:",x1)
	fmt.Println("Address of x1:",p1)
	fmt.Println("--------------------")

	// dereferencing means accessing the value of pointer address points
	x2 := 10
	p2 := &x2
	fmt.Println("value of p2:", *p2)
	fmt.Println("--------------------")
	
	// mdifying pointer value. it also changes original
	x3 := 10
	p3 := &x3
	fmt.Println("befor value of x3:",x3)
	*p3 = 50
	fmt.Println("after value of x3:",x3)
}