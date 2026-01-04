package main

import "fmt"

func main(){
	x := 20
	p := &x

	fmt.Println("Value of p:", *p)
}