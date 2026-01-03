package main

import "fmt"

type Person struct{
	Name string
	Age int
	IsAlive bool
}

func (p *Person)SetAll(name string, age int, alive bool){
	*p = Person{
		Name: name,
		Age: age,
		IsAlive: alive,
	}
}

func (p *Person) Reset(){
	*p = Person{}
}

func main(){
	person := Person{
		Name: "Akhil",
		Age: 24,
		IsAlive: true,
	}
	fmt.Println("befor setAge:", person)

	person.SetAll("Junaid", 20,true)
	fmt.Println("After setAge:", person)
	
	person.SetAll("manu", 55, false)
	fmt.Println("After setAge:", person)

	person.Reset()
	fmt.Println("After reset:",person)

	person.SetAll("manu", 55, true)
	fmt.Println("After setAge:", person)
}