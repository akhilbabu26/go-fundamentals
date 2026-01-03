package main

import "fmt"

//struct
type Person struct{
	Name string
	Age int
}

//method
func (p Person)Details(){ 
	fmt.Printf(" Name: %s\n Age: %d", p.Name,p.Age)
}

func main(){
	person := Person{
		Name: "Akhil",
		Age: 24,
	}

	person.Details()
}