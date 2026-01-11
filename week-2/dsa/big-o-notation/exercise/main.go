package main

import "fmt"

func findMax(arr []int)int{
	max := arr[0]

	for i := range arr{
		if arr[i] > max{
			max = arr[i]// the timr complexity is O(n) beacuse the algorithm must check every element at least once.
		}
	}
	return max
}

func main(){
	arr := []int{5,8,1,15,100,122,51}
	result := findMax(arr)
	fmt.Println("Result:",result)
}