package main

import "fmt"

func main(){
	//1. break
	for i:=1; i<=5; i++{
		if i == 4{
			break
		}
		fmt.Println(i)
	}
	fmt.Println("-------")

	//2. continue
	for j:=1; j<=5; j++{
		if j == 2{
			continue
		}
		fmt.Println(j)
	}
	fmt.Println("-------")

	//3. goto

	x := 1
	loop:
	if x <=3{
		fmt.Println(x)
		x++
		goto loop
	}
	fmt.Println("-------")
}