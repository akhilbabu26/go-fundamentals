package main

// import (
// 	"fmt"
//  "errors"	
// )

// func Mark(a float64)(string, error){
// 	if a < 0 || a > 100{
// 		return "fail",errors.New("Error: invalid input")
// 	}
// 	if a <= 50 {
// 		return "fail",nil
// 	}
// 	if a >= 50{
// 		return "pass",nil
// 	}
// 	return "pass",nil
// }

// func main(){
// 	r, err := Mark(11)
// 	if err !=nil{
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(r)
// }

// package main 

// import "fmt"

// func main(){
// 	arr := [5]int{1,2,3,4,5}

// 	max := arr[0]

// 	for i := 0; i< len(arr); i++{
// 		if arr[i] > max{
// 			max = arr[i]
// 		}
// 	}
// 	fmt.Println(max)
// }