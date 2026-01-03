package main

import "fmt"

func main(){
	// creating map
	fruits := make(map[string]int)

	// add key and values
	fruits["apple"] = 10
	fruits["orange"] = 50
	fruits["mango"] = 70
	fruits["banana"] = 20
	fruits["grape"] = 300

	// update value with key
	fruits["apple"] = 100

	fmt.Println(fruits)
	fmt.Println("--------------------")

	// iterating
	for k, v := range fruits{
		fmt.Println(k,v)
	}
	fmt.Println("--------------------")

	//deleting
	delete(fruits,"banana")
	fmt.Println(fruits)

	for k, v := range fruits{
		fmt.Println(k,v)
	}

	fmt.Println("length:",len(fruits))

}