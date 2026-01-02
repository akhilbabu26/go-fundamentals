package main

import "fmt"

func main(){

	// classic for loop
	for i :=0; i<5; i++{
		fmt.Println(i)
	}
	fmt.Println("---------")

	// for as a while loop
	w := 0
	for w <5{
		fmt.Println("Hello")
		w++
	}
	fmt.Println("---------")

	// infinite loop(it will go infinitely if we dont use return or break)
	for{
		fmt.Println("Running...")
		break
	}
	fmt.Println("---------")

	// for with range;
	   // 1. slice
	   nums := []int{10,20,30}

	   for index, value := range nums{
		   fmt.Println(index, value)
	   }
	   fmt.Println("---------")
	   //2. map
	   m := map[string]int{"a": 1, "b": 2}
	   for key, value := range m{
		   fmt.Println(key, value)
	   }
	   fmt.Println("---------")

	//loop control inside "for"
	   //1. break
	for x :=0; x <=5; x++{
		if x == 3{
			break
		}
		fmt.Println(x)
	}
	fmt.Println("---------")
		//2. comtinue 
	for j :=0; j <=5; j++{
		if j == 3{
			continue // skips 3 and continue
		}
		fmt.Println(j)
	}
	fmt.Println("---------")
		//3. return
	for x :=0; x <=5; x++{
		if x == 4{
			return //it will return output and end main
		}
		fmt.Println(x)
	}
	fmt.Println("---------")
}