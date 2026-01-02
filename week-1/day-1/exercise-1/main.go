package main

import "fmt"

func main(){
	num := 2

	if num %2 == 0{
		fmt.Println(num, "is an Even number")
	}else{
		fmt.Println(num, "is an Odd number")
	}
}