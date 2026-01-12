package main

import "fmt"

func main(){
	arr := []int{1,4,5,8,9,25}

	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}