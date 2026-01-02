package main

import (
	"fmt"
	"strconv"
)

func main(){
	//int to float
	var a int = 15
	var b float64 = float64(a)

	fmt.Println(b)

	// Float to int
	var x float64 = 10.99
	var y int = int(x)

	fmt.Println(y)

	//different int types
	var f int = 5
	var g int64 = int64(f)

		fmt.Println(g)

	// byte <-> string

	   //1. str to int (use strconv.Atoi)
	var s string = "123"

	var n int 
	var err error 
	n, err = strconv.Atoi(s)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(n)

	   //2. int to str (use strconv.Itoa)
	var m int = 456
	var str string = strconv.Itoa(m)

	fmt.Println(str)
}