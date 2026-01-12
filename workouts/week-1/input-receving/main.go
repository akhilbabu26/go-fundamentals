package main

import (
	"fmt"
	"errors"
)

type Students struct{
	Name string
	Age int
	Class string
}

var students []Students

func AddStudents(name string, age int, class string)error{
	if name == "" || age <= 0 || class == ""{
		return errors.New("ERROR: Invalid Input!!")
	}

	students = append(students, Students{
		Name: name,
		Age: age,
		Class: class,
	})

	return nil
}

func main(){

	var name string
	var age int
	var class string
	for{
		fmt.Println("Enter name age class(if want to exit type q)")

		fmt.Scan(&name)

		if name == "q"{
			fmt.Println("Exixting...")
			break
		}

		fmt.Scan(&age, &class)

		err := AddStudents(name,age,class)
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println("Current student list:")
		fmt.Println(students)
	}
	
}