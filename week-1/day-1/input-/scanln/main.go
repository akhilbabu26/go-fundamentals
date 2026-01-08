package main

import "fmt"

func main(){
	var name string
	var age int

	fmt.Scan(&name, &age)
	fmt.Println("Name and age:",name, age)
}