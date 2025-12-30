package main

import "fmt"

func divide(a float64, b float64)(result float64, ok bool){
	if b == 0{
		result = 0
		ok = false
		return
	}
	result = a/b
	ok = true
	return
}

func main(){
	r, success := divide(81,2)

	if success{
		fmt.Println("Result",r)
	}else{
		fmt.Println("Can't divisible by zero")
	}
	//
	r, success = divide(5,0) // = is an assingment operator

	if success{
		fmt.Println("Result",r)
	}else{
		fmt.Println("Can't divisible by zero")
	}

}