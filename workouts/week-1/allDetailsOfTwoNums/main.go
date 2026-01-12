package main

import "fmt"

func allDetails(a float64, b float64)(add float64, sub float64, mul float64, div float64, ok bool){
	add = a+b
	sub = a-b
	mul = a*b

	if b == 0{
		div = 0
		ok = false
		return
	}

	div = a/b
	ok = true
	return
}

func main(){
	add, sub, mul, div, success := allDetails(5,3)

	fmt.Println("Added:",add)
	fmt.Println("Difference:",sub)
	fmt.Println("Multiplied:",mul)
	
	if success{
		fmt.Println("Divided:",div)
		}else{
		fmt.Println("Can't divided by zero")
	}
}