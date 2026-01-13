package main

import "fmt"

func deleteAt(arr []int, index int)[]int{
	for i := index; i< len(arr)-1; i++{
		arr[i] = arr[i+1]
	}
	arr = arr[:len(arr)-1]
	return arr
}

func main(){
	arr := []int{5,8,7,9,6,1,10}
	fmt.Println(arr)

	result := deleteAt(arr, 2)
	fmt.Println(result)
}