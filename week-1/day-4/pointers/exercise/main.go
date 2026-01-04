package main

import "fmt"

func Add(x *int, y int)int{
	*x = *x + y
	return *x
}

func main(){
	value := 20
	fmt.Println("Value brfore update:", value)

	Add(&value,50)
	fmt.Println("Value after update:", value)
}