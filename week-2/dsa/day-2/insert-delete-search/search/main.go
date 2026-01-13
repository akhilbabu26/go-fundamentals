package main

import "fmt"

// binary search (it only work on sorted list O(log n))
func binarySearch(arr []int, target int)int{
	left := 0
	right := len(arr) - 1

	for left <= right{
		mid := (left+right)/2

		if arr[mid] == target{
			return mid
		}else if arr[mid] < target{
			left = mid + 1
		}else{
			right = mid -1
		}
	}
	return -1
}

// Linear search (it work on any list but time O(n))
func linearSearch(arr []int, target int)int{
	for i, v:= range arr{
		if v == target{
			return i
		}
	}
	return -1
}

func main(){
	// binary search
	arr1 := []int{10,15,17,19,32,55}
	result1 := binarySearch(arr1, 19)
	fmt.Println("result:",result1)

	// Linear search
	arr2 := []int{10,15,17,19,32,55}
	result2 := linearSearch(arr2, 19)
	fmt.Println("result:",result2)
}