package main

import "fmt"

// func insert(arr []int, index int, val int)[]int{
// 	if index > len(arr) || index <0{
// 		fmt.Println("Error: invelid input")
// 		return nil
// 	}

// 	arr = append(arr, 0)
// 	copy(arr[index+1:], arr[index:])
// 	arr[index] = val
// 	return arr
// }

func insert(arr []int, index int, val int)[]int{
	arr = append(arr, 0)

	for i := len(arr)-1; i > index; i--{
		arr[i] = arr[i-1]
	}
	arr[index]=val
	return arr
}

func main(){
	arr := []int{5,7,9,1,3,2,5,4}
	fmt.Println(arr)
	result := insert(arr, 5,100)
	fmt.Println(result)
}