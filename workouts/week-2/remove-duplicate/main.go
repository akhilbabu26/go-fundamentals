package main

import "fmt"

func removeDuplicate(arr []int)[]int{
	seen := make(map[int]bool)
	result := []int{}

	for _,v := range arr{
		if !seen[v]{
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main(){
	arr := []int{1,4,5,7,7,8,2,1,5,5,11,4,8,9,7,4,11}
	fmt.Println(arr)
	result := removeDuplicate(arr)
	fmt.Println(result)
}