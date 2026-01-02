package main

import "fmt"

func main(){
	//1. defer
	defer fmt.Println("1")
	defer fmt.Println("2")// LEFO order output = 3,2,1
	defer fmt.Println("3")
	defer fmt.Println("first")
	fmt.Println("last")
	// return // it will givea warning
}