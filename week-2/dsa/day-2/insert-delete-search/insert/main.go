package main

import "fmt"

func insert(arr []int, index int, value int)[]int{
	if index < 0 || index > len(arr) {
		return arr
	}
	arr = append(arr, 0)

	for i := len(arr)-1; i > index; i--{
		arr[i] = arr[i-1]
	}

	arr[index] = value
	return arr
}

func main(){
	arr := []int{80,30,70,10,20}
	fmt.Println("Befor insert:", arr)
	
	result := insert(arr,2,1000)
	fmt.Println("After insert:", result)
}
