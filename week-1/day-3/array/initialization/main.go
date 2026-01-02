package main

import "fmt"

func main(){
	// initial value will be zero
	a := [5]int{}
	fmt.Println(a)

	//with values & updationg values
	b := [2]int{10,20}
	fmt.Println(b)
	b[0] = 55
	fmt.Println(b)

	// go inter size means it automatically deside the size by how many elements in there using [...]int
	c := [...]int{10,20,55}
	fmt.Println(c)

	//indexed initialiation
	d := [5]int{0:5,2:10}// it will fill automatically fill by initial value
	fmt.Println(d)

}