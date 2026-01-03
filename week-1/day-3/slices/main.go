package main

import "fmt"

func main(){
	// Create a slice using a literal
	s1 := []int{1,3,5,7,9}
	fmt.Println("Slice:",s1,"Len(s1):",len(s1),"Cap(s1):",cap(s1))
	fmt.Println("-----------------------")

	// Create a slice from an array
	arr2 := [5]int{5,1,5,2,3}
	s2 := arr2[1:4] //it copies arr2's index 1 to 4
	fmt.Println("Array:", arr2) 
	fmt.Println("s2:", s2,"Len(s1):",len(s2),"Cap(s1):",cap(s2))// cap is 4 beacuse it count form 1 to 4
	fmt.Println("-----------------------")

	// Reslicing (same underlying array) 
	s3 := s2[:2]
	fmt.Println("s2:", s2) 
	fmt.Println("s3:", s3) 
	s3[0]=100 // changing s3 also change s2 beacuse it share the same memory and array
	fmt.Println("s2:", s2) 
	fmt.Println("s3:", s3) 
	fmt.Println("arr2:", arr2) // <-- this is the original array it also changed 
	fmt.Println("-----------------------")

	// Append without reallocation
	s4 := make([]int,4,6)
	fmt.Println("s4:", s4)
	s4[0], s4[1] = 5,8
	fmt.Println("s4:", s4)
	fmt.Println("-----------------------")

	//  Append WITH reallocation
	s5 := []int{1,2,3,4}
	s6 := s5
	fmt.Println("s6:", s6)
	//Force reallocation (exceed capacity)
	s5 = append(s6,8,11,10,20)
	fmt.Println("s5:", s5, "new array")//if capacity exceeds it crates a new array
	fmt.Println("s6:", s6, "old array")// it will connect ot old array
	fmt.Println("-----------------------")

	// coping slice to avoid sharing
	src := []int{2,8,6,4,7,9}
	dst := make([]int, len(src))
	copy(dst,src)

	dst[0] = 999 
	fmt.Println("src =", src)
	fmt.Println("dst =", dst)
	fmt.Println("-----------------------")

	// ------- 7. Nil slice vs Empty slice -------
	var nilSlice []int
	emptySlice := []int{}

	fmt.Println("nilSlice   =", nilSlice, "len =", len(nilSlice), "cap =", cap(nilSlice))
	fmt.Println("emptySlice =", emptySlice, "len =", len(emptySlice), "cap =", cap(emptySlice))
}