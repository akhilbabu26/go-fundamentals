package main

import "fmt"

func main(){
	//normal for loop
	a := [...]int{1,2,3,4,5,6}

	for i := 0; i < len(a); i++{
		fmt.Println("index:",i, "value:",a[i])
	}
	fmt.Println("____________________")

	// for loop with range
	b := [...]int{9,8,7,6,5,4,3,2,1}

	for index , value := range b{
		fmt.Println("index:",index, "value:",value)
	}
	fmt.Println("____________________")
	
	//or:
	d := [...]int{80,70,60}

	for _,value := range d{
		fmt.Println("value:",value)
	}
	fmt.Println("____________________")
	
	//or:
	g := [...]int{80,70,60}

	for index,_ := range g{
		fmt.Println("value:",index)
	}
	fmt.Println("____________________")

	//or simplyfied method:
	c := [...]int{10,20,30,40,50,60}

	for f := range c{
		fmt.Println("index:",f, "value:",c[f])
	}
}