package main

import "fmt"

func reverse(arr []int){
	first := 0
	last := len(arr)-1

	for first < last{
		arr[first],arr[last] = arr[last], arr[first]
		first++
		last--
	}
}

func main(){
	arr := []int{1,2,4,5,6,8,9}
	fmt.Println("Befor reverse:",arr)
	reverse(arr)
	fmt.Println("After reverse:",arr)
}