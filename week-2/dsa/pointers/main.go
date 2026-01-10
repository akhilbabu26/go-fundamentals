package main

import "fmt"

func main(){
	x := 50
	fmt.Println("Value of x =",x)

	p := &x
	*p = 100

	fmt.Println("After Value of x =",x)
}