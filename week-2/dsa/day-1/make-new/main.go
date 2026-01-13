package main

import "fmt"

type Students struct{
	Name string
	Age int
}

func main(){
	// using new it gives zero value
		// int
	n := new(int)
	fmt.Println(*n)
		// string
	n1 := new(string)
	fmt.Println(*n1)
		// struct
	n2 := new(Students)// dynamic allocation
	n2.Name = "Akhil"
	n2.Age = 24
	fmt.Println("Student:", *n2)

	// using make give ready-to-use values
		// map
	m := make(map[string]int)
	m["Apple"] = 40
	m["orange"] = 100
	fmt.Println("map:",m)
		//channel
	m1 := make(chan int)
	go func(){
		m1 <- 30
	}()
	value := <- m1
	fmt.Println("channel:",value)
		// slice
	m2 := make([]int,3)
	for i := range m2{
		m2[i] = i+1
	}
	fmt.Println("slice:",m2)
}