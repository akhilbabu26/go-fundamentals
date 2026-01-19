package main

import "fmt"

func delete(arr []int, index int)[]int{
	if index > len(arr) || index < 0{
		fmt.Println("Enter valid index!!")
		return arr
	}

	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

func main(){
	arr := []int{1,7,5,9,8}
	fmt.Println(arr)
	result := delete(arr, 1)
	fmt.Println(result)
}