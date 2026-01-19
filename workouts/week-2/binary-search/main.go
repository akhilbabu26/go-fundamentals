package main

import "fmt"

func binarySearch(arr []int, target int)bool{
	head := 0
	tail := len(arr)-1

	for head <= tail{
		mid := (head + tail) /2

		if arr[mid] == target{
			return true
		}else if arr[mid] < target{
			head = mid + 1
		}else {
			tail = mid - 1
		}
	}
	return false
}

func main(){
	arr := []int{1,5,6,7,8,9,15,17,19}

	result := binarySearch(arr, 9)
	fmt.Println(result)
}