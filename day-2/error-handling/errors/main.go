package main

import (
	"fmt"
	"errors"
)

func div(a, b float64)(float64, error){
	if b==0{
		return 0, errors.New("Can't divisible by zero")
	}
	return  a/b,nil
}

func main(){
	result, err := div(2.5,0)

	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println(result)
}