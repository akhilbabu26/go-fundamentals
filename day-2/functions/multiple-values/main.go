package main

import "fmt"

func divide(a float64, b float64)(float64, bool){
	if b == 0{
		return 0, false
	}
	return a/b, true
}

func main(){
	result, ok := divide(85,2)

	if ok{
		fmt.Println("Result:",result)
	}else{
		fmt.Println("can't devisible by zero")
	}
}