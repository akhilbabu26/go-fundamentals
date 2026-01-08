package main

import (
	"fmt"
	// "errors"
)

type Person struct{
	Name string
	Age int
}

var person []Person

func AddPerson(name string, age int){

	person = append(person, Person{
		Name: name,
		Age: age,
	})
	
}

func UpdatePerson(i int, name string, age int){
	person[i]= Person{
		Name: name,
		Age: age,
	}
}

func Delete(i int){
	person = append(person[:i],person[i+1:]...)
}

func ListOfStudents(){
	fmt.Println(person)
}

func main(){
	AddPerson("Akhil",24)
	AddPerson("Akhil",24)
	AddPerson("Akhil",24)
	ListOfStudents()

	UpdatePerson(1,"manu",34)
	ListOfStudents()

	Delete(2)
	ListOfStudents()

}