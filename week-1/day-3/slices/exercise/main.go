package main

import (
	"fmt"
)

func removeDuplicates(nums []int)[]int{
	result := []int{}

	for _, v1 := range nums{

		duplicate := false
		for _, v2 := range result{
			if v1 == v2{
				duplicate = true
				break
			}
		}

		if !duplicate{
			result = append(result, v1)
		}
	}
	return result
}

func main(){
	numbres := []int{1,1,7,8,10,8,2,3,2}

	noDuplicates := removeDuplicates(numbres)

	fmt.Println("Numbres:",numbres)
	fmt.Println("Numbres:",noDuplicates)

}