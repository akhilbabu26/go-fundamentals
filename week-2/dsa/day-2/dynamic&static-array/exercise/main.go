package main

import "fmt"

func reverse(arr []int){
	left := 0
	right := len(arr)-1

	for left < right{
		arr[left], arr[right] = arr[right], arr[left]
		left ++
		right --
	}
}

func main(){
	arr := []int{1,2,3,4,5,6,7,8,9}
	reverse(arr)
	fmt.Println(arr)
}