package main

import "fmt"

type Person struct{
	Name string
	Age int
	IsAlive bool
}

func (p Person)Greet(){
	fmt.Printf("Name is %s and he is %d years old",p.Name,p.Age)
}

func main(){
	person := Person{
		Name: "Akhil",
		Age: 24,
		IsAlive: true,
	}

	person.Greet()
}