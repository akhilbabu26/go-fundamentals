package main

import "fmt"

func main(){
	// arr := [6]int{5,10,55,22,7,100}

	// min := arr[0]
	// max := arr[0]

	// for i := 0; i < len(arr); i++{
	// 	if arr[i] < min{
	// 		min = arr[i]
	// 	}
	// 	if arr[i] > max{
	// 		max = arr[i]
	// 	}
	// }

	//using range;

	arr := [6]int{10,51,1,8,60,67}

	min := arr[0]
	max := arr[0]

	for i := range arr{
		if arr[i] < min{
			min = arr[i]
		}
		if arr[i] > max{
			max = arr[i]
		}
	}
	
	fmt.Println("Array:",arr)
	fmt.Println("Minimum:",min)
	fmt.Println("Maximium:",max)
}